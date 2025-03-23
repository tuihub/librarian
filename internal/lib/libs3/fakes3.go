package libs3

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3afero"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/afero"
)

func newFakeS3Adapter(c *conf.S3) (*minio.Client, error) {
	var backend gofakes3.Backend
	switch c.GetDriver() {
	case "memory":
		backend = s3mem.New()
	case "file":
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		backend, err = s3afero.MultiBucket(
			afero.NewBasePathFs(
				afero.NewOsFs(),
				path.Join(dir, "data"),
			),
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported driver: %s", c.GetDriver())
	}
	faker := gofakes3.New(backend)
	handler := faker.Server()
	// Create custom transport that redirects to our FakeS3 handler
	transport := &fakeS3Transport{handler: handler}

	// Create Minio client with our custom transport
	minioClient, err := minio.New("localhost:9000", &minio.Options{ //nolint:exhaustruct // no need
		Secure:    false,
		Transport: transport,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

// Custom transport that routes requests to the GoFakeS3 handler directly.
type fakeS3Transport struct {
	handler http.Handler
}

func (t *fakeS3Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request with our handler
	t.handler.ServeHTTP(w, req)

	// Get the response from the recorder
	resp := w.Result()

	// Set the request URL in the response to maintain compatibility
	if resp.Request == nil {
		resp.Request = req
	}

	return resp, nil
}
