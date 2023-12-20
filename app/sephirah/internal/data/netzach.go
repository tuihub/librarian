package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"
)

type netzachRepo struct {
	data *Data
}

func NewNetzachRepo(data *Data) biznetzach.NetzachRepo {
	return &netzachRepo{
		data: data,
	}
}

func (n *netzachRepo) CreateNotifyTarget(ctx context.Context, id model.InternalID, t *modelnetzach.NotifyTarget) error {
	q := n.data.db.NotifyTarget.Create().
		SetOwnerID(id).
		SetID(t.ID).
		SetName(t.Name).
		SetDescription(t.Description).
		SetToken(t.Token).
		SetType(converter.ToEntNotifyTargetType(t.Type)).
		SetStatus(converter.ToEntNotifyTargetStatus(t.Status))
	return q.Exec(ctx)
}

func (n *netzachRepo) UpdateNotifyTarget(
	ctx context.Context,
	userID model.InternalID,
	t *modelnetzach.NotifyTarget,
) error {
	q := n.data.db.NotifyTarget.Update().Where(
		notifytarget.HasOwnerWith(user.IDEQ(userID)),
		notifytarget.IDEQ(t.ID),
	)
	if len(t.Name) > 0 {
		q.SetName(t.Name)
	}
	if len(t.Description) > 0 {
		q.SetDescription(t.Description)
	}
	if len(t.Token) > 0 {
		q.SetToken(t.Token)
	}
	if t.Type != modelnetzach.NotifyTargetTypeUnspecified {
		q.SetType(converter.ToEntNotifyTargetType(t.Type))
	}
	if t.Status != modelnetzach.NotifyTargetStatusUnspecified {
		q.SetStatus(converter.ToEntNotifyTargetStatus(t.Status))
	}
	return q.Exec(ctx)
}

func (n *netzachRepo) ListNotifyTargets(
	ctx context.Context,
	paging model.Paging,
	userID model.InternalID,
	ids []model.InternalID,
	types []modelnetzach.NotifyTargetType,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, error) {
	q := n.data.db.NotifyTarget.Query().Where(
		notifytarget.HasOwnerWith(user.IDEQ(userID)),
	)
	if len(ids) > 0 {
		q.Where(notifytarget.IDIn(ids...))
	}
	if len(types) > 0 {
		q.Where(notifytarget.TypeIn(converter.ToEntNotifyTargetTypeList(types)...))
	}
	if len(statuses) > 0 {
		q.Where(notifytarget.StatusIn(converter.ToEntNotifyTargetStatusList(statuses)...))
	}
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	res, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizNotifyTargetList(res), int64(total), nil
}

func (n *netzachRepo) GetNotifyTarget(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyTarget, error) {
	res, err := n.data.db.NotifyTarget.Query().Where(notifytarget.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizNotifyTarget(res), nil
}

func (n *netzachRepo) CreateNotifyFlow(ctx context.Context, userID model.InternalID, f *modelnetzach.NotifyFlow) error {
	err := n.data.WithTx(ctx, func(tx *ent.Tx) error {
		flowSources := make([]*ent.NotifyFlowSourceCreate, len(f.Sources))
		for i, source := range f.Sources {
			flowSources[i] = tx.NotifyFlowSource.Create().
				SetNotifyFlowID(f.ID).
				SetNotifySourceID(source.SourceID).
				SetFilterExcludeKeywords(source.Filter.ExcludeKeywords).
				SetFilterIncludeKeywords(source.Filter.IncludeKeywords)
		}
		source, err := tx.NotifyFlowSource.CreateBulk(flowSources...).Save(ctx)
		if err != nil {
			return err
		}
		flowTargets := make([]*ent.NotifyFlowTargetCreate, len(f.Targets))
		for i, target := range f.Targets {
			flowTargets[i] = tx.NotifyFlowTarget.Create().
				SetNotifyFlowID(f.ID).
				SetNotifyTargetID(target.TargetID).
				SetChannelID(target.ChannelID).
				SetFilterExcludeKeywords(target.Filter.ExcludeKeywords).
				SetFilterIncludeKeywords(target.Filter.IncludeKeywords)
		}
		target, err := tx.NotifyFlowTarget.CreateBulk(flowTargets...).Save(ctx)
		if err != nil {
			return err
		}
		q := n.data.db.NotifyFlow.Create().
			SetID(f.ID).
			SetOwnerID(userID).
			SetName(f.Name).
			SetDescription(f.Description).
			SetStatus(converter.ToEntNotifySourceSource(f.Status)).
			AddNotifyFlowSource(source...).
			AddNotifyFlowTarget(target...)
		return q.Exec(ctx)
	})
	if err != nil {
		return err
	}
	return nil
}

func (n *netzachRepo) UpdateNotifyFlow( //nolint:gocognit // TODO
	ctx context.Context,
	userID model.InternalID,
	f *modelnetzach.NotifyFlow,
) error {
	err := n.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.NotifyFlow.Update().Where(
			notifyflow.HasOwnerWith(user.IDEQ(userID)),
			notifyflow.IDEQ(f.ID),
		)
		if len(f.Name) > 0 {
			q.SetName(f.Name)
		}
		if len(f.Description) > 0 {
			q.SetDescription(f.Description)
		}
		if f.Sources != nil {
			_, err := tx.NotifyFlowSource.Delete().Where(
				notifyflowsource.HasNotifyFlowWith(
					notifyflow.IDEQ(f.ID),
				),
			).Exec(ctx)
			if err != nil {
				return err
			}
			flowSources := make([]*ent.NotifyFlowSourceCreate, len(f.Sources))
			for i, source := range f.Sources {
				flowSources[i] = tx.NotifyFlowSource.Create().
					SetNotifyFlowID(f.ID).
					SetNotifySourceID(source.SourceID).
					SetFilterExcludeKeywords(source.Filter.ExcludeKeywords).
					SetFilterIncludeKeywords(source.Filter.IncludeKeywords)
			}
			source, err := tx.NotifyFlowSource.CreateBulk(flowSources...).Save(ctx)
			if err != nil {
				return err
			}
			q.ClearFeedConfig().AddNotifyFlowSource(source...)
		}
		if f.Targets != nil {
			_, err := tx.NotifyFlowTarget.Delete().Where(
				notifyflowtarget.HasNotifyFlowWith(
					notifyflow.IDEQ(f.ID),
				),
			).Exec(ctx)
			if err != nil {
				return err
			}
			flowTargets := make([]*ent.NotifyFlowTargetCreate, len(f.Targets))
			for i, target := range f.Targets {
				flowTargets[i] = tx.NotifyFlowTarget.Create().
					SetNotifyFlowID(f.ID).
					SetNotifyTargetID(target.TargetID).
					SetChannelID(target.ChannelID).
					SetFilterExcludeKeywords(target.Filter.ExcludeKeywords).
					SetFilterIncludeKeywords(target.Filter.IncludeKeywords)
			}
			targets, err := tx.NotifyFlowTarget.CreateBulk(flowTargets...).Save(ctx)
			if err != nil {
				return err
			}
			q.ClearNotifyTarget().AddNotifyFlowTarget(targets...)
		}
		if f.Status != modelnetzach.NotifyFlowStatusUnspecified {
			q.SetStatus(converter.ToEntNotifySourceSource(f.Status))
		}
		return q.Exec(ctx)
	})
	if err != nil {
		return err
	}
	return nil
}

func (n *netzachRepo) ListNotifyFlows(
	ctx context.Context,
	paging model.Paging,
	userID model.InternalID,
	ids []model.InternalID,
) ([]*modelnetzach.NotifyFlow, int64, error) {
	q := n.data.db.NotifyFlow.Query().Where(
		notifyflow.HasOwnerWith(user.IDEQ(userID)),
	)
	if len(ids) > 0 {
		q.Where(notifyflow.IDIn(ids...))
	}
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	flows, err := q.
		WithNotifyFlowSource().
		WithNotifyFlowTarget().
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*modelnetzach.NotifyFlow, len(flows))
	for i := range flows {
		res[i] = converter.ToBizNotifyFlow(flows[i])
	}
	return res, int64(total), nil
}

func (n *netzachRepo) GetNotifyFlow(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyFlow, error) {
	res, err := n.data.db.NotifyFlow.Query().
		Where(notifyflow.IDEQ(id)).
		WithFeedConfig().
		WithNotifyFlowTarget().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizNotifyFlow(res), nil
}

func (n *netzachRepo) GetNotifyFlowIDsWithFeed(ctx context.Context, id model.InternalID) ([]model.InternalID, error) {
	ids, err := n.data.db.NotifyFlow.Query().Where(
		notifyflow.HasFeedConfigWith(feedconfig.IDEQ(id)),
	).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}
