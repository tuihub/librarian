//nolint:exhaustruct // no need
package tuihubrss

import (
	"context"
	"net"
	"os"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-rss/internal"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

const (
	rssServerURLPrefix = "RSS_SERVER_URL_PREFIX"
	rssServerHost      = "RSS_SERVER_HOST"
	rssServerPort      = "RSS_SERVER_PORT"
)

func NewPorter(version string) (*tuihub.Porter, error) {
	config := &porter.GetPorterInformationResponse{
		BinarySummary: &librarian.PorterBinarySummary{
			SourceCodeAddress: "https://github.com/tuihub/librarian",
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
	rssServer, err := internal.NewServer(nil, net.JoinHostPort(os.Getenv(rssServerHost), os.Getenv(rssServerPort)))
	if err != nil {
		return nil, err
	}
	porterServer, err := tuihub.NewPorter(
		context.Background(),
		config,
		internal.NewHandler(),
		tuihub.WithAsUser(),
		tuihub.WithBackgroundServer(rssServer),
	)
	if err != nil {
		return nil, err
	}
	rssServer.SetPorter(porterServer)

	return porterServer, nil
}
