package biznetzach

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/logger"
	"github.com/tuihub/librarian/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
)

type NetzachRepo interface {
	CreateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	UpdateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	ListNotifyTargets(context.Context, model.Paging, model.InternalID, []model.InternalID,
		[]modelnetzach.NotifyTargetType, []modelnetzach.NotifyTargetStatus) (
		[]*modelnetzach.NotifyTarget, int64, error)
	GetNotifyTarget(context.Context, model.InternalID) (*modelnetzach.NotifyTarget, error)
	CreateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	UpdateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	ListNotifyFlows(context.Context, model.Paging, model.InternalID, []model.InternalID) (
		[]*modelnetzach.NotifyFlow, int64, error)
	GetNotifyFlow(context.Context, model.InternalID) (*modelnetzach.NotifyFlow, error)
	GetNotifyFlowIDsWithFeed(context.Context, model.InternalID) ([]model.InternalID, error)
}

type Netzach struct {
	repo              NetzachRepo
	searcher          *client.Searcher
	notifySourceCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue]
	notifyFlowCache   *libcache.Map[model.InternalID, modelnetzach.NotifyFlow]
	notifyTargetCache *libcache.Map[model.InternalID, modelnetzach.NotifyTarget]
}

func NewNetzach(
	repo NetzachRepo,
	sClient *client.Searcher,
	notifySourceCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue],
	notifyFlowCache *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	notifyTargetCache *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
) *Netzach {
	y := &Netzach{
		repo,
		sClient,
		notifySourceCache,
		notifyFlowCache,
		notifyTargetCache,
	}
	return y
}

func (n *Netzach) CreateNotifyTarget(ctx context.Context, target *modelnetzach.NotifyTarget) (
	model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	if target == nil {
		return 0, pb.ErrorErrorReasonBadRequest("notify target required")
	}
	id, err := n.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	target.ID = id
	err = n.repo.CreateNotifyTarget(ctx, claims.InternalID, target)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return target.ID, nil
}

func (n *Netzach) UpdateNotifyTarget(ctx context.Context, target *modelnetzach.NotifyTarget) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	err := n.repo.UpdateNotifyTarget(ctx, claims.InternalID, target)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	err = n.notifyTargetCache.Delete(ctx, target.ID)
	if err != nil {
		logger.Errorf("failed to delete cache %s", err.Error())
	}
	return nil
}

func (n *Netzach) ListNotifyTargets(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
	types []modelnetzach.NotifyTargetType,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	targets, total, err := n.repo.ListNotifyTargets(ctx, paging, claims.InternalID, ids, types, statuses)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return targets, total, nil
}

func (n *Netzach) CreateNotifyFlow(ctx context.Context, flow *modelnetzach.NotifyFlow) (
	model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	if flow == nil {
		return 0, pb.ErrorErrorReasonBadRequest("notify target required")
	}
	id, err := n.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	flow.ID = id
	for _, source := range flow.Sources {
		if source.Filter == nil {
			source.Filter = new(modelnetzach.NotifyFilter)
		}
	}
	for _, target := range flow.Targets {
		if target.Filter == nil {
			target.Filter = new(modelnetzach.NotifyFilter)
		}
	}
	err = n.repo.CreateNotifyFlow(ctx, claims.InternalID, flow)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return flow.ID, nil
}

func (n *Netzach) UpdateNotifyFlow(ctx context.Context, flow *modelnetzach.NotifyFlow) *errors.Error {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return bizutils.NoPermissionError()
	}
	for _, source := range flow.Sources {
		if source.Filter == nil {
			source.Filter = new(modelnetzach.NotifyFilter)
		}
	}
	for _, target := range flow.Targets {
		if target.Filter == nil {
			target.Filter = new(modelnetzach.NotifyFilter)
		}
	}
	err := n.repo.UpdateNotifyFlow(ctx, claims.InternalID, flow)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if flow.Sources != nil && len(flow.Sources) > 0 {
		for _, source := range flow.Sources {
			err = n.notifySourceCache.Delete(ctx, source.SourceID)
			if err != nil {
				logger.Errorf("failed to delete cache %s", err.Error())
			}
		}
	}
	err = n.notifyFlowCache.Delete(ctx, flow.ID)
	if err != nil {
		logger.Errorf("failed to delete cache %s", err.Error())
	}
	return nil
}

func (n *Netzach) ListNotifyFlows(
	ctx context.Context,
	paging model.Paging,
	ids []model.InternalID,
) ([]*modelnetzach.NotifyFlow, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	res, total, err := n.repo.ListNotifyFlows(ctx, paging, claims.InternalID, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}
