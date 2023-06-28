package bizangela

import (
	"context"
	"net/url"
	"sort"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"

	"golang.org/x/exp/slices"
)

func NewPullFeedTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	notify *libmq.Topic[modelangela.NotifyRouter],
) *libmq.Topic[modelyesod.PullFeed] {
	return libmq.NewTopic[modelyesod.PullFeed](
		"PullFeed",
		func(ctx context.Context, p *modelyesod.PullFeed) error {
			resp, err := a.porter.PullFeed(ctx, &porter.PullFeedRequest{
				Source:    porter.FeedSource_FEED_SOURCE_COMMON,
				ChannelId: p.URL,
			})
			if err != nil {
				return err
			}
			feed := modelfeed.NewConverter().FromPBFeed(resp.GetData())
			feed.ID = p.InternalID
			ids, err := a.searcher.NewBatchIDs(ctx, &searcher.NewBatchIDsRequest{Num: int32(len(feed.Items))})
			if err != nil {
				return err
			}
			for i, item := range feed.Items {
				item.ID = converter.ToBizInternalID(ids.GetIds()[i])
				// generate publish_platform
				if len(item.Link) > 0 {
					var linkParsed *url.URL
					linkParsed, err = url.Parse(item.Link)
					if err == nil {
						item.PublishPlatform = linkParsed.Host
					}
				}
				// generate published_parsed
				if item.PublishedParsed == nil {
					t := time.Now()
					item.PublishedParsed = &t
				}
			}
			sort.Sort(feed)
			if err = a.repo.UpsertFeed(ctx, feed); err != nil {
				return err
			}
			newItemGUIDs, err := a.repo.UpsertFeedItems(ctx, feed.Items, feed.ID)
			if err != nil {
				return err
			}
			newItems := make([]*modelfeed.Item, 0, len(newItemGUIDs))
			for _, item := range feed.Items {
				if slices.Contains(newItemGUIDs, item.GUID) {
					newItems = append(newItems, item)
				}
			}
			if len(newItems) > 0 {
				err = notify.Publish(ctx, modelangela.NotifyRouter{
					FeedID:   feed.ID,
					Messages: newItems,
				})
			}
			return err
		},
	)
}
