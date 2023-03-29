package main

import (
	"os"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libcron"
	"github.com/tuihub/librarian/internal/lib/libmq"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name string //nolint:gochecknoglobals //TODO
	// version is the version of the compiled software.
	version string

	id, _ = os.Hostname() //nolint:gochecknoglobals //TODO
)

func newApp(gs *grpc.Server, hs *http.Server, mq *libmq.MQ, cron *libcron.Cron) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs, hs, mq, cron),
	)
}

func main() {
	appSettings, err := libapp.NewAppSettings(id, name, version)
	if err != nil {
		panic(err)
	}

	var bc conf.Sephirah
	appSettings.LoadConfig(&bc)

	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Auth, bc.Mq, appSettings)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
