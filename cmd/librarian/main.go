package main

import (
	"os"

	"github.com/tuihub/librarian/internal/client"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/inprocgrpc"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	miner "github.com/tuihub/protos/pkg/librarian/miner/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name string //nolint:gochecknoglobals //TODO
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
	mapperClientSelector,
	searcherClientSelector,
	porterClientSelector,
	minerClientSelector,
)

func newApp(gs *grpc.Server, hs *http.Server, mq *libmq.MQ, cron *libcron.Cron) *kratos.App {
	return kratos.New(
		kratos.ID(id+name),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, mq, cron),
	)
}

func main() {
	appSettings, err := libapp.NewAppSettings(id, name, version, protoVersion, date)
	if err != nil {
		panic(err)
	}

	var bc conf.Librarian
	appSettings.LoadConfig(&bc)

	if bc.EnableServiceDiscovery == nil {
		bc.EnableServiceDiscovery = new(conf.Librarian_EnableServiceDiscovery)
	}

	app, cleanup, err := wireApp(
		bc.EnableServiceDiscovery,
		bc.Sephirah.Server,
		bc.Sephirah.Data,
		bc.Mapper.Data,
		bc.Searcher.Data,
		bc.Porter.Data,
		bc.Miner.Data,
		bc.Sephirah.Auth,
		bc.Sephirah.Mq,
		bc.Sephirah.Cache,
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

func mapperClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	inproc *inprocgrpc.InprocClients,
) (mapper.LibrarianMapperServiceClient, error) {
	if conf.Mapper {
		return client.NewMapperClient()
	}
	return inproc.Mapper, nil
}

func searcherClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	inproc *inprocgrpc.InprocClients,
) (searcher.LibrarianSearcherServiceClient, error) {
	if conf.Searcher {
		return client.NewSearcherClient()
	}
	return inproc.Searcher, nil
}

func porterClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	inproc *inprocgrpc.InprocClients,
) (porter.LibrarianPorterServiceClient, error) {
	if conf.Porter {
		return client.NewPorterClient()
	}
	return inproc.Porter, nil
}

func minerClientSelector(
	conf *conf.Librarian_EnableServiceDiscovery,
	inproc *inprocgrpc.InprocClients,
) (miner.LibrarianMinerServiceClient, error) {
	if conf.Miner {
		return client.NewMinerClient()
	}
	return inproc.Miner, nil
}
