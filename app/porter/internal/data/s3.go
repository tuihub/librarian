package data

import (
	"context"
	"io"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"

	"github.com/minio/minio-go/v7"
)

type s3Repo struct {
	data *Data
}

func NewS3Repo(data *Data) bizs3.S3Repo {
	return &s3Repo{
		data: data,
	}
}

func (s *s3Repo) PutObject(ctx context.Context, r io.Reader, bucket bizs3.Bucket, objectName string) error {
	_, err := s.data.mc.PutObject(
		ctx,
		s.data.buckets[bucket],
		objectName,
		r,
		-1,
		minio.PutObjectOptions{}, //nolint:exhaustruct // default value
	)
	return err
}
