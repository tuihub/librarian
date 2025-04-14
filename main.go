package main

import (
	"os"

	"github.com/tuihub/librarian/cmd"
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

func main() {
	app := cmd.NewCmd(name, version, id, date, protoVersion)
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
