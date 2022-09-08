package service

import (
	"context"
	"io"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	uc *biz.GreeterUsecase
}

func NewLibrarianSephirahServiceService(uc *biz.GreeterUsecase) pb.LibrarianSephirahServiceServer {
	return &LibrarianSephirahServiceService{
		uc: uc,
	}
}

func (s *LibrarianSephirahServiceService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (
	*pb.GetTokenResponse, error) {
	return &pb.GetTokenResponse{}, nil
}
func (s *LibrarianSephirahServiceService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (
	*pb.RefreshTokenResponse, error) {
	return &pb.RefreshTokenResponse{}, nil
}
func (s *LibrarianSephirahServiceService) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (
	*pb.GenerateTokenResponse, error) {
	return &pb.GenerateTokenResponse{}, nil
}
func (s *LibrarianSephirahServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (
	*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (
	*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (
	*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}
func (s *LibrarianSephirahServiceService) LinkAccount(ctx context.Context, req *pb.LinkAccountRequest) (
	*pb.LinkAccountResponse, error) {
	return &pb.LinkAccountResponse{}, nil
}
func (s *LibrarianSephirahServiceService) UnLinkAccount(ctx context.Context, req *pb.UnLinkAccountRequest) (
	*pb.UnLinkAccountResponse, error) {
	return &pb.UnLinkAccountResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListLinkAccount(ctx context.Context, req *pb.ListLinkAccountRequest) (
	*pb.ListLinkAccountResponse, error) {
	return &pb.ListLinkAccountResponse{}, nil
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
	return &pb.CreateAppResponse{}, nil
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
