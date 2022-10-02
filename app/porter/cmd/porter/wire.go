//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/porter/internal/biz"
	"github.com/tuihub/librarian/app/porter/internal/client"
	"github.com/tuihub/librarian/app/porter/internal/data"
	"github.com/tuihub/librarian/app/porter/internal/server"
	"github.com/tuihub/librarian/app/porter/internal/service"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Porter_Server, *conf.Porter_Data) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
