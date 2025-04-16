package libs3

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioAdapter struct {
	client *minio.Client
	ch     chan struct{}
}

func newMinioAdapter(c *conf.S3) (*minioAdapter, error) {
	minioClient, err := minio.New(c.GetEndPoint(), &minio.Options{ //nolint:exhaustruct // no need
		Creds:  credentials.NewStaticV4(c.GetAccessKey(), c.GetSecretKey(), ""),
		Secure: c.GetUseSsl(),
	})
	if err != nil {
		return nil, err
	}

	return &minioAdapter{
		client: minioClient,
		ch:     make(chan struct{}),
	}, nil
}

func (m *minioAdapter) Start(_ context.Context) error {
	<-m.ch
	return nil
}

func (m *minioAdapter) Stop(_ context.Context) error {
	m.ch <- struct{}{}
	return nil
}

func (m *minioAdapter) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return m.client.BucketExists(ctx, bucketName)
}

func (m *minioAdapter) MakeBucket(ctx context.Context, bucketName string) error {
	return m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}) //nolint:exhaustruct // no need
}

func (m *minioAdapter) GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error) {
	return m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{}) //nolint:exhaustruct // no need
}

func (m *minioAdapter) PutObject(
	ctx context.Context,
	bucketName, objectName string,
	reader io.Reader,
	objectSize int64,
) (minio.UploadInfo, error) {
	return m.client.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		objectSize,
		minio.PutObjectOptions{}, //nolint:exhaustruct // no need
	)
}

func (m *minioAdapter) RemoveObject(ctx context.Context, bucketName, objectName string) error {
	return m.client.RemoveObject(
		ctx,
		bucketName,
		objectName,
		minio.RemoveObjectOptions{}, //nolint:exhaustruct // no need
	)
}

func (m *minioAdapter) ListObjects(ctx context.Context, bucketName string) <-chan minio.ObjectInfo {
	return m.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{}) //nolint:exhaustruct // no need
}

func (m *minioAdapter) PresignedGetObject(
	ctx context.Context,
	bucketName, objectName string,
	expires time.Duration,
	reqParams url.Values,
) (*url.URL, error) {
	return m.client.PresignedGetObject(ctx, bucketName, objectName, expires, reqParams)
}

func (m *minioAdapter) PresignedPutObject(
	ctx context.Context,
	bucketName, objectName string,
	expires time.Duration,
) (*url.URL, error) {
	return m.client.PresignedPutObject(ctx, bucketName, objectName, expires)
}
