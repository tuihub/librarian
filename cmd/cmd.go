package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	// name is the name of the compiled software.
	name string //nolint:gochecknoglobals //no need
	// version is the version of the compiled software.
	version string

	id string //nolint:gochecknoglobals //no need

	// date is the build date of the compiled software.
	date string //nolint:gochecknoglobals //no need

	// version is the proto version of the compiled software.
	protoVersion string //nolint:gochecknoglobals //no need
)

func NewCmd(_name, _version, _id, _date, _protoVersion string) *cli.App {
	name = _name
	version = _version
	id = _id
	date = _date
	protoVersion = _protoVersion
	return &cli.App{
		Name:  "TuiHub Librarian",
		Usage: "Librarian is the standard server implementation of TuiHub",
		Commands: []*cli.Command{
			newCmdServe(),
			newCmdAdmin(),
			newCmdConfig(),
		},
		Action: runCmdServe,
	}
}
