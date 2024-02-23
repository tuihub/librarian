package bizyesod

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type YesodRepo interface {
	CreateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	UpdateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	ListFeedCategories(context.Context, model.InternalID) ([]string, error)
	ListFeedPlatforms(context.Context, model.InternalID) ([]string, error)
	ListFeedConfigNeedPull(context.Context, []string, []modelyesod.FeedConfigStatus,
		modelyesod.ListFeedOrder, time.Time, int) ([]*modelyesod.FeedConfig, error)
	UpdateFeedConfigAsInQueue(context.Context, model.InternalID) error
	ListFeedConfigs(context.Context, model.InternalID, model.Paging, []model.InternalID, []model.InternalID,
		[]string, []modelyesod.FeedConfigStatus, []string) ([]*modelyesod.FeedWithConfig, int, error)
	ListFeedItems(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]model.InternalID, []string, *model.TimeRange, []string) ([]*modelyesod.FeedItemDigest, int, error)
	GroupFeedItems(context.Context, model.InternalID, []model.TimeRange, []model.InternalID,
		[]model.InternalID, []string, int, []string) (
		map[model.TimeRange][]*modelyesod.FeedItemDigest, error)
	GetFeedItems(context.Context, model.InternalID, []model.InternalID) ([]*modelfeed.Item, error)
	ReadFeedItem(context.Context, model.InternalID, model.InternalID) error
}

type Yesod struct {
	repo YesodRepo
	supv *supervisor.Supervisor
	// mapper   mapper.LibrarianMapperServiceClient
	searcher *client.Searcher
	pullFeed *libmq.Topic[modelyesod.PullFeed]
}

func NewYesod(
	repo YesodRepo,
	supv *supervisor.Supervisor,
	cron *libcron.Cron,
	// mClient mapper.LibrarianMapperServiceClient,
	sClient *client.Searcher,
	pullFeed *libmq.Topic[modelyesod.PullFeed],
) (*Yesod, error) {
	y := &Yesod{
		repo: repo,
		supv: supv,
		//mapper:   mClient,
		searcher: sClient,
		pullFeed: pullFeed,
	}
	err := cron.BySeconds("YesodPullFeeds", 60, y.PullFeeds, context.Background()) //nolint:gomnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return y, nil
}

func (y *Yesod) PullFeeds(ctx context.Context) error {
	configs, err := y.repo.ListFeedConfigNeedPull(ctx, nil,
		[]modelyesod.FeedConfigStatus{modelyesod.FeedConfigStatusActive},
		modelyesod.ListFeedOrderNextPull, time.Now(), 32) //nolint:gomnd // TODO
	if err != nil {
		logger.Errorf("%s", err.Error())
		return err
	}
	var errRes error
	for _, c := range configs {
		err = y.pullFeed.Publish(ctx, modelyesod.PullFeed{
			InternalID: c.ID,
			URL:        c.FeedURL,
			Source:     c.Source,
		})
		if err != nil {
			logger.Errorf("%s", err.Error())
			errRes = err
			continue
		}
		err = y.repo.UpdateFeedConfigAsInQueue(ctx, c.ID)
		if err != nil {
			logger.Errorf("%s", err.Error())
			errRes = err
		}
	}
	return errRes
}
