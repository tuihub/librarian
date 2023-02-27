package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/converter"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
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
	if req.GetUser() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	id, err := s.t.CreateUser(ctx, s.converter.ToBizUser(req.GetUser()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: &librarian.InternalID{Id: int64(*id)},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (
	*pb.UpdateUserResponse, error,
) {
	if req.GetUser() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	err := s.t.UpdateUser(ctx, s.converter.ToBizUser(req.GetUser()), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (
	*pb.ListUserResponse, error,
) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	u, total, err := s.t.ListUser(ctx,
		model.Paging{
			PageSize: int(req.GetPaging().GetPageSize()),
			PageNum:  int(req.GetPaging().GetPageNum()),
		},
		s.converter.ToLibAuthUserTypeList(req.GetTypeFilter()),
		s.converter.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{
		Paging:   &librarian.PagingResponse{Total: total},
		UserList: s.converter.ToPBUserList(u),
	}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error,
) {
	if req.GetAccountId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	a, err := s.t.LinkAccount(ctx, biztiphereth.Account{
		InternalID:        0,
		Platform:          converter.ToBizAccountPlatform(req.GetAccountId().GetPlatform()),
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
	if req.GetAccountId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UnLinkAccount(ctx, biztiphereth.Account{
		InternalID:        0,
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
func (s *LibrarianSephirahServiceService) ListLinkAccount(ctx context.Context, req *pb.ListLinkAccountRequest) (
	*pb.ListLinkAccountResponse, error,
) {
	res, total, err := s.t.ListLinkAccount(ctx, model.Paging{
		PageSize: int(req.GetPaging().GetPageSize()),
		PageNum:  int(req.GetPaging().GetPageNum()),
	}, model.InternalID(converter.ToBizInternalID(req.GetUserId())))
	if err != nil {
		return nil, err
	}
	return &pb.ListLinkAccountResponse{
		Paging:      &librarian.PagingResponse{Total: total},
		AccountList: s.converter.ToPBAccountList(res),
	}, nil
}
