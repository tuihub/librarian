package sephirah

import (
	"context"
	"errors"
	"io"

	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahService) UploadFile(conn sephirah.LibrarianSephirahService_UploadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahService) DownloadFile(conn sephirah.LibrarianSephirahService_DownloadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahService) SimpleUploadFile(
	conn sephirah.LibrarianSephirahService_SimpleUploadFileServer,
) error {
	ctx, err := libapp.NewStreamMiddlewareJwt(s.auth)(conn.Context())
	if err != nil {
		return err
	}
	file, bizErr := s.b.NewSimpleUploadFile(ctx)
	if bizErr != nil {
		return bizErr
	}
	for {
		var req *sephirah.SimpleUploadFileRequest
		if req, err = conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return file.Finish(ctx)
			}
			return err
		} else if _, err = file.Writer.Write(req.GetData()); err != nil {
			return err
		}
		if err = conn.Send(&sephirah.SimpleUploadFileResponse{
			Status: sephirah.FileTransferStatus_FILE_TRANSFER_STATUS_IN_PROGRESS,
		}); err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahService) SimpleDownloadFile(req *sephirah.SimpleDownloadFileRequest,
	conn sephirah.LibrarianSephirahService_SimpleDownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SimpleDownloadFile not implemented")
}

func (s *LibrarianSephirahService) PresignedUploadFile(ctx context.Context,
	req *sephirah.PresignedUploadFileRequest) (*sephirah.PresignedUploadFileResponse, error) {
	res, err := s.b.PresignedUploadFile(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.PresignedUploadFileResponse{UploadUrl: res}, nil
}

func (s *LibrarianSephirahService) PresignedUploadFileStatus(ctx context.Context,
	req *sephirah.PresignedUploadFileStatusRequest) (*sephirah.PresignedUploadFileStatusResponse, error) {
	if req.GetStatus() == sephirah.FileTransferStatus_FILE_TRANSFER_STATUS_SUCCESS {
		err := s.b.PresignedUploadFileComplete(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &sephirah.PresignedUploadFileStatusResponse{}, nil
}

func (s *LibrarianSephirahService) PresignedDownloadFile(ctx context.Context,
	req *sephirah.PresignedDownloadFileRequest) (*sephirah.PresignedDownloadFileResponse, error) {
	res, err := s.b.PresignedDownloadFile(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.PresignedDownloadFileResponse{DownloadUrl: res}, nil
}
