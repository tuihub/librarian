package libs3

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3afero"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/afero"
)

type fakeS3Adapter struct {
	client    *minio.Client
	faker     *gofakes3.GoFakeS3
	ch        chan struct{}
	endpoint  string
	accessKey string
	secretKey string
}

func newFakeS3Adapter(c *conf.Storage, app *libapp.Settings) (*fakeS3Adapter, error) {
	var backend gofakes3.Backend
	switch c.Driver {
	case conf.StorageDriverMemory:
		backend = s3mem.New()
	case conf.StorageDriverFile:
		var err error
		backend, err = s3afero.MultiBucket(
			afero.NewBasePathFs(
				afero.NewOsFs(),
				path.Join(app.DataPath, "data"),
			),
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported driver: %s", c.Driver)
	}
	faker := gofakes3.New(backend)
	handler := faker.Server()
	// Create custom transport that redirects to our FakeS3 handler
	transport := &fakeS3Transport{handler: handler}

	// Create Minio client with our custom transport
	minioClient, err := minio.New(
		net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
		&minio.Options{ //nolint:exhaustruct // no need
			Secure:    false,
			Transport: transport,
		},
	)
	if err != nil {
		return nil, err
	}

	return &fakeS3Adapter{
		client:    minioClient,
		faker:     faker,
		ch:        make(chan struct{}),
		endpoint:  net.JoinHostPort(c.Host, strconv.Itoa(int(c.Port))),
		accessKey: c.AccessKey,
		secretKey: c.SecretKey,
	}, nil
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

func (m *fakeS3Adapter) Start(ctx context.Context) error {
	// Create a handler wrapped with the token validation middleware
	handler := m.tokenValidationMiddleware(m.faker.Server())

	server := &http.Server{
		Addr:              m.endpoint,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second, //nolint: mnd // no need
	}

	// Channel to signal server shutdown complete
	serverClosed := make(chan struct{})

	go func() {
		<-m.ch // Wait for stop signal

		// Create a shutdown context with timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //nolint: mnd // no need
		defer cancel()

		_ = server.Shutdown(shutdownCtx)
		close(serverClosed)
	}()

	err := server.ListenAndServe()

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("FakeS3 server error: %v\n", err)
	}

	<-serverClosed

	return err
}

const (
	accessKeyName  = "AWSAccessKeyId"
	expirationName = "Expires"
	signatureName  = "Signature"
)

// tokenValidationMiddleware validates the security token in presigned URLs.
func (m *fakeS3Adapter) tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// Extract AWS style authentication parameters
		accessKey := query.Get(accessKeyName)
		expiration := query.Get(expirationName)
		signature := query.Get(signatureName)

		// Validate required parameters
		if expiration == "" || signature == "" {
			http.Error(w, "Unauthorized: Missing required authentication parameters", http.StatusUnauthorized)
			return
		}

		// Verify the access key matches our configured access key
		if accessKey != m.accessKey {
			http.Error(w, "Unauthorized: Invalid access key", http.StatusUnauthorized)
			return
		}

		// Parse the path to get bucket and object
		objectPath := strings.TrimPrefix(r.URL.Path, "/")
		parts := strings.SplitN(objectPath, "/", 2) //nolint: mnd // no need
		if len(parts) != 2 {                        //nolint: mnd // no need
			http.Error(w, "Invalid request path", http.StatusBadRequest)
			return
		}
		bucketName := parts[0]
		objectName := parts[1]

		// Compute the expected signature
		expectedSignature := m.generateS3Signature(r.Method, bucketName, objectName, expiration)

		// Validate the signature
		if signature != expectedSignature {
			http.Error(w, "Unauthorized: Invalid signature", http.StatusUnauthorized)
			return
		}

		// Authentication is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}

// generateS3Signature creates an AWS/Minio compatible signature for authentication.
func (m *fakeS3Adapter) generateS3Signature(method, bucketName, objectName, expires string) string {
	// Create a canonical string to sign (similar to how S3 builds it)
	// This is a simplified version focusing on the most essential components
	var stringToSign strings.Builder

	stringToSign.WriteString(method + "\n")
	stringToSign.WriteString("\n") // Content-MD5 (empty)
	stringToSign.WriteString("\n") // Content-Type (empty)
	stringToSign.WriteString(expires + "\n")

	// Add canonicalized resource
	stringToSign.WriteString("/" + bucketName + "/" + objectName)

	// Create signature using HMAC-SHA1 (AWS S3 standard)
	h := hmac.New(sha256.New, []byte(m.secretKey))
	h.Write([]byte(stringToSign.String()))
	return hex.EncodeToString(h.Sum(nil))
}

func (m *fakeS3Adapter) Stop(ctx context.Context) error {
	m.ch <- struct{}{}
	return nil
}

func (m *fakeS3Adapter) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return m.client.BucketExists(ctx, bucketName)
}

func (m *fakeS3Adapter) MakeBucket(ctx context.Context, bucketName string) error {
	return m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}) //nolint:exhaustruct // no need
}

func (m *fakeS3Adapter) GetObject(ctx context.Context, bucketName, objectName string) (*minio.Object, error) {
	return m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{}) //nolint:exhaustruct // no need
}

func (m *fakeS3Adapter) PutObject(
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

func (m *fakeS3Adapter) RemoveObject(ctx context.Context, bucketName, objectName string) error {
	return m.client.RemoveObject(
		ctx,
		bucketName,
		objectName,
		minio.RemoveObjectOptions{}, //nolint:exhaustruct // no need
	)
}

func (m *fakeS3Adapter) ListObjects(ctx context.Context, bucketName string) <-chan minio.ObjectInfo {
	return m.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{}) //nolint:exhaustruct // no need
}

func (m *fakeS3Adapter) PresignedGetObject(
	ctx context.Context,
	bucketName, objectName string,
	expires time.Duration,
	reqParams url.Values,
) (*url.URL, error) {
	// Get the presigned URL from MinIO client
	presignedURL, err := m.client.PresignedGetObject(ctx, bucketName, objectName, expires, reqParams)
	if err != nil {
		return nil, err
	}

	// Add AWS-compatible authentication parameters manually
	expiresStr := strconv.FormatInt(time.Now().Add(expires).Unix(), 10)
	signature := m.generateS3Signature("GET", bucketName, objectName, expiresStr)

	// Update the URL query parameters
	query := presignedURL.Query()
	query.Set(accessKeyName, m.accessKey)
	query.Set(expirationName, expiresStr)
	query.Set(signatureName, signature)
	presignedURL.RawQuery = query.Encode()

	return presignedURL, nil
}

func (m *fakeS3Adapter) PresignedPutObject(
	ctx context.Context,
	bucketName, objectName string,
	expires time.Duration,
) (*url.URL, error) {
	// Get the presigned URL from MinIO client
	presignedURL, err := m.client.PresignedPutObject(ctx, bucketName, objectName, expires)
	if err != nil {
		return nil, err
	}

	// Add AWS-compatible authentication parameters manually
	expiresStr := strconv.FormatInt(time.Now().Add(expires).Unix(), 10)
	signature := m.generateS3Signature("PUT", bucketName, objectName, expiresStr)

	// Update the URL query parameters
	query := presignedURL.Query()
	query.Set(accessKeyName, m.accessKey)
	query.Set(expirationName, expiresStr)
	query.Set(signatureName, signature)
	presignedURL.RawQuery = query.Encode()

	return presignedURL, nil
}
