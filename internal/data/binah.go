package data

import (
	"context"
	"errors"
	"io"
	"net/url"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libs3"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/minio/minio-go/v7"
)

type BinahRepo struct {
	mc      *minio.Client
	buckets map[Bucket]string
}

func NewBinahRepo(c *conf.S3) (*BinahRepo, error) {
	if c == nil || len(c.GetDriver()) == 0 {
		return new(BinahRepo), nil
	}
	minioClient, err := libs3.NewS3(c)
	if err != nil {
		return nil, err
	}

	bucketName := defaultBucketName()
	location := "us-east-1"
	for i, v := range bucketName {
		if i == BucketUnspecified {
			continue
		}
		if err = initBucket(minioClient, v, location); err != nil {
			return nil, err
		}
	}

	return &BinahRepo{
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

type Bucket int

const (
	BucketUnspecified Bucket = iota
	BucketDefault
)

func defaultBucketName() map[Bucket]string {
	return map[Bucket]string{
		BucketUnspecified: "",
		BucketDefault:     "default",
	}
}

func (s *BinahRepo) check() error {
	if s.mc != nil {
		return nil
	}
	return errors.New("storage feature is not enabled")
}

func (s *BinahRepo) PutObject(ctx context.Context, r io.Reader, bucket Bucket, objectName string) error {
	if err := s.check(); err != nil {
		return err
	}
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

func (s *BinahRepo) PresignedPutObject(
	ctx context.Context,
	bucket Bucket,
	objectName string,
	expires time.Duration,
) (string, error) {
	if err := s.check(); err != nil {
		return "", err
	}
	res, err := s.mc.PresignedPutObject(ctx, s.buckets[bucket], objectName, expires)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (s *BinahRepo) PresignedGetObject(
	ctx context.Context,
	bucket Bucket,
	objectName string,
	expires time.Duration,
) (string, error) {
	if err := s.check(); err != nil {
		return "", err
	}
	reqParams := make(url.Values)
	res, err := s.mc.PresignedGetObject(ctx, s.buckets[bucket], objectName, expires, reqParams)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
