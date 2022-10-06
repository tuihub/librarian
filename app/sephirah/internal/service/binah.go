package service

import (
	"errors"
	"io"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
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
