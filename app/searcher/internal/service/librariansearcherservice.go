package service

import (
	"context"

	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
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
		Id: &librarian.InternalID{Id: id},
	}, nil
}
func (s *LibrarianSearcherServiceService) NewBatchIDs(ctx context.Context, req *pb.NewBatchIDsRequest) (
	*pb.NewBatchIDsResponse, error) {
	ids, err := s.uc.NewBatchIDs(ctx, int(req.GetNum()))
	if err != nil {
		return nil, err
	}
	res := make([]*librarian.InternalID, len(ids))
	for i, id := range ids {
		res[i] = &librarian.InternalID{
			Id: id,
		}
	}
	return &pb.NewBatchIDsResponse{Ids: res}, nil
}
func (s *LibrarianSearcherServiceService) DescribeID(ctx context.Context, req *pb.DescribeIDRequest) (
	*pb.DescribeIDResponse, error) {
	err := s.uc.DescribeID(
		ctx,
		model.InternalID(req.GetId().GetId()),
		toBizIndex(req.GetIndex()),
		req.GetMode() == pb.DescribeIDRequest_DESCRIBE_MODE_APPEND,
		req.GetDescription(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.DescribeIDResponse{}, nil
}
func (s *LibrarianSearcherServiceService) SearchID(ctx context.Context, req *pb.SearchIDRequest) (
	*pb.SearchIDResponse, error) {
	res, err := s.uc.SearchID(
		ctx,
		model.ToBizPaging(req.GetPaging()),
		toBizIndex(req.GetIndex()),
		req.GetQuery(),
	)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.SearchIDResponse_Result, len(res))
	for i := range res {
		result[i] = &pb.SearchIDResponse_Result{
			Id:   &librarian.InternalID{Id: int64(res[i].ID)},
			Rank: res[i].Rank,
		}
	}
	return &pb.SearchIDResponse{
		Paging: nil,
		Result: result,
	}, nil
}
