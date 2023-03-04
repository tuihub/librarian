package service

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) CreateFeedConfig(
	ctx context.Context,
	req *pb.CreateFeedConfigRequest,
) (*pb.CreateFeedConfigResponse, error) {
	id, err := s.y.CreateFeedConfig(ctx, s.converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedConfigResponse{
		Id: &librarian.InternalID{Id: id},
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateFeedConfig(
	ctx context.Context,
	req *pb.UpdateFeedConfigRequest,
) (*pb.UpdateFeedConfigResponse, error) {
	err := s.y.UpdateFeedConfig(ctx, s.converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFeedConfigResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListFeeds(
	ctx context.Context,
	req *pb.ListFeedsRequest,
) (*pb.ListFeedsResponse, error) {
	_, err := s.y.ListFeeds(ctx, model.Paging{
		PageSize: int(req.GetPaging().GetPageSize()),
		PageNum:  int(req.GetPaging().GetPageSize()),
	}, s.converter.ToBizInternalIDList(
		req.GetIdFilter()),
		s.converter.ToBizInternalIDList(req.GetAuthorIdFilter()),
		s.converter.ToBizFeedConfigSourceList(req.GetSourceFilter()),
		s.converter.ToBizFeedConfigStatusList(req.GetStatusFilter()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedsResponse{
		Paging:          nil,
		FeedsWithConfig: nil,
	}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedItems(
	ctx context.Context,
	req *pb.ListFeedItemsRequest,
) (*pb.ListFeedItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeedItem not implemented")
}
func (s *LibrarianSephirahServiceService) GetFeedItem(
	ctx context.Context,
	req *pb.GetFeedItemRequest,
) (*pb.GetFeedItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedItem not implemented")
}
