package sentinel

import (
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sentinel"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLibrarianSentinelService)

type LibrarianSentinelService struct {
	pb.UnimplementedLibrarianSentinelServiceServer

	t *biztiphereth.Tiphereth
	g *bizgebura.Gebura
}

func NewLibrarianSentinelService(
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
) pb.LibrarianSentinelServiceServer {
	return &LibrarianSentinelService{
		UnimplementedLibrarianSentinelServiceServer: pb.UnimplementedLibrarianSentinelServiceServer{},
		t: t,
		g: g,
	}
}
