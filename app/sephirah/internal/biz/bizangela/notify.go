package bizangela

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func NewNotifyRouterTopic( //nolint:gocognit // TODO
	a *AngelaBase,
	flowMap *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	feedToFlowMap *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue],
	push *libmq.Topic[modelangela.NotifyPush],
) *libmq.Topic[modelangela.NotifyRouter] {
	return libmq.NewTopic[modelangela.NotifyRouter](
		"NotifyRouter",
		func(ctx context.Context, r *modelangela.NotifyRouter) error {
			flowIDs, err := feedToFlowMap.GetWithFallBack(ctx, r.FeedID, nil)
			if err != nil {
				return err
			}
			if flowIDs == nil {
				return errors.New("nil result from feedToFlowMap")
			}
			for _, flowID := range *flowIDs {
				var flow *modelnetzach.NotifyFlow
				flow, err = flowMap.GetWithFallBack(ctx, flowID, nil)
				if err != nil {
					return err
				}
				if flow.Status != modelnetzach.NotifyFlowStatusActive {
					continue
				}
				var messages []*modelfeed.Item
				for _, source := range flow.Sources {
					if source.FeedConfigID == r.FeedID {
						messages = applyFilter(r.Messages, source.Filter.IncludeKeywords, source.Filter.ExcludeKeywords)
					}
				}
				if len(messages) == 0 {
					continue
				}
				for _, target := range flow.Targets {
					if target == nil {
						continue
					}
					err = push.Publish(ctx, modelangela.NotifyPush{
						Target:   *target,
						Messages: applyFilter(messages, target.Filter.IncludeKeywords, target.Filter.ExcludeKeywords),
					})
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
	a *AngelaBase,
	targetMap *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
) *libmq.Topic[modelangela.NotifyPush] {
	return libmq.NewTopic[modelangela.NotifyPush](
		"NotifyPush",
		func(ctx context.Context, p *modelangela.NotifyPush) error {
			target, err := targetMap.GetWithFallBack(ctx, p.Target.TargetID, nil)
			if err != nil {
				return err
			}
			if !a.supv.CheckNotifyDestination(target.Destination) {
				return nil
			}
			if target.Status != modelnetzach.NotifyTargetStatusActive {
				return nil
			}
			_, err = a.porter.PushFeedItems(
				a.supv.CallNotifyDestination(ctx, target.Destination),
				&porter.PushFeedItemsRequest{
					Destination: target.Destination,
					ChannelId:   p.Target.ChannelID,
					Items:       converter.ToPBFeedItemList(p.Messages),
					Token:       target.Token,
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
	n biznetzach.NetzachRepo,
	store libcache.Store,
) *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue] {
	return libcache.NewMap[model.InternalID, modelangela.FeedToNotifyFlowValue](
		store,
		"FeedToNotifyFlow",
		func(k model.InternalID) string {
			return strconv.FormatInt(int64(k), 10)
		},
		func(ctx context.Context, id model.InternalID) (*modelangela.FeedToNotifyFlowValue, error) {
			res, err := n.GetNotifyFlowIDsWithFeed(ctx, id)
			if err != nil {
				return nil, err
			}
			return (*modelangela.FeedToNotifyFlowValue)(&res), nil
		},
		libcache.WithExpiration(libtime.SevenDays),
	)
}

// NewNotifyFlowCache Cache-Aside Pattern.
func NewNotifyFlowCache(
	n biznetzach.NetzachRepo,
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
	n biznetzach.NetzachRepo,
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
