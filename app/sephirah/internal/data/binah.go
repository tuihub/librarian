package data

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type binahRepo struct {
	mc      *minio.Client
	buckets map[bizbinah.Bucket]string
}

func NewBinahRepo(c *conf.S3) (bizbinah.BinahRepo, error) {
	if c == nil {
		return new(binahRepo), nil
	}
	minioClient, err := minio.New(c.GetEndPoint(), &minio.Options{ //nolint:exhaustruct //TODO
		Creds:  credentials.NewStaticV4(c.GetAccessKey(), c.GetSecretKey(), ""),
		Secure: c.GetUseSsl(),
	})
	if err != nil {
		return nil, err
	}

	bucketName := defaultBucketName()
	location := "us-east-1"
	for i, v := range bucketName {
		if i == bizbinah.BucketUnspecified {
			continue
		}
		if err = initBucket(minioClient, v, location); err != nil {
			return nil, err
		}
	}

	return &binahRepo{
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

func defaultBucketName() map[bizbinah.Bucket]string {
	return map[bizbinah.Bucket]string{
		bizbinah.BucketUnspecified: "",
		bizbinah.BucketDefault:     "default",
	}
}

func (s *binahRepo) FeatureEnabled() bool {
	return s.mc != nil
}

func (s *binahRepo) PutObject(ctx context.Context, r io.Reader, bucket bizbinah.Bucket, objectName string) error {
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

func (s *binahRepo) PresignedPutObject(
	ctx context.Context,
	bucket bizbinah.Bucket,
	objectName string,
	expires time.Duration,
) (string, error) {
	res, err := s.mc.PresignedPutObject(ctx, s.buckets[bucket], objectName, expires)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (s *binahRepo) PresignedGetObject(
	ctx context.Context,
	bucket bizbinah.Bucket,
	objectName string,
	expires time.Duration,
) (string, error) {
	reqParams := make(url.Values)
	res, err := s.mc.PresignedGetObject(ctx, s.buckets[bucket], objectName, expires, reqParams)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
