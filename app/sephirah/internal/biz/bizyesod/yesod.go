package bizyesod

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewYesod,
	NewFeedOwnerCache,
)

type YesodRepo interface {
	CreateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	UpdateFeedConfig(context.Context, model.InternalID, *modelyesod.FeedConfig) error
	ListFeedCategories(context.Context, model.InternalID) ([]string, error)
	ListFeedPlatforms(context.Context, model.InternalID) ([]string, error)
	ListFeedConfigNeedPull(context.Context, []string, []modelyesod.FeedConfigStatus,
		modelyesod.ListFeedOrder, time.Time, int) ([]*modelyesod.FeedConfig, error)
	UpdateFeedConfigAsInQueue(context.Context, model.InternalID) error
	ListFeedConfigs(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]modelyesod.FeedConfigStatus, []string) ([]*modelyesod.FeedWithConfig, int, error)
	ListFeedItems(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]string, []string, *model.TimeRange, []string) ([]*modelyesod.FeedItemDigest, int, error)
	GroupFeedItems(context.Context, model.InternalID, []model.TimeRange, []model.InternalID,
		[]string, []string, int, []string) (
		map[model.TimeRange][]*modelyesod.FeedItemDigest, error)
	GetFeedItems(context.Context, model.InternalID, []model.InternalID) ([]*modelfeed.Item, error)
	ReadFeedItem(context.Context, model.InternalID, model.InternalID) error
	CreateFeedItemCollection(context.Context, model.InternalID, *modelyesod.FeedItemCollection) error
	UpdateFeedItemCollection(context.Context, model.InternalID, *modelyesod.FeedItemCollection) error
	ListFeedItemCollections(context.Context, model.InternalID, model.Paging, []model.InternalID,
		[]string) ([]*modelyesod.FeedItemCollection, int, error)
	AddFeedItemToCollection(context.Context, model.InternalID, model.InternalID, model.InternalID) error
	RemoveFeedItemFromCollection(context.Context, model.InternalID, model.InternalID, model.InternalID) error
	ListFeedItemsInCollection(context.Context, model.InternalID, model.Paging, []model.InternalID, []string,
		[]string, []string, *model.TimeRange) ([]*modelyesod.FeedItemDigest, int, error)
	GetFeedOwner(context.Context, model.InternalID) (*modeltiphereth.User, error)
	CreateFeedActionSet(context.Context, model.InternalID, *modelyesod.FeedActionSet) error
	UpdateFeedActionSet(context.Context, model.InternalID, *modelyesod.FeedActionSet) error
	ListFeedActionSets(context.Context, model.InternalID, model.Paging) ([]*modelyesod.FeedActionSet, int, error)
}

type Yesod struct {
	repo               YesodRepo
	supv               *supervisor.Supervisor
	id                 *libidgenerator.IDGenerator
	search             libsearch.Search
	pullFeed           *libmq.Topic[modelyesod.PullFeed]
	systemNotify       *libmq.Topic[modelnetzach.SystemNotify]
	feedOwner          *libcache.Map[modelyesod.FeedConfig, modeltiphereth.User]
	builtinFeedActions []*modelsupervisor.FeatureFlag
}

func NewYesod(
	repo YesodRepo,
	supv *supervisor.Supervisor,
	cron *libcron.Cron,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
	feedOwner *libcache.Map[modelyesod.FeedConfig, modeltiphereth.User],
) (*Yesod, error) {
	builtinFeedActions, err := getBuiltinActionFeatureFlags()
	if err != nil {
		return nil, err
	}
	y := &Yesod{
		repo:               repo,
		supv:               supv,
		id:                 id,
		search:             search,
		pullFeed:           pullFeed,
		systemNotify:       systemNotify,
		feedOwner:          feedOwner,
		builtinFeedActions: builtinFeedActions,
	}
	err = cron.BySeconds("YesodPullFeeds", 60, y.PullFeeds, context.Background()) //nolint:mnd // hard code min interval
	if err != nil {
		return nil, err
	}
	return y, nil
}

func (y *Yesod) GetBuiltInFeedActions() []*modelsupervisor.FeatureFlag {
	return y.builtinFeedActions
}

func (y *Yesod) PullFeeds(ctx context.Context) error {
	configs, err := y.repo.ListFeedConfigNeedPull(ctx, nil,
		[]modelyesod.FeedConfigStatus{modelyesod.FeedConfigStatusActive},
		modelyesod.ListFeedOrderNextPull, time.Now(), 32) //nolint:mnd // TODO
	if err != nil {
		logger.Errorf("%s", err.Error())
		return err
	}
	var errRes error
	for _, c := range configs {
		doNotify := func() *modelnetzach.SystemNotify {
			var owner *modeltiphereth.User
			owner, err = y.feedOwner.Get(ctx, *c)
			if err != nil {
				return nil
			}
			un := modelnetzach.NewUserNotify(
				owner.ID,
				modelnetzach.SystemNotificationLevelOngoing,
				fmt.Sprintf("%s: Update Feed %s", modelnetzach.SystemNotifyTitleCronJob, c.Name),
				"Queued",
			)
			un.Notification.ID, err = y.id.New()
			if err != nil {
				return nil
			}
			err = y.systemNotify.PublishFallsLocalCall(ctx, un)
			if err != nil {
				return nil
			}
			return &un
		}
		err = y.pullFeed.Publish(ctx, modelyesod.PullFeed{
			InternalID:   c.ID,
			Source:       c.Source,
			SystemNotify: doNotify(),
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

func NewFeedOwnerCache(
	repo YesodRepo,
	store libcache.Store,
) *libcache.Map[modelyesod.FeedConfig, modeltiphereth.User] {
	return libcache.NewMap[modelyesod.FeedConfig, modeltiphereth.User](
		store,
		"FeedOwner",
		func(k modelyesod.FeedConfig) string {
			return strconv.FormatInt(int64(k.ID), 10)
		},
		func(ctx context.Context, fc modelyesod.FeedConfig) (*modeltiphereth.User, error) {
			res, err := repo.GetFeedOwner(ctx, fc.ID)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}
