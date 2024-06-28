package modelyesod

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type FeedItemDigest struct {
	FeedID              model.InternalID
	ItemID              model.InternalID
	AvatarURL           string
	Authors             string
	PublishedParsedTime time.Time
	Title               string
	ShortDescription    string
	ImageUrls           []string
	PublishPlatform     string
	FeedConfigName      string
	FeedAvatarURL       string
	ReadCount           int64
}

type FeedWithConfig struct {
	FeedConfig *FeedConfig
	Feed       *modelfeed.Feed
}

type FeedConfig struct {
	ID                model.InternalID
	Name              string
	FeedURL           string
	Category          string
	AuthorAccount     model.InternalID
	Source            string
	Status            FeedConfigStatus
	PullInterval      time.Duration
	LatestPullTime    time.Time
	LatestPullStatus  FeedConfigPullStatus
	LatestPullMessage string
	HideItems         bool
	ActionSets        []model.InternalID
}

type FeedConfigStatus int

const (
	FeedConfigStatusUnspecified FeedConfigStatus = iota
	FeedConfigStatusActive
	FeedConfigStatusSuspend
)

type FeedConfigPullStatus int

const (
	FeedConfigPullStatusUnspecified FeedConfigPullStatus = iota
	FeedConfigPullStatusProcessing
	FeedConfigPullStatusSuccess
	FeedConfigPullStatusFailed
)

type ListFeedOrder int

const (
	ListFeedOrderUnspecified ListFeedOrder = iota
	ListFeedOrderNextPull
)

type PullFeed struct {
	InternalID   model.InternalID
	URL          string
	Source       string
	SystemNotify *modelnetzach.SystemNotify
}

type GroupFeedItemsBy int

const (
	GroupFeedItemsByUnspecified GroupFeedItemsBy = iota
	GroupFeedItemsByYear
	GroupFeedItemsByMonth
	GroupFeedItemsByDay
	GroupFeedItemsByOverall
)

type FeedItemCollection struct {
	ID          model.InternalID
	Name        string
	Description string
	Category    string
}

type FeedActionSet struct {
	ID          model.InternalID
	Name        string
	Description string
	Actions     []*modeltiphereth.FeatureRequest
}
