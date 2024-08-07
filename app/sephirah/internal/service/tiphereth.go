package service

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
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
func (s *LibrarianSephirahServiceService) AcquireUserToken(ctx context.Context, req *pb.AcquireUserTokenRequest) (
	*pb.AcquireUserTokenResponse, error,
) {
	if req.GetUserId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	token, err := s.t.AcquireUserToken(ctx, converter.ToBizInternalID(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.AcquireUserTokenResponse{
		AccessToken: string(token),
	}, nil
}
func (s *LibrarianSephirahServiceService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (
	*pb.RegisterUserResponse, error,
) {
	var captchaAns *modeltiphereth.CaptchaAns
	if req.GetCaptcha() != nil {
		captchaAns = &modeltiphereth.CaptchaAns{
			ID:    req.GetCaptcha().GetId(),
			Value: req.GetCaptcha().GetValue(),
		}
	}
	captchaQue, refreshToken, err := s.t.RegisterUser(
		ctx,
		req.GetUsername(),
		req.GetPassword(),
		captchaAns,
	)
	if err != nil {
		return nil, err
	}
	if len(refreshToken) > 0 {
		return &pb.RegisterUserResponse{
			Stage: &pb.RegisterUserResponse_RefreshToken{RefreshToken: refreshToken},
		}, nil
	}
	return &pb.RegisterUserResponse{
		Stage: &pb.RegisterUserResponse_Captcha{Captcha: &pb.RegisterUserResponse_ImageCaptcha{
			Id:    captchaQue.ID,
			Image: captchaQue.Image,
		}},
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
func (s *LibrarianSephirahServiceService) ListRegisteredDevices(
	ctx context.Context,
	req *pb.ListRegisteredDevicesRequest,
) (
	*pb.ListRegisteredDevicesResponse, error,
) {
	devices, err := s.t.ListRegisteredDevices(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListRegisteredDevicesResponse{
		Devices: converter.ToPBDeviceInfoList(devices),
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
	porters, total, err := s.t.ListPorters(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	res := make([]*modelsupervisor.PorterInstanceController, len(porters))
	for i := range res {
		res[i] = s.s.GetInstanceController(ctx, porters[i].Address)
		if res[i] == nil {
			res[i] = new(modelsupervisor.PorterInstanceController)
			res[i].PorterInstance = *porters[i]
			res[i].ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
		}
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
		converter.ToBizUserStatus(req.GetStatus()),
	); err != nil {
		return nil, err
	}
	return &pb.UpdatePorterStatusResponse{}, nil
}

func (s *LibrarianSephirahServiceService) CreatePorterContext(
	ctx context.Context,
	req *pb.CreatePorterContextRequest,
) (*pb.CreatePorterContextResponse, error) {
	if req.GetContext() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	id, err := s.t.CreatePorterContext(ctx,
		converter.ToBizPorterContext(req.GetContext()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePorterContextResponse{
		ContextId: converter.ToPBInternalID(id),
	}, nil
}

func (s *LibrarianSephirahServiceService) ListPorterContexts(
	ctx context.Context,
	req *pb.ListPorterContextsRequest,
) (*pb.ListPorterContextsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	contexts, total, err := s.t.ListPorterContexts(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListPorterContextsResponse{
		Paging:   &librarian.PagingResponse{TotalSize: total},
		Contexts: converter.ToPBPorterContextList(contexts),
	}, nil
}

func (s *LibrarianSephirahServiceService) UpdatePorterContext(
	ctx context.Context,
	req *pb.UpdatePorterContextRequest,
) (*pb.UpdatePorterContextResponse, error) {
	if req.GetContext() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UpdatePorterContext(ctx,
		converter.ToBizPorterContext(req.GetContext()),
	); err != nil {
		return nil, err
	}
	return &pb.UpdatePorterContextResponse{}, nil
}

func (s *LibrarianSephirahServiceService) ListPorterGroups(
	ctx context.Context,
	req *pb.ListPorterGroupsRequest,
) (*pb.ListPorterGroupsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	groups, total, err := s.t.ListPorterGroups(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListPorterGroupsResponse{
		Paging:       &librarian.PagingResponse{TotalSize: total},
		PorterGroups: converter.ToPBPorterGroupList(groups),
	}, nil
}
