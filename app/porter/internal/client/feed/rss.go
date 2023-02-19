package feed

import (
	"fmt"
	"net/http"

	"github.com/tuihub/librarian/app/porter/internal/biz/bizfeed"
	"github.com/tuihub/librarian/internal/lib/libcodec"

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

func (s *rssRepo) Parse(data []byte) (*bizfeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(data))
	if err != nil {
		return nil, err
	}
	feedJSON, err := libcodec.Marshal(libcodec.JSON, feed)
	if err != nil {
		return nil, err
	}
	res := new(bizfeed.Feed)
	err = libcodec.Unmarshal(libcodec.JSON, feedJSON, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *rssRepo) Get(url string, data []byte) error {
	c := s.c.Clone()
	var err error
	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			err = fmt.Errorf("request %s failed with code %d", url, response.StatusCode)
		}
		data = response.Body
	})
	if err2 := c.Visit(url); err2 != nil {
		return err2
	}
	c.Wait()
	return err
}
