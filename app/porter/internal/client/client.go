package client

import (
	"github.com/tuihub/librarian/app/porter/internal/client/feed"
	"github.com/tuihub/librarian/app/porter/internal/client/steam"
	"github.com/tuihub/librarian/internal/lib/logger"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewColly, steam.ProviderSet, feed.ProviderSet)

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
