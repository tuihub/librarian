package bizyesod

import (
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
)

func ToBizFeedConfig(c *pb.FeedConfig) *FeedConfig {
	return &FeedConfig{
		InternalID:    c.GetId().GetId(),
		FeedURL:       c.GetFeedUrl(),
		AuthorAccount: c.GetAuthorAccount().GetId(),
		Source:        ToBizFeedConfigSource(c.GetSource()),
		Status:        ToBizFeedConfigStatus(c.GetStatus()),
		PullInterval:  c.GetPullInterval().AsTime(),
	}
}

func ToBizFeedConfigSource(s pb.FeedConfigSource) FeedConfigSource {
	switch s {
	case pb.FeedConfigSource_FEED_CONFIG_SOURCE_COMMON:
		return FeedConfigSourceCommon
	default:
		return FeedConfigSourceUnspecified
	}
}

func ToBizFeedConfigStatus(s pb.FeedConfigStatus) FeedConfigStatus {
	switch s {
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_ACTIVE:
		return FeedConfigStatusActive
	case pb.FeedConfigStatus_FEED_CONFIG_STATUS_SUSPEND:
		return FeedConfigStatusSuspend
	default:
		return FeedConfigStatusUnspecified
	}
}
