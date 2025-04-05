package sephirah

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahService) GetToken(ctx context.Context, req *sephirah.GetTokenRequest) (
	*sephirah.GetTokenResponse, error,
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
	return &sephirah.GetTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
func (s *LibrarianSephirahService) RefreshToken(ctx context.Context, req *sephirah.RefreshTokenRequest) (
	*sephirah.RefreshTokenResponse, error,
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
	return &sephirah.RefreshTokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	}, nil
}
func (s *LibrarianSephirahService) RegisterUser(ctx context.Context, req *sephirah.RegisterUserRequest) (
	*sephirah.RegisterUserResponse, error,
) {
	var captchaAns *model.CaptchaAns
	if req.GetCaptcha() != nil {
		captchaAns = &model.CaptchaAns{
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
		return &sephirah.RegisterUserResponse{
			Stage: &sephirah.RegisterUserResponse_RefreshToken{RefreshToken: refreshToken},
		}, nil
	}
	return &sephirah.RegisterUserResponse{
		Stage: &sephirah.RegisterUserResponse_Captcha{Captcha: &sephirah.RegisterUserResponse_ImageCaptcha{
			Id:    captchaQue.ID,
			Image: captchaQue.Image,
		}},
	}, nil
}
func (s *LibrarianSephirahService) RegisterDevice(ctx context.Context, req *sephirah.RegisterDeviceRequest) (
	*sephirah.RegisterDeviceResponse, error,
) {
	if req.GetDeviceInfo() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	localID := req.GetClientLocalId()
	id, err := s.t.RegisterDevice(ctx, converter.ToBizDeviceInfo(req.GetDeviceInfo()), &localID)
	if err != nil {
		return nil, err
	}
	return &sephirah.RegisterDeviceResponse{
		DeviceId: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahService) ListUserSessions(ctx context.Context, req *sephirah.ListUserSessionsRequest) (
	*sephirah.ListUserSessionsResponse, error,
) {
	sessions, err := s.t.ListUserSessions(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListUserSessionsResponse{
		Sessions: converter.ToPBUserSessionList(sessions),
	}, nil
}
func (s *LibrarianSephirahService) DeleteUserSession(ctx context.Context, req *sephirah.DeleteUserSessionRequest) (
	*sephirah.DeleteUserSessionResponse, error,
) {
	if req.GetSessionId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	err := s.t.DeleteUserSession(ctx, converter.ToBizInternalID(req.GetSessionId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.DeleteUserSessionResponse{}, nil
}
func (s *LibrarianSephirahService) UpdateUser(ctx context.Context, req *sephirah.UpdateUserRequest) (
	*sephirah.UpdateUserResponse, error,
) {
	if req.GetUser() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	err := s.t.UpdateUser(ctx, converter.ToBizUser(req.GetUser()), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahService) GetUser(ctx context.Context, req *sephirah.GetUserRequest) (
	*sephirah.GetUserResponse, error,
) {
	u, err := s.t.GetUser(ctx, converter.ToBizInternalIDPtr(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.GetUserResponse{User: converter.ToPBUser(u)}, nil
}
func (s *LibrarianSephirahService) LinkAccount(ctx context.Context, req *sephirah.LinkAccountRequest) (
	*sephirah.LinkAccountResponse, error,
) {
	a, err := s.t.LinkAccount(ctx, model.Account{
		ID:                0,
		Platform:          req.GetPlatform(),
		PlatformAccountID: req.GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
		LatestUpdateTime:  time.Time{},
	})
	if err != nil {
		return nil, err
	}
	return &sephirah.LinkAccountResponse{AccountId: converter.ToPBInternalID(a.ID)}, nil
}
func (s *LibrarianSephirahService) UnLinkAccount(ctx context.Context, req *sephirah.UnLinkAccountRequest) (
	*sephirah.UnLinkAccountResponse, error,
) {
	if err := s.t.UnLinkAccount(ctx, model.Account{
		ID:                0,
		Platform:          req.GetPlatform(),
		PlatformAccountID: req.GetPlatformAccountId(),
		Name:              "",
		ProfileURL:        "",
		AvatarURL:         "",
		LatestUpdateTime:  time.Time{},
	}); err != nil {
		return nil, err
	}
	return &sephirah.UnLinkAccountResponse{}, nil
}
func (s *LibrarianSephirahService) ListLinkAccounts(ctx context.Context, req *sephirah.ListLinkAccountsRequest) (
	*sephirah.ListLinkAccountsResponse, error,
) {
	res, err := s.t.ListLinkAccounts(ctx,
		converter.ToBizInternalID(req.GetUserId()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListLinkAccountsResponse{
		Accounts: converter.ToPBAccountList(res),
	}, nil
}

// func (s *LibrarianSephirahService) ListPorters(ctx context.Context, req *sephirah.ListPortersRequest) ( //nolint:dupl //no need
//	*sephirah.ListPortersResponse, error,
// ) {
//	if req.GetPaging() == nil {
//		return nil, pb.ErrorErrorReasonBadRequest("")
//	}
//	porters, total, err := s.t.ListPorters(ctx,
//		model.ToBizPaging(req.GetPaging()),
//	)
//	if err != nil {
//		return nil, err
//	}
//	res := make([]*modelsupervisor.PorterInstanceController, len(porters))
//	for i := range res {
//		res[i] = s.s.GetInstanceController(ctx, porters[i].Address)
//		if res[i] == nil {
//			res[i] = new(modelsupervisor.PorterInstanceController)
//			res[i].ConnectionStatus = modelsupervisor.PorterConnectionStatusDisconnected
//		}
//		res[i].PorterInstance = *porters[i]
//	}
//	return &sephirah.ListPortersResponse{
//		Paging:  &librarian.PagingResponse{TotalSize: total},
//		Porters: converter.ToPBPorterList(res),
//	}, nil
//}
//
// func (s *LibrarianSephirahService) UpdatePorterStatus(ctx context.Context, req *sephirah.UpdatePorterStatusRequest) (
//	*sephirah.UpdatePorterStatusResponse, error,
// ) {
//	if req.GetPorterId() == nil {
//		return nil, pb.ErrorErrorReasonBadRequest("")
//	}
//	if err := s.t.UpdatePorterStatus(ctx,
//		converter.ToBizInternalID(req.GetPorterId()),
//		converter.ToBizUserStatus(req.GetStatus()),
//	); err != nil {
//		return nil, err
//	}
//	return &sephirah.UpdatePorterStatusResponse{}, nil
//}

func (s *LibrarianSephirahService) CreatePorterContext(
	ctx context.Context,
	req *sephirah.CreatePorterContextRequest,
) (*sephirah.CreatePorterContextResponse, error) {
	if req.GetContext() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	id, err := s.t.CreatePorterContext(ctx,
		converter.ToBizPorterContext(req.GetContext()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.CreatePorterContextResponse{
		ContextId: converter.ToPBInternalID(id),
	}, nil
}

func (s *LibrarianSephirahService) ListPorterContexts(
	ctx context.Context,
	req *sephirah.ListPorterContextsRequest,
) (*sephirah.ListPorterContextsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	contexts, total, err := s.t.ListPorterContexts(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	res := make([]*modelsupervisor.PorterContextController, len(contexts))
	for i := range res {
		res[i] = s.s.GetContextController(ctx, contexts[i].ID)
		if res[i] == nil {
			res[i] = new(modelsupervisor.PorterContextController)
			res[i].HandleStatus = modelsupervisor.PorterContextHandleStatusBlocked
		}
		res[i].PorterContext = *contexts[i]
	}
	return &sephirah.ListPorterContextsResponse{
		Paging:   &librarian.PagingResponse{TotalSize: total},
		Contexts: converter.ToPBPorterContextList(res),
	}, nil
}

func (s *LibrarianSephirahService) UpdatePorterContext(
	ctx context.Context,
	req *sephirah.UpdatePorterContextRequest,
) (*sephirah.UpdatePorterContextResponse, error) {
	if req.GetContext() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	if err := s.t.UpdatePorterContext(ctx,
		converter.ToBizPorterContext(req.GetContext()),
	); err != nil {
		return nil, err
	}
	return &sephirah.UpdatePorterContextResponse{}, nil
}

func (s *LibrarianSephirahService) ListPorterDigests(
	ctx context.Context,
	req *sephirah.ListPorterDigestsRequest,
) (*sephirah.ListPorterDigestsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	groups, total, err := s.t.ListPorterDigests(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListPorterDigestsResponse{
		Paging:        &librarian.PagingResponse{TotalSize: total},
		PorterDigests: converter.ToPBPorterDigestList(groups),
	}, nil
}
