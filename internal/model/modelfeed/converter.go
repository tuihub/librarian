package modelfeed

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/internal/model/modelfeed --packageName modelfeed --output ./generated.go ./

// goverter:converter
// goverter:extend FromPBInternalID
type Converter interface {
	// goverter:matchIgnoreCase
	// goverter:mapIdentity Id
	ToPBFeed(*Feed) *librarian.Feed
	// goverter:map ID Id
	ToPBFeedInternalID(Feed) librarian.InternalID
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	// goverter:map UpdatedParsed | ToPBTime
	// goverter:map PublishedParsed | ToPBTime
	ToPBFeedItem(*Item) *librarian.FeedItem
	// goverter:matchIgnoreCase
	ToPBFeedImage(*Image) *librarian.FeedImage
	// goverter:matchIgnoreCase
	ToPBFeedEnclosure(*Enclosure) *librarian.FeedEnclosure

	// goverter:matchIgnoreCase
	// goverter:map Id ID
	// goverter:ignore FeedType
	// goverter:ignore FeedVersion
	FromPBFeed(*librarian.Feed) *Feed
	// goverter:matchIgnoreCase
	// goverter:map Id ID
	// goverter:map UpdatedParsed | FromPBTime
	// goverter:map PublishedParsed | FromPBTime
	// goverter:ignore DigestDescription
	// goverter:ignore DigestImages
	FromPBFeedItem(*librarian.FeedItem) *Item
	FromPBFeedItemList([]*librarian.FeedItem) []*Item
	// goverter:matchIgnoreCase
	FromPBFeedImage(*librarian.FeedImage) *Image
	// goverter:matchIgnoreCase
	FromPBFeedEnclosure(*librarian.FeedEnclosure) *Enclosure
}

func NewConverter() Converter {
	return &ConverterImpl{}
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
	return model.InternalID(id.Id)
}
