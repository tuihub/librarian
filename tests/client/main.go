package main

import (
	"context"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"golang.org/x/oauth2"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

func main() {
	cli := NewSephirahClient()
	ctx := context.Background()
	const (
		username = "testUser"
		password = "testPass"
		userType = pb.UserType_USER_TYPE_ADMIN
	)
	var userID *pb.InternalID
	var accessToken string
	if resp, err := cli.CreateUser(ctx, &pb.CreateUserRequest{
		Username: username,
		Password: password,
		Type:     userType,
	}); err != nil {
		panic(err)
	} else {
		userID = resp.Id
	}
	if resp, err := cli.GetToken(ctx, &pb.GetTokenRequest{
		Username: username,
		Password: password,
	}); err != nil {
		panic(err)
	} else {
		accessToken = resp.AccessToken
	}
	cred := grpc2.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{AccessToken: accessToken}))
	if resp, err := cli.ListUser(ctx, &pb.ListUserRequest{
		PageNum:      1,
		PageSize:     1,
		TypeFilter:   nil,
		StatusFilter: nil,
	}, cred); err != nil {
		panic(err)
	} else if len(resp.GetUserList()) != 1 || resp.GetUserList()[0].Id != userID {
		panic("inconsistent user id")
	}
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
