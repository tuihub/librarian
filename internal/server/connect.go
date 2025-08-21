package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"
	"syscall"
	"time"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1/v1connect"
	sentinel "github.com/tuihub/protos/pkg/librarian/sentinel/v1/v1connect"
	sephirah "github.com/tuihub/protos/pkg/librarian/sephirah/v1/v1connect"

	"connectrpc.com/grpcreflect"
)

const (
	ReadHeaderTimeout = 30 * time.Second
)

type ConnectServer struct {
	mux    *http.ServeMux
	host   string
	port   string
	server *http.Server
}

func NewConnectServer(
	c *conf.Server,
	auth *libauth.Auth,
	sephirahservice sephirah.LibrarianSephirahServiceHandler,
	sentinelservice sentinel.LibrarianSephirahSentinelServiceHandler,
	porterservice porter.LibrarianSephirahPorterServiceHandler,
	app *libapp.Settings,
) (*ConnectServer, error) {
	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		"librarian.sephirah.v1.LibrarianSephirahService",
		"librarian.porter.v1.LibrarianSephirahPorterService",
		"librarian.sentinel.v1.LibrarianSentinelService",
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	mux.Handle(sephirah.NewLibrarianSephirahServiceHandler(sephirahservice))
	mux.Handle(sentinel.NewLibrarianSephirahSentinelServiceHandler(sentinelservice))
	mux.Handle(porter.NewLibrarianSephirahPorterServiceHandler(porterservice))
	return &ConnectServer{
		mux:    mux,
		host:   c.Main.Host,
		port:   strconv.Itoa(int(c.Main.Port)),
		server: nil,
	}, nil
}

func (s *ConnectServer) Start(ctx context.Context) error {
	lc := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			var err error
			err2 := c.Control(func(fd uintptr) {
				// Enable SO_REUSEADDR for faster restart
				err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
			})
			if err2 != nil {
				return err2
			}
			return err
		},
		KeepAlive: 3 * time.Minute, //nolint:mnd // Enable TCP keep-alive and set idle time
		KeepAliveConfig: net.KeepAliveConfig{
			Enable:   true,
			Idle:     30 * time.Second, //nolint:mnd // Time before first probe
			Interval: 10 * time.Second, //nolint:mnd // Interval between probes
			Count:    3,                //nolint:mnd // Max probe attempts
		},
	}
	lis, err := lc.Listen(ctx, "tcp", net.JoinHostPort(s.host, s.port))
	if err != nil {
		return err
	}

	s.server = &http.Server{
		Handler:           s.mux,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	err = s.server.Serve(lis)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *ConnectServer) Stop(ctx context.Context) error {
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}
