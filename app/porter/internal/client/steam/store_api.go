package steam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tuihub/librarian/app/porter/internal/client/steam/model"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/google/go-querystring/query"
)

type StoreAPI struct {
	c *colly.Collector
}

func NewStoreAPI() (*StoreAPI, error) {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{
			Output: logger.NewWriter(),
			Prefix: "[colly]",
			Flag:   0,
		}),
		colly.AllowURLRevisit(),
	)
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*store.steampowered.com*",
		Parallelism: 1,
		Delay:       5 * time.Second, //nolint:gomnd // This API is now rate limited to 200 requests per 5 minutes
	})
	if err != nil {
		return nil, err
	}
	return &StoreAPI{
		c: c,
	}, nil
}

func (s *StoreAPI) GetAppDetails(
	ctx context.Context,
	req model.GetAppDetailsRequest,
) (map[string]model.AppDetailsBasic, error) {
	res := map[string]model.AppDetailsBasic{}
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "/appdetails", reqStr, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StoreAPI) Get(ctx context.Context, path string, query url.Values, data interface{}) error {
	u := "https://store.steampowered.com/api" + path + "?" + query.Encode()
	c := s.c.Clone()
	var err error
	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != http.StatusOK {
			err = fmt.Errorf("request %s failed with code %d", path, response.StatusCode)
		}
		err = libcodec.Unmarshal(libcodec.JSON, response.Body, data)
	})
	if err2 := c.Visit(u); err2 != nil {
		return err2
	}
	c.Wait()
	return err
}
