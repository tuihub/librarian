package libauth

import (
	"errors"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuth)

type Auth struct {
	config conf.Auth
}

func NewAuth(config *conf.Auth) (*Auth, error) {
	if config == nil {
		return nil, errors.New("")
	}
	return &Auth{config: conf.Auth{
		Salt:      config.Salt,
		Issuer:    config.Issuer,
		JwtSecret: config.JwtSecret,
	}}, nil
}
