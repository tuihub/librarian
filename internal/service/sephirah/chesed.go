package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/converter"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

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
		Status:      modelchesed.ImageStatusUploaded,
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
	res, total, err := s.c.ListImages(ctx, model.ToBizPaging(req.GetPaging()))
	if err != nil {
		return nil, err
	}
	return &pb.ListImagesResponse{
		Paging: &librarian.PagingResponse{TotalSize: total},
		Ids:    converter.ToPBInternalIDList(res),
	}, nil
}
func (s *LibrarianSephirahServiceService) SearchImages(ctx context.Context,
	req *pb.SearchImagesRequest) (*pb.SearchImagesResponse, error) {
	res, err := s.c.SearchImages(ctx, model.ToBizPaging(req.GetPaging()), req.GetKeywords())
	if err != nil {
		return nil, err
	}
	return &pb.SearchImagesResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(len(res))},
		Ids:    converter.ToPBInternalIDList(res),
	}, nil
}
func (s *LibrarianSephirahServiceService) GetImage(ctx context.Context, req *pb.GetImageRequest) (
	*pb.GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (s *LibrarianSephirahServiceService) DownloadImage(ctx context.Context, req *pb.DownloadImageRequest) (
	*pb.DownloadImageResponse, error) {
	token, err := s.c.DownloadImage(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.DownloadImageResponse{DownloadToken: token}, nil
}
