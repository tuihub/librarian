package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	converter2 "github.com/tuihub/librarian/internal/service/sephirah/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (s *LibrarianSephirahServiceService) PorterGetNotifyTargetItems(
	ctx context.Context,
	req *pb.PorterGetNotifyTargetItemsRequest,
) (*pb.PorterGetNotifyTargetItemsResponse, error) {
	fr, items, err := s.a.PorterGetNotifyTargetItems(ctx,
		converter2.ToBizInternalID(req.GetId()),
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.PorterGetNotifyTargetItemsResponse{
		Paging:      nil,
		Destination: converter2.ToPBFeatureRequest(fr),
		Items:       converter2.ToPBFeedItemList(items),
	}, nil
}
