package sentinel

import (
	"context"

	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	pb "github.com/tuihub/protos/pkg/librarian/sentinel/v1"
	"github.com/tuihub/protos/pkg/librarian/sentinel/v1/v1connect"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLibrarianSentinelService)

type LibrarianSentinelService struct {
	v1connect.UnimplementedLibrarianSephirahSentinelServiceHandler

	t *biztiphereth.Tiphereth
	g *bizgebura.Gebura
}

func NewLibrarianSentinelService(
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
) v1connect.LibrarianSephirahSentinelServiceHandler {
	return &LibrarianSentinelService{
		UnimplementedLibrarianSephirahSentinelServiceHandler: v1connect.UnimplementedLibrarianSephirahSentinelServiceHandler{},
		t: t,
		g: g,
	}
}

func (s *LibrarianSentinelService) Heartbeat(
	ctx context.Context,
	req *pb.HeartbeatRequest,
) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{}, nil
}
