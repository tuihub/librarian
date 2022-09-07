package libapp

import "os"

// go build -ldflags "-X main.version=x.y.z".
var (
	// name is the name of the compiled software.
	name string //nolint:gochecknoglobals //TODO
	// version is the version of the compiled software.
	version string

	id, _ = os.Hostname() //nolint:gochecknoglobals //TODO
)

type Metadata struct {
	Name    string
	Version string
	ID      string
}

func GetAppMetadata() Metadata {
	return Metadata{
		Name:    name,
		Version: version,
		ID:      id,
	}
}
