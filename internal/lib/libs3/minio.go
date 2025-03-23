package libs3

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func newMinioAdapter(c *conf.S3) (*minio.Client, error) {
	minioClient, err := minio.New(c.GetEndPoint(), &minio.Options{ //nolint:exhaustruct // no need
		Creds:  credentials.NewStaticV4(c.GetAccessKey(), c.GetSecretKey(), ""),
		Secure: c.GetUseSsl(),
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
