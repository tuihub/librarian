package service

import (
	"context"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LibrarianSephirahServiceService) CreateFeedConfig(
	ctx context.Context,
	req *pb.CreateFeedConfigRequest,
) (*pb.CreateFeedConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeedConfig not implemented")
}
func (s *LibrarianSephirahServiceService) UpdateFeedConfig(
	ctx context.Context,
	req *pb.UpdateFeedConfigRequest,
) (*pb.UpdateFeedConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeedConfig not implemented")
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
