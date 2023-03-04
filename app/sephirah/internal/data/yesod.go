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

func (y *yesodRepo) CreateFeedConfig(ctx context.Context, c *bizyesod.FeedConfig, owner model.InternalID) error {
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

func (y *yesodRepo) UpdateFeedConfig(ctx context.Context, c *bizyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Update().
		Where(feedconfig.IDEQ(c.ID))
	if len(c.FeedURL) > 0 {
		q.SetFeedURL(c.FeedURL)
	}
	if c.AuthorAccount > 0 {
		q.SetAuthorAccount(c.AuthorAccount)
	}
	if c.Source != bizyesod.FeedConfigSourceUnspecified {
		q.SetSource(converter.ToEntFeedConfigSource(c.Source))
	}
	if c.Status != bizyesod.FeedConfigStatusUnspecified {
		q.SetStatus(converter.ToEntFeedConfigStatus(c.Status))
	}
	if c.PullInterval > 0 {
		q.SetPullInterval(c.PullInterval)
	}
	return q.Exec(ctx)
}

func (y *yesodRepo) ListFeedConfigNeedPull(ctx context.Context, sources []bizyesod.FeedConfigSource,
	statuses []bizyesod.FeedConfigStatus, order bizyesod.ListFeedOrder,
	pullTime time.Time, i int) ([]*bizyesod.FeedConfig, error) {
	q := y.data.db.FeedConfig.Query()
	if len(sources) > 0 {
		q.Where(feedconfig.SourceIn(y.data.converter.ToEntFeedConfigSourceList(sources)...))
	}
	if len(statuses) > 0 {
		q.Where(feedconfig.StatusIn(y.data.converter.ToEntFeedConfigStatusList(statuses)...))
	}
	switch order {
	case bizyesod.ListFeedOrderUnspecified:
		{
		}
	case bizyesod.ListFeedOrderNextPull:
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
			Where(feedconfig.IDEQ(f.InternalID)).
			Only(ctx)
		if err != nil {
			return err
		}
		err = tx.Feed.Create().
			SetConfig(conf).
			SetID(f.InternalID).
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
			Where(feedconfig.IDEQ(f.InternalID)).
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
			SetID(item.InternalID).
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
			SetEnclosure(item.Enclosures).
			SetPublishPlatform(item.PublishPlatform)
	}
	return y.data.db.FeedItem.CreateBulk(il...).
		OnConflict(
			sql.ConflictColumns(feeditem.FieldFeedID, feeditem.FieldGUID),
			sql.ResolveWithNewValues(),
		).
		Exec(ctx)
}

func (y *yesodRepo) ListFeeds(ctx context.Context, id model.InternalID, paging model.Paging,
	ids []model.InternalID, ids2 []model.InternalID, sources []bizyesod.FeedConfigSource,
	statuses []bizyesod.FeedConfigStatus) ([]*bizyesod.FeedWithConfig, error) {
	// TODO implement me
	panic("implement me")
}
