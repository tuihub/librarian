//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/miner/internal/biz"
	"github.com/tuihub/librarian/app/miner/internal/data"
	"github.com/tuihub/librarian/app/miner/internal/server"
	"github.com/tuihub/librarian/app/miner/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Miner_Server, *conf.Miner_Data, *libapp.Settings) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		libapp.ProviderSet,
		newApp,
	))
}
