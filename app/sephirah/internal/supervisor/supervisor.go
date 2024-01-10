package supervisor

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libauth"
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
	featureSummary         *modeltiphereth.ServerFeatureSummary
	muFeatureSummary       sync.RWMutex
	trustedAddresses       []string
}

func NewSupervisor(
	c *conf.Sephirah_Porter,
	auth *libauth.Auth,
	porter *client.Porter,
) (*Supervisor, error) {
	if c == nil {
		c = new(conf.Sephirah_Porter)
	}
	return &Supervisor{
		porter:                 porter,
		auth:                   auth,
		aliveInstances:         map[string]*modeltiphereth.PorterInstance{},
		knownInstances:         nil, // init in UpdateKnownInstances
		knownInstancesOutdated: true,
		featureSummary:         new(modeltiphereth.ServerFeatureSummary),
		muFeatureSummary:       sync.RWMutex{},
		trustedAddresses:       c.GetTrustedAddress(),
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

func (s *Supervisor) RefreshAliveInstances( //nolint:gocognit // TODO
	ctx context.Context,
) ([]*modeltiphereth.PorterInstance, error) {
	if s.knownInstances == nil {
		return nil, fmt.Errorf("known instances not set")
	}
	addresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
		return nil, err
	}
	newInstances := make([]*modeltiphereth.PorterInstance, 0, len(addresses))
	aliveInstanceMap := make(map[string]*modeltiphereth.PorterInstance, len(addresses))
	var info *porter.GetPorterInformationResponse
	for _, address := range addresses {
		if address == "" {
			// bad address
			continue
		}
		info, err = s.porter.GetPorterInformation(
			client.WithPorterAddress(ctx, address),
			&porter.GetPorterInformationRequest{},
		)
		if err != nil {
			// bad instance
			logger.Infof("%s", err.Error())
			continue
		}
		if info == nil {
			// bad instance
			continue
		}
		feature := converter.ToBizPorterFeatureSummary(info.GetFeatureSummary())
		var ins *modeltiphereth.PorterInstance
		if s.knownInstances[address] != nil { //nolint:nestif // TODO
			// known instance
			if s.knownInstances[address].GlobalName != info.GetGlobalName() {
				// bad instance, global name changed
				continue
			}
			ins = s.knownInstances[address]
			if ins.Status == modeltiphereth.PorterInstanceStatusActive {
				// enable & check ownership
				if err2 := s.enablePorterInstance(ctx, ins); err2 != nil {
					logger.Errorf("%s", err2.Error())
					// bad instance, can't enable
					continue
				}
			}
			if reflect.DeepEqual(ins.FeatureSummary, feature) {
				// no change, but alive
				aliveInstanceMap[address] = ins
				continue
			}
			ins.FeatureSummary = feature
		} else {
			// new instance
			ins = &modeltiphereth.PorterInstance{
				ID:             0,
				Name:           info.GetName(),
				Version:        info.GetVersion(),
				GlobalName:     info.GetGlobalName(),
				Address:        address,
				FeatureSummary: feature,
				Status:         modeltiphereth.PorterInstanceStatusUnspecified,
			}
		}
		// new instance or feature changed
		newInstances = append(newInstances, ins)
		aliveInstanceMap[address] = ins
	}
	s.aliveInstances = aliveInstanceMap
	if len(newInstances) > 0 {
		go s.updateFeatureSummary()
	}
	return newInstances, nil
}

// EnablePorterInstance enable porter instance, can be called multiple times.
func (s *Supervisor) enablePorterInstance(ctx context.Context, instance *modeltiphereth.PorterInstance) error {
	if instance == nil {
		return fmt.Errorf("instance is nil")
	}
	refreshToken, err := s.auth.GenerateToken(
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
	return err
}
