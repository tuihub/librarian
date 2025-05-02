package sephirah

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/bizkether"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/model/modelangela"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	"github.com/tuihub/librarian/internal/service/supervisor"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/sephirah"

	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ProviderSet = wire.NewSet(NewLibrarianSephirahService)

type LibrarianSephirahService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	a    *bizangela.Angela
	k    *bizkether.Kether
	t    *biztiphereth.Tiphereth
	g    *bizgebura.Gebura
	b    *bizbinah.Binah
	y    *bizyesod.Yesod
	n    *biznetzach.Netzach
	c    *bizchesed.Chesed
	s    *supervisor.Supervisor
	app  *libapp.Settings
	auth *libauth.Auth
}

func NewLibrarianSephirahService(
	a *bizangela.Angela,
	k *bizkether.Kether,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	c *bizchesed.Chesed,
	s *supervisor.Supervisor,
	app *libapp.Settings,
	auth *libauth.Auth,
) pb.LibrarianSephirahServiceServer {
	t.CreateConfiguredAdmin()
	res := &LibrarianSephirahService{
		UnimplementedLibrarianSephirahServiceServer: pb.UnimplementedLibrarianSephirahServiceServer{},
		a:    a,
		k:    k,
		t:    t,
		g:    g,
		b:    b,
		y:    y,
		n:    n,
		c:    c,
		s:    s,
		app:  app,
		auth: auth,
	}
	return res
}

func (s *LibrarianSephirahService) GetServerInformation(ctx context.Context,
	_ *pb.GetServerInformationRequest) (*pb.GetServerInformationResponse, error) {
	featureSummary := converter.ToPBServerFeatureSummary(s.s.GetFeatureSummary())
	featureSummary.FeedItemActions = append(featureSummary.FeedItemActions,
		converter.ToPBFeatureFlagList(s.y.GetBuiltInFeedActions())...,
	)
	info, err := s.a.GetServerInstanceSummary(ctx)
	if err != nil {
		info = new(modelangela.ServerInstanceSummary)
	}
	return &pb.GetServerInformationResponse{
		ServerInformation: &pb.ServerInformation{
			ServerBinarySummary: &pb.ServerBinarySummary{
				SourceCodeAddress: s.app.SourceCodeAddress,
				BuildVersion:      s.app.Version,
				BuildDate:         s.app.BuildDate,
			},
			ProtocolSummary: &pb.ServerProtocolSummary{
				Version: s.app.ProtoVersion,
			},
			CurrentTime:           timestamppb.New(time.Now()),
			FeatureSummary:        featureSummary,
			ServerInstanceSummary: converter.ToPBServerInstanceSummary(info),
			StatusReport:          nil,
		},
	}, nil
}
