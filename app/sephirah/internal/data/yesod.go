package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent/dialect/sql"
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
		SetUserID(owner).
		SetID(c.ID).
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

func (y *yesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []modelyesod.FeedConfigSource,
	statuses []modelyesod.FeedConfigStatus, order modelyesod.ListFeedOrder,
	pullTime time.Time, i int) ([]*modelyesod.FeedConfig, error) {
	q := y.data.db.FeedConfig.Query()
	if len(sources) > 0 {
		q.Where(feedconfig.SourceIn(y.data.converter.ToEntFeedConfigSourceList(sources)...))
	}
	if len(statuses) > 0 {
		q.Where(feedconfig.StatusIn(y.data.converter.ToEntFeedConfigStatusList(statuses)...))
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
	return y.data.converter.ToBizFeedConfigList(feedConfigs), nil
}

func (y *yesodRepo) UpsertFeed(ctx context.Context, f *modelfeed.Feed) error {
	return y.data.WithTx(ctx, func(tx *ent.Tx) error {
		conf, err := tx.FeedConfig.Query().
			Where(feedconfig.IDEQ(f.ID)).
			Only(ctx)
		if err != nil {
			return err
		}
		err = tx.Feed.Create().
			SetConfig(conf).
			SetID(f.ID).
			SetTitle(f.Title).
			SetDescription(f.Description).
			SetLink(f.Link).
			SetAuthors(f.Authors).
			SetLanguage(f.Language).
			SetImage(f.Image).
			OnConflict(
				sql.ConflictColumns(feed.FieldID),
				sql.ResolveWithNewValues(),
			).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.FeedConfig.Update().
			Where(feedconfig.IDEQ(f.ID)).
			SetNextPullBeginAt(time.Now().Add(conf.PullInterval)).
			Exec(ctx)
		return err
	})
}

func (y *yesodRepo) UpsertFeedItems(ctx context.Context, items []*modelfeed.Item, feedID model.InternalID) error {
	il := make([]*ent.FeedItemCreate, len(items))
	for i, item := range items {
		il[i] = y.data.db.FeedItem.Create().
			SetFeedID(feedID).
			SetID(item.ID).
			SetTitle(item.Title).
			SetDescription(item.Description).
			SetContent(item.Content).
			SetLink(item.Link).
			SetUpdated(item.Updated).
			SetNillableUpdatedParsed(item.UpdatedParsed).
			SetPublished(item.Published).
			SetNillablePublishedParsed(item.PublishedParsed).
			SetAuthors(item.Authors).
			SetGUID(item.GUID).
			SetImage(item.Image).
			SetEnclosures(item.Enclosures).
			SetPublishPlatform(item.PublishPlatform)
	}
	return y.data.db.FeedItem.CreateBulk(il...).
		OnConflict(
			sql.ConflictColumns(feeditem.FieldFeedID, feeditem.FieldGUID),
			sql.ResolveWithNewValues(),
		).
		Exec(ctx)
}

func (y *yesodRepo) ListFeeds(
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
			q.Where(feedconfig.SourceIn(y.data.converter.ToEntFeedConfigSourceList(sources)...))
		}
		if len(statuses) > 0 {
			q.Where(feedconfig.StatusIn(y.data.converter.ToEntFeedConfigStatusList(statuses)...))
		}
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		configs, err := q.
			Limit(paging.PageSize).
			Offset((paging.PageNum - 1) * paging.PageSize).
			WithFeed().
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedWithConfig, 0, len(configs))
		for _, config := range configs {
			res = append(res, &modelyesod.FeedWithConfig{
				FeedConfig: y.data.converter.ToBizFeedConfig(config),
				Feed:       y.data.converter.ToBizFeed(config.Edges.Feed),
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
) ([]*modelyesod.FeedItemIDWithFeedID, int, error) {
	var res []*modelyesod.FeedItemIDWithFeedID
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
		total, err = iq.Count(ctx)
		if err != nil {
			return err
		}
		items, err := iq.
			Select(feeditem.FieldID, feeditem.FieldFeedID).
			Limit(paging.PageSize).
			Offset((paging.PageNum - 1) * paging.PageSize).
			All(ctx)
		if err != nil {
			return err
		}
		res = make([]*modelyesod.FeedItemIDWithFeedID, 0, len(items))
		for _, item := range items {
			res = append(res, &modelyesod.FeedItemIDWithFeedID{
				FeedID: item.ID,
				ItemID: item.FeedID,
			})
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
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
		res = y.data.converter.ToBizFeedItemList(items)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
