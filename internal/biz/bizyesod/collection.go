package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	pb "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) CreateFeedItemCollection(
	ctx context.Context,
	collection *modelyesod.FeedItemCollection,
) (model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	id, err := y.id.New()
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	collection.ID = id
	err = y.repo.CreateFeedItemCollection(ctx, claims.UserID, collection)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return collection.ID, nil
}

func (y *Yesod) UpdateFeedItemCollection(
	ctx context.Context,
	collection *modelyesod.FeedItemCollection,
) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := y.repo.UpdateFeedItemCollection(ctx, claims.UserID, collection)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) ListFeedItemCollections(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	categories []string,
) ([]*modelyesod.FeedItemCollection, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	collections, i, err := y.repo.ListFeedItemCollections(ctx, claims.UserID, paging, ids, categories)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return collections, i, nil
}

func (y *Yesod) AddFeedItemToCollection(ctx context.Context, collectionID model.InternalID,
	itemID model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := y.repo.AddFeedItemToCollection(ctx, claims.UserID, collectionID, itemID)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) RemoveFeedItemFromCollection(ctx context.Context, collectionID model.InternalID,
	itemID model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := y.repo.RemoveFeedItemFromCollection(ctx, claims.UserID, collectionID, itemID)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}

func (y *Yesod) ListFeedItemsInCollection(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	authors []string,
	platforms []string,
	categories []string,
	timeRange *model.TimeRange,
) ([]*modelyesod.FeedItemDigest, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	items, i, err := y.repo.ListFeedItemsInCollection(ctx, claims.UserID, paging, ids, authors,
		platforms, categories, timeRange)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, i, nil
}
