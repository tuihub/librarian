package converter

import (
	"time"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"

	"github.com/mmcdole/gofeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --packagePath github.com/tuihub/librarian/app/porter/internal/client/feed/converter --packageName converter --output ./generated.go ./

// goverter:converter
type Converter interface {
	// goverter:matchIgnoreCase
	ToPBFeed(t *gofeed.Feed) *bizfeed.Feed
	// goverter:matchIgnoreCase
	// goverter:map UpdatedParsed | TimeToTime
	// goverter:map PublishedParsed | TimeToTime
	ToPBFeedItem(t *gofeed.Item) *bizfeed.Item
}

func NewConverter() Converter {
	return &ConverterImpl{}
}

func TimeToTime(t *time.Time) *time.Time {
	return t
}
