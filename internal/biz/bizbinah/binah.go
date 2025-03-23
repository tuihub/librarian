package bizbinah

import (
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelbinah"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewBinah,
	NewControlBlock,
)

type Binah struct {
	repo     *data.BinahRepo
	callback *modelbinah.ControlBlock
	auth     *libauth.Auth
}

func NewBinah(
	repo *data.BinahRepo,
	callback *modelbinah.ControlBlock,
	auth *libauth.Auth,
) *Binah {
	return &Binah{
		repo:     repo,
		callback: callback,
		auth:     auth,
	}
}

func NewControlBlock(a *libauth.Auth) *modelbinah.ControlBlock {
	return modelbinah.NewControlBlock(a)
}
