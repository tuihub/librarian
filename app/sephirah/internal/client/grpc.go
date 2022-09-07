package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

func NewMapperClient() (*mapper.LibrarianMapperServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(""),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := mapper.NewLibrarianMapperServiceClient(conn)
	return &cli, err
}

func NewSearcherClient() (*searcher.LibrarianSearcherServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(""),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := searcher.NewLibrarianSearcherServiceClient(conn)
	return &cli, err
}

func NewPorterClient() (*porter.LibrarianPorterServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(""),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	cli := porter.NewLibrarianPorterServiceClient(conn)
	return &cli, err
}
