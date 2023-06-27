package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"google.golang.org/protobuf/types/known/durationpb"
)

const feedURL = "https://github.com/TuiHub/Librarian/releases.atom"

func (c *Client) TestYesod(ctx context.Context) { //nolint:gocognit // no need
	var feedConfigID, feedItemID int64
	if resp, err := c.cli.CreateFeedConfig(ctx, &pb.CreateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id:             nil,
			Name:           "",
			FeedUrl:        feedURL,
			AuthorAccount:  nil,
			Source:         pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON,
			Status:         pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND,
			PullInterval:   durationpb.New(time.Hour),
			Tags:           nil,
			LatestPullTime: nil,
		},
	}); err != nil {
		panic(err)
	} else {
		feedConfigID = resp.GetId().GetId()
	}
	if _, err := c.cli.UpdateFeedConfig(ctx, &pb.UpdateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id: &librarian.InternalID{
				Id: feedConfigID,
			},
			Name:           "",
			FeedUrl:        feedURL,
			AuthorAccount:  nil,
			Source:         pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON,
			Status:         pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE,
			PullInterval:   durationpb.New(time.Hour),
			Tags:           nil,
			LatestPullTime: nil,
		},
	}); err != nil {
		panic(err)
	}
	time.Sleep(time.Minute * 2) //nolint:gomnd // waiting
	if resp, err := c.cli.ListFeedConfigs(ctx, &pb.ListFeedConfigsRequest{
		Paging:         defaultPaging,
		IdFilter:       nil,
		AuthorIdFilter: nil,
		SourceFilter:   nil,
		StatusFilter:   nil,
	}); err != nil {
		return
	} else if resp.GetPaging().GetTotalSize() != 1 ||
		len(resp.GetFeedsWithConfig()) != 1 ||
		resp.GetFeedsWithConfig()[0].GetConfig().GetId().GetId() != feedConfigID ||
		resp.GetFeedsWithConfig()[0].GetFeed().GetId().GetId() != feedConfigID {
		panic(fmt.Sprintf("unexpected ListFeeds response, %+v", resp))
	}
	if resp, err := c.cli.ListFeedItems(ctx, &pb.ListFeedItemsRequest{
		Paging:                defaultPaging,
		FeedIdFilter:          nil,
		AuthorIdFilter:        nil,
		PublishPlatformFilter: nil,
		PublishTimeRange:      nil,
		TagFilter:             nil,
	}); err != nil {
		panic(err)
	} else if resp.GetPaging().GetTotalSize() < 1 ||
		len(resp.GetItems()) < 1 ||
		resp.GetItems()[0].GetFeedId().GetId() != feedConfigID {
		panic(fmt.Sprintf("unexpected ListFeedItems response, %+v, %v", resp, feedConfigID))
	} else {
		feedItemID = resp.GetItems()[0].GetItemId().GetId()
	}
	if resp, err := c.cli.GetFeedItem(ctx, &pb.GetFeedItemRequest{
		Id: &librarian.InternalID{Id: feedItemID},
	}); err != nil {
		panic(err)
	} else if resp.GetItem().GetId().GetId() != feedItemID {
		panic(fmt.Sprintf("unexpected GetFeedItem response, %+v", resp))
	}
}
