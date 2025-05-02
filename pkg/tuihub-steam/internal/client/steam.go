package client

import (
	"github.com/tuihub/librarian/pkg/tuihub-go/logger"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

type Steam struct {
	*StoreAPI
	*WebAPI
}

func NewSteam(apiKey string) (*Steam, error) {
	c := NewColly()
	s, err := NewStoreAPI(c)
	if err != nil {
		return nil, err
	}
	w, err := NewWebAPI(c, apiKey)
	if err != nil {
		return nil, err
	}
	return &Steam{
		StoreAPI: s,
		WebAPI:   w,
	}, nil
}

func NewColly() *colly.Collector {
	return colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{
			Output: logger.NewWriter(log.LevelDebug),
			Prefix: "[colly]",
			Flag:   0,
		}),
		colly.AllowURLRevisit(),
	)
}
