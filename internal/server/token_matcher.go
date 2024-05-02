package server

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewTokenMatcher(auth *libauth.Auth) []middleware.Middleware {
	return []middleware.Middleware{
		selector.Server(
			libapp.NewLeakyBucketMiddleware(10), //nolint:gomnd // 10 requests per second
		).Match(NewAllowAnonymousMatcher()).Build(),
		selector.Server(
			libapp.NewLeakyBucketMiddleware(1),
		).Match(NewRegisterMatcher()).Build(),
		selector.Server(
			libapp.NewLeakyBucketMiddleware(1),
		).Match(NewLoginMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeAccessToken),
				jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewAccessTokenMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeRefreshToken),
				jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewRefreshTokenMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeUploadToken),
				jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewUploadTokenMatcher()).Build(),
		selector.Server(
			jwt.Server(
				auth.KeyFunc(libauth.ClaimsTypeDownloadToken),
				jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(libauth.NewClaims),
			),
		).Match(NewDownloadTokenMatcher()).Build(),
	}
}

func allowAnonymous() map[string]bool {
	return map[string]bool{
		"/grpc.health.v1.Health/Check":                                         true,
		"/grpc.health.v1.Health/Watch":                                         true,
		"/librarian.sephirah.v1.LibrarianSephirahService/GetServerInformation": true,
	}
}

func register() map[string]bool {
	return map[string]bool{
		"/librarian.sephirah.v1.LibrarianSephirahService/RegisterUser": true,
	}
}

func login() map[string]bool {
	return map[string]bool{
		"/librarian.sephirah.v1.LibrarianSephirahService/GetToken": true,
	}
}

func refreshTokenProtected() map[string]bool {
	return map[string]bool{
		"/librarian.sephirah.v1.LibrarianSephirahService/RefreshToken": true,
	}
}

func uploadTokenProtected() map[string]bool {
	return map[string]bool{
		"/librarian.sephirah.v1.LibrarianSephirahService/SimpleUploadFile":          true,
		"/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFile":       true,
		"/librarian.sephirah.v1.LibrarianSephirahService/PresignedUploadFileStatus": true,
	}
}

func downloadTokenProtected() map[string]bool {
	return map[string]bool{
		"/librarian.sephirah.v1.LibrarianSephirahService/SimpleDownloadFile":    true,
		"/librarian.sephirah.v1.LibrarianSephirahService/PresignedDownloadFile": true,
	}
}

func mergeMap(maps ...map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func NewAllowAnonymousMatcher() selector.MatchFunc {
	list := allowAnonymous()
	return func(ctx context.Context, operation string) bool {
		if ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewRegisterMatcher() selector.MatchFunc {
	list := register()
	return func(ctx context.Context, operation string) bool {
		if ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewLoginMatcher() selector.MatchFunc {
	list := login()
	return func(ctx context.Context, operation string) bool {
		if ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewAccessTokenMatcher() selector.MatchFunc {
	whiteList := mergeMap(
		allowAnonymous(),
		register(),
		login(),
		refreshTokenProtected(),
		uploadTokenProtected(),
		downloadTokenProtected(),
	)
	return func(ctx context.Context, operation string) bool {
		if ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func NewRefreshTokenMatcher() selector.MatchFunc {
	list := refreshTokenProtected()
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewUploadTokenMatcher() selector.MatchFunc {
	list := uploadTokenProtected()
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}

func NewDownloadTokenMatcher() selector.MatchFunc {
	list := downloadTokenProtected()
	return func(ctx context.Context, operation string) bool {
		if _, ok := list[operation]; ok {
			return true
		}
		return false
	}
}
