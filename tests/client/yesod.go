package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
)

const feedURL = "https://github.com/TuiHub/Librarian/releases.atom"

func (c *Client) TestYesod(ctx context.Context) {
	var feedConfigID, feedItemID int64
	if resp, err := c.cli.CreateFeedConfig(ctx, &pb.CreateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id:          nil,
			Name:        "",
			Description: feedURL,
			Source: &librarian.FeatureRequest{
				Id:         "rss",
				Region:     "",
				ConfigJson: "",
				ContextId:  nil,
			},
			Status:            pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND,
			PullInterval:      durationpb.New(time.Hour),
			Category:          "",
			HideItems:         false,
			LatestPullTime:    nil,
			LatestPullStatus:  nil,
			LatestPullMessage: nil,
			ActionSets:        nil,
		},
	}); err != nil {
		log.Fatal(err)
	} else {
		feedConfigID = resp.GetId().GetId()
	}
	if _, err := c.cli.UpdateFeedConfig(ctx, &pb.UpdateFeedConfigRequest{
		Config: &pb.FeedConfig{
			Id: &librarian.InternalID{
				Id: feedConfigID,
			},
			Name:        "",
			Description: feedURL,
			Source: &librarian.FeatureRequest{
				Id:         "rss",
				Region:     "",
				ConfigJson: "",
				ContextId:  nil,
			},
			ActionSets:        nil,
			Status:            pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE,
			PullInterval:      durationpb.New(time.Hour),
			Category:          "",
			HideItems:         false,
			LatestPullTime:    nil,
			LatestPullStatus:  nil,
			LatestPullMessage: nil,
		},
	}); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Minute * 2) //nolint:mnd // waiting
	if resp, err := c.cli.ListFeedConfigs(ctx, &pb.ListFeedConfigsRequest{
		Paging:         defaultPaging,
		IdFilter:       nil,
		AuthorIdFilter: nil,
		SourceFilter:   nil,
		StatusFilter:   nil,
		CategoryFilter: nil,
	}); err != nil {
		return
	} else if resp.GetPaging().GetTotalSize() != 1 ||
		len(resp.GetFeedsWithConfig()) != 1 ||
		resp.GetFeedsWithConfig()[0].GetConfig().GetId().GetId() != feedConfigID ||
		resp.GetFeedsWithConfig()[0].GetFeed().GetId().GetId() != feedConfigID {
		log.Fatal(fmt.Sprintf("unexpected ListFeeds response, %+v", resp))
	}
	if resp, err := c.cli.ListFeedItems(ctx, &pb.ListFeedItemsRequest{
		Paging:                defaultPaging,
		FeedIdFilter:          nil,
		AuthorFilter:          nil,
		PublishPlatformFilter: nil,
		PublishTimeRange:      nil,
		CategoryFilter:        nil,
	}); err != nil {
		log.Fatal(err)
	} else if resp.GetPaging().GetTotalSize() < 1 ||
		len(resp.GetItems()) < 1 ||
		resp.GetItems()[0].GetFeedId().GetId() != feedConfigID {
		log.Fatal(fmt.Sprintf("unexpected ListFeedItems response, %+v, %v", resp, feedConfigID))
	} else {
		feedItemID = resp.GetItems()[0].GetItemId().GetId()
	}
	if resp, err := c.cli.GetFeedItem(ctx, &pb.GetFeedItemRequest{
		Id: &librarian.InternalID{Id: feedItemID},
	}); err != nil {
		log.Fatal(err)
	} else if resp.GetItem().GetId().GetId() != feedItemID {
		log.Fatal(fmt.Sprintf("unexpected GetFeedItem response, %+v", resp))
	}
}
