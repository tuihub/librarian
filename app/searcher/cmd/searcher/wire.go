//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/app/searcher/internal/server"
	"github.com/tuihub/librarian/app/searcher/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Searcher_Server, *conf.Searcher_Data, *conf.Consul, *libapp.Settings) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		libapp.ProviderSet,
		newApp,
	))
}
