//go:build !goverter

package modelfeed

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen .

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:format function
// goverter:output:file ./generated.go
// goverter:output:package github.com/tuihub/librarian/internal/model/modelfeed
// goverter:matchIgnoreCase
// goverter:ignoreUnexported
// goverter:extend FromPBInternalID
type Converter interface {
	// goverter:map . Id
	ToPBFeed(*Feed) *librarian.Feed
	// goverter:map ID Id
	ToPBFeedInternalID(Feed) librarian.InternalID
	// goverter:ignore Id
	// goverter:map UpdatedParsed | ToPBTime
	// goverter:map PublishedParsed | ToPBTime
	ToPBFeedItem(*Item) *librarian.FeedItem
	ToPBFeedImage(*Image) *librarian.FeedImage
	ToPBFeedEnclosure(*Enclosure) *librarian.FeedEnclosure

	// goverter:map Id ID
	// goverter:ignore FeedType
	// goverter:ignore FeedVersion
	FromPBFeed(*librarian.Feed) *Feed
	// goverter:map Id ID
	// goverter:map UpdatedParsed | FromPBTime
	// goverter:map PublishedParsed | FromPBTime
	// goverter:ignore DigestDescription
	// goverter:ignore DigestImages
	FromPBFeedItem(*librarian.FeedItem) *Item
	FromPBFeedItemList([]*librarian.FeedItem) []*Item
	FromPBFeedImage(*librarian.FeedImage) *Image
	FromPBFeedEnclosure(*librarian.FeedEnclosure) *Enclosure
}

func ToPBTime(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func FromPBTime(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	res := t.AsTime()
	return &res
}

func FromPBInternalID(id *librarian.InternalID) model.InternalID {
	if id == nil {
		return 0
	}
	return model.InternalID(id.GetId())
}
