//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/app/mapper/internal/data"
	"github.com/tuihub/librarian/app/mapper/internal/server"
	"github.com/tuihub/librarian/app/mapper/internal/service"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Mapper_Server, *conf.Mapper_Data) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
