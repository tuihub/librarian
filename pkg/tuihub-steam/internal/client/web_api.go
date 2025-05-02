package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-steam/internal/model"

	"github.com/gocolly/colly/v2"
	"github.com/google/go-querystring/query"
)

type WebAPI struct {
	key string
	c   *colly.Collector
}

func NewWebAPI(c *colly.Collector, apiKey string) (*WebAPI, error) {
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "*api.steampowered.com*",
		Delay:        time.Second,
		RandomDelay:  0,
		Parallelism:  1,
	})
	if err != nil {
		return nil, err
	}
	return &WebAPI{
		key: apiKey,
		c:   c,
	}, nil
}

func (s *WebAPI) GetPlayerSummary(
	ctx context.Context,
	req model.GetPlayerSummariesRequest,
) (*model.PlayerSummary, error) {
	res := new(model.GetPlayerSummariesResponse)
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "ISteamUser/GetPlayerSummaries/v2", reqStr, res)
	if err != nil {
		return nil, err
	}
	if len(res.Response.Players) == 0 {
		return nil, errors.New("empty response")
	}

	return &res.Response.Players[0], nil
}

func (s *WebAPI) GetOwnedGames(ctx context.Context, req model.GetOwnedGamesRequest) (*model.OwnedGames, error) {
	res := new(model.GetOwnedGamesResponse)
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "IPlayerService/GetOwnedGames/v1", reqStr, res)
	if err != nil {
		return nil, err
	}

	return &res.Response, nil
}

func (s *WebAPI) GetAppList(ctx context.Context, req model.GetAppListRequest) (*model.AppList, error) {
	res := new(model.GetAppListResponse)
	reqStr, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	err = s.Get(ctx, "ISteamApps/GetAppList/v2", reqStr, res)
	if err != nil {
		return nil, err
	}

	return &res.AppList, nil
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
		err = json.Unmarshal(response.Body, data)
	})
	if err2 := c.Visit(u); err2 != nil {
		return err2
	}
	c.Wait()
	return err
}
