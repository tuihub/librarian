package libauth

import (
	"errors"

	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
	"github.com/samber/lo"
)

var ProviderSet = wire.NewSet(NewAuth)

type Auth struct {
	config conf.Auth
}

func NewAuth(config *conf.Auth) (*Auth, error) {
	if config == nil {
		return nil, errors.New("auth config is required")
	}
	return &Auth{config: lo.FromPtr(config)}, nil
}
