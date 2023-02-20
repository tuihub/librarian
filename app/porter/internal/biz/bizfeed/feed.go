package bizfeed

import "context"

type FeedUseCase struct {
	rss RSSRepo
}

type RSSRepo interface {
	Parse(string) (*Feed, error)
	Get(string) (string, error)
}

func NewFeed(rss RSSRepo) *FeedUseCase {
	return &FeedUseCase{
		rss,
	}
}

func (f *FeedUseCase) GetFeed(ctx context.Context, url string) (*Feed, error) {
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
