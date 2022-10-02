package client

import (
	"github.com/tuihub/librarian/app/porter/internal/client/steam"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(steam.NewSteam, steam.NewStoreAPI, steam.NewWebAPI)
