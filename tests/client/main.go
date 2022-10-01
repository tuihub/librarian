package main

import (
	"context"
	"fmt"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	pb2 "github.com/tuihub/protos/pkg/librarian/v1"

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
	userID *pb2.InternalID
}

func main() {
	ctx := context.Background()
	c := Client{
		cli:    NewSephirahClient(),
		userID: nil,
	}
	ctx = c.CreateDefaultUserAndLogin(ctx)
	c.TestUser(ctx)
	c.TestApp(ctx)
}

func NewSephirahClient() pb.LibrarianSephirahServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
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

func (c *Client) CreateDefaultUserAndLogin(ctx context.Context) context.Context {
	var accessToken, refreshToken string
	if resp, err := c.cli.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Username: username,
			Password: password,
			Type:     userType,
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
		PageNum:      1,
		PageSize:     1,
		TypeFilter:   nil,
		StatusFilter: nil,
	}); err != nil {
		panic(err)
	} else if len(resp.GetUserList()) != 1 || resp.GetUserList()[0].Id.GetId() != c.userID.GetId() {
		panic("inconsistent user id")
	}
}

func (c *Client) TestApp(ctx context.Context) {
	var appID *pb2.InternalID
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &pb2.App{
		Id:               nil,
		Source:           pb2.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             pb2.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	} else {
		appID = resp.Id
	}
	if resp, err := c.cli.ListApp(ctx, &pb.ListAppRequest{
		PageNum:        1,
		PageSize:       1,
		SourceFilter:   nil,
		TypeFilter:     nil,
		IdFilter:       nil,
		ContainDetails: false,
		WithBind:       false,
	}); err != nil {
		panic(err)
	} else if len(resp.GetWithoutBind().GetAppList()) != 1 ||
		resp.GetWithoutBind().GetAppList()[0].GetId().GetId() != appID.GetId() {
		panic("inconsistent app id")
	}
	if _, err := c.cli.UpdateApp(ctx, &pb.UpdateAppRequest{App: &pb2.App{
		Id:               appID,
		Source:           pb2.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             pb2.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description update",
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	}
}
