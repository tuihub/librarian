package converter

import (
	"time"

	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/mmcdole/gofeed"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --packagePath github.com/tuihub/librarian/app/porter/internal/client/feed/converter --packageName converter --output ./generated.go ./

// goverter:converter
type Converter interface {
	// goverter:matchIgnoreCase
	// goverter:ignore ID
	ToPBFeed(t *gofeed.Feed) *modelfeed.Feed
	// goverter:matchIgnoreCase
	// goverter:ignore ID
	// goverter:map UpdatedParsed | TimeToTime
	// goverter:map PublishedParsed | TimeToTime
	// goverter:ignore PublishPlatform
	// goverter:ignore DigestDescription
	// goverter:ignore DigestImages
	ToPBFeedItem(t *gofeed.Item) *modelfeed.Item
}

func NewConverter() Converter {
	return &ConverterImpl{}
}

func TimeToTime(t *time.Time) *time.Time {
	return t
}
