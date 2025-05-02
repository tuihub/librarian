package converter

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

import (
	"time"

	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/pkg/tuihub-rss/internal/converter
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend ToPBTime
// goverter:extend FromPBTime
type converter interface { //nolint:unused // used by generator
	// goverter:ignore Id
	ToPBFeed(t *gofeed.Feed) *librarian.Feed
	// goverter:ignore Id
	// goverter:ignore PublishPlatform
	// goverter:ignore ReadCount
	ToPBFeedItem(t *gofeed.Item) *librarian.FeedItem

	// goverter:ignore Id
	// goverter:ignore Link
	// goverter:ignore Author
	// goverter:ignore Updated
	// goverter:ignore Created
	// goverter:ignore Subtitle
	// goverter:ignore Copyright
	FromPBFeed(t *librarian.Feed) *feeds.Feed
	// goverter:ignore Id
	// goverter:ignore Link
	// goverter:ignore Source
	// goverter:ignore Author
	// goverter:ignore IsPermaLink
	// goverter:map UpdatedParsed Updated
	// goverter:map PublishedParsed Created
	// goverter:ignore Enclosure
	FromPBFeedItem(t *librarian.FeedItem) *feeds.Item
	FromPBFeedItems(t []*librarian.FeedItem) []*feeds.Item
	// goverter:ignore Link
	// goverter:ignore Width
	// goverter:ignore Height
	FromPBFeedImage(t *librarian.FeedImage) *feeds.Image
}

func ToPBTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func FromPBTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime()
}
