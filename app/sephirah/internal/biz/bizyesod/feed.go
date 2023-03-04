package bizyesod

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) ListFeeds(
	ctx context.Context, paging model.Paging, ids []model.InternalID,
	authorIDs []model.InternalID, sources []modelyesod.FeedConfigSource, statuses []modelyesod.FeedConfigStatus,
) ([]*modelyesod.FeedWithConfig, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	feeds, i, err := y.repo.ListFeeds(ctx, c.InternalID, paging, ids, authorIDs, sources, statuses)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return feeds, i, nil
}
func (y *Yesod) ListFeedItems(
	ctx context.Context,
	paging model.Paging,
	feedIDs []model.InternalID,
	authorIDs []model.InternalID,
	platforms []string,
) ([]*modelyesod.FeedItemIDWithFeedID, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	items, i, err := y.repo.ListFeedItems(ctx, c.InternalID, paging, feedIDs, authorIDs, platforms)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, i, nil
}

func (y *Yesod) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	items, err := y.repo.GetFeedItems(ctx, c.InternalID, []model.InternalID{id})
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(items) != 1 {
		return nil, pb.ErrorErrorReasonBadRequest("no such item")
	}
	return items[0], nil
}

func (y *Yesod) GetFeedItems(ctx context.Context, ids []model.InternalID) ([]*modelfeed.Item, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	items, err := y.repo.GetFeedItems(ctx, c.InternalID, ids)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, nil
}
