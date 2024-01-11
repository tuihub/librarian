package bizangela

import (
	"context"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"
)

func NewPullFeedTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	notify *libmq.Topic[modelangela.NotifyRouter],
	parse *libmq.Topic[modelangela.ParseFeedItemDigest],
) *libmq.Topic[modelyesod.PullFeed] {
	return libmq.NewTopic[modelyesod.PullFeed](
		"PullFeed",
		func(ctx context.Context, p *modelyesod.PullFeed) error {
			if !a.supv.CheckFeedSource(p.Source) {
				return nil
			}
			resp, err := a.porter.PullFeed(
				a.supv.CallFeedSource(ctx, p.Source),
				&porter.PullFeedRequest{
					Source:    p.Source,
					ChannelId: p.URL,
				},
			)
			if err != nil {
				return err
			}
			feed := modelfeed.NewConverter().FromPBFeed(resp.GetData())
			feed.ID = p.InternalID
			ids, err := a.searcher.NewBatchIDs(ctx, len(feed.Items))
			if err != nil {
				return err
			}
			for i, item := range feed.Items {
				item.ID = ids[i]
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
					_ = parse.Publish(ctx, modelangela.ParseFeedItemDigest{ID: item.ID})
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

func NewParseFeedItemDigestTopic( //nolint:gocognit // TODO
	a *AngelaBase,
) *libmq.Topic[modelangela.ParseFeedItemDigest] {
	return libmq.NewTopic[modelangela.ParseFeedItemDigest](
		"ParseFeedItemDigest",
		func(ctx context.Context, p *modelangela.ParseFeedItemDigest) error {
			const maxImgNum = 9
			const maxDescLen = 128
			item, err := a.repo.GetFeedItem(ctx, p.ID)
			if err != nil {
				return err
			}
			content := item.Content
			if len(content) == 0 {
				content = item.Description
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
			if err != nil {
				return err
			}
			digestDesc := doc.Text()
			digestDesc = strings.ReplaceAll(digestDesc, " ", "")
			digestDesc = strings.ReplaceAll(digestDesc, "\n", "")
			digestDescRune := []rune(digestDesc)
			if len(digestDescRune) > maxDescLen {
				digestDescRune = digestDescRune[:maxDescLen]
			}
			digestDesc = string(digestDescRune)
			item.DigestDescription = digestDesc

			for i, n := range doc.Find("img").Nodes {
				if i == maxImgNum {
					break
				}
				image := new(modelfeed.Image)
				for _, attr := range n.Attr {
					if attr.Key == "src" {
						image.URL = attr.Val
					}
					if attr.Key == "alt" {
						image.Title = attr.Val
					}
				}
				item.DigestImages = append(item.DigestImages, image)
			}
			err = a.repo.UpdateFeedItemDigest(ctx, item)
			if err != nil {
				return err
			}
			return nil
		},
	)
}
