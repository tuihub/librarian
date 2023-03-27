package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahServiceService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (
	*pb.GetTokenResponse, error,
) {
	accessToken, refreshToken, err := s.t.GetToken(ctx, &modeltiphereth.User{
		ID:       0,
		UserName: req.GetUsername(),
		PassWord: req.GetPassword(),
		Type:     0,
		Status:   0,
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
	if req.GetUser() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	id, err := s.t.CreateUser(ctx, converter.ToBizUser(req.GetUser()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: converter.ToPBInternalID(*id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (
	*pb.UpdateUserResponse, error,
) {
	if req.GetUser() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	err := s.t.UpdateUser(ctx, converter.ToBizUser(req.GetUser()), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (
	*pb.ListUsersResponse, error,
) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	u, total, err := s.t.ListUser(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		converter.ToLibAuthUserTypeList(req.GetTypeFilter()),
		converter.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListUsersResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Users:  converter.ToPBUserList(u),
	}, nil
}
func (s *LibrarianSephirahServiceService) GetUser(ctx context.Context, req *pb.GetUserRequest) (
	*pb.GetUserResponse, error,
) {
	u, err := s.t.GetUser(ctx, converter.ToBizInternalIDPtr(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: converter.ToPBUser(u)}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error,
) {
	if req.GetAccountId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	a, err := s.t.LinkAccount(ctx, modeltiphereth.Account{
		ID:                0,
		Platform:          converter.ToBizAccountPlatform(req.GetAccountId().GetPlatform()),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
	})
	if err != nil {
		return nil, err
	}
	return &pb.LinkAccountResponse{AccountId: converter.ToPBInternalID(a.ID)}, nil
}
func (s *LibrarianSephirahServiceService) UnLinkAccount(ctx context.Context, req *pb.UnLinkAccountRequest) (
	*pb.UnLinkAccountResponse, error,
) {
	if req.GetAccountId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UnLinkAccount(ctx, modeltiphereth.Account{
		ID:                0,
		Platform:          converter.ToBizAccountPlatform(req.GetAccountId().GetPlatform()),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
	}); err != nil {
		return nil, err
	}
	return &pb.UnLinkAccountResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListLinkAccount(ctx context.Context, req *pb.ListLinkAccountsRequest) (
	*pb.ListLinkAccountsResponse, error,
) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	res, total, err := s.t.ListLinkAccount(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		converter.ToBizInternalID(req.GetUserId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListLinkAccountsResponse{
		Paging:   &librarian.PagingResponse{TotalSize: total},
		Accounts: converter.ToPBAccountList(res),
	}, nil
}
