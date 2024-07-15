package bizangela

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizyesod"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelyesod"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"

	"golang.org/x/exp/slices"
)

func NewPullFeedTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	parse *libmq.Topic[modelangela.FeedItemPostprocess],
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
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
				fc.LatestPullMessage = fmt.Sprintf("Pull %s feature not activate", p.Source.ID)
				return nil
			}

			// Pull feed and upsert
			resp, err := a.porter.PullFeed(
				a.supv.CallFeedSource(ctx, p.Source),
				&porter.PullFeedRequest{
					Source: converter.ToPBFeatureRequest(p.Source),
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
				// generate published_parsed, used for sorting
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
			newItemGUIDs, err := a.repo.CheckNewFeedItems(ctx, feed.Items, feed.ID)
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("UpsertFeedItems failed: %s", err.Error())
				return err
			}
			fc.LatestPullStatus = modelyesod.FeedConfigPullStatusSuccess

			// Queue FeedItemPostprocess
			for _, item := range feed.Items {
				if slices.Contains(newItemGUIDs, item.GUID) {
					err = parse.Publish(ctx, modelangela.FeedItemPostprocess{
						FeedID:       feed.ID,
						Item:         item,
						SystemNotify: p.SystemNotify,
					})
				}
			}
			if err != nil {
				fc.LatestPullMessage = fmt.Sprintf("Queue FeedItemPostprocess failed: %s", err.Error())
			}
			return nil
		},
	)
}

func NewFeedItemPostprocessTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	notify *libmq.Topic[modelangela.NotifyRouter],
	systemNotify *libmq.Topic[modelnetzach.SystemNotify],
) *libmq.Topic[modelangela.FeedItemPostprocess] {
	return libmq.NewTopic[modelangela.FeedItemPostprocess](
		"FeedItemPostprocess",
		func(ctx context.Context, p *modelangela.FeedItemPostprocess) error {
			notifyMsg := p.SystemNotify
			if notifyMsg == nil {
				notifyMsg = new(modelnetzach.SystemNotify)
			}
			notifyMsg.Notification.Content = ""
			notifyMsg.Notification.Level = modelnetzach.SystemNotificationLevelError
			defer func() {
				if p.SystemNotify != nil && len(p.SystemNotify.Notification.Content) > 0 {
					p.SystemNotify.Notification.Content = fmt.Sprintf(
						"[%d] %s",
						p.Item.ID,
						p.SystemNotify.Notification.Content,
					)
					_ = systemNotify.PublishFallsLocalCall(ctx, *p.SystemNotify)
				}
			}()

			item := p.Item
			actionSets, err := a.repo.GetFeedActions(ctx, p.FeedID)
			if err != nil {
				notifyMsg.Notification.Content = fmt.Sprintf("GetFeedActions failed: %s", err.Error())
				return err
			}
			builtin := bizyesod.GetBuiltinActionMap(ctx)
			item, err = bizyesod.RequiredStartAction(ctx, item)
			if err != nil {
				notifyMsg.Notification.Content = fmt.Sprintf("RequiredStartAction failed: %s", err.Error())
				return err
			}
			for _, actions := range actionSets {
				for _, action := range actions.Actions {
					if err != nil {
						notifyMsg.Notification.Content = fmt.Sprintf("%s Unmarshal failed: %s", action.ID, err.Error())
						return err
					}
					if f, ok := builtin[action.ID]; ok { //nolint:nestif // TODO
						item, err = f(ctx, action, item)
						if err != nil {
							notifyMsg.Notification.Content = fmt.Sprintf("%s Exec failed: %s", action.ID, err.Error())
							return err
						}
						if item == nil {
							notifyMsg.Notification.Content = fmt.Sprintf("%s Filtered", action.ID)
							notifyMsg.Notification.Level = modelnetzach.SystemNotificationLevelWarning
							return nil
						}
					} else if a.supv.CheckFeedItemAction(action) {
						var resp *porter.ExecFeedItemActionResponse
						resp, err = a.porter.ExecFeedItemAction(
							a.supv.CallFeedItemAction(ctx, action),
							&porter.ExecFeedItemActionRequest{
								Action: converter.ToPBFeatureRequest(action),
								Item:   converter.ToPBFeedItem(item),
							},
						)
						if err != nil {
							notifyMsg.Notification.Content = fmt.Sprintf("%s Exec failed: %s", action.ID, err.Error())
							return err
						}
						if resp.GetItem() != nil {
							notifyMsg.Notification.Content = fmt.Sprintf("%s Filtered", action.ID)
							notifyMsg.Notification.Level = modelnetzach.SystemNotificationLevelWarning
							return nil
						}
						item = converter.ToBizFeedItem(resp.GetItem())
					}
				}
			}
			item, err = bizyesod.RequiredEndAction(ctx, item)
			if err != nil {
				notifyMsg.Notification.Content = fmt.Sprintf("RequiredEndAction failed: %s", err.Error())
				return err
			}
			err = a.repo.UpsertFeedItems(ctx, []*modelfeed.Item{item}, p.FeedID)
			if err != nil {
				notifyMsg.Notification.Content = fmt.Sprintf("UpsertFeedItems failed: %s", err.Error())
				return err
			}
			_ = notify.Publish(ctx, modelangela.NotifyRouter{
				FeedID:   p.FeedID,
				Messages: []*modelfeed.Item{item},
			})
			return nil
		},
	)
}
