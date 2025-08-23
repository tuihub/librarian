//nolint:exhaustruct // no need
package tuihubsteam

import (
	"context"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal"
	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPorter(version string) (*tuihub.Porter, error) {
	contextSchema := tuihub.MustReflectJSONSchema(new(model.PorterContext))

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
					RequireContext:   true,
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
					RequireContext:   true,
				},
			},
		},
		ContextJsonSchema: &contextSchema,
	}
	return tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
	)
}
