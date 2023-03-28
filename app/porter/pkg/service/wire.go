//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/porter/internal/biz"
	"github.com/tuihub/librarian/app/porter/internal/client"
	"github.com/tuihub/librarian/app/porter/internal/data"
	"github.com/tuihub/librarian/app/porter/internal/service"
	"github.com/tuihub/librarian/internal/conf"

	"github.com/google/wire"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func NewPorterService(*conf.Porter_Data) (pb.LibrarianPorterServiceServer, func(), error) {
	panic(wire.Build(
		data.ProviderSet,
		biz.ProviderSet,
		client.ProviderSet,
		service.ProviderSet,
	))
}
