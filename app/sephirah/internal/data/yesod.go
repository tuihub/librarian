package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/exp/slices"
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
		SetUserID(int64(owner)).
		SetID(c.InternalID).
		SetFeedURL(c.FeedURL).
		SetAuthorAccount(c.AuthorAccount).
		SetSource(converter.ToEntFeedConfigSource(c.Source)).
		SetStatus(converter.ToEntFeedConfigStatus(c.Status)).
		SetPullInterval(c.PullInterval)
	return q.Exec(ctx)
}

func (y *yesodRepo) UpdateFeedConfig(ctx context.Context, c *bizyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Update().
		Where(feedconfig.IDEQ(c.InternalID))
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

func (y *yesodRepo) ListFeedConfig(ctx context.Context, int64s []int64, int64s2 []int64,
	sources []bizyesod.FeedConfigSource, statuses []bizyesod.FeedConfigStatus,
	order bizyesod.ListFeedOrder, paging bizyesod.Paging) ([]*bizyesod.FeedConfig, error) {
	// TODO implement me
	panic("implement me")
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
				sql.ResolveWith(func(u *sql.UpdateSet) {
					ignores := []string{
						feed.FieldID,
					}
					for _, c := range u.Columns() {
						if slices.Contains(ignores, c) {
							u.SetIgnore(c)
						}
						u.SetExcluded(c)
					}
				}),
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
