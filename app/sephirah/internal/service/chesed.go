package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) UploadImage(ctx context.Context, req *pb.UploadImageRequest) (
	*pb.UploadImageResponse, error) {
	fm := converter.ToBizFileMetadata(req.GetFileMetadata())
	if fm == nil {
		return nil, pb.ErrorErrorReasonBadRequest("app required")
	}
	token, err := s.c.UploadImage(ctx, modelchesed.Image{
		ID:          0,
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}, *fm)
	if err != nil {
		return nil, err
	}
	return &pb.UploadImageResponse{UploadToken: token}, nil
}
func (s *LibrarianSephirahServiceService) UpdateImage(ctx context.Context, req *pb.UpdateImageRequest) (
	*pb.UpdateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImage not implemented")
}
func (s *LibrarianSephirahServiceService) ListImages(ctx context.Context, req *pb.ListImagesRequest) (
	*pb.ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (s *LibrarianSephirahServiceService) GetImage(ctx context.Context, req *pb.GetImageRequest) (
	*pb.GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (s *LibrarianSephirahServiceService) DownloadImage(ctx context.Context, req *pb.DownloadImageRequest) (
	*pb.DownloadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadImage not implemented")
}