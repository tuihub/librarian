package cmd

import (
	"context"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libdiscovery"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libs3"
	"github.com/tuihub/librarian/internal/lib/libzap"
	"github.com/tuihub/librarian/internal/service/angelaweb"
	"github.com/tuihub/librarian/internal/service/supervisor"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/urfave/cli/v3"
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
)

func newApp(
	gs *grpc.Server,
	hs *http.Server,
	aw *angelaweb.AngelaWeb,
	sv *supervisor.SupervisorService,
	mq *libmq.MQ,
	cron *libcron.Cron,
	consul *conf.Consul,
	s3 libs3.S3,
	inprocPorter *client.InprocPorter,
	observe *libobserve.Observe,
) (*kratos.App, error) {
	options := []kratos.Option{
		kratos.ID(id + name),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(append([]transport.Server{
			gs, hs, aw, sv, mq, cron, s3, observe,
		}, inprocPorter.Servers...)...),
	}
	r, err := libdiscovery.NewRegistrar(consul)
	if err == nil {
		options = append(options, kratos.Registrar(r))
	}
	return kratos.New(options...), nil
}

func runCmdServe(ctx context.Context, cmd *cli.Command) error {
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
		cmd.String(cmdServeFlagConfig),
		cmd.String(cmdServeFlagData),
	)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}

	bc, err := conf.Load(appSettings.ConfPath)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	bc, err = conf.ApplyDeployMode(bc, stdLogger)
	if err != nil {
		stdLogger.Fatalf("Apply deploy mode failed: %v", err)
	}
	digests := conf.GenConfigDigest(bc)
	logConfigDigest(digests, stdLogger)

	if bc.EnableServiceDiscovery == nil {
		bc.EnableServiceDiscovery = new(conf.EnableServiceDiscovery)
	}

	stdLogger.Infof("=== Initializing ===")

	app, cleanup, err := wireServe(
		digests,
		bc,
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

// func minerClientSelector(
//	conf *conf.EnableServiceDiscovery,
//	c *conf.Consul,
//	inproc *inprocgrpc.InprocClients,
//	app *libapp.Settings,
// ) (miner.LibrarianMinerServiceClient, error) {
//	if conf.GetMiner() {
//		return client.NewMinerClient(c, app)
//	}
//	return inproc.Miner, nil
//}

func logConfigDigest(digests []*conf.ConfigDigest, logger *zap.SugaredLogger) {
	for _, d := range digests {
		logger.Info(d.String())
	}
}
