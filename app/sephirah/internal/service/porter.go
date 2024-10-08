package service

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (s *LibrarianSephirahServiceService) PorterGetNotifyTargetItems(
	ctx context.Context,
	req *pb.PorterGetNotifyTargetItemsRequest,
) (*pb.PorterGetNotifyTargetItemsResponse, error) {
	fr, items, err := s.a.PorterGetNotifyTargetItems(ctx,
		converter.ToBizInternalID(req.GetId()),
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.PorterGetNotifyTargetItemsResponse{
		Paging:      nil,
		Destination: converter.ToPBFeatureRequest(fr),
		Items:       converter.ToPBFeedItemList(items),
	}, nil
}
