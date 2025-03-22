package sephirah

import (
	"context"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (s *LibrarianSephirahService) CreateFeedConfig(
	ctx context.Context,
	req *sephirah.CreateFeedConfigRequest,
) (*sephirah.CreateFeedConfigResponse, error) {
	id, err := s.y.CreateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateFeedConfigResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahService) UpdateFeedConfig(
	ctx context.Context,
	req *sephirah.UpdateFeedConfigRequest,
) (*sephirah.UpdateFeedConfigResponse, error) {
	err := s.y.UpdateFeedConfig(ctx, converter.ToBizFeedConfig(req.GetConfig()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateFeedConfigResponse{}, nil
}
func (s *LibrarianSephirahService) ListFeedConfigs(
	ctx context.Context,
	req *sephirah.ListFeedConfigsRequest,
) (*sephirah.ListFeedConfigsResponse, error) {
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
	return &sephirah.ListFeedConfigsResponse{
		Paging:          &librarian.PagingResponse{TotalSize: int64(total)},
		FeedsWithConfig: converter.ToPBFeedWithConfigList(feeds),
	}, nil
}
func (s *LibrarianSephirahService) CreateFeedActionSet(
	ctx context.Context,
	req *sephirah.CreateFeedActionSetRequest,
) (*sephirah.CreateFeedActionSetResponse, error) {
	id, err := s.y.CreateFeedActionSet(ctx, converter.ToBizFeedActionSet(req.GetActionSet()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateFeedActionSetResponse{
		Id: converter.ToPBInternalID(id),
	}, nil
}
func (s *LibrarianSephirahService) UpdateFeedActionSet(
	ctx context.Context,
	req *sephirah.UpdateFeedActionSetRequest,
) (*sephirah.UpdateFeedActionSetResponse, error) {
	err := s.y.UpdateFeedActionSet(ctx, converter.ToBizFeedActionSet(req.GetActionSet()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateFeedActionSetResponse{}, nil
}
func (s *LibrarianSephirahService) ListFeedActionSets(
	ctx context.Context,
	req *sephirah.ListFeedActionSetsRequest,
) (*sephirah.ListFeedActionSetsResponse, error) {
	if req.GetPaging() == nil {
		return nil, pb.ErrorErrorReasonBadRequest("")
	}
	actions, total, err := s.y.ListFeedActionSets(ctx,
		model.ToBizPaging(req.GetPaging()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListFeedActionSetsResponse{
		Paging:     &librarian.PagingResponse{TotalSize: int64(total)},
		ActionSets: converter.ToPBFeedActionSetList(actions),
	}, nil
}
func (s *LibrarianSephirahService) ListFeedCategories(
	ctx context.Context,
	req *sephirah.ListFeedCategoriesRequest,
) (*sephirah.ListFeedCategoriesResponse, error) {
	res, err := s.y.ListFeedCategories(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListFeedCategoriesResponse{Categories: res}, nil
}
func (s *LibrarianSephirahService) ListFeedPlatforms(
	ctx context.Context,
	req *sephirah.ListFeedPlatformsRequest,
) (*sephirah.ListFeedPlatformsResponse, error) {
	res, err := s.y.ListFeedPlatforms(ctx)
	if err != nil {
		return nil, err
	}
	return &sephirah.ListFeedPlatformsResponse{Platforms: res}, nil
}

func (s *LibrarianSephirahService) ListFeedItems(
	ctx context.Context,
	req *sephirah.ListFeedItemsRequest,
) (*sephirah.ListFeedItemsResponse, error) {
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
	return &sephirah.ListFeedItemsResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Items:  converter.ToPBFeedItemDigestList(items),
	}, nil
}

func (s *LibrarianSephirahService) GetFeedItem(
	ctx context.Context,
	req *sephirah.GetFeedItemRequest,
) (*sephirah.GetFeedItemResponse, error) {
	item, err := s.y.GetFeedItem(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.GetFeedItemResponse{
		Item: converter.ToPBFeedItem(item),
	}, nil
}

func (s *LibrarianSephirahService) GetBatchFeedItems(
	ctx context.Context,
	req *sephirah.GetBatchFeedItemsRequest,
) (*sephirah.GetBatchFeedItemsResponse, error) {
	items, err := s.y.GetFeedItems(ctx, converter.ToBizInternalIDList(req.GetIds()))
	if err != nil {
		return nil, err
	}
	return &sephirah.GetBatchFeedItemsResponse{
		Items: converter.ToPBFeedItemList(items),
	}, nil
}

func (s *LibrarianSephirahService) ReadFeedItem(
	ctx context.Context,
	req *sephirah.ReadFeedItemRequest,
) (*sephirah.ReadFeedItemResponse, error) {
	err := s.y.ReadFeedItem(ctx, converter.ToBizInternalID(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &sephirah.ReadFeedItemResponse{}, nil
}

func (s *LibrarianSephirahService) CreateFeedItemCollection(
	ctx context.Context,
	req *sephirah.CreateFeedItemCollectionRequest,
) (*sephirah.CreateFeedItemCollectionResponse, error) {
	_, err := s.y.CreateFeedItemCollection(ctx, converter.ToBizFeedItemCollection(req.GetCollection()))
	if err != nil {
		return nil, err
	}
	return &sephirah.CreateFeedItemCollectionResponse{}, nil
}

func (s *LibrarianSephirahService) UpdateFeedItemCollection(
	ctx context.Context,
	req *sephirah.UpdateFeedItemCollectionRequest,
) (*sephirah.UpdateFeedItemCollectionResponse, error) {
	err := s.y.UpdateFeedItemCollection(ctx, converter.ToBizFeedItemCollection(req.GetCollection()))
	if err != nil {
		return nil, err
	}
	return &sephirah.UpdateFeedItemCollectionResponse{}, nil
}

func (s *LibrarianSephirahService) ListFeedItemCollections(
	ctx context.Context,
	req *sephirah.ListFeedItemCollectionsRequest,
) (*sephirah.ListFeedItemCollectionsResponse, error) {
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
	return &sephirah.ListFeedItemCollectionsResponse{
		Paging:      &librarian.PagingResponse{TotalSize: int64(total)},
		Collections: converter.ToPBFeedItemCollectionList(collections),
	}, nil
}

func (s *LibrarianSephirahService) AddFeedItemToCollection(
	ctx context.Context,
	req *sephirah.AddFeedItemToCollectionRequest,
) (*sephirah.AddFeedItemToCollectionResponse, error) {
	err := s.y.AddFeedItemToCollection(ctx,
		converter.ToBizInternalID(req.GetCollectionId()),
		converter.ToBizInternalID(req.GetFeedItemId()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.AddFeedItemToCollectionResponse{}, nil
}

func (s *LibrarianSephirahService) RemoveFeedItemFromCollection(
	ctx context.Context,
	req *sephirah.RemoveFeedItemFromCollectionRequest,
) (*sephirah.RemoveFeedItemFromCollectionResponse, error) {
	err := s.y.RemoveFeedItemFromCollection(
		ctx,
		converter.ToBizInternalID(req.GetCollectionId()),
		converter.ToBizInternalID(req.GetFeedItemId()),
	)
	if err != nil {
		return nil, err
	}
	return &sephirah.RemoveFeedItemFromCollectionResponse{}, nil
}

func (s *LibrarianSephirahService) ListFeedItemsInCollection(
	ctx context.Context,
	req *sephirah.ListFeedItemsInCollectionRequest,
) (*sephirah.ListFeedItemsInCollectionResponse, error) {
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
	return &sephirah.ListFeedItemsInCollectionResponse{
		Paging: &librarian.PagingResponse{TotalSize: int64(total)},
		Items:  converter.ToPBFeedItemDigestList(items),
	}, nil
}
