//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/sephirah/internal/biz"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/data"
	"github.com/tuihub/librarian/app/sephirah/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/server"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/google/wire"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
)

func NewSephirahService(
	*conf.Sephirah_Data,
	*libauth.Auth,
	*libmq.MQ,
	*libcron.Cron,
	libcache.Store,
	*libapp.Settings,
	mapper.LibrarianMapperServiceClient,
	searcher.LibrarianSearcherServiceClient,
	porter.LibrarianPorterServiceClient,
	miner.LibrarianMinerServiceClient,
) (pb.LibrarianSephirahServiceServer, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		service.ProviderSet,
		server.ProviderSet,
	))
}
