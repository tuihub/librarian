package service

import (
	"context"
	"errors"
	"io"

	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) UploadFile(conn pb.LibrarianSephirahService_UploadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) DownloadFile(conn pb.LibrarianSephirahService_DownloadFileServer) error {
	return pb.ErrorErrorReasonNotImplemented("impl in next version")
}
func (s *LibrarianSephirahServiceService) SimpleUploadFile(
	conn pb.LibrarianSephirahService_SimpleUploadFileServer,
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
		var req *pb.SimpleUploadFileRequest
		if req, err = conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return file.Finish(ctx)
			}
			return err
		} else if _, err = file.Writer.Write(req.Data); err != nil {
			return err
		}
		if err = conn.Send(&pb.SimpleUploadFileResponse{
			Status: pb.FileTransferStatus_FILE_TRANSFER_STATUS_IN_PROGRESS,
		}); err != nil {
			return err
		}
	}
}
func (s *LibrarianSephirahServiceService) SimpleDownloadFile(req *pb.SimpleDownloadFileRequest,
	conn pb.LibrarianSephirahService_SimpleDownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SimpleDownloadFile not implemented")
}

func (s *LibrarianSephirahServiceService) PresignedUploadFile(ctx context.Context,
	req *pb.PresignedUploadFileRequest) (*pb.PresignedUploadFileResponse, error) {
	res, err := s.b.PresignedUploadFile(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.PresignedUploadFileResponse{UploadUrl: res}, nil
}

func (s *LibrarianSephirahServiceService) PresignedUploadFileStatus(ctx context.Context,
	req *pb.PresignedUploadFileStatusRequest) (*pb.PresignedUploadFileStatusResponse, error) {
	if req.GetStatus() == pb.FileTransferStatus_FILE_TRANSFER_STATUS_SUCCESS {
		err := s.b.PresignedUploadFileComplete(ctx)
		if err != nil {
			return nil, err
		}
	}
	return &pb.PresignedUploadFileStatusResponse{}, nil
}

func (s *LibrarianSephirahServiceService) PresignedDownloadFile(ctx context.Context,
	req *pb.PresignedDownloadFileRequest) (*pb.PresignedDownloadFileResponse, error) {
	res, err := s.b.PresignedDownloadFile(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.PresignedDownloadFileResponse{DownloadUrl: res}, nil
}
