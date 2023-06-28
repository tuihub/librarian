package bizangela

import (
	"context"
	"errors"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
)

func NewNotifyRouterTopic(
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
				for _, target := range flow.Targets {
					if target == nil {
						continue
					}
					err = push.Publish(ctx, modelangela.NotifyPush{
						Target:   *target,
						Messages: r.Messages,
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
			_, err = a.porter.PushFeedItems(ctx, &porter.PushFeedItemsRequest{
				Destination: converter.ToPBFeedDestination(target.Type),
				ChannelId:   p.Target.ChannelID,
				Items:       converter.ToPBFeedItemList(p.Messages),
				Token:       target.Token,
			})
			if err != nil {
				return err
			}
			return nil
		},
	)
}

// NewFeedToNotifyFlowMap Cache-Aside Pattern.
func NewFeedToNotifyFlowMap(
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
