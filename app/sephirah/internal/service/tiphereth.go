package service

import (
	"context"
	"time"

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
	accessToken, refreshToken, err := s.t.GetToken(ctx,
		req.GetUsername(),
		req.GetPassword(),
		converter.ToBizInternalIDPtr(req.GetDeviceId()),
	)
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
	var deviceID *model.InternalID
	if req.GetDeviceId() != nil {
		id := converter.ToBizInternalID(req.GetDeviceId())
		deviceID = &id
	}
	accessToken, refreshToken, err := s.t.RefreshToken(ctx, deviceID)
	if err != nil {
		logger.Infof("GetToken failed: %s", err.Error())
		return nil, err
	}
	return &pb.RefreshTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
func (s *LibrarianSephirahServiceService) GainUserPrivilege(ctx context.Context, req *pb.GainUserPrivilegeRequest) (
	*pb.GainUserPrivilegeResponse, error,
) {
	if req.GetUserId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	token, err := s.t.GainUserPrivilege(ctx, converter.ToBizInternalID(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.GainUserPrivilegeResponse{
		AccessToken: string(token),
	}, nil
}
func (s *LibrarianSephirahServiceService) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (
	*pb.RegisterDeviceResponse, error,
) {
	if req.GetDeviceInfo() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	id, err := s.t.RegisterDevice(ctx, converter.ToBizDeviceInfo(req.GetDeviceInfo()))
	if err != nil {
		return nil, err
	}
	return &pb.RegisterDeviceResponse{
		DeviceId: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) ListUserSessions(ctx context.Context, req *pb.ListUserSessionsRequest) (
	*pb.ListUserSessionsResponse, error,
) {
	sessions, err := s.t.ListUserSessions(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserSessionsResponse{
		Sessions: converter.ToPBUserSessionList(sessions),
	}, nil
}
func (s *LibrarianSephirahServiceService) DeleteUserSession(ctx context.Context, req *pb.DeleteUserSessionRequest) (
	*pb.DeleteUserSessionResponse, error,
) {
	if req.GetSessionId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	err := s.t.DeleteUserSession(ctx, converter.ToBizInternalID(req.GetSessionId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserSessionResponse{}, nil
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
	u, total, err := s.t.ListUsers(ctx,
		model.ToBizPaging(req.GetPaging()),
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
		Platform:          req.GetAccountId().GetPlatform(),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
		LatestUpdateTime:  time.Time{},
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
		Platform:          req.GetAccountId().GetPlatform(),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
		LatestUpdateTime:  time.Time{},
	}); err != nil {
		return nil, err
	}
	return &pb.UnLinkAccountResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListLinkAccounts(ctx context.Context, req *pb.ListLinkAccountsRequest) (
	*pb.ListLinkAccountsResponse, error,
) {
	res, err := s.t.ListLinkAccounts(ctx,
		converter.ToBizInternalID(req.GetUserId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListLinkAccountsResponse{
		Accounts: converter.ToPBAccountList(res),
	}, nil
}

func (s *LibrarianSephirahServiceService) ListPorters(ctx context.Context, req *pb.ListPortersRequest) (
	*pb.ListPortersResponse, error,
) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	res, total, err := s.t.ListPorters(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	for i := range res {
		res[i].ConnectionStatus = s.s.GetInstanceConnectionStatus(ctx, res[i].Address)
	}
	return &pb.ListPortersResponse{
		Paging:  &librarian.PagingResponse{TotalSize: total},
		Porters: converter.ToPBPorterList(res),
	}, nil
}

func (s *LibrarianSephirahServiceService) UpdatePorterStatus(ctx context.Context, req *pb.UpdatePorterStatusRequest) (
	*pb.UpdatePorterStatusResponse, error,
) {
	if req.GetPorterId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UpdatePorterStatus(ctx,
		converter.ToBizInternalID(req.GetPorterId()),
		converter.ToBizPorterStatus(req.GetStatus()),
	); err != nil {
		return nil, err
	}
	return &pb.UpdatePorterStatusResponse{}, nil
}

func (s *LibrarianSephirahServiceService) UpdatePorterPrivilege(
	ctx context.Context,
	req *pb.UpdatePorterPrivilegeRequest,
) (*pb.UpdatePorterPrivilegeResponse, error) {
	if req.GetPorterId() == nil || req.GetPrivilege() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UpdatePorterPrivilege(ctx,
		converter.ToBizInternalID(req.GetPorterId()),
		converter.ToBizPorterPrivilege(req.GetPrivilege()),
	); err != nil {
		return nil, err
	}
	return &pb.UpdatePorterPrivilegeResponse{}, nil
}
