package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/internal/lib/logger"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t *biz.TipherethUseCase
	g *biz.GeburaUseCase
}

func NewLibrarianSephirahServiceService(
	t *biz.TipherethUseCase, g *biz.GeburaUseCase) pb.LibrarianSephirahServiceServer {
	return &LibrarianSephirahServiceService{
		t: t,
		g: g,
	}
}

func (s *LibrarianSephirahServiceService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (
	*pb.GetTokenResponse, error) {
	accessToken, refreshToken, err := s.t.GetToken(ctx, &biz.User{
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
	*pb.RefreshTokenResponse, error) {
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
	*pb.GenerateTokenResponse, error) {
	return nil, pb.ErrorErrorReasonNotImplemented("")
}
func (s *LibrarianSephirahServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (
	*pb.CreateUserResponse, error) {
	u, err := s.t.AddUser(ctx, &biz.User{
		UserName: req.GetUsername(),
		PassWord: req.GetPassword(),
		UserType: toLibAuthUserType(req.GetType()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id: &pb.InternalID{Id: u.InternalID},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (
	*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (
	*pb.ListUserResponse, error) {
	u, err := s.t.ListUser(ctx,
		biz.Paging{
			PageSize: int(req.GetPageSize()),
			PageNum:  int(req.GetPageNum()),
		},
		toLibAuthUserTypeList(req.GetTypeFilter()),
		toBizUserStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListUserResponse{
		UserList: toPBUserList(u),
	}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error) {
	return nil, pb.ErrorErrorReasonNotImplemented("")
}
func (s *LibrarianSephirahServiceService) UnLinkAccount(ctx context.Context, req *pb.UnLinkAccountRequest) (
	*pb.UnLinkAccountResponse, error) {
	return nil, pb.ErrorErrorReasonNotImplemented("")
}
func (s *LibrarianSephirahServiceService) ListLinkAccount(ctx context.Context, req *pb.ListLinkAccountRequest) (
	*pb.ListLinkAccountResponse, error) {
	return nil, pb.ErrorErrorReasonNotImplemented("")
}
func (s *LibrarianSephirahServiceService) UploadFile(conn pb.LibrarianSephirahService_UploadFileServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		err = conn.Send(&pb.UploadFileResponse{})
		if err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) DownloadFile(conn pb.LibrarianSephirahService_DownloadFileServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		err = conn.Send(&pb.DownloadFileResponse{})
		if err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) SimpleUploadFile(
	conn pb.LibrarianSephirahService_SimpleUploadFileServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		err = conn.Send(&pb.SimpleUploadFileResponse{})
		if err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) SimpleDownloadFile(
	conn pb.LibrarianSephirahService_SimpleDownloadFileServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		err = conn.Send(&pb.SimpleDownloadFileResponse{})
		if err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (
	*pb.CreateAppResponse, error) {
	app := req.GetApp()
	if app == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app required")
	}
	a, err := s.g.CreateApp(ctx, &biz.App{
		Name:            app.GetName(),
		Type:            toBizAppType(app.GetType()),
		ShorDescription: app.GetShortDescription(),
		ImageURL:        app.GetImageUrl(),
		Details:         toBizAppDetail(app.GetDetails()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateAppResponse{
		Id: &pb.InternalID{Id: a.InternalID},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (
	*pb.UpdateAppResponse, error) {
	return &pb.UpdateAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListApp(ctx context.Context, req *pb.ListAppRequest) (
	*pb.ListAppResponse, error) {
	return &pb.ListAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) BindApp(ctx context.Context, req *pb.BindAppRequest) (
	*pb.BindAppResponse, error) {
	return &pb.BindAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnBindApp(ctx context.Context, req *pb.UnBindAppRequest) (
	*pb.UnBindAppResponse, error) {
	return &pb.UnBindAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) RefreshApp(ctx context.Context, req *pb.RefreshAppRequest) (
	*pb.RefreshAppResponse, error) {
	return &pb.RefreshAppResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UploadArtifacts(ctx context.Context, req *pb.UploadArtifactsRequest) (
	*pb.UploadArtifactsResponse, error) {
	return &pb.UploadArtifactsResponse{}, nil
}
func (s *LibrarianSephirahServiceService) DownloadArtifacts(ctx context.Context, req *pb.DownloadArtifactsRequest) (
	*pb.DownloadArtifactsResponse, error) {
	return &pb.DownloadArtifactsResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListArtifacts(ctx context.Context, req *pb.ListArtifactsRequest) (
	*pb.ListArtifactsResponse, error) {
	return &pb.ListArtifactsResponse{}, nil
}
