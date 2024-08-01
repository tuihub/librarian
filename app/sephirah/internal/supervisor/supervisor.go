package supervisor

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/lib/logger"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/uuid"
	"github.com/google/wire"
)

const (
	defaultHeartbeatInterval  = libtime.Second * 10
	defaultHeartbeatDowngrade = libtime.Second * 30
	defaultHeartbeatTimeout   = libtime.Second * 60
)

var ProviderSet = wire.NewSet(NewSupervisor)

type Supervisor struct {
	UUID         int64
	porter       *client.Porter
	auth         *libauth.Auth
	systemNotify *libmq.Topic[modelnetzach.SystemNotify]

	refreshMu          sync.Mutex
	trustedAddresses   []string
	instanceController *libtype.SyncMap[modeltiphereth.PorterInstanceController]
	instanceCache      *libcache.Map[string, modeltiphereth.PorterInstance]

	featureSummary     *modeltiphereth.ServerFeatureSummary
	featureSummaryMap  *modeltiphereth.ServerFeatureSummaryMap
	featureSummaryRWMu sync.RWMutex
}

func NewSupervisor(
	c *conf.Porter,
	auth *libauth.Auth,
	porter *client.Porter,
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
	instanceCache *libcache.Map[string, modeltiphereth.PorterInstance],
) (*Supervisor, error) {
	if c == nil {
		c = new(conf.Porter)
	}
	return &Supervisor{
		UUID:               int64(uuid.New().ID()),
		porter:             porter,
		auth:               auth,
		instanceController: libtype.NewSyncMap[modeltiphereth.PorterInstanceController](),
		instanceCache:      instanceCache,
		refreshMu:          sync.Mutex{},
		featureSummary:     new(modeltiphereth.ServerFeatureSummary),
		featureSummaryMap:  modeltiphereth.NewServerFeatureSummaryMap(),
		featureSummaryRWMu: sync.RWMutex{},
		trustedAddresses:   c.GetTrusted(),
		systemNotify:       systemNotify,
	}, nil
}

func (s *Supervisor) GetHeartbeatInterval() time.Duration {
	return defaultHeartbeatInterval
}

func (s *Supervisor) GetInstanceConnectionStatus(
	ctx context.Context,
	address string,
) modeltiphereth.PorterConnectionStatus {
	value := s.instanceController.Load(address)
	if value == nil {
		return modeltiphereth.PorterConnectionStatusDisconnected
	}
	return value.ConnectionStatus
}

func (s *Supervisor) RefreshAliveInstances( //nolint:gocognit,funlen // TODO
	ctx context.Context,
) ([]*modeltiphereth.PorterInstance, error) {
	if !s.refreshMu.TryLock() {
		return nil, errors.New("refresh in progress")
	}
	defer s.refreshMu.Unlock()

	discoveredAddresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, err
	}
	newInstances := make([]*modeltiphereth.PorterInstance, 0, len(discoveredAddresses))
	newInstancesMu := sync.Mutex{}
	hasError := false
	notification := modelnetzach.NewSystemNotify(
		modelnetzach.SystemNotificationLevelOngoing,
		fmt.Sprintf("%s: Refresh Porter Instances", modelnetzach.SystemNotifyTitleCronJob),
		fmt.Sprintf("Found %d Porter Instances", len(discoveredAddresses)),
	)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(ctx, libtime.Minute)
	defer cancel()
	notifyCh := make(chan string)
	done := make(chan struct{})

	// Discover new instances and Refresh disconnected instances
	for _, address := range discoveredAddresses {
		if ic := s.instanceController.Load(address); ic != nil &&
			ic.ConnectionStatus != modeltiphereth.PorterConnectionStatusDisconnected {
			continue
		}

		wg.Add(1)
		go func(ctx context.Context, address string) {
			defer wg.Done()
			var ins *modeltiphereth.PorterInstance
			ins, err = s.evaluatePorterInstance(ctx, address)
			if err != nil {
				logger.Errorf("%s", err.Error())
				hasError = true
				notifyCh <- fmt.Sprintf("Error on %s: %s", address, err.Error())
				return
			}

			if ic := s.instanceController.Load(address); ic == nil ||
				(ic.GlobalName != ins.GlobalName || ic.Version != ins.Version) {
				newInstancesMu.Lock()
				newInstances = append(newInstances, ins)
				newInstancesMu.Unlock()
			}
			s.instanceController.Store(address, modeltiphereth.PorterInstanceController{
				PorterInstance:   *ins,
				ConnectionStatus: modeltiphereth.PorterConnectionStatusConnected,
				LastHeartbeat:    time.Now(),
			})
		}(ctx, address)
	}

	// Heartbeat
	s.instanceController.Range(func(address string, ctl modeltiphereth.PorterInstanceController) bool {
		var ins *modeltiphereth.PorterInstance
		if ins, err = s.instanceCache.Get(ctx, address); err != nil ||
			ins.Status != modeltiphereth.PorterInstanceStatusActive {
			return true
		}

		wg.Add(1)
		go func(ctx context.Context, ins *modeltiphereth.PorterInstance) {
			defer wg.Done()
			err = s.enablePorterInstance(ctx, ins)
			if err != nil {
				logger.Errorf("%s", err.Error())
				hasError = true
				notifyCh <- fmt.Sprintf("Error on %s: %s", ins.Address, err.Error())
				// no return
			}

			now := time.Now()
			if err != nil {
				if ctl.LastHeartbeat.Add(defaultHeartbeatTimeout).Before(now) {
					ctl.ConnectionStatus = modeltiphereth.PorterConnectionStatusDisconnected
				} else if ctl.LastHeartbeat.Add(defaultHeartbeatDowngrade).Before(now) {
					ctl.ConnectionStatus = modeltiphereth.PorterConnectionStatusActivationFailed
				} else {
					ctl.ConnectionStatus = modeltiphereth.PorterConnectionStatusActive
				}
			} else {
				ctl.ConnectionStatus = modeltiphereth.PorterConnectionStatusActive
				ctl.LastHeartbeat = now
			}

			s.instanceController.Store(ins.Address, ctl)
		}(ctx, ins)

		return true
	})

	// Save notification messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-notifyCh:
				notification.Notification.Content += "\n" + msg
			}
		}
	}()

	// Trigger done
	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait for goroutines or timeout
	select {
	case <-done:
		notifyCh <- fmt.Sprintf("Found %d Porter Instances", len(newInstances))
	case <-ctx.Done():
	}

	if hasError {
		notification.Notification.Level = modelnetzach.SystemNotificationLevelError
		_ = s.systemNotify.PublishFallsLocalCall(ctx, notification)
	}

	if len(newInstances) > 0 {
		go s.updateFeatureSummary(ctx)
	}
	return newInstances, nil
}

func (s *Supervisor) evaluatePorterInstance(
	ctx context.Context,
	address string,
) (*modeltiphereth.PorterInstance, error) {
	if address == "" {
		// bad address
		return nil, errors.New("address is empty")
	}
	info, err := s.porter.GetPorterInformation(
		client.WithPorterAddress(ctx, []string{address}),
		&porter.GetPorterInformationRequest{},
	)
	if err != nil {
		// bad instance
		logger.Infof("%s", err.Error())
		return nil, err
	}
	if info == nil {
		// bad instance
		return nil, errors.New("info is nil")
	}
	feature := converter.ToBizPorterFeatureSummary(info.GetFeatureSummary())
	return &modeltiphereth.PorterInstance{
		ID:                0,
		Name:              info.GetName(),
		Version:           info.GetVersion(),
		GlobalName:        info.GetGlobalName(),
		Address:           address,
		Region:            info.GetRegion(),
		FeatureSummary:    feature,
		Status:            modeltiphereth.PorterInstanceStatusUnspecified,
		ContextJSONSchema: info.GetContextJsonSchema(),
	}, nil
}

// enablePorterInstance enable porter instance, can be called multiple times.
func (s *Supervisor) enablePorterInstance(ctx context.Context, instance *modeltiphereth.PorterInstance) error {
	if instance == nil {
		return errors.New("instance is nil")
	}
	_, err := s.porter.EnablePorter(client.WithPorterAddress(ctx, []string{instance.Address}), &porter.EnablePorterRequest{
		SephirahId:   s.UUID,
		RefreshToken: nil,
	})
	if err != nil {
		var refreshToken string
		refreshToken, err = s.auth.GenerateToken(
			instance.ID,
			0,
			libauth.ClaimsTypeRefreshToken,
			libauth.UserTypePorter,
			nil,
			libtime.Hour,
		)
		if err != nil {
			return err
		}
		_, err = s.porter.EnablePorter(client.WithPorterAddress(ctx, []string{instance.Address}), &porter.EnablePorterRequest{
			SephirahId:   s.UUID,
			RefreshToken: &refreshToken,
		})
	}
	return err
}
