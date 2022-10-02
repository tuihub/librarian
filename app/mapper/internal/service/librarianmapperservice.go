package service

import (
	"context"

	"github.com/tuihub/librarian/app/mapper/internal/biz"
	pb "github.com/tuihub/protos/pkg/librarian/mapper/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type LibrarianMapperServiceService struct {
	pb.UnimplementedLibrarianMapperServiceServer

	m *biz.Mapper
}

func NewLibrarianMapperServiceService(m *biz.Mapper) pb.LibrarianMapperServiceServer {
	return &LibrarianMapperServiceService{
		m: m,
	}
}

func (s *LibrarianMapperServiceService) InsertVertex(ctx context.Context, req *pb.InsertVertexRequest) (
	*pb.InsertVertexResponse, error) {
	if len(req.GetVertexList()) != 1 {
		return nil, errors.BadRequest("", "")
	}
	err := s.m.InsertVertex(ctx, biz.Vertex{
		ID:   req.GetVertexList()[0].GetVid(),
		Type: toBizVertexType(req.GetVertexList()[0].GetType()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.InsertVertexResponse{}, nil
}
func (s *LibrarianMapperServiceService) DeleteVertex(ctx context.Context, req *pb.DeleteVertexRequest) (
	*pb.DeleteVertexResponse, error) {
	return &pb.DeleteVertexResponse{}, nil
}
func (s *LibrarianMapperServiceService) UpdateVertex(ctx context.Context, req *pb.UpdateVertexRequest) (
	*pb.UpdateVertexResponse, error) {
	return &pb.UpdateVertexResponse{}, nil
}
func (s *LibrarianMapperServiceService) InsertEdge(ctx context.Context, req *pb.InsertEdgeRequest) (
	*pb.InsertEdgeResponse, error) {
	if len(req.GetEdgeList()) != 1 {
		return nil, errors.BadRequest("", "")
	}
	err := s.m.InsertEdge(ctx, biz.Edge{
		SourceID:      req.GetEdgeList()[0].GetSrcVid(),
		DestinationID: req.GetEdgeList()[0].GetDstVid(),
		Type:          toBizEdgeType(req.GetEdgeList()[0].GetType()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.InsertEdgeResponse{}, nil
}
func (s *LibrarianMapperServiceService) DeleteEdge(ctx context.Context, req *pb.DeleteEdgeRequest) (
	*pb.DeleteEdgeResponse, error) {
	return &pb.DeleteEdgeResponse{}, nil
}
func (s *LibrarianMapperServiceService) UpdateEdge(ctx context.Context, req *pb.UpdateEdgeRequest) (
	*pb.UpdateEdgeResponse, error) {
	return &pb.UpdateEdgeResponse{}, nil
}
func (s *LibrarianMapperServiceService) GoFromVertex(ctx context.Context, req *pb.GoFromVertexRequest) (
	*pb.GoFromVertexResponse, error) {
	return &pb.GoFromVertexResponse{}, nil
}
func (s *LibrarianMapperServiceService) FetchEqualVertex(ctx context.Context, req *pb.FetchEqualVertexRequest) (
	*pb.FetchEqualVertexResponse, error) {
	vl, err := s.m.FetchEqualVertex(ctx, biz.Vertex{
		ID: req.GetSrcVid(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.FetchEqualVertexResponse{
		VertexList: toPBVertexList(vl),
	}, nil
}
func (s *LibrarianMapperServiceService) FetchEqualVertexNeighbor(ctx context.Context,
	req *pb.FetchEqualVertexNeighborRequest) (*pb.FetchEqualVertexNeighborResponse, error) {
	return &pb.FetchEqualVertexNeighborResponse{}, nil
}
func (s *LibrarianMapperServiceService) FindPath(ctx context.Context, req *pb.FindPathRequest) (
	*pb.FindPathResponse, error) {
	return &pb.FindPathResponse{}, nil
}
