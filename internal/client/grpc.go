package client

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libapp"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type DiscoverClients struct {
	Mapper   mapper.LibrarianMapperServiceClient
	Searcher searcher.LibrarianSearcherServiceClient
	Porter   porter.LibrarianPorterServiceClient
}

func NewDiscoverClients() (*DiscoverClients, error) {
	m, err := NewMapperClient()
	if err != nil {
		return nil, err
	}
	s, err := NewSearcherClient()
	if err != nil {
		return nil, err
	}
	p, err := NewPorterClient()
	if err != nil {
		return nil, err
	}
	return &DiscoverClients{
		Mapper:   m,
		Searcher: s,
		Porter:   p,
	}, nil
}

func NewMapperClient() (mapper.LibrarianMapperServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///mapper"),
		grpc.WithDiscovery(libapp.NewDiscovery()),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := mapper.NewLibrarianMapperServiceClient(conn)
	return cli, err
}

func NewSearcherClient() (searcher.LibrarianSearcherServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///searcher"),
		grpc.WithDiscovery(libapp.NewDiscovery()),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := searcher.NewLibrarianSearcherServiceClient(conn)
	return cli, err
}

func NewPorterClient() (porter.LibrarianPorterServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///porter"),
		grpc.WithDiscovery(libapp.NewDiscovery()),
		grpc.WithNodeFilter(libapp.NewNodeFilter()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := porter.NewLibrarianPorterServiceClient(conn)
	return cli, err
}
