package service

import (
	"context"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type LibrarianSearcherServiceService struct {
	pb.UnimplementedLibrarianSearcherServiceServer

	uc *biz.Searcher
}

func NewLibrarianSearcherServiceService(uc *biz.Searcher) pb.LibrarianSearcherServiceServer {
	return &LibrarianSearcherServiceService{
		UnimplementedLibrarianSearcherServiceServer: pb.UnimplementedLibrarianSearcherServiceServer{},
		uc: uc,
	}
}

func (s *LibrarianSearcherServiceService) NewID(ctx context.Context, req *pb.NewIDRequest) (
	*pb.NewIDResponse, error) {
	id, err := s.uc.NewID(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.NewIDResponse{
		Id: id,
	}, nil
}
