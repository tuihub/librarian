package libs3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

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

func NewS3(c *conf.Storage, app *libapp.Settings) (S3, error) {
	if c == nil {
		return nil, errors.New("storage configuration is required")
	}
	switch c.Driver {
	case conf.StorageDriverMemory:
		return newFakeS3Adapter(c, app)
	case conf.StorageDriverFile:
		return newFakeS3Adapter(c, app)
	case conf.StorageDriverS3:
		return newMinioAdapter(c)
	}
	return nil, fmt.Errorf("unsupported driver: %s", c.Driver)
}
