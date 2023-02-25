package feed

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/app/porter/internal/client/feed/converter"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"github.com/gocolly/colly/v2"
	"github.com/google/wire"
	"github.com/mmcdole/gofeed"
)

var ProviderSet = wire.NewSet(NewRSSRepo)

type rssRepo struct {
	c *colly.Collector
}

func NewRSSRepo(c *colly.Collector) (bizfeed.RSSRepo, error) {
	return &rssRepo{
		c: c,
	}, nil
}

func (s *rssRepo) Parse(data string) (*modelfeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(data)
	if err != nil {
		return nil, err
	}
	res := converter.NewConverter().ToPBFeed(feed)
	return res, nil
}

func (s *rssRepo) Get(url string) (string, error) {
	c := s.c.Clone()
	var err error
	var data string
	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			err = fmt.Errorf("request %s failed with code %d", url, response.StatusCode)
		}
		data = string(response.Body)
	})
	if err2 := c.Visit(url); err2 != nil {
		return "", err2
	}
	c.Wait()
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", errors.New("empty response")
	}
	return data, nil
}
