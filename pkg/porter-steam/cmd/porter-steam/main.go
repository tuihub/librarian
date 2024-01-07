package main

import (
	"os"

	portersdk "github.com/tuihub/librarian/pkg/porter-sdk"
	"github.com/tuihub/librarian/pkg/porter-steam/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	apiKey, exist := os.LookupEnv("STEAM_API_KEY")
	if !exist || apiKey == "" {
		panic("STEAM_API_KEY is required")
	}
	config := portersdk.PorterConfig{
		Name:       "porter-steam",
		Version:    version,
		GlobalName: "github.com/tuihub/librarian/pkg/porter-steam",
		FeatureSummary: &porter.PorterFeatureSummary{
			SupportedAccounts: []*porter.PorterFeatureSummary_Account{
				{
					Platform: portersdk.WellKnownToString(
						librarian.WellKnownAccountPlatform_WELL_KNOWN_ACCOUNT_PLATFORM_STEAM,
					),
					AppRelationTypes: []librarian.AccountAppRelationType{
						librarian.AccountAppRelationType_ACCOUNT_APP_RELATION_TYPE_OWN,
					},
				},
			},
			SupportedAppSources: []string{
				portersdk.WellKnownToString(librarian.WellKnownAppSource_WELL_KNOWN_APP_SOURCE_STEAM),
			},
			SupportedFeedSources:        nil,
			SupportedNotifyDestinations: nil,
		},
		Server: portersdk.ServerConfig{
			Network: "",
			Addr:    "",
			Timeout: nil,
		},
	}
	server, err := portersdk.New(config, internal.NewHandler(apiKey))
	if err != nil {
		return
	}
	if err = server.Run(); err != nil {
		return
	}
}
