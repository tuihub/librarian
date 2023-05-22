package main

import (
	"os"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizs3"
	"github.com/tuihub/librarian/app/porter/internal/biz/bizsteam"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	pb "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name = "porter" //nolint:gochecknoglobals //TODO
	// version is the version of the compiled software.
	version string

	id, _ = os.Hostname() //nolint:gochecknoglobals //TODO

	// date is the build date of the compiled software.
	date string //nolint:gochecknoglobals //TODO

	// version is the proto version of the compiled software.
	protoVersion string //nolint:gochecknoglobals //TODO
)

func newApp(gs *grpc.Server, r registry.Registrar, m metadata) *kratos.App {
	return kratos.New(
		kratos.ID(id+"porter"),
		kratos.Name(name),
		kratos.Version(version),
		kratos.Metadata(m),
		kratos.Server(gs),
		kratos.Registrar(r),
	)
}

func main() {
	appSettings, err := libapp.NewAppSettings(id, name, version, protoVersion, date)
	if err != nil {
		panic(err)
	}

	var bc conf.Porter
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

type metadata map[string]string

func newMetadata(steam *bizsteam.SteamUseCase, s3 *bizs3.S3) metadata {
	v := "enable"
	res := metadata{}
	if steam.FeatureEnabled() {
		res[pb.FeatureFlag_name[int32(pb.FeatureFlag_FEATURE_FLAG_SOURCE_STEAM)]] = v
	}
	if s3.FeatureEnabled() {
		res[pb.FeatureFlag_name[int32(pb.FeatureFlag_FEATURE_FLAG_DEFAULT_DATA_STORAGE)]] = v
	}
	return res
}
