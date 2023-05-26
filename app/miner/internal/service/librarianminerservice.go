package service

import (
	"context"

	"github.com/tuihub/librarian/app/miner/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LibrarianMinerServiceService struct {
	pb.UnimplementedLibrarianMinerServiceServer

	m *biz.Miner
}

func NewLibrarianMinerServiceService(m *biz.Miner) pb.LibrarianMinerServiceServer {
	return &LibrarianMinerServiceService{
		UnimplementedLibrarianMinerServiceServer: pb.UnimplementedLibrarianMinerServiceServer{},
		m:                                        m,
	}
}

func (s *LibrarianMinerServiceService) RecognizeImageBinary(
	req pb.LibrarianMinerService_RecognizeImageBinaryServer) error {
	return status.Errorf(codes.Unimplemented, "method RecognizeImageBinary not implemented")
}
func (s *LibrarianMinerServiceService) RecognizeImageURL(
	ctx context.Context, req *pb.RecognizeImageURLRequest) (*pb.RecognizeImageURLResponse, error) {
	res, err := s.m.RecognizeImageURL(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	return &pb.RecognizeImageURLResponse{Results: ToPBOCRResults(res)}, nil
}
