//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	globalclient "github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.Sephirah_Server,
	*conf.Sephirah_Data,
	*conf.Sephirah_Porter,
	*conf.Auth,
	*conf.MQ,
	*conf.Cache,
	*libapp.Settings,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		globalclient.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		supervisor.ProviderSet,
		service.ProviderSet,
		libauth.ProviderSet,
		libmq.ProviderSet,
		libcron.ProviderSet,
		libcache.ProviderSet,
		newApp,
	))
}
