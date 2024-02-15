package libapp

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	capi "github.com/hashicorp/consul/api"
)

type requiredFeatureKey struct{}

func NewDiscovery(c *conf.Consul) (registry.Discovery, error) {
	config := capi.DefaultConfig()
	if c != nil {
		config.Address = c.GetAddr()
		config.Token = c.GetToken()
	}
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	// check if the consul agent is available
	_, err = client.Status().Leader()
	if err != nil {
		if c == nil { // if consul is not configured, return an empty discovery
			return emptyDiscovery{}, nil
		}
		return nil, err
	}
	return consul.New(client), nil
}

func NewRegistrar(c *conf.Consul) (registry.Registrar, error) {
	config := capi.DefaultConfig()
	if c != nil {
		config.Address = c.GetAddr()
		config.Token = c.GetToken()
	}
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return consul.New(client), nil
}

func NewContext(ctx context.Context, requiredFeature string) context.Context {
	return context.WithValue(ctx, requiredFeatureKey{}, requiredFeature)
}

func NewNodeFilter() selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		rff, ok := ctx.Value(requiredFeatureKey{}).(string)
		if !ok {
			return nodes
		}
		newNodes := make([]selector.Node, 0)
		for _, n := range nodes {
			n.InitialWeight()
			if _, exist := n.Metadata()[rff]; exist {
				newNodes = append(newNodes, n)
			}
		}
		return newNodes
	}
}

type emptyDiscovery struct{}

func (e emptyDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return []*registry.ServiceInstance{}, nil
}

func (e emptyDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return &emptyWatcher{}, nil
}

type emptyWatcher struct{}

func (e emptyWatcher) Next() ([]*registry.ServiceInstance, error) {
	return []*registry.ServiceInstance{}, nil
}

func (e emptyWatcher) Stop() error {
	return nil
}
