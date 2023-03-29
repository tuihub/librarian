//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/tuihub/librarian/app/searcher/internal/biz"
	"github.com/tuihub/librarian/app/searcher/internal/data"
	"github.com/tuihub/librarian/app/searcher/internal/service"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

	"github.com/google/wire"
	pb "github.com/tuihub/protos/pkg/librarian/searcher/v1"
)

func NewSearcherService(*conf.Searcher_Data, *libapp.Settings) (pb.LibrarianSearcherServiceServer, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
