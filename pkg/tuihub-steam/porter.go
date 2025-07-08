//nolint:exhaustruct // no need
package tuihubsteam

import (
	"context"
	"os"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPorter(version string) (*tuihub.Porter, error) {
	apiKey, exist := os.LookupEnv("STEAM_API_KEY")
	if !exist || apiKey == "" {
		logger.Errorf("STEAM_API_KEY environment variable not set")
	}
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/librarian",
			BuildVersion:      version,
			BuildDate:         "",
			Name:              "tuihub-steam",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/librarian/pkg/tuihub-steam",
		Region:     "",
		FeatureSummary: &librarian.FeatureSummary{
			AccountPlatforms: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownAccountPlatform_WELL_KNOWN_ACCOUNT_PLATFORM_STEAM,
					),
					Name:             "Steam",
					Description:      "",
					ConfigJsonSchema: "",
					RequireContext:   false,
				},
			},
			AppInfoSources: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_STEAM,
					),
					Name:             "Steam",
					Description:      "",
					ConfigJsonSchema: "",
					RequireContext:   false,
				},
			},
		},
	}
	return tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(apiKey),
	)
}
