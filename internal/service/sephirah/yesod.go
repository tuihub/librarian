package sephirah

import (
	"context"
	"github.com/tuihub/librarian/internal/service/sephirah/internal/converter"

	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahServiceService) CreateFeedConfig(
	ctx context.Context,
	req *pb.CreateFeedConfigRequest,
) (*pb.CreateFeedConfigResponse, error) {
	id, err := s.y.CreateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedConfigResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateFeedConfig(
	ctx context.Context,
	req *pb.UpdateFeedConfigRequest,
) (*pb.UpdateFeedConfigResponse, error) {
	err := s.y.UpdateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFeedConfigResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedConfigs(
	ctx context.Context,
	req *pb.ListFeedConfigsRequest,
) (*pb.ListFeedConfigsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	feeds, total, err := s.y.ListFeeds(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		converter.ToBizFeedConfigStatusList(req.GetStatusFilter()),
		req.GetCategoryFilter(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedConfigsResponse{
		Paging:          &librarian.PagingResponse{TotalSize: int64(total)},
		FeedsWithConfig: converter.ToPBFeedWithConfigList(feeds),
	}, nil
}
func (s *LibrarianSephirahServiceService) CreateFeedActionSet(
	ctx context.Context,
	req *pb.CreateFeedActionSetRequest,
) (*pb.CreateFeedActionSetResponse, error) {
	id, err := s.y.CreateFeedActionSet(ctx, converter.ToBizFeedActionSet(req.GetActionSet()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedActionSetResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahServiceService) UpdateFeedActionSet(
	ctx context.Context,
	req *pb.UpdateFeedActionSetRequest,
) (*pb.UpdateFeedActionSetResponse, error) {
	err := s.y.UpdateFeedActionSet(ctx, converter.ToBizFeedActionSet(req.GetActionSet()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFeedActionSetResponse{}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedActionSets(
	ctx context.Context,
	req *pb.ListFeedActionSetsRequest,
) (*pb.ListFeedActionSetsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	actions, total, err := s.y.ListFeedActionSets(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedActionSetsResponse{
		Paging:     &librarian.PagingResponse{TotalSize: int64(total)},
		ActionSets: converter.ToPBFeedActionSetList(actions),
	}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedCategories(
	ctx context.Context,
	req *pb.ListFeedCategoriesRequest,
) (*pb.ListFeedCategoriesResponse, error) {
	res, err := s.y.ListFeedCategories(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedCategoriesResponse{Categories: res}, nil
}
func (s *LibrarianSephirahServiceService) ListFeedPlatforms(
	ctx context.Context,
	req *pb.ListFeedPlatformsRequest,
) (*pb.ListFeedPlatformsResponse, error) {
	res, err := s.y.ListFeedPlatforms(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedPlatformsResponse{Platforms: res}, nil
}

func (s *LibrarianSephirahServiceService) ListFeedItems(
	ctx context.Context,
	req *pb.ListFeedItemsRequest,
) (*pb.ListFeedItemsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	items, total, err := s.y.ListFeedItems(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetFeedIdFilter()),
		req.GetAuthorFilter(),
		req.GetPublishPlatformFilter(),
		converter.ToBizTimeRange(req.GetPublishTimeRange()),
		req.GetCategoryFilter(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedItemsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Items:  converter.ToPBFeedItemDigestList(items),
	}, nil
}
func (s *LibrarianSephirahServiceService) GroupFeedItems(
	ctx context.Context,
	req *pb.GroupFeedItemsRequest,
) (*pb.GroupFeedItemsResponse, error) {
	itemMap, err := s.y.GroupFeedItems(ctx,
		converter.ToBizGroupFeedItemsBy(req.GetPublishTimeAggregation().GetAggregationType()),
		converter.ToBizInternalIDList(req.GetFeedIdFilter()),
		req.GetAuthorFilter(),
		req.GetPublishPlatformFilter(),
		converter.ToBizTimeRange(req.GetPublishTimeAggregation().GetTimeRange()),
		int(req.GetGroupSize()),
		req.GetCategoryFilter(),
	)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.GroupFeedItemsResponse_FeedItemsGroup, 0, len(itemMap))
	for timeRange, items := range itemMap {
		res = append(res, &pb.GroupFeedItemsResponse_FeedItemsGroup{
			TimeRange: converter.ToPBTimeRange(&timeRange),
			Items:     converter.ToPBFeedItemDigestList(items),
		})
	}
	return &pb.GroupFeedItemsResponse{Groups: res}, nil
}
func (s *LibrarianSephirahServiceService) GetFeedItem(
	ctx context.Context,
	req *pb.GetFeedItemRequest,
) (*pb.GetFeedItemResponse, error) {
	item, err := s.y.GetFeedItem(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.GetFeedItemResponse{
		Item: converter.ToPBFeedItem(item),
	}, nil
}

func (s *LibrarianSephirahServiceService) GetBatchFeedItems(
	ctx context.Context,
	req *pb.GetBatchFeedItemsRequest,
) (*pb.GetBatchFeedItemsResponse, error) {
	items, err := s.y.GetFeedItems(ctx, converter.ToBizInternalIDList(req.GetIds()))
	if err != nil {
		return nil, err
	}
	return &pb.GetBatchFeedItemsResponse{
		Items: converter.ToPBFeedItemList(items),
	}, nil
}

func (s *LibrarianSephirahServiceService) ReadFeedItem(
	ctx context.Context,
	req *pb.ReadFeedItemRequest,
) (*pb.ReadFeedItemResponse, error) {
	err := s.y.ReadFeedItem(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.ReadFeedItemResponse{}, nil
}

func (s *LibrarianSephirahServiceService) CreateFeedItemCollection(
	ctx context.Context,
	req *pb.CreateFeedItemCollectionRequest,
) (*pb.CreateFeedItemCollectionResponse, error) {
	_, err := s.y.CreateFeedItemCollection(ctx, converter.ToBizFeedItemCollection(req.GetCollection()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateFeedItemCollectionResponse{}, nil
}

func (s *LibrarianSephirahServiceService) UpdateFeedItemCollection(
	ctx context.Context,
	req *pb.UpdateFeedItemCollectionRequest,
) (*pb.UpdateFeedItemCollectionResponse, error) {
	err := s.y.UpdateFeedItemCollection(ctx, converter.ToBizFeedItemCollection(req.GetCollection()))
	if err != nil {
		return nil, err
	}
	return &pb.UpdateFeedItemCollectionResponse{}, nil
}

func (s *LibrarianSephirahServiceService) ListFeedItemCollections(
	ctx context.Context,
	req *pb.ListFeedItemCollectionsRequest,
) (*pb.ListFeedItemCollectionsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	collections, total, err := s.y.ListFeedItemCollections(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetIdFilter()),
		req.GetCategoryFilter(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedItemCollectionsResponse{
		Paging:      &librarian.PagingResponse{TotalSize: int64(total)},
		Collections: converter.ToPBFeedItemCollectionList(collections),
	}, nil
}

func (s *LibrarianSephirahServiceService) AddFeedItemToCollection(
	ctx context.Context,
	req *pb.AddFeedItemToCollectionRequest,
) (*pb.AddFeedItemToCollectionResponse, error) {
	err := s.y.AddFeedItemToCollection(ctx,
		converter.ToBizInternalID(req.GetCollectionId()),
		converter.ToBizInternalID(req.GetFeedItemId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.AddFeedItemToCollectionResponse{}, nil
}

func (s *LibrarianSephirahServiceService) RemoveFeedItemFromCollection(
	ctx context.Context,
	req *pb.RemoveFeedItemFromCollectionRequest,
) (*pb.RemoveFeedItemFromCollectionResponse, error) {
	err := s.y.RemoveFeedItemFromCollection(
		ctx,
		converter.ToBizInternalID(req.GetCollectionId()),
		converter.ToBizInternalID(req.GetFeedItemId()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveFeedItemFromCollectionResponse{}, nil
}

func (s *LibrarianSephirahServiceService) ListFeedItemsInCollection(
	ctx context.Context,
	req *pb.ListFeedItemsInCollectionRequest,
) (*pb.ListFeedItemsInCollectionResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	items, total, err := s.y.ListFeedItemsInCollection(ctx,
		model.ToBizPaging(req.GetPaging()),
		converter.ToBizInternalIDList(req.GetCollectionIdFilter()),
		req.GetAuthorFilter(),
		req.GetPublishPlatformFilter(),
		req.GetCategoryFilter(),
		converter.ToBizTimeRange(req.GetPublishTimeRange()),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListFeedItemsInCollectionResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Items:  converter.ToPBFeedItemDigestList(items),
	}, nil
}
