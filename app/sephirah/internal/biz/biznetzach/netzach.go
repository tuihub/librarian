package biznetzach

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizutils"
	"github.com/tuihub/librarian/app/sephirah/internal/client"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelangela"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/supervisor"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libcache"
	"github.com/tuihub/librarian/internal/lib/libmq"
	"github.com/tuihub/librarian/internal/lib/logger"
	"github.com/tuihub/librarian/internal/model"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewNetzach,
	NewSystemNotificationTopic,
)

type NetzachRepo interface {
	CreateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	UpdateNotifyTarget(context.Context, model.InternalID, *modelnetzach.NotifyTarget) error
	ListNotifyTargets(context.Context, model.Paging, model.InternalID, []model.InternalID,
		[]modelnetzach.NotifyTargetStatus) (
		[]*modelnetzach.NotifyTarget, int64, error)
	GetNotifyTarget(context.Context, model.InternalID) (*modelnetzach.NotifyTarget, error)
	CreateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	UpdateNotifyFlow(context.Context, model.InternalID, *modelnetzach.NotifyFlow) error
	ListNotifyFlows(context.Context, model.Paging, model.InternalID, []model.InternalID) (
		[]*modelnetzach.NotifyFlow, int64, error)
	GetNotifyFlow(context.Context, model.InternalID) (*modelnetzach.NotifyFlow, error)
	GetNotifyFlowIDsWithFeed(context.Context, model.InternalID) ([]model.InternalID, error)

	UpsertSystemNotification(context.Context, model.InternalID, *modelnetzach.SystemNotification) error
	ListSystemNotifications(context.Context, model.Paging, *model.InternalID, []modelnetzach.SystemNotificationType,
		[]modelnetzach.SystemNotificationLevel, []modelnetzach.SystemNotificationStatus) (
		[]*modelnetzach.SystemNotification, int64, error)
}

type Netzach struct {
	repo              NetzachRepo
	supv              *supervisor.Supervisor
	searcher          *client.Searcher
	notifySourceCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue]
	notifyFlowCache   *libcache.Map[model.InternalID, modelnetzach.NotifyFlow]
	notifyTargetCache *libcache.Map[model.InternalID, modelnetzach.NotifyTarget]
}

func NewNetzach(
	repo NetzachRepo,
	supv *supervisor.Supervisor,
	sClient *client.Searcher,
	mq *libmq.MQ,
	notifySourceCache *libcache.Map[model.InternalID, modelangela.FeedToNotifyFlowValue],
	notifyFlowCache *libcache.Map[model.InternalID, modelnetzach.NotifyFlow],
	notifyTargetCache *libcache.Map[model.InternalID, modelnetzach.NotifyTarget],
	systemNotification *libmq.Topic[modelnetzach.SystemNotify],
) (*Netzach, error) {
	if err := mq.RegisterTopic(systemNotification); err != nil {
		return nil, err
	}
	y := &Netzach{
		repo,
		supv,
		sClient,
		notifySourceCache,
		notifyFlowCache,
		notifyTargetCache,
	}
	return y, nil
}

func (n *Netzach) CreateNotifyTarget(ctx context.Context, target *modelnetzach.NotifyTarget) (
	model.InternalID, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return 0, bizutils.NoPermissionError()
	}
	if !n.supv.HasNotifyDestination(target.Destination) {
		return 0, bizutils.UnsupportedFeatureError()
	}
	if target == nil {
		return 0, pb.ErrorErrorReasonBadRequest("notify target required")
	}
	id, err := n.searcher.NewID(ctx)
	if err != nil {
		return 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	target.ID = id
	err = n.repo.CreateNotifyTarget(ctx, claims.UserID, target)
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
	if !n.supv.HasNotifyDestination(target.Destination) {
		return bizutils.UnsupportedFeatureError()
	}
	err := n.repo.UpdateNotifyTarget(ctx, claims.UserID, target)
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
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	targets, total, err := n.repo.ListNotifyTargets(ctx, paging, claims.UserID, ids, statuses)
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
	err = n.repo.CreateNotifyFlow(ctx, claims.UserID, flow)
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
	err := n.repo.UpdateNotifyFlow(ctx, claims.UserID, flow)
	if err != nil {
		return pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	if len(flow.Sources) > 0 {
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
	res, total, err := n.repo.ListNotifyFlows(ctx, paging, claims.UserID, ids)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}

func (n *Netzach) ListSystemNotifications(
	ctx context.Context,
	paging model.Paging,
	types []modelnetzach.SystemNotificationType,
	levels []modelnetzach.SystemNotificationLevel,
	statuses []modelnetzach.SystemNotificationStatus,
) ([]*modelnetzach.SystemNotification, int64, *errors.Error) {
	claims := libauth.FromContextAssertUserType(ctx)
	if claims == nil {
		return nil, 0, bizutils.NoPermissionError()
	}
	var userID *model.InternalID
	if claims.UserType != libauth.UserTypeAdmin {
		types = []modelnetzach.SystemNotificationType{modelnetzach.SystemNotificationTypeUser}
		userID = &claims.UserID
	}
	res, total, err := n.repo.ListSystemNotifications(ctx, paging, userID, types, levels, statuses)
	if err != nil {
		return nil, 0, pb.ErrorErrorReasonUnspecified("%s", err.Error())
	}
	return res, total, nil
}
