//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/google/wire"
	"github.com/tuihub/librarian/internal/biz"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/client/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/supervisor"
)

func wireAdmin(
	[]*model.ConfigDigest,
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
) (*biztiphereth.Tiphereth, func(), error) {
	panic(
		wire.Build(
			data.ProviderSet,
			biz.ProviderSet,
			client.ProviderSet,
			supervisor.ProviderSet,
			libauth.ProviderSet,
			libmq.ProviderSet,
			libcron.ProviderSet,
			libcache.ProviderSet,
			libobserve.ProviderSet,
			libidgenerator.ProviderSet,
			libsearch.ProviderSet,
		),
	)
}
