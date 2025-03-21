package libauth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tuihub/librarian/internal/model"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID           model.InternalID `json:"uid,string"`
	PorterID         model.InternalID `json:"pid,string,omitempty"`
	Type             ClaimsType       `json:"ct"`
	UserType         model.UserType   `json:"ut"`
	TransferMetadata any              `json:"tm,omitempty"`
	jwtv5.RegisteredClaims
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
	return transport.NewServerContext(ctx, &Transport{
		kind:      "",
		endpoint:  "",
		operation: "",
		reqHeader: newTokenHeader("Authorization", fmt.Sprintf("Bearer %s", token)),
	})
}

func ValidateString(tokenString string, keyFunc jwtv5.Keyfunc) (bool, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, new(Claims), keyFunc)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func FromContextAssertUserType(ctx context.Context, userTypes ...model.UserType) *Claims {
	if userTypes == nil {
		userTypes = []model.UserType{model.UserTypeAdmin, model.UserTypeNormal}
	}
	c := FromContext(ctx)
	for _, ut := range userTypes {
		if c.UserType == ut {
			return c
		}
	}
	return nil
}

func (a *Auth) generateSecret(t ClaimsType) interface{} {
	return []byte(fmt.Sprintf("%s%d", a.config.GetJwtSecret(), t))
}

func (a *Auth) GenerateToken(
	uid model.InternalID,
	pid model.InternalID,
	claimsType ClaimsType,
	userType model.UserType,
	transferMetadata any,
	expire time.Duration,
) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		UserID:           uid,
		PorterID:         pid,
		Type:             claimsType,
		UserType:         userType,
		TransferMetadata: transferMetadata,
		RegisteredClaims: jwtv5.RegisteredClaims{
			Issuer:    a.config.GetJwtIssuer(),
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwtv5.NewNumericDate(expireTime),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}

	tokenClaims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(a.generateSecret(claimsType))
	return token, err
}

type headerCarrier http.Header

func (hc headerCarrier) Get(key string) string { return http.Header(hc).Get(key) }

func (hc headerCarrier) Set(key string, value string) { http.Header(hc).Set(key, value) }

func (hc headerCarrier) Add(key string, value string) { http.Header(hc).Add(key, value) }

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice value associated with the passed key.
func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}

func newTokenHeader(headerKey string, token string) *headerCarrier {
	header := &headerCarrier{}
	header.Set(headerKey, token)
	return header
}

type Transport struct {
	kind      transport.Kind
	endpoint  string
	operation string
	reqHeader transport.Header
}

func (tr *Transport) Kind() transport.Kind {
	return tr.kind
}

func (tr *Transport) Endpoint() string {
	return tr.endpoint
}

func (tr *Transport) Operation() string {
	return tr.operation
}

func (tr *Transport) RequestHeader() transport.Header {
	return tr.reqHeader
}

func (tr *Transport) ReplyHeader() transport.Header {
	return nil
}
