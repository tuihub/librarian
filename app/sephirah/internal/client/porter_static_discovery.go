package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2/registry"
)

type staticDiscovery struct {
	serviceInstances []*registry.ServiceInstance
	watcherCount     int
	watcherCh        chan struct{}
}

func newStaticDiscovery(c *conf.Porter) (*staticDiscovery, error) {
	if c == nil {
		c = new(conf.Porter)
	}
	var serviceInstances []*registry.ServiceInstance
	for i, addr := range c.GetAddress() {
		parsed, err := url.Parse(addr)
		if err != nil {
			return nil, fmt.Errorf("porter address %s invalid: %w", addr, err)
		}
		if parsed.Scheme != "grpc" && parsed.Scheme != "grpcs" {
			return nil, fmt.Errorf("porter address %s is not a valid gRPC address", addr)
		}
		serviceInstances = append(serviceInstances, &registry.ServiceInstance{
			ID:        fmt.Sprintf("porter-%d", i),
			Name:      "porter",
			Version:   "",
			Metadata:  nil,
			Endpoints: []string{addr},
		})
	}
	return &staticDiscovery{
		serviceInstances: serviceInstances,
		watcherCount:     0,
		watcherCh:        make(chan struct{}),
	}, nil
}

func (s *staticDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return s.serviceInstances, nil
}

func (s *staticDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return s, nil
}

func (s *staticDiscovery) Next() ([]*registry.ServiceInstance, error) {
	if s.watcherCount > 0 {
		<-s.watcherCh
	}
	s.watcherCount++
	return s.serviceInstances, nil
}

func (s *staticDiscovery) Stop() error {
	s.watcherCh <- struct{}{}
	return nil
}

func (s *staticDiscovery) GetAliveInstances() ([]string, error) {
	res := make([]string, 0, len(s.serviceInstances))
	for _, instance := range s.serviceInstances {
		parsed, err := url.Parse(instance.Endpoints[0])
		if err != nil {
			continue
		}
		res = append(res, fmt.Sprintf("%s:%s", parsed.Hostname(), parsed.Port()))
	}
	return res, nil
}
