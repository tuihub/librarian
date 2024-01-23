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

func NewMapperClient(c *conf.Consul) (mapper.LibrarianMapperServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mapper"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := mapper.NewLibrarianMapperServiceClient(conn)
	return cli, err
}

func NewSearcherClient(c *conf.Consul) (searcher.LibrarianSearcherServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///searcher"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := searcher.NewLibrarianSearcherServiceClient(conn)
	return cli, err
}

func NewMinerClient(c *conf.Consul) (miner.LibrarianMinerServiceClient, error) {
	r, err := libapp.NewDiscovery(c)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///porter"),
		grpc.WithDiscovery(r),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := miner.NewLibrarianMinerServiceClient(conn)
	return cli, err
}
