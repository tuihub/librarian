//nolint:exhaustruct // no need
package tuihubtelegram

import (
	"context"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-telegram/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func NewPorter(version string) (*tuihub.Porter, error) {
	contextSchema := tuihub.MustReflectJSONSchema(new(internal.PorterContext))
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/librarian",
			BuildVersion:      version,
			BuildDate:         "",
			Name:              "tuihub-telegram",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/librarian/pkg/tuihub-telegram",
		Region:     "",
		FeatureSummary: &librarian.FeatureSummary{
			NotifyDestinations: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownNotifyDestination_WELL_KNOWN_NOTIFY_DESTINATION_TELEGRAM,
					),
					Name:             "Telegram",
					Description:      "",
					ConfigJsonSchema: tuihub.MustReflectJSONSchema(new(internal.PushFeedItems)),
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
