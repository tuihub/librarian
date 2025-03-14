//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libidgenerator"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsearch"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/google/wire"
)

func NewSephirahService(
	*conf.SephirahServer,
	*conf.Database,
	*conf.S3,
	*conf.Porter,
	*conf.Consul,
	*libauth.Auth,
	*libmq.MQ,
	*libcron.Cron,
	libcache.Store,
	*libapp.Settings,
	*libidgenerator.IDGenerator,
	libsearch.Search,
	miner.LibrarianMinerServiceClient,
) (pb.LibrarianSephirahServiceServer, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		supervisor.ProviderSet,
		service.ProviderSet,
	))
}
