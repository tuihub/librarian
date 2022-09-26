package libauth

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID   int64
	Type ClaimsType
	jwtv4.RegisteredClaims
}

type ClaimsType string

const (
	ClaimsTypeUnspecified  = ""
	ClaimsTypeAccessToken  = "access_token"
	ClaimsTypeRefreshToken = "refresh_token"
)

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

func (a *Auth) GenerateToken(id int64, ty ClaimsType, expire time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		ID:   id,
		Type: ty,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(expireTime),
			Issuer:    a.config.Issuer,
		},
	}

	tokenClaims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(a.config.JwtSecret))
	return token, err
}
