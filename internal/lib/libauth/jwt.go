package libauth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tuihub/librarian/internal/lib/libkratos"
	"github.com/tuihub/librarian/internal/model"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Claims struct {
	UserID   model.InternalID `json:"uid,string"`
	PorterID model.InternalID `json:"pid,string,omitempty"`
	Type     ClaimsType       `json:"ct"`
	UserType model.UserType   `json:"ut"`
	ClaimsTransferExtra
	ClaimsSentinelExtra
	jwtv5.RegisteredClaims
}

type ClaimsExtra func(*ClaimsExtras)

type ClaimsExtras struct {
	ClaimsTransferExtra *ClaimsTransferExtra
	ClaimsSentinelExtra *ClaimsSentinelExtra
}

type ClaimsTransferExtra struct {
	TransferMetadata any `json:"tm,omitempty"`
}

func WithClaimsTransferExtra(extra *ClaimsTransferExtra) ClaimsExtra {
	return func(claims *ClaimsExtras) {
		claims.ClaimsTransferExtra = extra
	}
}

type ClaimsSentinelExtra struct {
	SentinelSessionID model.InternalID `json:"ssid,string"`
}

func WithClaimsSentinelExtra(extra *ClaimsSentinelExtra) ClaimsExtra {
	return func(claims *ClaimsExtras) {
		claims.ClaimsSentinelExtra = extra
	}
}

type ClaimsType int

const (
	ClaimsTypeUnspecified ClaimsType = iota
	ClaimsTypeAccessToken
	ClaimsTypeRefreshToken
	ClaimsTypeUploadToken
	ClaimsTypeDownloadToken
)

func (a *Auth) KeyFunc(t ClaimsType) jwtv5.Keyfunc {
	return func(token *jwtv5.Token) (interface{}, error) {
		return a.generateSecret(t), nil
	}
}

func NewClaims() jwtv5.Claims {
	return new(Claims)
}

func FromContext(ctx context.Context) *Claims {
	if token, ok := jwt.FromContext(ctx); ok {
		if claims, met := token.(*Claims); met {
			return claims
		}
	}
	return nil
}

func RawFromContext(ctx context.Context) string {
	if header, ok := transport.FromServerContext(ctx); ok {
		auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2) //nolint:mnd // exactly 2
		if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
			return ""
		}
		return auths[1]
	}
	return ""
}

func RawToContext(ctx context.Context, token string) context.Context {
	header := &libkratos.HTTPHeaderCarrier{}
	header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return transport.NewServerContext(ctx, &libkratos.Transport{
		K:   "",
		E:   "",
		O:   "",
		Req: header,
		Res: nil,
	})
}

func ValidateString(tokenString string, keyFunc jwtv5.Keyfunc) (bool, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, new(Claims), keyFunc)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func FromString(tokenString string, keyFunc jwtv5.Keyfunc) (*Claims, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, new(Claims), keyFunc)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid claims type")
}

func FromContextAssertUserType(ctx context.Context, userTypes ...model.UserType) *Claims {
	if userTypes == nil {
		userTypes = []model.UserType{model.UserTypeAdmin, model.UserTypeNormal}
	}
	c := FromContext(ctx)
	if c == nil {
		return nil
	}
	for _, ut := range userTypes {
		if c.UserType == ut {
			return c
		}
	}
	return nil
}

func (a *Auth) generateSecret(t ClaimsType) interface{} {
	return []byte(fmt.Sprintf("%s%d", a.config.TokenSecret, t))
}

func (a *Auth) GenerateToken(
	uid model.InternalID,
	pid model.InternalID,
	claimsType ClaimsType,
	userType model.UserType,
	expire time.Duration,
	extras ...ClaimsExtra,
) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	extra := new(ClaimsExtras)
	for _, e := range extras {
		e(extra)
	}

	claims := Claims{
		UserID:              uid,
		PorterID:            pid,
		Type:                claimsType,
		UserType:            userType,
		ClaimsTransferExtra: lo.FromPtr(extra.ClaimsTransferExtra),
		ClaimsSentinelExtra: lo.FromPtr(extra.ClaimsSentinelExtra),
		RegisteredClaims: jwtv5.RegisteredClaims{
			Issuer:    a.config.TokenIssuer,
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwtv5.NewNumericDate(expireTime),
			NotBefore: jwtv5.NewNumericDate(nowTime),
			IssuedAt:  jwtv5.NewNumericDate(nowTime),
			ID:        id.String(),
		},
	}

	tokenClaims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(a.generateSecret(claimsType))
	return token, err
}
