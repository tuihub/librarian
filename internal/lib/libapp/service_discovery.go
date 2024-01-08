package libapp

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	capi "github.com/hashicorp/consul/api"
)

const DiscoveryAddress = "localhost:8500"

type requiredFeatureKey struct{}

func NewDiscovery() (registry.Discovery, error) {
	config := capi.DefaultConfig()
	config.Address = DiscoveryAddress
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return consul.New(client), nil
}

func NewRegistrar() (registry.Registrar, error) {
	config := capi.DefaultConfig()
	config.Address = DiscoveryAddress
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
