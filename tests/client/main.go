package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	username = "admin"
	password = "admin"
)

type Client struct {
	cli    pb.LibrarianSephirahServiceClient
	userID *librarian.InternalID
}

func main() {
	ctx := context.Background()

	c := Client{
		cli:    NewSephirahClient(),
		userID: nil,
	}
	c.WaitServerOnline(ctx)
	ctx = c.LoginViaDefaultAdmin(ctx)
	c.TestUser(ctx)
	c.TestGebura(ctx)
	c.TestYesod(ctx)
}

func NewSephirahClient() pb.LibrarianSephirahServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:10000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	cli := pb.NewLibrarianSephirahServiceClient(conn)
	return cli
}

func (c *Client) WaitServerOnline(ctx context.Context) {
	_, err := c.cli.GetToken(ctx, new(pb.GetTokenRequest))
	i := 1
	for errors.IsServiceUnavailable(err) && i < 30 {
		time.Sleep(time.Second)
		i += 1
		log.Infof("Waiting server online %s", err.Error())
		_, err = c.cli.GetToken(ctx, new(pb.GetTokenRequest))
	}
}

func (c *Client) LoginViaDefaultAdmin(ctx context.Context) context.Context {
	var accessToken, refreshToken string
	if resp, err := c.cli.GetToken(ctx, &pb.GetTokenRequest{
		Username: username,
		Password: password,
	}); err != nil {
		panic(err)
	} else {
		refreshToken = resp.RefreshToken
	}
	ctxForRefresh := metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("authorization", fmt.Sprintf("bearer %s", refreshToken)),
	)
	if resp, err := c.cli.RefreshToken(ctxForRefresh, &pb.RefreshTokenRequest{}); err != nil {
		panic(err)
	} else {
		accessToken = resp.AccessToken
	}
	return metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("authorization", fmt.Sprintf("bearer %s", accessToken)),
	)
}

func (c *Client) TestUser(ctx context.Context) {
	if _, err := c.cli.ListUser(ctx, &pb.ListUserRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		TypeFilter:   nil,
		StatusFilter: nil,
	}); err != nil {
		panic(err)
	}
}
