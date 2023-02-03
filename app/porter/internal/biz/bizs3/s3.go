package bizs3

import (
	"context"
	"io"
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
	return &S3{
		repo,
	}
}

func (s *S3) NewPushData(ctx context.Context, bucket Bucket, objectName string) (*PutObject, error) {
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
