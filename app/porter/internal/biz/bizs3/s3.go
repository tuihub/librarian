package bizs3

import (
	"context"
	"io"

	"github.com/go-kratos/kratos/v2/errors"
)

type PutObject struct {
	ch     chan error
	writer *io.PipeWriter
}

func (p *PutObject) Close() error {
	err := p.writer.Close()
	if err != nil {
		return err
	}
	return <-p.ch
}

func (p *PutObject) Write(b []byte) (int, error) {
	return p.writer.Write(b)
}

type S3Repo interface {
	FeatureEnabled() bool
	PutObject(context.Context, io.Reader, Bucket, string) error
}

type S3 struct {
	repo S3Repo
}

type Bucket int

const (
	BucketUnspecified Bucket = iota
	BucketDefault
)

func NewS3(repo S3Repo) *S3 {
	if !repo.FeatureEnabled() {
		return new(S3)
	}
	return &S3{
		repo,
	}
}

func (s *S3) FeatureEnabled() bool {
	return s.repo != nil
}

func (s *S3) NewPushData(ctx context.Context, bucket Bucket, objectName string) (*PutObject, error) {
	if !s.FeatureEnabled() {
		return nil, errors.BadRequest("request disabled feature", "")
	}
	reader, writer := io.Pipe()
	ch := make(chan error)
	go func() {
		ch <- s.repo.PutObject(ctx, reader, bucket, objectName)
	}()
	return &PutObject{
		ch,
		writer,
	}, nil
}
