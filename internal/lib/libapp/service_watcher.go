package libapp

import (
	"fmt"
	"net"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/hashicorp/consul/api"
)

type HealthChecker interface {
	GetAliveInstances() ([]string, error)
}

type healthChecker struct {
	healthClient *api.Health
	serviceName  string
}

func NewHealthChecker(serviceName string, c *conf.Consul) (HealthChecker, error) {
	config := api.DefaultConfig()
	if c != nil {
		config.Address = net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port)))
		config.Token = c.Token
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	// check if the consul agent is available
	_, err = client.Status().Leader()
	if err != nil {
		if c == nil {
			return &healthChecker{
				healthClient: nil,
				serviceName:  serviceName,
			}, nil
		}
		return nil, err
	}

	return &healthChecker{
		healthClient: client.Health(),
		serviceName:  serviceName,
	}, nil
}

func (hc *healthChecker) GetAliveInstances() ([]string, error) {
	if hc.healthClient == nil {
		return []string{}, nil
	}
	instances, _, err := hc.healthClient.Service(hc.serviceName, "", true, nil)
	res := make([]string, 0, len(instances))
	for _, instance := range instances {
		res = append(res, fmt.Sprintf("%s:%d", instance.Service.Address, instance.Service.Port))
	}
	return res, err
}

func IsEmptyHealthChecker(hc HealthChecker) bool {
	h, ok := hc.(*healthChecker)
	return !ok || h == nil || h.healthClient == nil
}
