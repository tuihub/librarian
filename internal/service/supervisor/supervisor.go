package supervisor

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizsupervisor"
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"github.com/google/wire"
	"github.com/panjf2000/ants/v2"
)

const (
	defaultHeartbeatInterval  = libtime.Second * 5
	defaultHeartbeatDowngrade = libtime.Second * 30
	defaultHeartbeatTimeout   = libtime.Second * 60
	defaultPoolSize           = 32
)

var ProviderSet = wire.NewSet(NewSupervisorService)

type SupervisorService struct {
	s            *bizsupervisor.Supervisor
	cron         *libcron.Cron
	job          *libcron.Job
	porter       *client.Porter
	t            *data.TipherethRepo
	auth         *libauth.Auth
	systemNotify *libmq.Topic[modelnetzach.SystemNotify]

	refreshMu                 sync.Mutex
	trustedAddresses          []string
	instanceController        *libtype.SyncMap[string, *bizsupervisor.PorterInstanceController]
	instanceContextController *libtype.SyncMap[model.InternalID, *bizsupervisor.PorterContextController]
	featureController         *bizsupervisor.PorterFeatureController
}

func NewSupervisorService(
	s *bizsupervisor.Supervisor,
	c *conf.Porter,
	auth *libauth.Auth,
	porter *client.Porter,
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
	cron *libcron.Cron,
	t *data.TipherethRepo,
	featureController *bizsupervisor.PorterFeatureController,
) (*SupervisorService, error) {
	if c == nil {
		c = new(conf.Porter)
	}
	res := SupervisorService{
		s:            s,
		cron:         cron,
		job:          nil,
		porter:       porter,
		t:            t,
		auth:         auth,
		systemNotify: systemNotify,

		refreshMu:                 sync.Mutex{},
		trustedAddresses:          c.Addresses,
		instanceController:        libtype.NewSyncMap[string, *bizsupervisor.PorterInstanceController](),
		instanceContextController: libtype.NewSyncMap[model.InternalID, *bizsupervisor.PorterContextController](),
		featureController:         featureController,
	}
	return &res, nil
}

func (s *SupervisorService) Start(ctx context.Context) error {
	job, err := s.cron.NewJobByDuration(
		"SupervisorService Heartbeat",
		s.getHeartbeatInterval(),
		func() {
			err := s.heartbeat(context.Background())
			if err != nil {
				logger.Errorf("refresh alive instances failed: %s", err.Error())
			}
		},
	)
	if err != nil {
		return fmt.Errorf("failed to register cron: %w", err)
	}
	s.job = job
	return nil
}

func (s *SupervisorService) Stop(ctx context.Context) error {
	if s.job != nil {
		if err := s.job.Cancel(); err != nil {
			return fmt.Errorf("failed to stop cron job: %w", err)
		}
		s.job = nil
	}
	return nil
}

func (s *SupervisorService) getHeartbeatInterval() time.Duration {
	return defaultHeartbeatInterval
}

func (s *SupervisorService) heartbeat( //nolint:gocognit,funlen // TODO
	ctx context.Context,
) error {
	if !s.refreshMu.TryLock() {
		return errors.New("refresh in progress")
	}
	defer s.refreshMu.Unlock()

	discoveredAddresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return err
	}
	enabledContexts, err := s.t.GetEnabledPorterContexts(context.Background())
	if err != nil {
		logger.Errorf("get enabled porter contexts failed: %s", err.Error())
		return err
	}
	hasError := false
	notification := modelnetzach.NewSystemNotify(
		modelnetzach.SystemNotificationLevelOngoing,
		fmt.Sprintf("%s: Refresh Porter Instances", modelnetzach.SystemNotifyTitleCronJob),
		fmt.Sprintf("Found %d Porter Instances", len(discoveredAddresses)),
	)
	notificationMu := &sync.Mutex{}
	pool, err := ants.NewPool(defaultPoolSize)
	if err != nil {
		logger.Errorf("failed to create ants pool: %s", err.Error())
		return fmt.Errorf("failed to create ants pool: %w", err)
	}
	wg := sync.WaitGroup{}
	availableAddresses := map[string][]string{}
	availableAddressesMu := sync.Mutex{}

	// Discover new instances
	for _, address := range discoveredAddresses {
		s.instanceController.LoadOrStore(address, s.s.NewPorterInstanceController(address))
	}
	// Save new instance contexts
	for _, con := range enabledContexts {
		if con == nil {
			continue
		}
		s.instanceContextController.LoadOrStore(con.ID, s.s.NewPorterContextController(
			con,
		))
	}

	// Instance Heartbeat
	s.instanceController.Range(func(address string, ctl *bizsupervisor.PorterInstanceController) bool {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()
			ctxWithTimeout, cancel := context.WithTimeout(ctx, libtime.Minute)
			defer cancel()
			instance := ctl.UpdateStatus(ctxWithTimeout)
			if instance.ConnectionStatus != modelsupervisor.PorterConnectionStatusActive &&
				instance.ConnectionStatus != modelsupervisor.PorterConnectionStatusDowngraded {
				notificationMu.Lock()
				defer notificationMu.Unlock()
				hasError = true
				notification.Notification.Content += fmt.Sprintf(
					"\nInstance %s is not active: %s",
					instance.Address,
					instance.ConnectionStatusMessage,
				)
				return
			}
			s.featureController.Update(instance)
			key := fmt.Sprintf("%s:%s", instance.GlobalName, instance.Region)
			availableAddressesMu.Lock()
			defer availableAddressesMu.Unlock()
			if _, ok := availableAddresses[key]; !ok {
				availableAddresses[key] = []string{}
			}
			availableAddresses[key] = append(availableAddresses[key], instance.Address)
		})
		return true
	})
	wg.Wait()
	s.featureController.Commit()

	// Instance Context Heartbeat
	s.instanceContextController.Range(func(ctxID model.InternalID, ctl *bizsupervisor.PorterContextController) bool {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()
			ctxWithTimeout, cancel := context.WithTimeout(ctx, libtime.Minute)
			defer cancel()
			key := fmt.Sprintf("%s:%s", ctl.GetGlobalName(), ctl.GetRegion())
			if addresses, ok := availableAddresses[key]; !ok || len(addresses) == 0 {
				notificationMu.Lock()
				defer notificationMu.Unlock()
				hasError = true
				notification.Notification.Content += fmt.Sprintf(
					"\nContext %s do not has active instance",
					ctl.GetName(),
				)
				return
			} else {
				handlerAddress := ctl.UpdateStatus(ctxWithTimeout, addresses)
				if handlerAddress == "" {
					notificationMu.Lock()
					defer notificationMu.Unlock()
					hasError = true
					notification.Notification.Content += fmt.Sprintf("\nContext %s is not active", ctl.GetName())
					return
				}
			}
		})
		return true
	})
	wg.Wait()

	if hasError {
		notification.Notification.Level = modelnetzach.SystemNotificationLevelError
		_ = s.systemNotify.PublishFallsLocalCall(ctx, notification)
	}

	return nil
}
