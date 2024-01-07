package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSupervisor)

type Supervisor struct {
	porter         *client.Porter
	instances      []*modelsupervisor.PorterInstance
	featureSummary *modelsupervisor.ServerFeatureSummary
}

func NewSupervisor(porter *client.Porter) *Supervisor {
	return &Supervisor{
		porter:         porter,
		instances:      nil,
		featureSummary: nil,
	}
}

func (s *Supervisor) RefreshPorterInstances(ctx context.Context) error {
	addresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		return err
	}
	instances := make([]*modelsupervisor.PorterInstance, 0, len(addresses))
	features := make([]*modelsupervisor.PorterFeatureSummary, 0, len(addresses))
	var info *porter.GetPorterInformationResponse
	for _, address := range addresses {
		info, err = s.porter.GetPorterInformation(
			client.WithPorterAddress(ctx, address),
			&porter.GetPorterInformationRequest{},
		)
		if err != nil {
			return err
		}
		feature := new(modelsupervisor.PorterFeatureSummary)
		instances = append(instances, &modelsupervisor.PorterInstance{
			ID:             0,
			Name:           info.GetName(),
			Version:        info.GetVersion(),
			GlobalName:     info.GetGlobalName(),
			Address:        address,
			FeatureSummary: feature,
		})
		features = append(features, feature)
	}
	s.instances = instances
	s.featureSummary = modelsupervisor.Summarize(features)
	return nil
}
