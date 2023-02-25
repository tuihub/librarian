package bizfeed

import (
	"context"

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
