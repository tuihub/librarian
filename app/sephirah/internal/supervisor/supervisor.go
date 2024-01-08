package supervisor

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/logger"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSupervisor)

type Supervisor struct {
	porter           *client.Porter
	instances        []*modelsupervisor.PorterInstance
	featureSummary   *modelsupervisor.ServerFeatureSummary
	trustedAddresses []string
}

func NewSupervisor(
	c *conf.Sephirah_Porter,
	porter *client.Porter,
	cron *libcron.Cron,
) (*Supervisor, error) {
	if c == nil {
		c = new(conf.Sephirah_Porter)
	}
	supv := &Supervisor{
		porter:           porter,
		instances:        nil,
		featureSummary:   nil,
		trustedAddresses: c.GetTrustedAddress(),
	}
	err := cron.BySeconds(5, supv.RefreshPorterInstances, context.Background()) //nolint:gomnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return supv, nil
}

func (s *Supervisor) GetFeatureSummary() *modelsupervisor.ServerFeatureSummary {
	return s.featureSummary
}

func (s *Supervisor) RefreshPorterInstances(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, libtime.Minute)
	defer cancel()
	addresses, err := s.porter.GetServiceAddresses(ctx)
	if err != nil {
		logger.Errorf("%s", err.Error())
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
			logger.Infof("%s", err.Error())
			continue
		}
		feature := new(modelsupervisor.PorterFeatureSummary)
		instances = append(instances, &modelsupervisor.PorterInstance{
			ID:             0,
			Name:           info.GetName(),
			Version:        info.GetVersion(),
			GlobalName:     info.GetGlobalName(),
			Address:        address,
			FeatureSummary: feature,
			Status:         modelsupervisor.PorterInstanceStatusBlocked,
		})
		features = append(features, feature)
	}
	s.instances = instances
	s.featureSummary = modelsupervisor.Summarize(features)
	return nil
}
