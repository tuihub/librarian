package main

import (
	"os"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libsentry"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
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
	consul *conf.Consul,
) (*kratos.App, error) {
	options := []kratos.Option{
		kratos.ID(id + name),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, mq, cron),
	}
	r, err := libapp.NewRegistrar(consul)
	if err == nil {
		options = append(options, kratos.Registrar(r))
	}
	return kratos.New(options...), nil
}

func main() {
	appSettings, err := libapp.NewAppSettings(id, name, version, protoVersion, date)
	if err != nil {
		panic(err)
	}

	var bc conf.Librarian
	appSettings.LoadConfig(&bc)

	if bc.GetEnableServiceDiscovery() == nil {
		bc.EnableServiceDiscovery = new(conf.Librarian_EnableServiceDiscovery)
	}

	err = libsentry.InitSentry(bc.GetSentry())
	if err != nil {
		panic(err)
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
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
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
) (searcher.LibrarianSearcherServiceClient, error) {
	if conf.GetSearcher() {
		return client.NewSearcherClient(c)
	}
	return inproc.Searcher, nil
}

func minerClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	c *conf.Consul,
	inproc *inprocgrpc.InprocClients,
) (miner.LibrarianMinerServiceClient, error) {
	if conf.GetMiner() {
		return client.NewMinerClient(c)
	}
	return inproc.Miner, nil
}
