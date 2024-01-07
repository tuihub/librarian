package libapp

import (
	"github.com/hashicorp/consul/api"
)

type HealthChecker struct {
	healthClient *api.Health
	serviceName  string
}

func NewHealthChecker(serviceName string) (*HealthChecker, error) {
	config := api.DefaultConfig()
	config.Address = DiscoveryAddress
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &HealthChecker{
		healthClient: client.Health(),
		serviceName:  serviceName,
	}, nil
}

func (hc *HealthChecker) GetAliveInstances() ([]*api.ServiceEntry, error) {
	instances, _, err := hc.healthClient.Service(hc.serviceName, "", true, nil)
	return instances, err
}
