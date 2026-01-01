//nolint:exhaustruct // no need
package tuihubbangumi

import (
	"context"

	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal"
	"github.com/tuihub/librarian/pkg/tuihub-bangumi/internal/model"
	"github.com/tuihub/librarian/pkg/tuihub-go"
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
			Name:              "tuihub-bangumi",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/librarian/pkg/tuihub-bangumi",
		Region:     "",
		FeatureSummary: &librarian.FeatureSummary{
			AppInfoSources: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownAppInfoSource_WELL_KNOWN_APP_INFO_SOURCE_BANGUMI,
					),
					Name:             "Bangumi",
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
