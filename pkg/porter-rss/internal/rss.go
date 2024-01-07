package internal

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tuihub/librarian/logger"
	"github.com/tuihub/librarian/model/modelfeed"
	"github.com/tuihub/librarian/pkg/porter-rss/internal/converter"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/mmcdole/gofeed"
)

type RSS struct {
	c *colly.Collector
}

func NewRSS() RSS {
	return RSS{
		c: newColly(),
	}
}

func newColly() *colly.Collector {
	return colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{
			Output: logger.NewWriter(log.LevelDebug),
			Prefix: "[colly]",
			Flag:   0,
		}),
		colly.AllowURLRevisit(),
	)
}

func (s *RSS) Parse(data string) (*modelfeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(data)
	if err != nil {
		return nil, err
	}
	res := converter.NewConverter().ToPBFeed(feed)
	return res, nil
}

func (s *RSS) Get(url string) (string, error) {
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
