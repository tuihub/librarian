package client

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libdiscovery"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewMinerClient(c *conf.Consul, app *libapp.Settings) (miner.LibrarianMinerServiceClient, error) {
	r, err := libdiscovery.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///miner"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libdiscovery.NewNodeFilter()),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, grpc.WithMiddleware(recovery.Recovery()))
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		middlewares...,
	)
	cli := miner.NewLibrarianMinerServiceClient(conn)
	return cli, err
}
