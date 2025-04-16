package libs3

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
)

var ProviderSet = wire.NewSet(NewS3)

type S3 interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	MakeBucket(ctx context.Context, bucketName string) error
	GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error)
	PutObject(
		ctx context.Context,
		bucketName, objectName string,
		reader io.Reader,
		objectSize int64,
	) (minio.UploadInfo, error)
	RemoveObject(ctx context.Context, bucketName, objectName string) error
	ListObjects(ctx context.Context, bucketName string) <-chan minio.ObjectInfo
	PresignedGetObject(
		ctx context.Context,
		bucketName, objectName string,
		expires time.Duration,
		reqParams url.Values,
	) (*url.URL, error)
	PresignedPutObject(ctx context.Context, bucketName, objectName string, expires time.Duration) (*url.URL, error)
}

func NewS3(c *conf.S3) (S3, error) {
	if c == nil {
		c = new(conf.S3)
	}
	switch c.GetDriver() {
	case "memory":
		return newFakeS3Adapter(c)
	case "file":
		return newFakeS3Adapter(c)
	case "minio":
		return newMinioAdapter(c)
	}
	return nil, fmt.Errorf("unsupported driver: %s", c.GetDriver())
}
