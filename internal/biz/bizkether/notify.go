package bizkether

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/tuihub/librarian/internal/data"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	"github.com/tuihub/librarian/internal/model/modelkether"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/service/sephirah/converter"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func NewNotifyRouterTopic( //nolint:gocognit // TODO
	a *KetherBase,
	flowMap *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	feedToFlowMap *libcache.Map[model.InternalID, modelkether.FeedToNotifyFlowValue],
	push *libmq.Topic[modelkether.NotifyPush],
) *libmq.Topic[modelkether.NotifyRouter] {
	return libmq.NewTopic[modelkether.NotifyRouter](
		"NotifyRouter",
		func(ctx context.Context, r *modelkether.NotifyRouter) error {
			flowIDs, err := feedToFlowMap.Get(ctx, r.FeedID)
			if err != nil {
				return err
			}
			if flowIDs == nil {
				return errors.New("nil result from feedToFlowMap")
			}
			for _, flowID := range *flowIDs {
				var flow *modelnetzach.NotifyFlow
				flow, err = flowMap.Get(ctx, flowID)
				if err != nil {
					return err
				}
				if flow.Status != modelnetzach.NotifyFlowStatusActive {
					continue
				}
				var messages []*modelfeed.Item
				for _, source := range flow.Sources {
					if source.SourceID == r.FeedID {
						messages = applyFilter(r.Messages, source.Filter.IncludeKeywords, source.Filter.ExcludeKeywords)
					}
				}
				if len(messages) == 0 {
					continue
				}
				itemIDs := make([]model.InternalID, 0, len(messages))
				for _, item := range messages {
					if item == nil {
						continue
					}
					itemIDs = append(itemIDs, item.ID)
				}
				for _, target := range flow.Targets {
					if target == nil {
						continue
					}
					err = push.Publish(ctx, modelkether.NotifyPush{
						Target:   *target,
						Messages: applyFilter(messages, target.Filter.IncludeKeywords, target.Filter.ExcludeKeywords),
					})
					if err != nil {
						return err
					}
					err = a.repo.AddFeedItemsToCollection(ctx, flowID, itemIDs)
					if err != nil {
						return err
					}
				}
			}
			return nil
		},
	)
}

func NewNotifyPushTopic(
	a *KetherBase,
	targetMap *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
) *libmq.Topic[modelkether.NotifyPush] {
	return libmq.NewTopic[modelkether.NotifyPush](
		"NotifyPush",
		func(ctx context.Context, p *modelkether.NotifyPush) error {
			target, err := targetMap.Get(ctx, p.Target.TargetID)
			if err != nil {
				return err
			}
			if !a.supv.HasNotifyDestination(target.Destination) {
				return nil
			}
			if target.Status != modelnetzach.NotifyTargetStatusActive {
				return nil
			}
			_, err = a.porter.PushFeedItems(
				a.supv.WithNotifyDestination(ctx, target.Destination),
				&porter.PushFeedItemsRequest{
					Destination: converter.ToPBFeatureRequest(target.Destination),
					Items:       converter.ToPBFeedItemList(p.Messages),
				},
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

// NewFeedToNotifyFlowCache Cache-Aside Pattern.
func NewFeedToNotifyFlowCache(
	n *data.NetzachRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelkether.FeedToNotifyFlowValue] {
	return libcache.NewMap[model.InternalID, modelkether.FeedToNotifyFlowValue](
		store,
		"FeedToNotifyFlow",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, id model.InternalID) (*modelkether.FeedToNotifyFlowValue, error) {
			res, err := n.GetNotifyFlowIDsWithFeed(ctx, id)
			if err != nil {
				return nil, err
			}
			return (*modelkether.FeedToNotifyFlowValue)(&res), nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

// NewNotifyFlowCache Cache-Aside Pattern.
func NewNotifyFlowCache(
	n *data.NetzachRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelnetzach.NotifyFlow] {
	return libcache.NewMap[model.InternalID, modelnetzach.NotifyFlow](
		store,
		"NotifyFlows",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyFlow, error) {
			res, err := n.GetNotifyFlow(ctx, id)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

// NewNotifyTargetCache Cache-Aside Pattern.
func NewNotifyTargetCache(
	n *data.NetzachRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelnetzach.NotifyTarget] {
	return libcache.NewMap[model.InternalID, modelnetzach.NotifyTarget](
		store,
		"NotifyTargets",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyTarget, error) {
			res, err := n.GetNotifyTarget(ctx, id)
			if err != nil {
				return nil, err
			}
			return res, nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

func applyFilter( //nolint:gocognit // TODO
	items []*modelfeed.Item,
	includes []string,
	excludes []string,
) []*modelfeed.Item {
	if len(includes) == 0 && len(excludes) == 0 {
		return items
	}
	res := make([]*modelfeed.Item, 0, len(items))
	for _, item := range items {
		if item == nil {
			continue
		}
		excluded := false
		if len(excludes) > 0 {
			for _, exclude := range excludes {
				if strings.Contains(item.Title, exclude) ||
					strings.Contains(item.Description, exclude) ||
					strings.Contains(item.Content, exclude) {
					excluded = true
					break
				}
			}
		}
		if excluded {
			continue
		}
		included := true
		if len(includes) > 0 {
			for _, include := range includes {
				if !(strings.Contains(item.Title, include) ||
					strings.Contains(item.Description, include) ||
					strings.Contains(item.Content, include)) {
					included = false
					continue
				}
			}
		}
		if included {
			res = append(res, item)
		}
	}
	return res
}
