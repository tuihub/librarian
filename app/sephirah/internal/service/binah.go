package service

import (
	"errors"
	"io"

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
	ctx := conn.Context()
	file, bizErr := s.b.NewUploadFile(ctx)
	if bizErr != nil {
		return bizErr
	}
	for {
		if req, err := conn.Recv(); err != nil {
			if errors.Is(err, io.EOF) {
				return file.Finish(ctx)
			}
			return err
		} else if _, err = file.Writer.Write(req.Data); err != nil {
			return err
		}
		if err := conn.Send(&pb.SimpleUploadFileResponse{
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
