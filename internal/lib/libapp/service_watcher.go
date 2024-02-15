package libapp

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/hashicorp/consul/api"
)

type HealthChecker struct {
	healthClient *api.Health
	serviceName  string
}

func NewHealthChecker(serviceName string, c *conf.Consul) (*HealthChecker, error) {
	config := api.DefaultConfig()
	if c != nil {
		config.Address = c.GetAddr()
		config.Token = c.GetToken()
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	// check if the consul agent is available
	_, err = client.Status().Leader()
	if err != nil {
		if c == nil {
			return &HealthChecker{
				healthClient: nil,
				serviceName:  serviceName,
			}, nil
		}
		return nil, err
	}

	return &HealthChecker{
		healthClient: client.Health(),
		serviceName:  serviceName,
	}, nil
}

func (hc *HealthChecker) GetAliveInstances() ([]*api.ServiceEntry, error) {
	if hc.healthClient == nil {
		return []*api.ServiceEntry{}, nil
	}
	instances, _, err := hc.healthClient.Service(hc.serviceName, "", true, nil)
	return instances, err
}
