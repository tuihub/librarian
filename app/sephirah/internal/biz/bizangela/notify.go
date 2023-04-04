package bizangela

import (
	"context"
	"strconv"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libtime"
	"github.com/tuihub/librarian/internal/model"
)

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
