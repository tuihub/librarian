package libauth

import (
	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuth)

type Auth struct {
	config conf.Auth
}

func NewAuth(config *conf.Auth) (*Auth, error) {
	if config == nil {
		config = new(conf.Auth)
	}
	return &Auth{config: conf.Auth{
		PasswordSalt: config.GetPasswordSalt(),
		JwtIssuer:    config.GetJwtIssuer(),
		JwtSecret:    config.GetJwtSecret(),
	}}, nil
}
