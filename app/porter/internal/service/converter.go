package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

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
