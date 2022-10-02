package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LibrarianPorterServiceService struct {
	pb.UnimplementedLibrarianPorterServiceServer

	uc *bizsteam.SteamUseCase
}

func NewLibrarianPorterServiceService(uc *bizsteam.SteamUseCase) pb.LibrarianPorterServiceServer {
	return &LibrarianPorterServiceService{uc: uc}
}

func (s *LibrarianPorterServiceService) PullFeed(ctx context.Context, req *pb.PullFeedRequest) (
	*pb.PullFeedResponse, error) {
	return &pb.PullFeedResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullDB(ctx context.Context, req *pb.PullDBRequest) (
	*pb.PullDBResponse, error) {
	return &pb.PullDBResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullWiki(ctx context.Context, req *pb.PullWikiRequest) (
	*pb.PullWikiResponse, error) {
	return &pb.PullWikiResponse{}, nil
}
func (s *LibrarianPorterServiceService) PullData(req *pb.PullDataRequest,
	conn pb.LibrarianPorterService_PullDataServer) error {
	for {
		err := conn.Send(&pb.PullDataResponse{})
		if err != nil {
			return err
		}
	}
}
func (s *LibrarianPorterServiceService) PullAccount(
	ctx context.Context, req *pb.PullAccountRequest) (*pb.PullAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullAccount not implemented")
}
func (s *LibrarianPorterServiceService) PullApp(
	ctx context.Context, req *pb.PullAppRequest) (*pb.PullAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullApp not implemented")
}
func (s *LibrarianPorterServiceService) PullAccountAppRelation(
	ctx context.Context, req *pb.PullAccountAppRelationRequest) (*pb.PullAccountAppRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullAccountAppRelation not implemented")
}
func (s *LibrarianPorterServiceService) PushData(conn pb.LibrarianPorterService_PushDataServer) error {
	for {
		_, err := conn.Recv()
		if errors.Is(err, io.EOF) {
			return conn.SendAndClose(&pb.PushDataResponse{})
		}
		if err != nil {
			return err
		}
	}
}
