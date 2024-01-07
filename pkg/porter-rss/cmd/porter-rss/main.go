package main

import (
	"github.com/tuihub/librarian/pkg/porter-rss/internal"
	portersdk "github.com/tuihub/librarian/pkg/porter-sdk"
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
		Name:       "porter-rss",
		Version:    version,
		GlobalName: "github.com/tuihub/librarian/pkg/porter-rss",
		FeatureSummary: &porter.PorterFeatureSummary{
			SupportedAccounts:   nil,
			SupportedAppSources: nil,
			SupportedFeedSources: []string{
				portersdk.WellKnownToString(librarian.WellKnownFeedSource_WELL_KNOWN_FEED_SOURCE_RSS),
			},
			SupportedNotifyDestinations: nil,
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
