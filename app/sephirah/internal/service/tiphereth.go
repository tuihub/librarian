package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/logger"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahServiceService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (
	*pb.GetTokenResponse, error,
) {
	accessToken, refreshToken, err := s.t.GetToken(ctx, &biztiphereth.User{
		InternalID: 0,
		UserName:   req.GetUsername(),
		PassWord:   req.GetPassword(),
		Type:       0,
		Status:     0,
	})
	if err != nil {
		logger.Infof("GetToken failed: %s", err.Error())
		return nil, err
	}
	return &pb.GetTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
func (s *LibrarianSephirahServiceService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (
	*pb.RefreshTokenResponse, error,
) {
	accessToken, refreshToken, err := s.t.RefreshToken(ctx)
	if err != nil {
		logger.Infof("GetToken failed: %s", err.Error())
		return nil, err
	}
	return &pb.RefreshTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
func (s *LibrarianSephirahServiceService) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (
	*pb.GenerateTokenResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (
	*pb.CreateUserResponse, error,
) {
	u, err := s.t.AddUser(ctx, &biztiphereth.User{
		InternalID: 0,
		UserName:   req.GetUser().GetUsername(),
		PassWord:   req.GetUser().GetPassword(),
		Type:       biztiphereth.ToLibAuthUserType(req.GetUser().GetType()),
		Status:     0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: &librarian.InternalID{Id: u.InternalID},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (
	*pb.UpdateUserResponse, error,
) {
	if req.GetUser().GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("id required")
	}
	err := s.t.UpdateUser(ctx, &biztiphereth.User{
		InternalID: req.GetUser().GetId().GetId(),
		UserName:   req.GetUser().GetUsername(),
		PassWord:   req.GetUser().GetPassword(),
		Type:       0,
		Status:     biztiphereth.ToBizUserStatus(req.GetUser().GetStatus()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (
	*pb.ListUserResponse, error,
) {
	u, err := s.t.ListUser(ctx,
		biztiphereth.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		biztiphereth.ToLibAuthUserTypeList(req.GetTypeFilter()),
		biztiphereth.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{ // TODO
		Paging:   nil,
		UserList: biztiphereth.ToPBUserList(u),
	}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error,
) {
	a, err := s.t.LinkAccount(ctx, biztiphereth.Account{
		InternalID:        0,
		Platform:          biztiphereth.ToBizAccountPlatform(req.GetAccountId().GetPlatform()),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
	})
	if err != nil {
		return nil, err
	}
	return &pb.LinkAccountResponse{AccountId: &librarian.InternalID{Id: a.InternalID}}, nil
}
func (s *LibrarianSephirahServiceService) UnLinkAccount(ctx context.Context, req *pb.UnLinkAccountRequest) (
	*pb.UnLinkAccountResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) ListLinkAccount(ctx context.Context, req *pb.ListLinkAccountRequest) (
	*pb.ListLinkAccountResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
