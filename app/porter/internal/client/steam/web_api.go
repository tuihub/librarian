package steam

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tuihub/librarian/app/porter/internal/client/steam/model"
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libcodec"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/google/go-querystring/query"
)

type WebAPI struct {
	key string
	c   *colly.Collector
}

func NewWebAPI(config *conf.Porter_Data) (*WebAPI, error) {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{
			Output: logger.NewWriter(),
			Prefix: "[colly]",
			Flag:   0,
		}),
		colly.AllowURLRevisit(),
	)
	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*api.steampowered.com*",
		Parallelism: 1,
		Delay:       time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &WebAPI{
		key: config.SteamApiKey,
		c:   c,
	}, nil
}

func (s *WebAPI) GetPlayerSummary(
	ctx context.Context,
	req model.GetPlayerSummariesRequest,
) (*model.PlayerSummary, error) {
	res := model.GetPlayerSummariesResponse{}
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "ISteamUser/GetPlayerSummaries/v2", reqStr, &res)
	if err != nil {
		return nil, err
	}

	return &res.Response.Players[0], nil
}

func (s *WebAPI) GetOwnedGames(ctx context.Context, req model.GetOwnedGamesRequest) (*model.OwnedGames, error) {
	res := model.GetOwnedGamesResponse{}
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "IPlayerService/GetOwnedGames/v1", reqStr, &res)
	if err != nil {
		return nil, err
	}

	return &res.Response, nil
}

func (s *WebAPI) Get(ctx context.Context, path string, query url.Values, data interface{}) error {
	query.Set("format", "json")
	query.Set("key", s.key)
	u := "https://api.steampowered.com/" + path + "?" + query.Encode()
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
