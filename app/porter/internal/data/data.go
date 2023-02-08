package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewS3Repo)

// Data .
type Data struct {
	mc      *minio.Client
	buckets map[bizs3.Bucket]string
}

// NewData .
func NewData(c *conf.Porter_Data) (*Data, error) {
	if c.S3 == nil {
		return nil, errors.New("missing s3 config")
	}
	minioClient, err := minio.New(c.S3.EndPoint, &minio.Options{
		Creds:           credentials.NewStaticV4(c.S3.AccessKey, c.S3.SecretKey, ""),
		Secure:          c.S3.UseSsl,
		Transport:       nil,
		Region:          "",
		BucketLookup:    0,
		TrailingHeaders: false,
		CustomMD5:       nil,
		CustomSHA256:    nil,
	})
	if err != nil {
		return nil, err
	}

	bucketName := defaultBucketName()
	location := "us-east-1"
	for i, v := range bucketName {
		if i == bizs3.BucketUnspecified {
			continue
		}
		if err = initBucket(minioClient, v, location); err != nil {
			return nil, err
		}
	}

	return &Data{
		mc:      minioClient,
		buckets: bucketName,
	}, nil
}

func initBucket(mc *minio.Client, bucketName, location string) error {
	err := mc.MakeBucket(
		context.Background(),
		bucketName,
		minio.MakeBucketOptions{
			Region:        location,
			ObjectLocking: false,
		},
	)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := mc.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			logger.Infof("We already own %s\n", bucketName)
		} else {
			logger.Error(err)
			return err
		}
	} else {
		logger.Infof("Successfully created %s\n", bucketName)
	}
	return nil
}

func defaultBucketName() map[bizs3.Bucket]string {
	return map[bizs3.Bucket]string{
		bizs3.BucketUnspecified: "",
		bizs3.BucketDefault:     "default",
	}
}
