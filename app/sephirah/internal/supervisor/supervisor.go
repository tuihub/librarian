package supervisor

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSupervisor)

type Supervisor struct {
	porter                 *client.Porter
	auth                   *libauth.Auth
	aliveInstances         map[string]*modeltiphereth.PorterInstance
	knownInstances         map[string]*modeltiphereth.PorterInstance
	knownInstancesOutdated bool
	refreshMu              sync.Mutex
	featureSummary         *modeltiphereth.ServerFeatureSummary
	muFeatureSummary       sync.RWMutex
	trustedAddresses       []string
	systemNotify           *libmq.Topic[modelnetzach.SystemNotify]
}

func NewSupervisor(
	c *conf.Porter,
	auth *libauth.Auth,
	porter *client.Porter,
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
) (*Supervisor, error) {
	if c == nil {
		c = new(conf.Porter)
	}
	return &Supervisor{
		porter:                 porter,
		auth:                   auth,
		aliveInstances:         map[string]*modeltiphereth.PorterInstance{},
		knownInstances:         nil, // init in UpdateKnownInstances
		knownInstancesOutdated: true,
		refreshMu:              sync.Mutex{},
		featureSummary:         new(modeltiphereth.ServerFeatureSummary),
		muFeatureSummary:       sync.RWMutex{},
		trustedAddresses:       c.GetTrustedAddress(),
		systemNotify:           systemNotify,
	}, nil
}

func (s *Supervisor) KnownInstancesOutdated() {
	s.knownInstancesOutdated = true
}

func (s *Supervisor) KnownInstancesRequireUpdate() bool {
	return s.knownInstances == nil || s.knownInstancesOutdated
}

func (s *Supervisor) UpdateKnownInstances(instances []*modeltiphereth.PorterInstance) {
	if s.knownInstances == nil {
		s.knownInstances = map[string]*modeltiphereth.PorterInstance{}
	}
	for _, instance := range instances {
		s.knownInstances[instance.Address] = instance
	}
	go s.updateFeatureSummary()
}

func (s *Supervisor) RefreshAliveInstances(
	ctx context.Context,
) ([]*modeltiphereth.PorterInstance, error) {
	s.refreshMu.Lock()
	defer s.refreshMu.Unlock()

	if s.knownInstances == nil {
		return nil, errors.New("known instances not set")
	}
	addresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, err
	}
	newInstances := make([]*modeltiphereth.PorterInstance, 0, len(addresses))
	aliveInstanceMap := make(map[string]*modeltiphereth.PorterInstance, len(addresses))
	hasError := false
	notification := modelnetzach.NewSystemNotify(
		modelnetzach.SystemNotificationLevelOngoing,
		fmt.Sprintf("%s: Refresh Porter Instances", modelnetzach.SystemNotifyTitleCronJob),
		fmt.Sprintf("Found %d Porter Instances", len(addresses)),
	)

	for _, address := range addresses {
		var ins *modeltiphereth.PorterInstance
		var isNew bool
		ins, isNew, err = s.evaluatePorterInstance(ctx, address)
		if err != nil {
			logger.Errorf("%s", err.Error())
			hasError = true
			notification.Notification.Content += "\n" + fmt.Sprintf("Error on %s: %s", address, err.Error())
		}
		if ins != nil {
			aliveInstanceMap[address] = ins
			if isNew {
				newInstances = append(newInstances, ins)
			}
		}
	}

	if hasError {
		notification.Notification.Level = modelnetzach.SystemNotificationLevelError
		_ = s.systemNotify.PublishFallsLocalCall(ctx, notification)
	}

	s.aliveInstances = aliveInstanceMap
	if len(newInstances) > 0 {
		go s.updateFeatureSummary()
	}
	return newInstances, nil
}

func (s *Supervisor) GetInstanceConnectionStatus(
	ctx context.Context,
	address string,
) modeltiphereth.PorterConnectionStatus {
	if s.aliveInstances[address] == nil {
		return modeltiphereth.PorterConnectionStatusDisconnected
	}
	return s.aliveInstances[address].ConnectionStatus
}

func (s *Supervisor) evaluatePorterInstance(
	ctx context.Context,
	address string,
) (*modeltiphereth.PorterInstance, bool, error) {
	if address == "" {
		// bad address
		return nil, false, errors.New("address is empty")
	}
	info, err := s.porter.GetPorterInformation(
		client.WithPorterAddress(ctx, address),
		&porter.GetPorterInformationRequest{},
	)
	if err != nil {
		// bad instance
		logger.Infof("%s", err.Error())
		return nil, false, err
	}
	if info == nil {
		// bad instance
		return nil, false, errors.New("info is nil")
	}
	feature := converter.ToBizPorterFeatureSummary(info.GetFeatureSummary())
	var ins *modeltiphereth.PorterInstance
	if s.knownInstances[address] != nil { //nolint:nestif // TODO
		// known instance
		if s.knownInstances[address].GlobalName != info.GetGlobalName() {
			// bad instance, global name changed
			return nil, false, errors.New("global name changed")
		}
		ins = s.knownInstances[address]
		ins.ConnectionStatus = modeltiphereth.PorterConnectionStatusConnected
		if ins.Status == modeltiphereth.PorterInstanceStatusActive {
			// enable & check ownership
			if err2 := s.enablePorterInstance(ctx, ins); err2 != nil {
				logger.Errorf("%s", err2.Error())
				ins.ConnectionStatus = modeltiphereth.PorterConnectionStatusActivationFailed
				// bad instance, can't enable
				return ins, false, err2
			}
			ins.ConnectionStatus = modeltiphereth.PorterConnectionStatusActive
		}
		if reflect.DeepEqual(ins.FeatureSummary, feature) {
			// no change, but alive
			return ins, false, nil
		}
		ins.FeatureSummary = feature
	} else {
		// new instance
		ins = &modeltiphereth.PorterInstance{
			ID:               0,
			Name:             info.GetName(),
			Version:          info.GetVersion(),
			GlobalName:       info.GetGlobalName(),
			Address:          address,
			FeatureSummary:   feature,
			Status:           modeltiphereth.PorterInstanceStatusUnspecified,
			ConnectionStatus: modeltiphereth.PorterConnectionStatusConnected,
		}
	}
	// new instance or feature changed
	return ins, true, nil
}

// EnablePorterInstance enable porter instance, can be called multiple times.
func (s *Supervisor) enablePorterInstance(ctx context.Context, instance *modeltiphereth.PorterInstance) error {
	if instance == nil {
		return errors.New("instance is nil")
	}
	_, err := s.porter.EnablePorter(client.WithPorterAddress(ctx, instance.Address), &porter.EnablePorterRequest{
		SephirahId:   0,
		RefreshToken: "",
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
		_, err = s.porter.EnablePorter(client.WithPorterAddress(ctx, instance.Address), &porter.EnablePorterRequest{
			SephirahId:   0,
			RefreshToken: refreshToken,
		})
	}
	return err
}
