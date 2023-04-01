package data

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type s3Repo struct {
	mc      *minio.Client
	buckets map[bizs3.Bucket]string
}

func NewS3Repo(c *conf.Porter_Data) (bizs3.S3Repo, error) {
	if c == nil || c.S3 == nil {
		return new(s3Repo), nil
	}
	minioClient, err := minio.New(c.S3.GetEndPoint(), &minio.Options{
		Creds:           credentials.NewStaticV4(c.S3.GetAccessKey(), c.S3.GetSecretKey(), ""),
		Secure:          c.S3.GetUseSsl(),
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

	return &s3Repo{
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

func (s *s3Repo) FeatureEnabled() bool {
	return s.mc != nil
}

func (s *s3Repo) PutObject(ctx context.Context, r io.Reader, bucket bizs3.Bucket, objectName string) error {
	_, err := s.mc.PutObject(
		ctx,
		s.buckets[bucket],
		objectName,
		r,
		-1,
		minio.PutObjectOptions{}, //nolint:exhaustruct // default value
	)
	return err
}
