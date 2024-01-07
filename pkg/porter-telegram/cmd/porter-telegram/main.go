package main

import (
	portersdk "github.com/tuihub/librarian/pkg/porter-sdk"
	"github.com/tuihub/librarian/pkg/porter-telegram/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	config := portersdk.PorterConfig{
		Name:       "porter-telegram",
		Version:    version,
		GlobalName: "github.com/tuihub/librarian/pkg/porter-telegram",
		FeatureSummary: &porter.PorterFeatureSummary{
			SupportedAccounts:    nil,
			SupportedAppSources:  nil,
			SupportedFeedSources: nil,
			SupportedNotifyDestinations: []string{
				portersdk.WellKnownToString(
					librarian.WellKnownNotifyDestination_WELL_KNOWN_NOTIFY_DESTINATION_TELEGRAM,
				),
			},
		},
		Server: portersdk.ServerConfig{
			Network: "",
			Addr:    "",
			Timeout: nil,
		},
	}
	server, err := portersdk.New(config, internal.NewHandler())
	if err != nil {
		return
	}
	if err = server.Run(); err != nil {
		return
	}
}
