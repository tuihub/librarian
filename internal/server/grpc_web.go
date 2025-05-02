package server

import (
	"net"
	http2 "net/http"
	"strconv"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libobserve"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func NewGrpcWebServer(
	s *grpc.Server,
	c *conf.Server,
	auth *libauth.Auth,
	app *libapp.Settings,
	observer *libobserve.BuiltInObserver,
) (*http.Server, error) {
	validator, err := libapp.NewValidator()
	if err != nil {
		return nil, err
	}
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		ratelimit.Server(),
		tracing.Server(),
		validator,
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	middlewares = append(middlewares, libobserve.Server(observer))
	middlewares = append(middlewares, NewTokenMatcher(auth)...)
	var opts = []http.ServerOption{
		http.Middleware(middlewares...),
	}
	opts = append(opts, http.Address(net.JoinHostPort(c.MainWeb.Host, strconv.Itoa(int(c.MainWeb.Port)))))
	if c.MainWeb.Timeout > 0 {
		opts = append(opts, http.Timeout(c.MainWeb.Timeout))
	}
	srv := http.NewServer(opts...)
	wrappedGrpc := grpcweb.WrapServer(s.Server)
	srv.HandlePrefix("/", http2.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
		}
	}))
	return srv, nil
}
