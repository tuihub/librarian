package server

import (
	http2 "net/http"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func NewGrpcWebServer(
	s *grpc.Server,
	c *conf.Sephirah_Server,
	auth *libauth.Auth,
	app *libapp.Settings,
) *http.Server {
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		ratelimit.Server(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeAccessToken),
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewWhiteListMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeRefreshToken),
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewRefreshTokenMatcher()).Build(),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	var opts = []http.ServerOption{
		http.Middleware(middlewares...),
	}
	if c.GrpcWeb.Network != "" {
		opts = append(opts, http.Network(c.GrpcWeb.Network))
	}
	if c.GrpcWeb.Addr != "" {
		opts = append(opts, http.Address(c.GrpcWeb.Addr))
	}
	if c.GrpcWeb.Timeout != nil {
		opts = append(opts, http.Timeout(c.GrpcWeb.Timeout.AsDuration()))
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
	return srv
}
