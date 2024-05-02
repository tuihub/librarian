package libapp

import (
	"context"

	"github.com/tuihub/librarian/internal/lib/libauth"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

func NewStreamMiddlewareJwt(auth *libauth.Auth) func(ctx context.Context) (context.Context, error) {
	m := jwt.Server(
		auth.KeyFunc(libauth.ClaimsTypeUploadToken),
		jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
		jwt.WithClaims(libauth.NewClaims),
	)
	return func(ctx context.Context) (context.Context, error) {
		resp := context.Background()
		_, err := (m(func(ctx context.Context, _ interface{}) (interface{}, error) {
			resp = ctx
			return nil, nil
		}))(ctx, nil)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
