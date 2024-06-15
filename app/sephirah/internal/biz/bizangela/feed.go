package bizangela

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
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
	systemNotify *libmq.Topic[modelangela.SystemNotify],
) *libmq.Topic[modelyesod.PullFeed] {
	return libmq.NewTopic[modelyesod.PullFeed](
		"PullFeed",
		func(ctx context.Context, p *modelyesod.PullFeed) error {
			// Prepare for updating feed pull status
			fc := new(modelyesod.FeedConfig)
			fc.ID = p.InternalID
			fc.LatestPullTime = time.Now()
			fc.LatestPullStatus = modelyesod.FeedConfigPullStatusFailed

			defer func() {
				_ = a.repo.UpdateFeedPullStatus(ctx, fc)
				if p.SystemNotify != nil {
					un := *p.SystemNotify
					if fc.LatestPullMessage == "" {
						un.Notification.Level = modelnetzach.SystemNotificationLevelInfo
					} else if fc.LatestPullStatus == modelyesod.FeedConfigPullStatusSuccess {
						un.Notification.Level = modelnetzach.SystemNotificationLevelWarning
					} else {
						un.Notification.Level = modelnetzach.SystemNotificationLevelError
					}
					un.Notification.Content = fc.LatestPullMessage
					_ = systemNotify.PublishFallsLocalCall(ctx, un)
				}
			}()

			// Check porter availability
			if !a.supv.CheckFeedSource(p.Source) {
				fc.LatestPullMessage = fmt.Sprintf("Pull %s feature not activate", p.Source)
				return nil
			}

			// Pull feed and upsert
			resp, err := a.porter.PullFeed(
				a.supv.CallFeedSource(ctx, p.Source),
				&porter.PullFeedRequest{
					Source:    p.Source,
					ChannelId: p.URL,
				},
			)
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("PullFeed failed: %s", err.Error())
				return err
			}
			feed := modelfeed.NewConverter().FromPBFeed(resp.GetData())
			feed.ID = p.InternalID
			ids, err := a.searcher.NewBatchIDs(ctx, len(feed.Items))
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("Generate IDs failed: %s", err.Error())
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
				fc.LatestPullMessage = fmt.Sprintf("UpsertFeed failed: %s", err.Error())
				return err
			}
			newItemGUIDs, err := a.repo.UpsertFeedItems(ctx, feed.Items, feed.ID)
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("UpsertFeedItems failed: %s", err.Error())
				return err
			}
			fc.LatestPullStatus = modelyesod.FeedConfigPullStatusSuccess

			// Queue ParseFeedItemDigest and NotifyRouter
			newItems := make([]*modelfeed.Item, 0, len(newItemGUIDs))
			for _, item := range feed.Items {
				if slices.Contains(newItemGUIDs, item.GUID) {
					newItems = append(newItems, item)
					err = parse.Publish(ctx, modelangela.ParseFeedItemDigest{ID: item.ID})
				}
			}
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("Queue ParseFeedItemDigest failed: %s", err.Error())
			}
			if len(newItems) > 0 {
				err = notify.Publish(ctx, modelangela.NotifyRouter{
					FeedID:   feed.ID,
					Messages: newItems,
				})
				if err != nil {
					fc.LatestPullMessage = fmt.Sprintf("Queue NotifyRouter failed: %s", err.Error())
				}
			}
			return nil
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
