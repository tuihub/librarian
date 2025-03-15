//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/internal/biz"
	globalclient "github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/client/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/server"
	"github.com/tuihub/librarian/internal/service/sephirah"
	"github.com/tuihub/librarian/internal/service/supervisor"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.SephirahServer,
	*conf.Database,
	*conf.S3,
	*conf.Porter,
	*conf.Auth,
	*conf.MQ,
	*conf.Cache,
	*conf.Consul,
	*conf.Search,
	*libapp.Settings,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		globalclient.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		supervisor.ProviderSet,
		sephirah.ProviderSet,
		libauth.ProviderSet,
		libmq.ProviderSet,
		libcron.ProviderSet,
		libcache.ProviderSet,
		libobserve.ProviderSet,
		libapp.ProviderSet,
		libidgenerator.ProviderSet,
		libsearch.ProviderSet,
		newApp,
	))
}
