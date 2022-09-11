package libauth

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id   int64
	Type int64
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
		if claims, met := token.(Claims); met {
			return &claims, true
		}
	}
	return nil, false
}

func (a *Auth) GenerateToken(id int64, ty int64, expire time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		Id:   id,
		Type: ty,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(expireTime),
			Issuer:    a.config.Issuer,
		},
	}

	tokenClaims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(a.config.JwtSecret)
	return token, err
}
