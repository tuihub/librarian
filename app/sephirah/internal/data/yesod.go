package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type yesodRepo struct {
	data *Data
}

// NewYesodRepo .
func NewYesodRepo(data *Data) bizyesod.YesodRepo {
	return &yesodRepo{
		data: data,
	}
}

func (y *yesodRepo) CreateFeedConfig(ctx context.Context, c *modelyesod.FeedConfig, owner model.InternalID) error {
	q := y.data.db.FeedConfig.Create().
		SetOwnerID(owner).
		SetID(c.ID).
		SetName(c.Name).
		SetFeedURL(c.FeedURL).
		SetAuthorAccount(c.AuthorAccount).
		SetSource(converter.ToEntFeedConfigSource(c.Source)).
		SetStatus(converter.ToEntFeedConfigStatus(c.Status)).
		SetPullInterval(c.PullInterval)
	return q.Exec(ctx)
}

func (y *yesodRepo) UpdateFeedConfig(ctx context.Context, c *modelyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Update().
		Where(feedconfig.IDEQ(c.ID))
	if len(c.Name) > 0 {
		q.SetName(c.Name)
	}
	if len(c.FeedURL) > 0 {
		q.SetFeedURL(c.FeedURL)
	}
	if c.AuthorAccount > 0 {
		q.SetAuthorAccount(c.AuthorAccount)
	}
	if c.Source != modelyesod.FeedConfigSourceUnspecified {
		q.SetSource(converter.ToEntFeedConfigSource(c.Source))
	}
	if c.Status != modelyesod.FeedConfigStatusUnspecified {
		q.SetStatus(converter.ToEntFeedConfigStatus(c.Status))
	}
	if c.PullInterval > 0 {
		q.SetPullInterval(c.PullInterval)
	}
	return q.Exec(ctx)
}

// UpdateFeedConfigAsInQueue set SetNextPullBeginAt to one day later to avoid repeat queue.
// While pull success, UpsertFeed will set correct value.
// While pull failed, server will retry task next day.
func (y *yesodRepo) UpdateFeedConfigAsInQueue(ctx context.Context, id model.InternalID) error {
	q := y.data.db.FeedConfig.UpdateOneID(id).
		SetNextPullBeginAt(time.Now().Add(libtime.Day))
	return q.Exec(ctx)
}

func (y *yesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []modelyesod.FeedConfigSource,
	statuses []modelyesod.FeedConfigStatus, order modelyesod.ListFeedOrder,
	pullTime time.Time, i int) ([]*modelyesod.FeedConfig, error) {
	q := y.data.db.FeedConfig.Query()
	if len(sources) > 0 {
		q.Where(feedconfig.SourceIn(converter.ToEntFeedConfigSourceList(sources)...))
	}
	if len(statuses) > 0 {
		q.Where(feedconfig.StatusIn(converter.ToEntFeedConfigStatusList(statuses)...))
	}
	switch order {
	case modelyesod.ListFeedOrderUnspecified:
		{
		}
	case modelyesod.ListFeedOrderNextPull:
		q.Where(feedconfig.NextPullBeginAtLT(pullTime))
	}
	q.Limit(i)
	feedConfigs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizFeedConfigList(feedConfigs), nil
}

func (y *yesodRepo) ListFeedConfigs(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	authorIDs []model.InternalID,
	sources []modelyesod.FeedConfigSource,
	statuses []modelyesod.FeedConfigStatus,
) ([]*modelyesod.FeedWithConfig, int, error) {
	var res []*modelyesod.FeedWithConfig
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		user, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		q := tx.User.QueryFeedConfig(user)
		if len(ids) > 0 {
			q.Where(feedconfig.IDIn(ids...))
		}
		if len(authorIDs) > 0 {
			q.Where(feedconfig.AuthorAccountIn(authorIDs...))
		}
		if len(sources) > 0 {
			q.Where(feedconfig.SourceIn(converter.ToEntFeedConfigSourceList(sources)...))
		}
		if len(statuses) > 0 {
			q.Where(feedconfig.StatusIn(converter.ToEntFeedConfigStatusList(statuses)...))
		}
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		configs, err := q.
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			WithFeed().
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedWithConfig, 0, len(configs))
		for _, config := range configs {
			res = append(res, &modelyesod.FeedWithConfig{
				FeedConfig: converter.ToBizFeedConfig(config),
				Feed:       converter.ToBizFeed(config.Edges.Feed),
			})
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) ListFeedItems(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	feedIDs []model.InternalID,
	authorIDs []model.InternalID,
	platforms []string,
	timeRange *model.TimeRange,
) ([]*modelyesod.FeedItemDigest, int, error) {
	var res []*modelyesod.FeedItemDigest
	var total int
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		user, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		fq := tx.User.QueryFeedConfig(user).QueryFeed()
		if len(feedIDs) > 0 {
			fq.Where(feed.IDIn(feedIDs...))
		}
		iq := fq.QueryItem()
		if len(platforms) > 0 {
			iq.Where(feeditem.PublishPlatformIn(platforms...))
		}
		if timeRange != nil {
			iq.
				Where(feeditem.PublishedParsedGTE(timeRange.StartTime)).
				Where(feeditem.PublishedParsedLT(timeRange.StartTime.Add(timeRange.Duration)))
		}
		total, err = iq.Count(ctx)
		if err != nil {
			return err
		}
		items, err := iq.
			Order(ent.Desc(feeditem.FieldPublishedParsed)).
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedItemDigest, 0, len(items))
		for _, item := range items {
			res = append(res, converter.ToBizFeedItemDigest(item))
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func (y *yesodRepo) GroupFeedItems(
	ctx context.Context,
	userID model.InternalID,
	groups []model.TimeRange,
	feedIDs []model.InternalID,
	authorIDs []model.InternalID,
	platforms []string,
	groupSize int,
) (map[model.TimeRange][]*modelyesod.FeedItemDigest, error) {
	res := make(map[model.TimeRange][]*modelyesod.FeedItemDigest)
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		user, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		for _, timeRange := range groups {
			fq := tx.User.QueryFeedConfig(user).QueryFeed()
			if len(feedIDs) > 0 {
				fq.Where(feed.IDIn(feedIDs...))
			}
			iq := fq.QueryItem()
			if len(platforms) > 0 {
				iq.Where(feeditem.PublishPlatformIn(platforms...))
			}
			var items []*ent.FeedItem
			items, err = iq.
				Where(feeditem.PublishedParsedGTE(timeRange.StartTime)).
				Where(feeditem.PublishedParsedLT(timeRange.StartTime.Add(timeRange.Duration))).
				Order(ent.Desc(feeditem.FieldPublishedParsed)).
				Limit(groupSize).
				All(ctx)
			if err != nil {
				return err
			}
			if len(items) == 0 {
				continue
			}
			il := make([]*modelyesod.FeedItemDigest, 0, len(items))
			for _, item := range items {
				il = append(il, converter.ToBizFeedItemDigest(item))
			}
			res[timeRange] = il
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (y *yesodRepo) GetFeedItems(
	ctx context.Context,
	userID model.InternalID,
	ids []model.InternalID,
) ([]*modelfeed.Item, error) {
	var res []*modelfeed.Item
	err := y.data.WithTx(ctx, func(tx *ent.Tx) error {
		user, err := tx.User.Get(ctx, userID)
		if err != nil {
			return err
		}
		items, err := tx.User.
			QueryFeedConfig(user).
			QueryFeed().
			QueryItem().
			Where(feeditem.IDIn(ids...)).
			All(ctx)
		if err != nil {
			return err
		}
		res = converter.ToBizFeedItemList(items)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
