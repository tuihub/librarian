package server

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewGrpcWebServer, NewAuthMiddleware)

// NewAuthMiddleware https://github.com/go-kratos/kratos/issues/2617
func NewAuthMiddleware(auth *libauth.Auth) func(context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		var newContext context.Context
		mw := jwt.Server(
			auth.KeyFunc(libauth.ClaimsTypeAccessToken),
			jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(libauth.NewClaims),
		)
		handler := mw(func(ctx context.Context, req interface{}) (interface{}, error) {
			newContext = ctx
			return nil, nil
		})
		_, err := handler(ctx, nil)
		if err != nil {
			return ctx, err
		}
		return newContext, nil
	}
}
