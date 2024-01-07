package internal

import (
	"context"

	"github.com/tuihub/librarian/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/muzhou233/go-favicon"
)

type Handler struct {
	rss     RSS
	favicon *favicon.Finder
}

func NewHandler() *Handler {
	return &Handler{
		NewRSS(),
		favicon.New(favicon.IgnoreManifest),
	}
}

func (h Handler) PullAccount(ctx context.Context, req *porter.PullAccountRequest) (
	*porter.PullAccountResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullApp(ctx context.Context, req *porter.PullAppRequest) (
	*porter.PullAppResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullAccountAppRelation(ctx context.Context, req *porter.PullAccountAppRelationRequest) (
	*porter.PullAccountAppRelationResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}

func (h Handler) PullFeed(ctx context.Context, req *porter.PullFeedRequest) (
	*porter.PullFeedResponse, error) {
	data, err := h.rss.Get(req.GetChannelId())
	if err != nil {
		return nil, err
	}
	feed, err := h.rss.Parse(data)
	if err != nil {
		return nil, err
	}
	if len(feed.Link) > 0 {
		if icons, err1 := h.favicon.Find(feed.Link); err1 == nil && len(icons) > 0 {
			for _, icon := range icons {
				if icon.Height > 0 && icon.Width > 0 {
					feed.Image = &modelfeed.Image{
						URL:   icons[0].URL,
						Title: "",
					}
					break
				}
			}
		}
	}
	res := modelfeed.NewConverter().ToPBFeed(feed)
	return &porter.PullFeedResponse{Data: res}, nil
}

func (h Handler) PushFeedItems(ctx context.Context, req *porter.PushFeedItemsRequest) (
	*porter.PushFeedItemsResponse, error) {
	return nil, errors.BadRequest("not supported", "")
}
