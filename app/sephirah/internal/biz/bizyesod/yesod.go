package bizyesod

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type YesodRepo interface {
	CreateFeedConfig(context.Context, *modelyesod.FeedConfig, model.InternalID) error
	UpdateFeedConfig(context.Context, *modelyesod.FeedConfig) error
	ListFeedConfigNeedPull(context.Context, []modelyesod.FeedConfigSource, []modelyesod.FeedConfigStatus,
		modelyesod.ListFeedOrder, time.Time, int) ([]*modelyesod.FeedConfig, error)
	UpsertFeed(context.Context, *modelfeed.Feed) error
	UpsertFeedItems(context.Context, []*modelfeed.Item, model.InternalID) error
	ListFeeds(context.Context, model.InternalID, model.Paging, []model.InternalID, []model.InternalID,
		[]modelyesod.FeedConfigSource, []modelyesod.FeedConfigStatus) ([]*modelyesod.FeedWithConfig, int, error)
	ListFeedItems(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]model.InternalID, []string) ([]*modelyesod.FeedItemIDWithFeedID, int, error)
	GetFeedItems(context.Context, model.InternalID, []model.InternalID) ([]*modelfeed.Item, error)
}

type Yesod struct {
	repo     YesodRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
	pullFeed *libmq.TopicImpl[modelyesod.PullFeed]
}

func NewYesod(
	repo YesodRepo,
	cron *libcron.Cron,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	pullFeed *libmq.TopicImpl[modelyesod.PullFeed],
) (*Yesod, error) {
	y := &Yesod{
		repo:     repo,
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
		pullFeed: pullFeed,
	}
	err := cron.BySeconds(60, y.PullFeeds, context.Background()) //nolint:gomnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return y, nil
}

func (y *Yesod) PullFeeds(ctx context.Context) {
	configs, err := y.repo.ListFeedConfigNeedPull(ctx, nil,
		[]modelyesod.FeedConfigStatus{modelyesod.FeedConfigStatusActive},
		modelyesod.ListFeedOrderNextPull, time.Now(), 32) //nolint:gomnd // TODO
	if err != nil {
		logger.Errorf("%s", err.Error())
		return
	}
	for _, c := range configs {
		err = y.pullFeed.Publish(ctx, modelyesod.PullFeed{
			InternalID: c.ID,
			URL:        c.FeedURL,
			Source:     c.Source,
		})
		if err != nil {
			logger.Errorf("%s", err.Error())
		}
	}
}
