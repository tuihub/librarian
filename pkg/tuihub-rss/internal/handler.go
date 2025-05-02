package internal

import (
	"context"
	"encoding/json"

	"github.com/tuihub/librarian/pkg/tuihub-rss/internal/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/mmcdole/gofeed"
	"github.com/muzhou233/go-favicon"
)

type Handler struct {
	porter.UnimplementedLibrarianPorterServiceServer
	rss     RSS
	favicon *favicon.Finder
}

func NewHandler() *Handler {
	return &Handler{
		porter.UnimplementedLibrarianPorterServiceServer{},
		NewRSS(),
		favicon.New(favicon.IgnoreManifest),
	}
}

func (h Handler) PullFeed(ctx context.Context, req *porter.PullFeedRequest) (
	*porter.PullFeedResponse, error) {
	var config PullRSSConfig
	err := json.Unmarshal([]byte(req.GetSource().GetConfigJson()), &config)
	if err != nil {
		return nil, err
	}
	data, err := h.rss.Get(config.URL)
	if err != nil {
		return nil, err
	}
	feed, err := h.rss.Parse(data)
	if err != nil {
		return nil, err
	}
	if feed.Image == nil && len(feed.Link) > 0 {
		h.findFavicon(feed)
	}
	res := converter.ToPBFeed(feed)
	return &porter.PullFeedResponse{Data: res}, nil
}

func (h Handler) findFavicon(feed *gofeed.Feed) {
	icons, err := h.favicon.Find(feed.Link)
	if err != nil || len(icons) == 0 {
		return
	}
	for _, icon := range icons {
		if icon.Height > 0 && icon.Width > 0 {
			feed.Image = &gofeed.Image{
				URL:   icon.URL,
				Title: "",
			}
			break
		}
	}
	if feed.Image == nil {
		feed.Image = &gofeed.Image{
			URL:   icons[0].URL,
			Title: "",
		}
	}
}
