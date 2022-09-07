package libjwt

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwtv4.RegisteredClaims
}

func KeyFunc(key string) jwtv4.Keyfunc {
	return func(token *jwtv4.Token) (interface{}, error) {
		return key, nil
	}
}

func NewClaims() jwtv4.Claims {
	return Claims{}
}

func FromContext(ctx context.Context) (*Claims, bool) {
	if token, ok := jwt.FromContext(ctx); ok {
		if claims, ok := token.(Claims); ok {
			return &claims, true
		}
	}
	return nil, false
}
