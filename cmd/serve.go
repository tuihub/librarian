package cmd

import (
	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libs3"
	"github.com/tuihub/librarian/internal/lib/libzap"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/service/angelaweb"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const (
	cmdServeFlagConfig = "config"
	cmdServeFlagData   = "data"
)

func newCmdServe() *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "Run the Librarian service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    cmdServeFlagConfig,
				Aliases: []string{"c"},
				Value:   "config.toml",
				Usage:   "Path to the configuration file",
			},
			&cli.StringFlag{
				Name:    cmdServeFlagData,
				Aliases: []string{"d"},
				Value:   "data",
				Usage:   "Path to the data directory",
			},
		},
		Action: runCmdServe,
	}
}

var ProviderSet = wire.NewSet(
	newApp,
	minerClientSelector,
)

func newApp(
	gs *grpc.Server,
	hs *http.Server,
	aw *angelaweb.AngelaWeb,
	mq *libmq.MQ,
	cron *libcron.Cron,
	obs *libobserve.BuiltInObserver,
	consul *conf.Consul,
	s3 libs3.S3,
) (*kratos.App, error) {
	options := []kratos.Option{
		kratos.ID(id + name),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, aw, mq, cron, obs, s3),
	}
	r, err := libapp.NewRegistrar(consul)
	if err == nil {
		options = append(options, kratos.Registrar(r))
	}
	return kratos.New(options...), nil
}

func runCmdServe(ctx *cli.Context) error {
	stdLogger := libzap.NewStdout(libzap.InfoLevel).Sugar()
	stdLogger.Infof("=== Configuring ===")
	stdLogger.Infof("[Service\t] Name: %s", name)
	stdLogger.Infof("[Service\t] Version: %s", version)
	appSettings, err := libapp.NewAppSettings(
		id,
		name,
		version,
		protoVersion,
		date,
		ctx.String(cmdServeFlagConfig),
		ctx.String(cmdServeFlagData),
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}

	var bc conf.Librarian
	err = appSettings.LoadConfig(&bc)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	digests := genConfigDigest(&bc)
	logConfigDigest(digests, stdLogger)

	if bc.GetEnableServiceDiscovery() == nil {
		bc.EnableServiceDiscovery = new(conf.Librarian_EnableServiceDiscovery)
	}

	stdLogger.Infof("=== Initializing ===")

	err = libobserve.InitOTEL(bc.GetOtlp())
	if err != nil {
		stdLogger.Fatalf("Initialize OTLP client failed: %v", err)
	}

	app, cleanup, err := wireServe(
		digests,
		bc.GetEnableServiceDiscovery(),
		bc.GetServer(),
		bc.GetDatabase(),
		bc.GetS3(),
		bc.GetPorter(),
		bc.GetMiner().GetData(),
		bc.GetAuth(),
		bc.GetMq(),
		bc.GetCache(),
		bc.GetConsul(),
		bc.GetSearch(),
		appSettings,
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}
	defer cleanup()

	// start and wait for stop signal
	stdLogger.Infof("=== Start Service ===")
	if err = app.Run(); err != nil {
		stdLogger.Fatalf("Exit with error: %v", err)
	}
	return nil
}

// func mapperClientSelector(
//	conf *conf.Librarian_EnableServiceDiscovery,
//	c *conf.Consul,
//	inproc *inprocgrpc.InprocClients,
// ) (mapper.LibrarianMapperServiceClient, error) {
//	if conf.GetMapper() {
//		return client.NewMapperClient(c)
//	}
//	return inproc.Mapper, nil
//}

func minerClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	c *conf.Consul,
	inproc *inprocgrpc.InprocClients,
	app *libapp.Settings,
) (miner.LibrarianMinerServiceClient, error) {
	if conf.GetMiner() {
		return client.NewMinerClient(c, app)
	}
	return inproc.Miner, nil
}

func genConfigDigest(c *conf.Librarian) []*model.ConfigDigest {
	var digests []*model.ConfigDigest

	digests = append(digests, &model.ConfigDigest{
		Name:    "Server gRPC",
		Enabled: lo.ToPtr(c.GetServer() != nil && c.GetServer().GetGrpc() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.GetServer().GetGrpc().GetAddr()),
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "Server gRPC-Web",
		Enabled: lo.ToPtr(c.GetServer() != nil && c.GetServer().GetGrpcWeb() != nil),
		Driver:  nil,
		Listen:  lo.ToPtr(c.GetServer().GetGrpcWeb().GetAddr()),
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "DB",
		Enabled: lo.ToPtr(c.GetDatabase() != nil && len(c.GetDatabase().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetDatabase().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "MQ",
		Enabled: lo.ToPtr(c.GetMq() != nil && len(c.GetMq().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetMq().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "Cache",
		Enabled: lo.ToPtr(c.GetCache() != nil && len(c.GetCache().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetCache().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "S3",
		Enabled: lo.ToPtr(c.GetS3() != nil && len(c.GetS3().GetDriver()) != 0),
		Driver:  lo.ToPtr(c.GetS3().GetDriver()),
		Listen:  nil,
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "Consul",
		Enabled: lo.ToPtr(c.GetConsul() != nil && len(c.GetConsul().GetAddr()) != 0),
		Driver:  nil,
		Listen:  nil,
	})
	digests = append(digests, &model.ConfigDigest{
		Name:    "OTLP",
		Enabled: lo.ToPtr(c.GetOtlp() != nil && len(c.GetOtlp().GetProtocol()) != 0),
		Driver:  nil,
		Listen:  nil,
	})

	return digests
}

func logConfigDigest(digests []*model.ConfigDigest, logger *zap.SugaredLogger) {
	for _, d := range digests {
		logger.Info(d.String())
	}
}
