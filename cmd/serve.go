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
	"github.com/tuihub/librarian/internal/service/angelaweb"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
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

	var bc conf.Config
	err = appSettings.LoadConfig(&bc)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	digests := conf.GenConfigDigest(&bc)
	logConfigDigest(digests, stdLogger)

	if bc.GetEnableServiceDiscovery() == nil {
		bc.EnableServiceDiscovery = new(conf.EnableServiceDiscovery)
	}

	stdLogger.Infof("=== Initializing ===")

	err = libobserve.InitOTEL(bc.GetOtlp())
	if err != nil {
		stdLogger.Fatalf("Initialize OTLP client failed: %v", err)
	}

	app, cleanup, err := wireServe(
		digests,
		&bc,
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
	conf *conf.EnableServiceDiscovery,
	c *conf.Consul,
	inproc *inprocgrpc.InprocClients,
	app *libapp.Settings,
) (miner.LibrarianMinerServiceClient, error) {
	if conf.GetMiner() {
		return client.NewMinerClient(c, app)
	}
	return inproc.Miner, nil
}

func logConfigDigest(digests []*conf.ConfigDigest, logger *zap.SugaredLogger) {
	for _, d := range digests {
		logger.Info(d.String())
	}
}
