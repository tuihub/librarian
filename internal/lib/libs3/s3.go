package libs3

import (
	"fmt"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/minio/minio-go/v7"
)

func NewS3(c *conf.S3) (*minio.Client, error) {
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
