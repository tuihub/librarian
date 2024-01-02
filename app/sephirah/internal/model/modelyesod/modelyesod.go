package modelyesod

import (
	"time"

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
	ID               model.InternalID
	Name             string
	FeedURL          string
	Category         string
	AuthorAccount    model.InternalID
	Source           FeedConfigSource
	Status           FeedConfigStatus
	PullInterval     time.Duration
	LatestUpdateTime time.Time
	HideItems        bool
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

type PullFeed struct {
	InternalID model.InternalID
	URL        string
	Source     FeedConfigSource
}

type GroupFeedItemsBy int

const (
	GroupFeedItemsByUnspecified GroupFeedItemsBy = iota
	GroupFeedItemsByYear
	GroupFeedItemsByMonth
	GroupFeedItemsByDay
	GroupFeedItemsByOverall
)
