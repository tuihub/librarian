package service

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/porter/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type LibrarianPorterServiceService struct {
	pb.UnimplementedLibrarianPorterServiceServer

	uc *biz.GreeterUsecase
}

func NewLibrarianPorterServiceService(uc *biz.GreeterUsecase) pb.LibrarianPorterServiceServer {
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
