package server

import (
	"context"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Sephirah_Server,
	auth *libauth.Auth,
	greeter pb.LibrarianSephirahServiceServer,
	app *libapp.Settings,
) *grpc.Server {
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
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeUploadToken),
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewUploadTokenMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeDownloadToken),
				jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewDownloadTokenMatcher()).Build(),
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianSephirahServiceServer(srv, greeter)
	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/grpc.health.v1.Health/Check"] = struct{}{}
	whiteList["/grpc.health.v1.Health/Watch"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/GetServerInformation"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/GetToken"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/RefreshToken"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/SimpleUploadFile"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/SimpleDownloadFile"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFile"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFileStatus"] = struct{}{}
	whiteList["/librarian.sephirah.v1.LibrarianSephirahService/PresignedDownloadFile"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewRefreshTokenMatcher() selector.MatchFunc {
	list := make(map[string]struct{})
	list["/librarian.sephirah.v1.LibrarianSephirahService/RefreshToken"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewUploadTokenMatcher() selector.MatchFunc {
	list := make(map[string]struct{})
	list["/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFile"] = struct{}{}
	list["/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFileStatus"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewDownloadTokenMatcher() selector.MatchFunc {
	list := make(map[string]struct{})
	list["/librarian.sephirah.v1.LibrarianSephirahService/PresignedDownloadFile"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}
