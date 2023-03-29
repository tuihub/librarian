//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/mapper/internal/biz"
	"github.com/tuihub/librarian/app/mapper/internal/data"
	"github.com/tuihub/librarian/app/mapper/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

	pb "github.com/tuihub/protos/pkg/librarian/mapper/v1"

	"github.com/google/wire"
)

func NewMapperService(*conf.Mapper_Data, *libapp.Settings) (pb.LibrarianMapperServiceServer, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
