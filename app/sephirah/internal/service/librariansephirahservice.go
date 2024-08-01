package service

import (
	"context"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizangela"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizbinah"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type LibrarianSephirahServiceService struct {
	pb.UnimplementedLibrarianSephirahServiceServer

	t        *biztiphereth.Tiphereth
	g        *bizgebura.Gebura
	b        *bizbinah.Binah
	y        *bizyesod.Yesod
	n        *biznetzach.Netzach
	c        *bizchesed.Chesed
	s        *supervisor.Supervisor
	app      *libapp.Settings
	auth     *libauth.Auth
	authFunc func(context.Context) (context.Context, error)
	info     *libcache.Key[pb.GetServerInformationResponse]
}

func NewLibrarianSephirahServiceService(
	_ *bizangela.Angela,
	t *biztiphereth.Tiphereth,
	g *bizgebura.Gebura,
	b *bizbinah.Binah,
	y *bizyesod.Yesod,
	n *biznetzach.Netzach,
	c *bizchesed.Chesed,
	s *supervisor.Supervisor,
	app *libapp.Settings,
	auth *libauth.Auth,
	authFunc func(context.Context) (context.Context, error),
	config *conf.SephirahServer,
	store libcache.Store,
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
		t:        t,
		g:        g,
		b:        b,
		y:        y,
		n:        n,
		c:        c,
		s:        s,
		app:      app,
		auth:     auth,
		authFunc: authFunc,
		info:     nil,
	}
	res.info = newServerInfromationCache(res, store, &pb.ServerInstanceSummary{
		Name:          config.GetInfo().GetName(),
		Description:   config.GetInfo().GetDescription(),
		WebsiteUrl:    config.GetInfo().GetWebsiteUrl(),
		LogoUrl:       config.GetInfo().GetLogoUrl(),
		BackgroundUrl: config.GetInfo().GetBackgroundUrl(),
	})
	return res
}

func newServerInfromationCache(
	s *LibrarianSephirahServiceService,
	store libcache.Store,
	serverSummary *pb.ServerInstanceSummary,
) *libcache.Key[pb.GetServerInformationResponse] {
	return libcache.NewKey[pb.GetServerInformationResponse](
		store,
		"GetServerInformationResponse",
		func(ctx context.Context) (*pb.GetServerInformationResponse, error) {
			featureSummary := s.s.GetFeatureSummary()
			featureSummary.FeedItemActions = append(featureSummary.FeedItemActions, s.y.GetBuiltInFeedActions()...)
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
				FeatureSummary:        converter.ToPBServerFeatureSummary(featureSummary),
				ServerInstanceSummary: serverSummary,
				StatusReport:          nil,
			}, nil
		},
		libcache.WithExpiration(time.Minute),
	)
}

func (s *LibrarianSephirahServiceService) GetServerInformation(ctx context.Context,
	_ *pb.GetServerInformationRequest) (*pb.GetServerInformationResponse, error) {
	return s.info.Get(ctx)
}
