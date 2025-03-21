package sephirah

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/biz/bizangela"
	"github.com/tuihub/librarian/internal/biz/bizbinah"
	"github.com/tuihub/librarian/internal/biz/bizchesed"
	"github.com/tuihub/librarian/internal/biz/bizgebura"
	"github.com/tuihub/librarian/internal/biz/biznetzach"
	"github.com/tuihub/librarian/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/biz/bizyesod"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	"github.com/tuihub/librarian/internal/service/supervisor"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	a    *bizangela.Angela
	t    *biztiphereth.Tiphereth
	g    *bizgebura.Gebura
	b    *bizbinah.Binah
	y    *bizyesod.Yesod
	n    *biznetzach.Netzach
	c    *bizchesed.Chesed
	s    *supervisor.Supervisor
	app  *libapp.Settings
	auth *libauth.Auth
	info *pb.ServerInstanceSummary
}

func NewLibrarianSephirahServiceService(
	a *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	c *bizchesed.Chesed,
	s *supervisor.Supervisor,
	app *libapp.Settings,
	auth *libauth.Auth,
	config *conf.SephirahServer,
) pb.LibrarianSephirahServiceServer {
	t.CreateConfiguredAdmin()
	if config == nil {
		config = new(conf.SephirahServer)
	}
	if config.GetInfo() == nil {
		config.Info = new(conf.SephirahServer_Info)
	}
	res := &LibrarianSephirahServiceService{
		UnimplementedLibrarianSephirahServiceServer: pb.UnimplementedLibrarianSephirahServiceServer{},
		a:    a,
		t:    t,
		g:    g,
		b:    b,
		y:    y,
		n:    n,
		c:    c,
		s:    s,
		app:  app,
		auth: auth,
		info: nil,
	}
	res.info = &pb.ServerInstanceSummary{
		Name:          config.GetInfo().GetName(),
		Description:   config.GetInfo().GetDescription(),
		WebsiteUrl:    config.GetInfo().GetWebsiteUrl(),
		LogoUrl:       config.GetInfo().GetLogoUrl(),
		BackgroundUrl: config.GetInfo().GetBackgroundUrl(),
	}
	return res
}

func (s *LibrarianSephirahServiceService) GetServerInformation(ctx context.Context,
	_ *pb.GetServerInformationRequest) (*pb.GetServerInformationResponse, error) {
	featureSummary := converter.ToPBServerFeatureSummary(s.s.GetFeatureSummary())
	featureSummary.FeedItemActions = append(featureSummary.FeedItemActions,
		converter.ToPBFeatureFlagList(s.y.GetBuiltInFeedActions())...,
	)
	return &pb.GetServerInformationResponse{
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
		ServerInstanceSummary: s.info,
		StatusReport:          nil,
	}, nil
}
