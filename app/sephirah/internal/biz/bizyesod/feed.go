package bizyesod

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func (y *Yesod) ListFeeds(
	ctx context.Context, paging model.Paging, ids []model.InternalID,
	authorIDs []model.InternalID, sources []modelyesod.FeedConfigSource, statuses []modelyesod.FeedConfigStatus,
) ([]*modelyesod.FeedWithConfig, int, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	return y.repo.ListFeeds(ctx, claims.InternalID, paging, ids, authorIDs, sources, statuses)
}
func (y *Yesod) ListFeedItems(
	ctx context.Context,
	paging model.Paging,
	feedIDs []model.InternalID,
	authorIDs []model.InternalID,
	platforms []string,
) ([]*modelyesod.FeedItemIDWithFeedID, int, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	return y.repo.ListFeedItems(ctx, claims.InternalID, paging, feedIDs, authorIDs, platforms)
}

func (y *Yesod) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	items, err := y.repo.GetFeedItems(ctx, claims.InternalID, []model.InternalID{id})
	if err != nil {
		return nil, err
	}
	if len(items) != 1 {
		return nil, errors.New("item not found")
	}
	return items[0], nil
}

func (y *Yesod) GetFeedItems(ctx context.Context, ids []model.InternalID) ([]*modelfeed.Item, error) {
	claims, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	return y.repo.GetFeedItems(ctx, claims.InternalID, ids)
}
