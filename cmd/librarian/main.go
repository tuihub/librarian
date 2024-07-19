package main

import (
	"os"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsentry"
	"github.com/tuihub/librarian/internal/lib/libzap"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name = "sephirah" //nolint:gochecknoglobals //TODO
	// version is the version of the compiled software.
	version string

	id, _ = os.Hostname() //nolint:gochecknoglobals //TODO

	// date is the build date of the compiled software.
	date string //nolint:gochecknoglobals //TODO

	// version is the proto version of the compiled software.
	protoVersion string //nolint:gochecknoglobals //TODO
)

var ProviderSet = wire.NewSet(
	newApp,
	// mapperClientSelector,
	searcherClientSelector,
	minerClientSelector,
)

func newApp(
	gs *grpc.Server,
	hs *http.Server,
	mq *libmq.MQ,
	cron *libcron.Cron,
	obs *libobserve.BuiltInObserver,
	consul *conf.Consul,
) (*kratos.App, error) {
	options := []kratos.Option{
		kratos.ID(id + name),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, mq, cron, obs),
	}
	r, err := libapp.NewRegistrar(consul)
	if err == nil {
		options = append(options, kratos.Registrar(r))
	}
	return kratos.New(options...), nil
}

func main() {
	stdLogger := libzap.NewStdout(libzap.InfoLevel).Sugar()
	stdLogger.Infof("=== Configuring ===")
	stdLogger.Infof("[Service] Name: %s", name)
	stdLogger.Infof("[Service] Version: %s", version)
	appSettings, err := libapp.NewAppSettings(id, name, version, protoVersion, date)
	if err != nil {
		stdLogger.Fatalf("Initialize failed: %v", err)
	}

	var bc conf.Librarian
	err = appSettings.LoadConfig(&bc)
	if err != nil {
		stdLogger.Fatalf("Load config failed: %v", err)
	}
	logConfigDigest(&bc, stdLogger)

	if bc.GetEnableServiceDiscovery() == nil {
		bc.EnableServiceDiscovery = new(conf.Librarian_EnableServiceDiscovery)
	}

	stdLogger.Infof("=== Initializing ===")

	err = libobserve.InitOTEL(bc.GetOtlp())
	if err != nil {
		stdLogger.Fatalf("Initialize OTLP client failed: %v", err)
	}

	err = libsentry.InitSentry(bc.GetSentry())
	if err != nil {
		stdLogger.Fatalf("Initialize Sentry client failed: %v", err)
	}

	app, cleanup, err := wireApp(
		bc.GetEnableServiceDiscovery(),
		bc.GetServer(),
		bc.GetDatabase(),
		bc.GetS3(),
		bc.GetPorter(),
		bc.GetMapper().GetData(),
		bc.GetSearcher().GetData(),
		bc.GetMiner().GetData(),
		bc.GetAuth(),
		bc.GetMq(),
		bc.GetCache(),
		bc.GetConsul(),
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

func searcherClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	c *conf.Consul,
	inproc *inprocgrpc.InprocClients,
	app *libapp.Settings,
) (searcher.LibrarianSearcherServiceClient, error) {
	if conf.GetSearcher() {
		return client.NewSearcherClient(c, app)
	}
	return inproc.Searcher, nil
}

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

func logConfigDigest(bc *conf.Librarian, logger *zap.SugaredLogger) { //nolint:gocognit // no need
	if bc == nil {
		logger.Warnf("Config not specified")
		return
	}
	if bc.GetServer() == nil {
		logger.Warnf("[Server\t] Not specified")
	} else {
		if bc.GetServer().GetInfo() != nil && len(bc.GetServer().GetInfo().GetName()) > 0 {
			logger.Infof("[Server\t] Name: %s", bc.GetServer().GetInfo().GetName())
		}
		if bc.GetServer().GetGrpc() != nil {
			logger.Infof("[Server\t] Listen gRPC on: %s", bc.GetServer().GetGrpc().GetAddr())
		}
		if bc.GetServer().GetGrpcWeb() != nil {
			logger.Infof("[Server\t] Listen gRPC-Web on: %s", bc.GetServer().GetGrpcWeb().GetAddr())
		}
	}
	if bc.GetDatabase() == nil || len(bc.GetDatabase().GetDriver()) == 0 {
		logger.Warnf("[DB\t\t] Not specified")
	} else {
		logger.Infof("[DB\t\t] Configured - Driver %s", bc.GetDatabase().GetDriver())
	}
	if bc.GetMq() == nil || len(bc.GetMq().GetDriver()) == 0 {
		logger.Warnf("[MQ\t\t] Not specified")
	} else {
		logger.Infof("[MQ\t\t] Configured - Driver %s", bc.GetMq().GetDriver())
	}
	if bc.GetCache() == nil || len(bc.GetCache().GetDriver()) == 0 {
		logger.Warnf("[Cache\t] Not specified")
	} else {
		logger.Infof("[Cache\t] Configured - Driver %s", bc.GetCache().GetDriver())
	}
	if bc.GetConsul() == nil || len(bc.GetConsul().GetAddr()) == 0 {
		logger.Warnf("[Consul\t] Not specified")
	} else {
		logger.Infof("[Consul\t] Configured")
	}
	if bc.GetSentry() == nil || len(bc.GetSentry().GetDsn()) == 0 {
		logger.Warnf("[Sentry\t] Not specified")
	} else {
		logger.Infof("[Sentry\t] Configured")
	}
	if bc.GetOtlp() == nil || len(bc.GetOtlp().GetProtocol()) == 0 {
		logger.Warnf("[OTLP\t] Not specified")
	} else {
		logger.Infof("[OTLP\t] Configured - Protocol %s", bc.GetOtlp().GetProtocol())
	}
}
