package bizfeed

type FeedUseCase struct {
	rss *RSSRepo
}

type RSSRepo interface {
	Parse([]byte) (*Feed, error)
	Get(string, []byte) error
}

func NewFeed(rss *RSSRepo) *FeedUseCase {
	return &FeedUseCase{
		rss,
	}
}

func (f *FeedUseCase) GetFeed() {

}
