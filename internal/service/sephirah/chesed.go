package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	pb "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahService) UploadImage(ctx context.Context, req *sephirah.UploadImageRequest) (
	*sephirah.UploadImageResponse, error) {
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
	return &sephirah.UploadImageResponse{UploadToken: token}, nil
}
func (s *LibrarianSephirahService) UpdateImage(ctx context.Context, req *sephirah.UpdateImageRequest) (
	*sephirah.UpdateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateImage not implemented")
}
func (s *LibrarianSephirahService) ListImages(ctx context.Context, req *sephirah.ListImagesRequest) (
	*sephirah.ListImagesResponse, error) {
	res, total, err := s.c.ListImages(ctx, model.ToBizPaging(req.GetPaging()))
	if err != nil {
		return nil, err
	}
	return &sephirah.ListImagesResponse{
		Paging: &pb.PagingResponse{TotalSize: total},
		Ids:    converter.ToPBInternalIDList(res),
	}, nil
}
func (s *LibrarianSephirahService) SearchImages(ctx context.Context,
	req *sephirah.SearchImagesRequest) (*sephirah.SearchImagesResponse, error) {
	res, err := s.c.SearchImages(ctx, model.ToBizPaging(req.GetPaging()), req.GetKeywords())
	if err != nil {
		return nil, err
	}
	return &sephirah.SearchImagesResponse{
		Paging: &pb.PagingResponse{TotalSize: int64(len(res))},
		Ids:    converter.ToPBInternalIDList(res),
	}, nil
}
func (s *LibrarianSephirahService) GetImage(ctx context.Context, req *sephirah.GetImageRequest) (
	*sephirah.GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (s *LibrarianSephirahService) DownloadImage(ctx context.Context, req *sephirah.DownloadImageRequest) (
	*sephirah.DownloadImageResponse, error) {
	token, err := s.c.DownloadImage(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.DownloadImageResponse{DownloadToken: token}, nil
}
