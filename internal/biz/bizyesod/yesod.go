package bizyesod

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model/modelyesod"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewYesod,
	NewFeedOwnerCache,
)

type Yesod struct {
	repo               *data.YesodRepo
	supv               *data.SupervisorRepo
	id                 *libidgenerator.IDGenerator
	search             libsearch.Search
	pullFeed           *libmq.Topic[modelyesod.PullFeed]
	systemNotify       *libmq.Topic[modelnetzach.SystemNotify]
	feedOwner          *libcache.Map[modelyesod.FeedConfig, model.User]
	builtinFeedActions []*modelsupervisor.FeatureFlag
}

func NewYesod(
	repo *data.YesodRepo,
	supv *data.SupervisorRepo,
	cron *libcron.Cron,
	id *libidgenerator.IDGenerator,
	search libsearch.Search,
	pullFeed *libmq.Topic[modelyesod.PullFeed],
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
	feedOwner *libcache.Map[modelyesod.FeedConfig, model.User],
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
	_, err = cron.NewJobBySeconds(
		"YesodPullFeeds",
		60, //nolint:mnd // hard code min interval
		y.PullFeeds,
		context.Background(),
	)
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
			var owner *model.User
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
	repo *data.YesodRepo,
	store libcache.Store,
) *libcache.Map[modelyesod.FeedConfig, model.User] {
	return libcache.NewMap[modelyesod.FeedConfig, model.User](
		store,
		"FeedOwner",
		func(k modelyesod.FeedConfig) string {
			return strconv.FormatInt(int64(k.ID), 10)
		},
		func(ctx context.Context, fc modelyesod.FeedConfig) (*model.User, error) {
			res, err := repo.GetFeedOwner(ctx, fc.ID)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}
