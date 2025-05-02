package main

import (
	"context"
	"os"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-telegram/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

func main() {
	contextSchema := tuihub.MustReflectJSONSchema(new(internal.PorterContext))
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/librarian/pkg/tuihub-telegram",
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
	server, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if err = server.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
