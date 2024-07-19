package client

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewMapperClient(c *conf.Consul, app *libapp.Settings) (mapper.LibrarianMapperServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///mapper"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, grpc.WithMiddleware(recovery.Recovery()))
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		middlewares...,
	)
	cli := mapper.NewLibrarianMapperServiceClient(conn)
	return cli, err
}

func NewSearcherClient(c *conf.Consul, app *libapp.Settings) (searcher.LibrarianSearcherServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///searcher"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, grpc.WithMiddleware(recovery.Recovery()))
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		middlewares...,
	)
	cli := searcher.NewLibrarianSearcherServiceClient(conn)
	return cli, err
}

func NewMinerClient(c *conf.Consul, app *libapp.Settings) (miner.LibrarianMinerServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	middlewares := []grpc.ClientOption{
		grpc.WithEndpoint("discovery:///miner"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
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
