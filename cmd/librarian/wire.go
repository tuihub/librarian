//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	minerService "github.com/tuihub/librarian/app/miner/pkg/service"
	sephirahService "github.com/tuihub/librarian/app/sephirah/pkg/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/server"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(
	*conf.Librarian_EnableServiceDiscovery,
	*conf.SephirahServer,
	*conf.Database,
	*conf.S3,
	*conf.Porter,
	*conf.Miner_Data,
	*conf.Auth,
	*conf.MQ,
	*conf.Cache,
	*conf.Consul,
	*conf.Search,
	*libapp.Settings,
) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			sephirahService.ProviderSet,
			minerService.ProviderSet,
			server.ProviderSet,
			inprocgrpc.ProviderSet,
			libauth.ProviderSet,
			libmq.ProviderSet,
			libcron.ProviderSet,
			libcache.ProviderSet,
			libobserve.ProviderSet,
			libidgenerator.ProviderSet,
			libsearch.ProviderSet,
			ProviderSet,
		),
	)
}
