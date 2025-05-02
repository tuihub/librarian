package internal

import (
	"context"
	"time"

	porter "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	capi "github.com/hashicorp/consul/api"
)

func NewSephirahClient(
	ctx context.Context,
	config *capi.Config,
	serviceName string,
) (sephirah.LibrarianSephirahServiceClient, error) {
	r, err := NewRegistry(config)
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
	r, err := NewRegistry(config)
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

func NewRegistry(config *capi.Config) (*consul.Registry, error) {
	if config == nil {
		config = capi.DefaultConfig()
	}
	client, err := capi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return consul.New(client), nil
}
