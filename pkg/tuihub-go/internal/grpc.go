package internal

import (
	"context"
	"time"

	porter "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	capi "github.com/hashicorp/consul/api"
)

func NewSephirahClient(
	ctx context.Context,
	config *capi.Config,
	serviceName string,
) (sephirah.LibrarianSephirahServiceClient, error) {
	r, err := NewDiscovery(config)
	if err != nil {
		return nil, err
	}
	if serviceName == "" {
		serviceName = "librarian"
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(time.Minute),
	)
	cli := sephirah.NewLibrarianSephirahServiceClient(conn)
	return cli, err
}

func NewPorterClient(
	ctx context.Context,
	config *capi.Config,
	serviceName string,
) (porter.LibrarianSephirahPorterServiceClient, error) {
	r, err := NewDiscovery(config)
	if err != nil {
		return nil, err
	}
	if serviceName == "" {
		serviceName = "librarian"
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+serviceName),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		return nil, err
	}
	return porter.NewLibrarianSephirahPorterServiceClient(conn), nil
}

func NewDiscovery(c *capi.Config) (registry.Discovery, error) {
	config := c
	if config == nil {
		config = capi.DefaultConfig()
	}
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	_, err = client.Status().Leader()
	if err != nil {
		if c == nil {
			return emptyDiscovery{}, nil
		}
		return nil, err
	}
	return consul.New(client), nil
}

func NewRegistrar(c *capi.Config) (registry.Registrar, error) {
	config := c
	if config == nil {
		config = capi.DefaultConfig()
	}
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	_, err = client.Status().Leader()
	if err != nil {
		if c == nil {
			return emptyRegistrar{}, nil
		}
		return nil, err
	}
	return consul.New(client), nil
}

type emptyDiscovery struct{}

func (e emptyDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return []*registry.ServiceInstance{}, nil
}

func (e emptyDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return emptyWatcher{}, nil
}

type emptyRegistrar struct{}

func (e emptyRegistrar) Register(ctx context.Context, service *registry.ServiceInstance) error {
	return nil
}

func (e emptyRegistrar) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	return nil
}

type emptyWatcher struct{}

func (e emptyWatcher) Next() ([]*registry.ServiceInstance, error) {
	return []*registry.ServiceInstance{}, nil
}

func (e emptyWatcher) Stop() error {
	return nil
}
