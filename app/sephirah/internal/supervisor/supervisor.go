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
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/libtype"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
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

	refreshMu                 sync.Mutex
	trustedAddresses          []string
	instanceController        *libtype.SyncMap[string, modelsupervisor.PorterInstanceController]
	instanceContextController *libtype.SyncMap[model.InternalID, modelsupervisor.PorterContextController]
	instanceCache             *libcache.Map[string, modelsupervisor.PorterInstance]
	instanceContextCache      *libcache.Map[model.InternalID, modelsupervisor.PorterContext]
	enableContextTopic        *libmq.Topic[modelsupervisor.PorterContext]

	// featureSummary contains unique enabled features
	featureSummary *modelsupervisor.ServerFeatureSummary
	// featureSummaryMap is the map from enabled feature ID to instance addresses
	featureSummaryMap  *modelsupervisor.ServerFeatureSummaryMap
	featureSummaryRWMu sync.RWMutex
}

func NewSupervisor(
	c *conf.Porter,
	mq *libmq.MQ,
	auth *libauth.Auth,
	porter *client.Porter,
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
	instanceCache *libcache.Map[string, modelsupervisor.PorterInstance],
	instanceContextCache *libcache.Map[model.InternalID, modelsupervisor.PorterContext],
) (*Supervisor, error) {
	if c == nil {
		c = new(conf.Porter)
	}
	res := Supervisor{
		UUID:         int64(uuid.New().ID()),
		porter:       porter,
		auth:         auth,
		systemNotify: systemNotify,

		refreshMu:                 sync.Mutex{},
		trustedAddresses:          c.GetTrusted(),
		instanceController:        libtype.NewSyncMap[string, modelsupervisor.PorterInstanceController](),
		instanceContextController: libtype.NewSyncMap[model.InternalID, modelsupervisor.PorterContextController](),
		instanceCache:             instanceCache,
		instanceContextCache:      instanceContextCache,
		enableContextTopic:        nil,

		featureSummary:     new(modelsupervisor.ServerFeatureSummary),
		featureSummaryMap:  modelsupervisor.NewServerFeatureSummaryMap(),
		featureSummaryRWMu: sync.RWMutex{},
	}
	res.enableContextTopic = newEnablePorterContextTopic(&res)
	err := mq.RegisterTopic(res.enableContextTopic)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Supervisor) GetHeartbeatInterval() time.Duration {
	return defaultHeartbeatInterval
}

func (s *Supervisor) GetInstanceController(
	ctx context.Context,
	address string,
) *modelsupervisor.PorterInstanceController {
	if c, ok := s.instanceController.Load(address); ok {
		return &c
	}
	return nil
}

func (s *Supervisor) GetContextController(
	ctx context.Context,
	id model.InternalID,
) *modelsupervisor.PorterContextController {
	if c, ok := s.instanceContextController.Load(id); ok {
		return &c
	}
	return nil
}

func (s *Supervisor) RefreshAliveInstances( //nolint:gocognit,funlen // TODO
	ctx context.Context,
) ([]*modelsupervisor.PorterInstance, error) {
	if !s.refreshMu.TryLock() {
		return nil, errors.New("refresh in progress")
	}
	defer s.refreshMu.Unlock()

	discoveredAddresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, err
	}
	newInstances := make([]*modelsupervisor.PorterInstance, 0, len(discoveredAddresses))
	newInstancesMu := sync.Mutex{}
	updateFeatureSummary := false
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
		if ic, ok := s.instanceController.Load(address); ok &&
			ic.ConnectionStatus != modelsupervisor.PorterConnectionStatusDisconnected {
			continue
		}

		wg.Add(1)
		go func(ctx context.Context, address string) {
			defer wg.Done()
			var ins *modelsupervisor.PorterInstance
			ins, err = s.evaluatePorterInstance(ctx, address)
			if err != nil {
				logger.Errorf("%s", err.Error())
				hasError = true
				notifyCh <- fmt.Sprintf("Error on %s: %s", address, err.Error())
				return
			}

			if ic, ok := s.instanceController.Load(address); ok ||
				(ic.GlobalName != ins.GlobalName || ic.BinarySummary.BuildVersion != ins.BinarySummary.Version) {
				newInstancesMu.Lock()
				newInstances = append(newInstances, ins)
				newInstancesMu.Unlock()
			}
			s.instanceController.Store(address, modelsupervisor.PorterInstanceController{
				PorterInstance:          *ins,
				ConnectionStatus:        modelsupervisor.PorterConnectionStatusConnected,
				ConnectionStatusMessage: "",
				LastHeartbeat:           time.Now(),
				LastEnabledContext:      nil,
			})
		}(ctx, address)
	}

	// Heartbeat
	s.instanceController.Range(func(address string, ctl modelsupervisor.PorterInstanceController) bool {
		var ins *modelsupervisor.PorterInstance
		if ins, err = s.instanceCache.Get(ctx, address); err != nil ||
			ins == nil ||
			ins.Status != modeltiphereth.UserStatusActive {
			return true
		}

		wg.Add(1)
		go func(ctx context.Context, ins *modelsupervisor.PorterInstance) {
			defer wg.Done()
			var resp *porter.EnablePorterResponse
			resp, err = s.enablePorterInstance(ctx, ins)
			if err != nil {
				logger.Errorf("%s", err.Error())
				hasError = true
				notifyCh <- fmt.Sprintf("Error on %s: %s", ins.Address, err.Error())
				// no return
			}

			if err != nil {
				ctl.ConnectionStatusMessage += fmt.Sprintf("Error: %s", err.Error())
			} else if resp != nil {
				ctl.ConnectionStatusMessage = resp.GetStatusMessage()
			}
			lastConnectionStatus := ctl.ConnectionStatus
			err = s.updateControllers(ctx, &ctl, resp)
			if err != nil {
				logger.Errorf("%s", err.Error())
				hasError = true
				notifyCh <- fmt.Sprintf("Error on %s: %s", ins.Address, err.Error())
			}
			if lastConnectionStatus != ctl.ConnectionStatus {
				updateFeatureSummary = true
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

	if updateFeatureSummary {
		go s.updateFeatureSummary(ctx)
	}
	return newInstances, nil
}

func (s *Supervisor) evaluatePorterInstance(
	ctx context.Context,
	address string,
) (*modelsupervisor.PorterInstance, error) {
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
	if info == nil || info.GetBinarySummary() == nil {
		// bad instance
		return nil, errors.New("bad instance info")
	}
	feature := converter.ToBizPorterFeatureSummary(info.GetFeatureSummary())
	return &modelsupervisor.PorterInstance{
		ID:                0,
		BinarySummary:     converter.ToBizPorterBinarySummary(info.GetBinarySummary()),
		GlobalName:        info.GetGlobalName(),
		Address:           address,
		Region:            info.GetRegion(),
		FeatureSummary:    feature,
		Status:            modeltiphereth.UserStatusUnspecified,
		ContextJSONSchema: info.GetContextJsonSchema(),
	}, nil
}

// enablePorterInstance enable porter instance, can be called multiple times.
func (s *Supervisor) enablePorterInstance(
	ctx context.Context,
	instance *modelsupervisor.PorterInstance,
) (*porter.EnablePorterResponse, error) {
	if instance == nil {
		return nil, errors.New("instance is nil")
	}
	resp, err := s.porter.EnablePorter(client.WithPorterAddress(ctx, []string{instance.Address}), &porter.EnablePorterRequest{
		SephirahId:   s.UUID,
		RefreshToken: nil,
	})
	if err != nil {
		return nil, err
	}
	if resp.GetNeedRefreshToken() {
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
			return resp, err
		}
		resp, err = s.porter.EnablePorter(client.WithPorterAddress(ctx, []string{instance.Address}), &porter.EnablePorterRequest{
			SephirahId:   s.UUID,
			RefreshToken: &refreshToken,
		})
	}
	return resp, err
}

func (s *Supervisor) updateControllers(
	ctx context.Context,
	ctl *modelsupervisor.PorterInstanceController,
	resp *porter.EnablePorterResponse,
) error {
	now := time.Now()
	if resp == nil { //nolint:nestif // TODO
		if ctl.LastHeartbeat.Add(defaultHeartbeatTimeout).Before(now) {
			ctl.ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
		} else if ctl.LastHeartbeat.Add(defaultHeartbeatDowngrade).Before(now) {
			ctl.ConnectionStatus = modelsupervisor.PorterConnectionStatusDowngraded
		} else if ctl.ConnectionStatus == modelsupervisor.PorterConnectionStatusActive {
			ctl.ConnectionStatus = modelsupervisor.PorterConnectionStatusActive
		} else {
			ctl.ConnectionStatus = modelsupervisor.PorterConnectionStatusActivationFailed
		}
	} else {
		ctl.ConnectionStatus = modelsupervisor.PorterConnectionStatusActive
		ctl.LastHeartbeat = now
	}
	if ctl.ConnectionStatus == modelsupervisor.PorterConnectionStatusActive && resp == nil {
		return nil
	}
	var reQueueContext []model.InternalID
	if ctl.ConnectionStatus == modelsupervisor.PorterConnectionStatusActive {
		onlyLast, _ := libtype.DiffSlices(
			ctl.LastEnabledContext,
			converter.ToBizInternalIDList(resp.GetEnablesSummary().GetContextIds()),
		)
		reQueueContext = onlyLast
		ctl.LastEnabledContext = converter.ToBizInternalIDList(resp.GetEnablesSummary().GetContextIds())
	} else {
		reQueueContext = ctl.LastEnabledContext
		if ctl.ConnectionStatus == modelsupervisor.PorterConnectionStatusDowngraded {
			ctl.LastEnabledContext = nil
		}
	}
	for _, id := range reQueueContext {
		ic, ok := s.instanceContextController.Load(id)
		if !ok {
			continue
		}
		ic.HandleStatus = modelsupervisor.PorterContextHandleStatusQueueing
		ic.HandleStatusMessage = ""
		ic.HandlerAddress = ""
		s.instanceContextController.Store(id, ic)
		err := s.QueuePorterContext(ctx, ic.PorterContext)
		if err != nil {
			return err
		}
	}
	return nil
}
