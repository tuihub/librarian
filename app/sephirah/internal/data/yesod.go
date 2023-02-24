package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
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

func (y *yesodRepo) CreateFeedConfig(ctx context.Context, c *bizyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Create().
		SetInternalID(c.InternalID).
		SetFeedURL(c.FeedURL).
		SetAuthorAccount(c.AuthorAccount).
		SetSource(converter.ToEntFeedConfigSource(c.Source)).
		SetStatus(converter.ToEntFeedConfigStatus(c.Status)).
		SetPullInterval(c.PullInterval)
	return q.Exec(ctx)
}

func (y *yesodRepo) UpdateFeedConfig(ctx context.Context, c *bizyesod.FeedConfig) error {
	q := y.data.db.FeedConfig.Update().
		Where(feedconfig.InternalIDEQ(c.InternalID)).
		SetFeedURL(c.FeedURL).
		SetAuthorAccount(c.AuthorAccount).
		SetSource(converter.ToEntFeedConfigSource(c.Source)).
		SetStatus(converter.ToEntFeedConfigStatus(c.Status)).
		SetPullInterval(c.PullInterval)
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
	time time.Time, i int) ([]*bizyesod.FeedConfig, error) {
	// TODO implement me
	panic("implement me")
}
