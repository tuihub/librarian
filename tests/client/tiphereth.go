package main

import (
	"context"
	"fmt"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/metadata"
)

func (c *Client) LoginViaDefaultAdmin(ctx context.Context) context.Context {
	const (
		adminUsername = "admin"
		adminPassword = "admin"
	)

	var accessToken, refreshToken string
	if resp, err := c.cli.GetToken(ctx, &pb.GetTokenRequest{
		Username: adminUsername,
		Password: adminPassword,
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

func (c *Client) TestTiphereth(ctx context.Context) {
	c.testUser(ctx)
	c.testAccount(ctx)
}

func (c *Client) testUser(ctx context.Context) {
	user1 := &pb.User{
		Id:       nil,
		Username: "user1",
		Password: "user1",
		Type:     pb.UserType_USER_TYPE_ADMIN,
		Status:   pb.UserStatus_USER_STATUS_BLOCKED,
	}
	user1Password := "user1"
	user2 := &pb.User{
		Id:       nil,
		Username: "user2",
		Password: "user2",
		Type:     pb.UserType_USER_TYPE_NORMAL,
		Status:   pb.UserStatus_USER_STATUS_BLOCKED,
	}
	user2Password := "user2"
	if user1ID, err := c.cli.CreateUser(ctx, &pb.CreateUserRequest{User: user1}); err != nil {
		panic(err)
	} else {
		user1.Id = user1ID.GetId()
	}
	if user2ID, err := c.cli.CreateUser(ctx, &pb.CreateUserRequest{User: user2}); err != nil {
		panic(err)
	} else {
		user2.Id = user2ID.GetId()
	}
	user1.Password = ""
	user2.Password = ""

	c.assertListUser(
		ctx, []pb.UserType{pb.UserType_USER_TYPE_NORMAL}, nil,
		func(resp *pb.ListUsersResponse) bool {
			return !cmp.Equal(resp.GetPaging().GetTotalSize(), 1) ||
				!cmp.Equal(resp.GetUsers()[0], user2)
		})
	c.assertListUser(
		ctx, []pb.UserType{pb.UserType_USER_TYPE_NORMAL}, []pb.UserStatus{pb.UserStatus_USER_STATUS_BLOCKED},
		func(resp *pb.ListUsersResponse) bool {
			return !cmp.Equal(resp.GetPaging().GetTotalSize(), 1) ||
				!cmp.Equal(resp.GetUsers()[0], user2)
		})
	c.assertListUser(
		ctx, nil, []pb.UserStatus{pb.UserStatus_USER_STATUS_BLOCKED},
		func(resp *pb.ListUsersResponse) bool {
			return !cmp.Equal(resp.GetPaging().GetTotalSize(), 2) //nolint:gomnd // definite
		})

	user1.Password = "user1newPass"
	user1.Type = pb.UserType_USER_TYPE_NORMAL
	user1.Status = pb.UserStatus_USER_STATUS_ACTIVE
	if _, err := c.cli.UpdateUser(ctx, &pb.UpdateUserRequest{User: user1, Password: nil}); err == nil {
		panic("err expected")
	}
	if _, err := c.cli.UpdateUser(ctx, &pb.UpdateUserRequest{User: user1, Password: &user1Password}); err != nil {
		panic(err)
	}
	if _, err := c.cli.UpdateUser(ctx, &pb.UpdateUserRequest{User: user2, Password: nil}); err != nil {
		panic(err)
	}

	c.assertListUser(
		ctx, []pb.UserType{pb.UserType_USER_TYPE_NORMAL}, nil,
		func(resp *pb.ListUsersResponse) bool {
			return !cmp.Equal(resp.GetPaging().GetTotalSize(), 2) //nolint:gomnd // definite
		})
	c.assertListUser(
		ctx, nil, []pb.UserStatus{pb.UserStatus_USER_STATUS_BLOCKED},
		func(resp *pb.ListUsersResponse) bool {
			return !cmp.Equal(resp.GetPaging().GetTotalSize(), 1) ||
				!cmp.Equal(resp.GetUsers()[0], user2)
		})

	user2.Password = user2Password
	if _, err := NewSephirahClient().GetToken(ctx, &pb.GetTokenRequest{
		Username: user1.Username,
		Password: user1.Password,
	}); err != nil {
		panic(err)
	}
	if _, err := NewSephirahClient().GetToken(ctx, &pb.GetTokenRequest{
		Username: user2.Username,
		Password: user2.Password,
	}); err == nil {
		panic("err expected")
	}
	user2.Password = user2Password
	if _, err := NewSephirahClient().GetToken(ctx, &pb.GetTokenRequest{
		Username: user2.Username,
		Password: user2.Password,
	}); err == nil {
		panic("err expected")
	}
}

func (c *Client) testAccount(ctx context.Context) {
	if _, err := c.cli.LinkAccount(ctx, &pb.LinkAccountRequest{
		AccountId: &librarian.AccountID{
			Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
			PlatformAccountId: "0",
		},
	}); err != nil {
		panic(err)
	}
	if _, err := c.cli.LinkAccount(ctx, &pb.LinkAccountRequest{
		AccountId: &librarian.AccountID{
			Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
			PlatformAccountId: "1",
		},
	}); err == nil {
		panic("err expected")
	}
	if resp, err := c.cli.ListLinkAccounts(ctx, &pb.ListLinkAccountsRequest{
		UserId: nil,
	}); err != nil {
		panic(err)
	} else if len(resp.GetAccounts()) != 1 || resp.GetAccounts()[0].GetPlatformAccountId() != "0" {
		panic(fmt.Sprintf("unexpected ListLinkAccounts response, %+v", resp))
	}
	if _, err := c.cli.UnLinkAccount(ctx, &pb.UnLinkAccountRequest{
		AccountId: &librarian.AccountID{
			Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
			PlatformAccountId: "0",
		},
	}); err != nil {
		panic(err)
	}
	if resp, err := c.cli.ListLinkAccounts(ctx, &pb.ListLinkAccountsRequest{
		UserId: nil,
	}); err != nil {
		panic(err)
	} else if len(resp.GetAccounts()) != 0 {
		panic(fmt.Sprintf("unexpected ListLinkAccounts response, %+v", resp))
	}
	if _, err := c.cli.LinkAccount(ctx, &pb.LinkAccountRequest{
		AccountId: &librarian.AccountID{
			Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
			PlatformAccountId: "1",
		},
	}); err != nil {
		panic(err)
	}
	if _, err := c.cli.LinkAccount(ctx, &pb.LinkAccountRequest{
		AccountId: &librarian.AccountID{
			Platform:          librarian.AccountPlatform_ACCOUNT_PLATFORM_STEAM,
			PlatformAccountId: "0",
		},
	}); err == nil {
		panic("err expected")
	}
}

func (c *Client) assertListUser(
	ctx context.Context, types []pb.UserType, statuses []pb.UserStatus, assertFunc func(*pb.ListUsersResponse) bool,
) {
	resp, err := c.cli.ListUsers(ctx, &pb.ListUsersRequest{
		Paging:       defaultPaging,
		TypeFilter:   types,
		StatusFilter: statuses,
	})
	if err != nil {
		panic(err)
	}
	if !assertFunc(resp) {
		panic(fmt.Sprintf("unexpected ListUser response, %+v", resp))
	}
}
