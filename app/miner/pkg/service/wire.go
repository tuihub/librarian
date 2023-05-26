//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/miner/internal/biz"
	"github.com/tuihub/librarian/app/miner/internal/data"
	"github.com/tuihub/librarian/app/miner/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"github.com/google/wire"
)

func NewMinerService(*conf.Miner_Data, *libapp.Settings) (pb.LibrarianMinerServiceServer, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
