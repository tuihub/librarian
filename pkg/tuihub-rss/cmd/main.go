package main

import (
	"context"
	"os"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"
	"github.com/tuihub/librarian/pkg/tuihub-rss/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

// go build -ldflags "-X main.version=x.y.z".
var (
	// version is the version of the compiled software.
	version string
)

const (
	rssServerURLPrefix = "RSS_SERVER_URL_PREFIX"
	rssServerHost      = "RSS_SERVER_HOST"
	rssServerPort      = "RSS_SERVER_PORT"
)

func main() {
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/librarian/pkg/tuihub-rss",
			BuildVersion:      version,
			BuildDate:         "",
			Name:              "tuihub-rss",
			Version:           version,
			Description:       "",
		},
		GlobalName: "github.com/tuihub/librarian/pkg/tuihub-rss",
		Region:     "",
		FeatureSummary: &librarian.FeatureSummary{
			FeedSources: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownFeedSource_WELL_KNOWN_FEED_SOURCE_RSS,
					),
					Name:             "RSS",
					Description:      "",
					ConfigJsonSchema: tuihub.MustReflectJSONSchema(new(internal.PullRSSConfig)),
				},
			},
			NotifyDestinations: []*librarian.FeatureFlag{
				{
					Id: tuihub.WellKnownToString(
						librarian.WellKnownFeedSource_WELL_KNOWN_FEED_SOURCE_RSS,
					),
					Name:             "RSS",
					Description:      "",
					ConfigJsonSchema: tuihub.MustReflectJSONSchema(new(internal.ServeRSSConfig)),
					Extra: map[string]string{
						"URLPrefix": os.Getenv(rssServerURLPrefix),
					},
				},
			},
		},
		ContextJsonSchema: nil,
	}
	porterServer, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
		tuihub.WithAsUser(),
	)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	rssServer, err := internal.NewServer(porterServer)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	go func() {
		err = rssServer.Run(os.Getenv(rssServerHost) + ":" + os.Getenv(rssServerPort))
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
	}()
	if err = porterServer.Run(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
