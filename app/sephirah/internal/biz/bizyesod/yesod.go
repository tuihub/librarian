package bizyesod

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

type YesodRepo interface {
	CreateFeedConfig(context.Context, *FeedConfig) error
	UpdateFeedConfig(context.Context, *FeedConfig) error
	ListFeedConfig(context.Context, []int64, []int64, []FeedConfigSource,
		[]FeedConfigStatus, ListFeedOrder, Paging) ([]*FeedConfig, error)
	ListFeedConfigNeedPull(context.Context, []FeedConfigSource, []FeedConfigStatus,
		ListFeedOrder, time.Time, int) ([]*FeedConfig, error)
}

type Yesod struct {
	repo     YesodRepo
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
	pullFeed *libmq.TopicImpl[PullFeed]
}

func NewYesod(
	repo YesodRepo,
	cron *libcron.Cron,
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
	pullFeed *libmq.TopicImpl[PullFeed],
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

type FeedConfig struct {
	InternalID    int64
	FeedURL       string
	AuthorAccount int64
	Source        FeedConfigSource
	Status        FeedConfigStatus
	PullInterval  time.Time
}

type FeedConfigSource int

const (
	FeedConfigSourceUnspecified FeedConfigSource = iota
	FeedConfigSourceCommon
)

type FeedConfigStatus int

const (
	FeedConfigStatusUnspecified FeedConfigStatus = iota
	FeedConfigStatusActive
	FeedConfigStatusSuspend
)

type ListFeedOrder int

const (
	ListFeedOrderUnspecified ListFeedOrder = iota
	ListFeedOrderNextPull
)

type Paging struct {
	PageSize int
	PageNum  int
}

type PullFeed struct {
	URL    string
	Source FeedConfigSource
}

func (y *Yesod) PullFeeds(ctx context.Context) {
	configs, err := y.repo.ListFeedConfigNeedPull(ctx, nil,
		[]FeedConfigStatus{FeedConfigStatusActive}, ListFeedOrderNextPull, time.Now(), 32) //nolint:gomnd // TODO
	if err != nil {
		logger.Errorf("%s", err.Error())
		return
	}
	for _, c := range configs {
		err = y.pullFeed.Publish(ctx, PullFeed{
			URL:    c.FeedURL,
			Source: c.Source,
		})
		if err != nil {
			logger.Errorf("%s", err.Error())
		}
	}
}
