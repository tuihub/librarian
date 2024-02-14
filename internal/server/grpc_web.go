package server

import (
	http2 "net/http"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func NewGrpcWebServer(
	s *grpc.Server,
	c *conf.SephirahServer,
	auth *libauth.Auth,
	app *libapp.Settings,
) (*http.Server, error) {
	validator, err := libapp.NewValidator()
	if err != nil {
		return nil, err
	}
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		ratelimit.Server(),
		validator,
	}
	middlewares = append(middlewares, NewTokenMatcher(auth)...)
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	var opts = []http.ServerOption{
		http.Middleware(middlewares...),
	}
	if c.GetGrpcWeb().GetNetwork() != "" {
		opts = append(opts, http.Network(c.GetGrpcWeb().GetNetwork()))
	}
	if c.GetGrpcWeb().GetAddr() != "" {
		opts = append(opts, http.Address(c.GetGrpcWeb().GetAddr()))
	}
	if c.GetGrpcWeb().GetTimeout() != nil {
		opts = append(opts, http.Timeout(c.GetGrpcWeb().GetTimeout().AsDuration()))
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
