package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
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
