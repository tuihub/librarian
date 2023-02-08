package service

import (
	"context"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

type LibrarianMapperServiceService struct {
	pb.UnimplementedLibrarianMapperServiceServer

	m *biz.Mapper
}

func NewLibrarianMapperServiceService(m *biz.Mapper) pb.LibrarianMapperServiceServer {
	return &LibrarianMapperServiceService{
		UnimplementedLibrarianMapperServiceServer: pb.UnimplementedLibrarianMapperServiceServer{},
		m: m,
	}
}

func (s *LibrarianMapperServiceService) InsertVertex(ctx context.Context, req *pb.InsertVertexRequest) (
	*pb.InsertVertexResponse, error) {
	err := s.m.InsertVertex(ctx, toBizVertexList(req.GetVertexList()))
	if err != nil {
		return nil, err
	}
	return &pb.InsertVertexResponse{}, nil
}
func (s *LibrarianMapperServiceService) InsertEdge(ctx context.Context, req *pb.InsertEdgeRequest) (
	*pb.InsertEdgeResponse, error) {
	err := s.m.InsertEdge(ctx, toBizEdgeList(req.GetEdgeList()))
	if err != nil {
		return nil, err
	}
	return &pb.InsertEdgeResponse{}, nil
}
func (s *LibrarianMapperServiceService) FetchEqualVertex(ctx context.Context, req *pb.FetchEqualVertexRequest) (
	*pb.FetchEqualVertexResponse, error) {
	vl, err := s.m.FetchEqualVertex(ctx, biz.Vertex{
		ID:   req.GetSrcVid(),
		Type: 0,
	})
	if err != nil {
		return nil, err
	}
	return &pb.FetchEqualVertexResponse{
		VertexList: toPBVertexList(vl),
	}, nil
}
