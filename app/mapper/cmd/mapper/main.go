package main

import (
	"os"

	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name = "mapper" //nolint:gochecknoglobals //TODO
	// version is the version of the compiled software.
	version string

	id, _ = os.Hostname() //nolint:gochecknoglobals //TODO
)

func newApp(gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(gs),
		kratos.Registrar(libapp.NewRegistrar()),
	)
}

func main() {
	appSettings, err := libapp.NewAppSettings(id, name, version)
	if err != nil {
		panic(err)
	}

	var bc conf.Mapper
	appSettings.LoadConfig(&bc)

	app, cleanup, err := wireApp(bc.Server, bc.Data, appSettings)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
