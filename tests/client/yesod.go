package main

import (
	"context"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
)

const feedURL = "https://github.com/TuiHub/Librarian/releases.atom"

func (c *Client) TestYesod(ctx context.Context) {
	var feedConfigID int64
	if resp, err := c.cli.CreateFeedConfig(ctx, &pb.CreateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id:             nil,
			FeedUrl:        feedURL,
			AuthorAccount:  nil,
			Source:         pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON,
			Status:         pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE,
			PullInterval:   durationpb.New(time.Hour),
			LatestPullTime: nil,
		},
	}); err != nil {
		panic(err)
	} else {
		feedConfigID = resp.GetId().GetId()
	}
	time.Sleep(time.Minute * 2) //nolint:gomnd // waiting
	if _, err := c.cli.UpdateFeedConfig(ctx, &pb.UpdateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id: &librarian.InternalID{
				Id: feedConfigID,
			},
			FeedUrl:        feedURL,
			AuthorAccount:  nil,
			Source:         pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON,
			Status:         pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE,
			PullInterval:   durationpb.New(time.Hour),
			LatestPullTime: nil,
		},
	}); err != nil {
		panic(err)
	}
}
