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
	username = "testUser"
	password = "testPass"
	userType = pb.UserType_USER_TYPE_ADMIN
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
	ctx = c.CreateDefaultUserAndLogin(ctx)
	c.TestUser(ctx)
	c.TestApp(ctx)
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

func (c *Client) CreateDefaultUserAndLogin(ctx context.Context) context.Context {
	var accessToken, refreshToken string
	if resp, err := c.cli.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Id:       nil,
			Username: username,
			Password: password,
			Type:     userType,
			Status:   0,
		},
	}); err != nil {
		panic(err)
	} else {
		c.userID = resp.Id
	}
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
	if resp, err := c.cli.ListUser(ctx, &pb.ListUserRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		TypeFilter:   nil,
		StatusFilter: nil,
	}); err != nil {
		panic(err)
	} else if len(resp.GetUserList()) != 1 || resp.GetUserList()[0].Id.GetId() != c.userID.GetId() {
		panic("inconsistent user id")
	}
}

func (c *Client) TestApp(ctx context.Context) {
	var appID *librarian.InternalID
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	} else {
		appID = resp.Id
	}
	if resp, err := c.cli.ListApp(ctx, &pb.ListAppRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		SourceFilter:   nil,
		TypeFilter:     nil,
		IdFilter:       nil,
		ContainDetails: false,
	}); err != nil {
		panic(err)
	} else if len(resp.GetAppList()) != 1 ||
		resp.GetAppList()[0].GetId().GetId() != appID.GetId() {
		panic("inconsistent app id")
	}
	if _, err := c.cli.UpdateApp(ctx, &pb.UpdateAppRequest{App: &librarian.App{
		Id:               appID,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description update",
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	}
}
