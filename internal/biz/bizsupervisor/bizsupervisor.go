package bizsupervisor

import (
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSupervisor,
	NewPorterFeatureController,
)

type Supervisor struct {
	app    *libapp.Settings
	repo   *data.SupervisorRepo
	porter *client.Porter
	auth   *libauth.Auth
	id     *libidgenerator.IDGenerator
}

func NewSupervisor(
	app *libapp.Settings,
	repo *data.SupervisorRepo,
	porter *client.Porter,
	auth *libauth.Auth,
	id *libidgenerator.IDGenerator,
) *Supervisor {
	return &Supervisor{
		app:    app,
		repo:   repo,
		porter: porter,
		auth:   auth,
		id:     id,
	}
}
