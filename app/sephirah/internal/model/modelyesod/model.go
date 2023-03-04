package modelyesod

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type FeedItemIDWithFeedID struct {
	FeedID model.InternalID
	ItemID model.InternalID
}

type FeedWithConfig struct {
	FeedConfig *FeedConfig
	Feed       *modelfeed.Feed
}

type FeedConfig struct {
	ID             model.InternalID
	FeedURL        string
	AuthorAccount  model.InternalID
	Source         FeedConfigSource
	Status         FeedConfigStatus
	PullInterval   time.Duration
	LatestPullTime time.Time
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
