package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/logger"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t *biztiphereth.Tiphereth
	g *bizgebura.Gebura
	b *bizbinah.Binah
}

func NewLibrarianSephirahServiceService(
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
) pb.LibrarianSephirahServiceServer {
	return &LibrarianSephirahServiceService{
		t: t,
		g: g,
		b: b,
	}
}

func (s *LibrarianSephirahServiceService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (
	*pb.GetTokenResponse, error,
) {
	accessToken, refreshToken, err := s.t.GetToken(ctx, &biztiphereth.User{
		UserName: req.GetUsername(),
		PassWord: req.GetPassword(),
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
		UserName: req.GetUser().GetUsername(),
		PassWord: req.GetUser().GetPassword(),
		Type:     biztiphereth.ToLibAuthUserType(req.GetUser().GetType()),
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
			PageSize: int(req.GetPageSize()),
			PageNum:  int(req.GetPageNum()),
		},
		biztiphereth.ToLibAuthUserTypeList(req.GetTypeFilter()),
		biztiphereth.ToBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{
		UserList: biztiphereth.ToPBUserList(u),
	}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error,
) {
	a, err := s.t.LinkAccount(ctx, biztiphereth.Account{
		Platform:          biztiphereth.ToBizAccountPlatform(req.GetAccountId().GetPlatform()),
		PlatformAccountID: req.GetAccountId().GetPlatformAccountId(),
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
func (s *LibrarianSephirahServiceService) UploadFile(conn pb.LibrarianSephirahService_UploadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) DownloadFile(conn pb.LibrarianSephirahService_DownloadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) SimpleUploadFile(
	conn pb.LibrarianSephirahService_SimpleUploadFileServer,
) error {
	file, bizErr := s.b.NewUploadFile(conn.Context())
	if bizErr != nil {
		return bizErr
	}
	for {
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return file.Finish()
			}
			return err
		} else if _, err = file.Writer.Write(req.Data); err != nil {
			return err
		}
		if err := conn.Send(&pb.SimpleUploadFileResponse{}); err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) SimpleDownloadFile(
	conn pb.LibrarianSephirahService_SimpleDownloadFileServer,
) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (
	*pb.CreateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app required")
	}
	a, err := s.g.CreateApp(ctx, &bizgebura.App{
		Source:          bizgebura.AppSourceInternal,
		SourceAppID:     "",
		SourceURL:       "",
		Name:            app.GetName(),
		Type:            bizgebura.ToBizAppType(app.GetType()),
		ShorDescription: app.GetShortDescription(),
		ImageURL:        app.GetImageUrl(),
		Details:         bizgebura.ToBizAppDetail(app.GetDetails()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{
		Id: &librarian.InternalID{Id: a.InternalID},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (
	*pb.UpdateAppResponse, error,
) {
	app := req.GetApp()
	if app == nil || app.GetId() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app and internal_id required")
	}
	err := s.g.UpdateApp(ctx, &bizgebura.App{
		InternalID:      app.GetId().GetId(),
		Source:          bizgebura.AppSourceInternal,
		SourceAppID:     "",
		SourceURL:       "",
		Name:            app.GetName(),
		Type:            bizgebura.ToBizAppType(app.GetType()),
		ShorDescription: app.GetShortDescription(),
		ImageURL:        app.GetImageUrl(),
		Details:         bizgebura.ToBizAppDetail(app.GetDetails()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApp(ctx context.Context, req *pb.ListAppRequest) (
	*pb.ListAppResponse, error,
) {
	a, err := s.g.ListApp(ctx,
		bizgebura.Paging{
			PageSize: int(req.GetPageSize()),
			PageNum:  int(req.GetPageNum()),
		},
		bizgebura.ToBizAppSourceList(req.GetSourceFilter()),
		bizgebura.ToBizAppTypeList(req.GetTypeFilter()),
		req.GetIdFilter(),
		req.GetContainDetails(),
		req.GetWithBind())
	if err != nil {
		return nil, err
	}
	return &pb.ListAppResponse{
		Content: &pb.ListAppResponse_WithoutBind{
			WithoutBind: &pb.ListAppResponse_AppList{
				AppList: bizgebura.ToPBAppList(a, req.GetContainDetails()),
			},
		},
	}, nil
}
func (s *LibrarianSephirahServiceService) BindApp(ctx context.Context, req *pb.BindAppRequest) (
	*pb.BindAppResponse, error,
) {
	a, err := s.g.BindApp(ctx,
		bizgebura.App{
			InternalID: req.GetInternalAppId().GetId(),
		},
		bizgebura.App{
			Source:      bizgebura.ToBizAppSource(req.GetBindAppId().GetSource()),
			SourceAppID: req.GetBindAppId().GetSourceAppId(),
		})
	if err != nil {
		return nil, err
	}
	return &pb.BindAppResponse{BindAppId: &librarian.InternalID{
		Id: a.InternalID,
	}}, nil
}
func (s *LibrarianSephirahServiceService) UnBindApp(ctx context.Context, req *pb.UnBindAppRequest) (
	*pb.UnBindAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) RefreshApp(ctx context.Context, req *pb.RefreshAppRequest) (
	*pb.RefreshAppResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) UploadArtifacts(ctx context.Context, req *pb.UploadArtifactsRequest) (
	*pb.UploadArtifactsResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) DownloadArtifacts(ctx context.Context, req *pb.DownloadArtifactsRequest) (
	*pb.DownloadArtifactsResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) ListArtifacts(ctx context.Context, req *pb.ListArtifactsRequest) (
	*pb.ListArtifactsResponse, error,
) {
	return nil, pb.ErrorErrorReasonNotImplemented("impl in next version")
}
