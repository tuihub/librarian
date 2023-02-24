package service

import (
	"context"

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
func (s *LibrarianSephirahServiceService) ListFeed(
	ctx context.Context,
	req *pb.ListFeedRequest,
) (*pb.ListFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeed not implemented")
}
func (s *LibrarianSephirahServiceService) ListFeedItem(
	ctx context.Context,
	req *pb.ListFeedItemRequest,
) (*pb.ListFeedItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeedItem not implemented")
}
func (s *LibrarianSephirahServiceService) GetFeedItem(
	ctx context.Context,
	req *pb.GetFeedItemRequest,
) (*pb.GetFeedItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedItem not implemented")
}
