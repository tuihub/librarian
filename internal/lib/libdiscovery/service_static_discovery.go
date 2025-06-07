package libdiscovery

import (
	"context"
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/v2/registry"
)

type StaticDiscovery struct {
	serviceInstances []*registry.ServiceInstance
	watcherCount     int
	watcherCh        chan struct{}
}

func NewStaticDiscovery(addresses []string, name, version string) (*StaticDiscovery, error) {
	var serviceInstances []*registry.ServiceInstance
	for i, addr := range addresses {
		parsed, err := url.Parse(addr)
		if err != nil {
			return nil, fmt.Errorf("static discovery address %s invalid: %w", addr, err)
		}
		if parsed.Scheme != "grpc" && parsed.Scheme != "grpcs" {
			return nil, fmt.Errorf("static discovery address %s is not a valid gRPC address", addr)
		}
		serviceInstances = append(serviceInstances, &registry.ServiceInstance{
			ID:        fmt.Sprintf("%s-%d", name, i),
			Name:      name,
			Version:   version,
			Metadata:  nil,
			Endpoints: []string{addr},
		})
	}
	return &StaticDiscovery{
		serviceInstances: serviceInstances,
		watcherCount:     0,
		watcherCh:        make(chan struct{}),
	}, nil
}

func (s *StaticDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return s.serviceInstances, nil
}

func (s *StaticDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return s, nil
}

func (s *StaticDiscovery) Next() ([]*registry.ServiceInstance, error) {
	if s.watcherCount > 0 {
		<-s.watcherCh
	}
	s.watcherCount++
	return s.serviceInstances, nil
}

func (s *StaticDiscovery) Stop() error {
	s.watcherCh <- struct{}{}
	return nil
}

func (s *StaticDiscovery) GetAliveInstances() ([]string, error) {
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
