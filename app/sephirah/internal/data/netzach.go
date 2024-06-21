package data

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biznetzach"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/systemnotification"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent/dialect/sql"
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
		SetDestination(t.Destination).
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
	if len(t.Destination) > 0 {
		q.SetDestination(t.Destination)
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
	destinations []string,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, error) {
	q := n.data.db.NotifyTarget.Query().Where(
		notifytarget.HasOwnerWith(user.IDEQ(userID)),
	)
	if len(ids) > 0 {
		q.Where(notifytarget.IDIn(ids...))
	}
	if len(destinations) > 0 {
		q.Where(notifytarget.DestinationIn(destinations...))
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
		err := tx.NotifyFlow.Create().
			SetID(f.ID).
			SetOwnerID(userID).
			SetName(f.Name).
			SetDescription(f.Description).
			SetStatus(converter.ToEntNotifySourceSource(f.Status)).
			Exec(ctx)
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
		err = tx.NotifyFlowSource.CreateBulk(flowSources...).Exec(ctx)
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
		err = tx.NotifyFlowTarget.CreateBulk(flowTargets...).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
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
			err = tx.NotifyFlowSource.CreateBulk(flowSources...).Exec(ctx)
			if err != nil {
				return err
			}
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
			err = tx.NotifyFlowTarget.CreateBulk(flowTargets...).Exec(ctx)
			if err != nil {
				return err
			}
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
		WithNotifyFlowSource().
		WithNotifyFlowTarget().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizNotifyFlow(res), nil
}

func (n *netzachRepo) GetNotifyFlowIDsWithFeed(ctx context.Context, id model.InternalID) ([]model.InternalID, error) {
	ids, err := n.data.db.NotifyFlow.Query().Where(
		notifyflow.HasNotifySourceWith(notifysource.FeedConfigIDEQ(id)),
	).IDs(ctx)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (n *netzachRepo) UpsertSystemNotification(
	ctx context.Context,
	userID model.InternalID,
	notification *modelnetzach.SystemNotification,
) error {
	old, err := n.data.db.SystemNotification.Get(ctx, notification.ID)
	if err == nil && n != nil && len(old.Content) > 0 {
		notification.Content = fmt.Sprintf("%s\n%s", old.Content, notification.Content)
	}
	q := n.data.db.SystemNotification.Create().
		SetID(notification.ID).
		SetType(converter.ToEntSystemNotificationType(notification.Type)).
		SetLevel(converter.ToEntSystemNotificationLevel(notification.Level)).
		SetStatus(converter.ToEntSystemNotificationStatus(notification.Status)).
		SetTitle(notification.Title).
		SetContent(notification.Content)
	if notification.Type == modelnetzach.SystemNotificationTypeUser {
		q.SetUserID(userID)
	}
	return q.OnConflict(
		sql.ConflictColumns(systemnotification.FieldID),
		resolveWithIgnores([]string{
			systemnotification.FieldID,
			systemnotification.FieldUserID,
			systemnotification.FieldType,
		}),
	).Exec(ctx)
}

func (n *netzachRepo) ListSystemNotifications(ctx context.Context, paging model.Paging, userID *model.InternalID, types []modelnetzach.SystemNotificationType, levels []modelnetzach.SystemNotificationLevel, statuses []modelnetzach.SystemNotificationStatus) ([]*modelnetzach.SystemNotification, int64, error) {
	q := n.data.db.SystemNotification.Query().
		Order(ent.Desc(systemnotification.FieldUpdatedAt))
	if userID != nil {
		q.Where(systemnotification.UserIDEQ(*userID))
	}
	if len(types) > 0 {
		q.Where(systemnotification.TypeIn(converter.ToEntSystemNotificationTypeList(types)...))
	}
	if len(levels) > 0 {
		q.Where(systemnotification.LevelIn(converter.ToEntSystemNotificationLevelList(levels)...))
	}
	if len(statuses) > 0 {
		q.Where(systemnotification.StatusIn(converter.ToEntSystemNotificationStatusList(statuses)...))
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
	return converter.ToBizSystemNotificationList(res), int64(total), nil
}
