package converter

import (
	"time"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter --ignoreUnexportedFields --packagePath github.com/tuihub/librarian/app/porter/internal/service/converter --packageName converter --output ./generated.go ./

// goverter:converter
type Converter interface {
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	ToPBFeed(t *bizfeed.Feed) *librarian.Feed
	// goverter:matchIgnoreCase
	// goverter:ignore Id
	// goverter:map UpdatedParsed | ToPBTime
	// goverter:map PublishedParsed | ToPBTime
	// goverter:map Enclosures Enclosure
	ToPBFeedItem(t *bizfeed.Item) *librarian.FeedItem
	// goverter:matchIgnoreCase
	ToPBFeedImage(t *bizfeed.Image) *librarian.FeedImage
	// goverter:matchIgnoreCase
	ToPBFeedEnclosure(t *bizfeed.Enclosure) *librarian.FeedEnclosure
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

func ToPBAppType(t bizsteam.AppType) librarian.AppType {
	switch t { //nolint:exhaustive //TODO
	case bizsteam.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func ToBizBucket(t pb.DataSource) bizs3.Bucket {
	switch t {
	case pb.DataSource_DATA_SOURCE_UNSPECIFIED:
		return bizs3.BucketUnspecified
	case pb.DataSource_DATA_SOURCE_INTERNAL_DEFAULT:
		return bizs3.BucketDefault
	default:
		return bizs3.BucketUnspecified
	}
}
