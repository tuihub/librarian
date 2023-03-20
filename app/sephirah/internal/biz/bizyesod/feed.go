package bizyesod

import (
	"context"
	"time"

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
	feeds, i, err := y.repo.ListFeedConfigs(ctx, c.InternalID, paging, ids, authorIDs, sources, statuses)
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
	timeRange *model.TimeRange,
) ([]*modelyesod.FeedItemIDWithFeedID, int, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, 0, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, 0, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	items, i, err := y.repo.ListFeedItems(ctx, c.InternalID, paging, feedIDs, authorIDs, platforms, timeRange)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, i, nil
}

func (y *Yesod) GroupFeedItems(
	ctx context.Context,
	groupBy modelyesod.GroupFeedItemsBy,
	feedIDs []model.InternalID,
	authorIDs []model.InternalID,
	platforms []string,
	timeRange *model.TimeRange,
	groupSize int,
) (map[model.TimeRange][]*modelyesod.FeedItemIDWithFeedID, *errors.Error) {
	if !libauth.FromContextAssertUserType(ctx, libauth.UserTypeAdmin, libauth.UserTypeNormal) {
		return nil, pb.ErrorErrorReasonForbidden("no permission")
	}
	c, exist := libauth.FromContext(ctx)
	if !exist {
		return nil, pb.ErrorErrorReasonUnauthorized("empty token")
	}
	// set default value
	if timeRange == nil {
		timeRange = &model.TimeRange{
			StartTime: time.Now(),
			Duration:  -time.Since(time.UnixMilli(0)),
		}
	}
	if timeRange.Duration > 0 {
		timeRange = &model.TimeRange{
			StartTime: timeRange.StartTime.Add(timeRange.Duration),
			Duration:  -timeRange.Duration,
		}
	}
	var groups []model.TimeRange
	var currentTime time.Time
	{
		year, month, day := timeRange.StartTime.Date()
		switch groupBy {
		case modelyesod.GroupFeedItemsByUnspecified:
			return nil, pb.ErrorErrorReasonBadRequest("invalid group_by")
		case modelyesod.GroupFeedItemsByYear:
			currentTime = time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		case modelyesod.GroupFeedItemsByMonth:
			currentTime = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		case modelyesod.GroupFeedItemsByDay:
			currentTime = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		default:
			return nil, pb.ErrorErrorReasonBadRequest("invalid group_by")
		}
	}
	for i := 0; i < 100 && currentTime.After(timeRange.StartTime.Add(timeRange.Duration)); i++ {
		var nextTime time.Time
		switch groupBy {
		case modelyesod.GroupFeedItemsByUnspecified:
			return nil, pb.ErrorErrorReasonBadRequest("invalid group_by")
		case modelyesod.GroupFeedItemsByYear:
			nextTime = currentTime.AddDate(-1, 0, 0)
		case modelyesod.GroupFeedItemsByMonth:
			nextTime = currentTime.AddDate(0, -1, 0)
		case modelyesod.GroupFeedItemsByDay:
			nextTime = currentTime.AddDate(0, 0, -1)
		}
		groups = append(groups, model.TimeRange{
			StartTime: currentTime,
			Duration:  currentTime.Sub(nextTime),
		})
		currentTime = nextTime
	}
	if groupSize <= 0 || groupSize > 100 {
		groupSize = 100
	}
	items, err := y.repo.GroupFeedItems(ctx, c.InternalID, groups, feedIDs, authorIDs,
		platforms, groupSize)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, nil
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
