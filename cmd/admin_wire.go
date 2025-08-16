//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/google/wire"
	"github.com/tuihub/librarian/internal/biz"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libsearch"
)

func wireAdmin(
	[]*conf.ConfigDigest,
	*conf.Config,
	*libapp.Settings,
) (*biztiphereth.Tiphereth, func(), error) {
	panic(
		wire.Build(
			conf.ProviderSet,
			data.ProviderSet,
			biz.ProviderSet,
			client.ProviderSet,
			libauth.ProviderSet,
			libcache.ProviderSet,
			libidgenerator.ProviderSet,
			libsearch.ProviderSet,
		),
	)
}
