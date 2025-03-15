package bizyesod

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizutils"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelyesod"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

func (y *Yesod) ListFeeds(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	statuses []modelyesod.FeedConfigStatus,
	categories []string,
) ([]*modelyesod.FeedWithConfig, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	feeds, i, err := y.repo.ListFeedConfigs(ctx, claims.UserID, paging, ids, statuses, categories)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return feeds, i, nil
}
func (y *Yesod) ListFeedItems(
	ctx context.Context,
	paging model.Paging,
	feedIDs []model.InternalID,
	authors []string,
	platforms []string,
	timeRange *model.TimeRange,
	categories []string,
) ([]*modelyesod.FeedItemDigest, int, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	items, i, err := y.repo.ListFeedItems(ctx,
		claims.UserID, paging, feedIDs, authors, platforms, timeRange, categories)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, i, nil
}

func (y *Yesod) GroupFeedItems(
	ctx context.Context,
	groupBy modelyesod.GroupFeedItemsBy,
	feedIDs []model.InternalID,
	authors []string,
	platforms []string,
	timeRange *model.TimeRange,
	groupSize int,
	categories []string,
) (map[model.TimeRange][]*modelyesod.FeedItemDigest, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
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
		default:
			return nil, pb.ErrorErrorReasonBadRequest("invalid group_by")
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
	items, err := y.repo.GroupFeedItems(ctx, claims.UserID, groups, feedIDs, authors,
		platforms, groupSize, categories)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, nil
}

func (y *Yesod) GetFeedItem(ctx context.Context, id model.InternalID) (*modelfeed.Item, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	items, err := y.repo.GetFeedItems(ctx, claims.UserID, []model.InternalID{id})
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(items) != 1 {
		return nil, pb.ErrorErrorReasonBadRequest("no such item")
	}
	return items[0], nil
}

func (y *Yesod) GetFeedItems(ctx context.Context, ids []model.InternalID) ([]*modelfeed.Item, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, bizutils.NoPermissionError()
	}
	items, err := y.repo.GetFeedItems(ctx, claims.UserID, ids)
	if err != nil {
		return nil, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return items, nil
}

func (y *Yesod) ReadFeedItem(ctx context.Context, id model.InternalID) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	if err := y.repo.ReadFeedItem(ctx, claims.UserID, id); err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return nil
}
