package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
)

func main() {
	ts := NewFakeS3()
	err := ts.Start(context.Background())
	if err != nil {
		return
	}
}

func NewFakeS3() *http.Server {
	backend := s3mem.New()
	faker := gofakes3.New(backend, gofakes3.WithGlobalLog())
	srv := http.NewServer(http.Address("127.0.0.1:9000"), http.Middleware(logging.Server(log.GetLogger())))
	srv.HandlePrefix("/", faker.Server())
	return srv
}
