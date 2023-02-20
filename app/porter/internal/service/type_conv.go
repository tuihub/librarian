package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func toPBAppType(t bizsteam.AppType) librarian.AppType {
	switch t {
	case bizsteam.AppTypeGame:
		return librarian.AppType_APP_TYPE_GAME
	default:
		return librarian.AppType_APP_TYPE_UNSPECIFIED
	}
}

func toBizBucket(t pb.DataSource) bizs3.Bucket {
	switch t { //nolint:gocritic // TODO
	default:
		return bizs3.BucketUnspecified
	}
}

func toPBFeed(t *bizfeed.Feed) (*librarian.Feed, error) {
	feedJSON, err := libcodec.Marshal(libcodec.JSON, t)
	if err != nil {
		return nil, err
	}
	res := new(librarian.Feed)
	err = libcodec.Unmarshal(libcodec.JSON, feedJSON, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
