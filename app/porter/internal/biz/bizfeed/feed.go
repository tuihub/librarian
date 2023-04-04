package bizfeed

import (
	"context"
	"errors"
	"strconv"

	"github.com/tuihub/librarian/app/porter/internal/client/telegram"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type FeedUseCase struct {
	rss RSSRepo
}

type RSSRepo interface {
	Parse(string) (*modelfeed.Feed, error)
	Get(string) (string, error)
}

func NewFeed(rss RSSRepo) *FeedUseCase {
	return &FeedUseCase{
		rss,
	}
}

type FeedDestination int

const (
	FeedDestinationUnspecified FeedDestination = iota
	FeedDestinationTelegram
)

func (f *FeedUseCase) GetFeed(ctx context.Context, url string) (*modelfeed.Feed, error) {
	data, err := f.rss.Get(url)
	if err != nil {
		return nil, err
	}
	feed, err := f.rss.Parse(data)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (f *FeedUseCase) PushFeedItems(ctx context.Context, dest FeedDestination, items []*modelfeed.Item,
	channelID, token string) error {
	switch dest {
	case FeedDestinationUnspecified:
		return errors.New("invalid destination")
	case FeedDestinationTelegram:
		messages := make(map[string]string)
		for _, item := range items {
			messages[item.Title] = item.Link
		}
		channelIDInt64, err := strconv.ParseInt(channelID, 10, 64)
		if err != nil {
			return errors.New("invalid channel_id")
		}
		err = telegram.SendBatch(ctx, token, channelIDInt64, messages)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported destination")
	}
	return nil
}
