package porter

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"
)

func (s *LibrarianSephirahPorterService) GetNotifyTargetItems(
	ctx context.Context,
	req *porter.GetNotifyTargetItemsRequest,
) (*porter.GetNotifyTargetItemsResponse, error) {
	fr, items, err := s.a.PorterGetNotifyTargetItems(ctx,
		converter.ToBizInternalID(req.GetId()),
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &porter.GetNotifyTargetItemsResponse{
		Paging:      nil,
		Destination: converter.ToPBFeatureRequest(fr),
		Items:       converter.ToPBFeedItemList(items),
	}, nil
}
