package data

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/orm/model"
	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelnetzach"

	"gorm.io/gorm/clause"
)

type NetzachRepo struct {
	data *Data
}

func NewNetzachRepo(data *Data) *NetzachRepo {
	return &NetzachRepo{
		data: data,
	}
}

func (n *NetzachRepo) CreateNotifyTarget(
	ctx context.Context,
	id libmodel.InternalID,
	t *modelnetzach.NotifyTarget,
) error {
	q := query.Use(n.data.db).NotifyTarget
	return q.WithContext(ctx).Create(&model.NotifyTarget{
		ID:          t.ID,
		OwnerID:     id,
		Name:        t.Name,
		Description: t.Description,
		Destination: t.Destination,
		Status:      converter.ToORMNotifyTargetStatus(t.Status),
	})
}

func (n *NetzachRepo) UpdateNotifyTarget(
	ctx context.Context,
	userID libmodel.InternalID,
	t *modelnetzach.NotifyTarget,
) error {
	q := query.Use(n.data.db).NotifyTarget
	u := q.WithContext(ctx)

	// Where clauses
	u = u.Where(q.OwnerID.Eq(int64(userID)), q.ID.Eq(int64(t.ID)))

	updates := make(map[string]interface{})
	if len(t.Name) > 0 {
		updates["name"] = t.Name
	}
	if len(t.Description) > 0 {
		updates["description"] = t.Description
	}
	if t.Destination != nil {
		updates["destination"] = t.Destination
	}
	if t.Status != modelnetzach.NotifyTargetStatusUnspecified {
		updates["status"] = converter.ToORMNotifyTargetStatus(t.Status)
	}

	if len(updates) == 0 {
		return nil
	}

	_, err := u.Updates(updates)
	return err
}

func (n *NetzachRepo) ListNotifyTargets(
	ctx context.Context,
	paging libmodel.Paging,
	userID libmodel.InternalID,
	ids []libmodel.InternalID,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, error) {
	q := query.Use(n.data.db).NotifyTarget
	u := q.WithContext(ctx).Where(q.OwnerID.Eq(int64(userID)))

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		u = u.Where(q.ID.In(castIDs...))
	}
	if len(statuses) > 0 {
		s := make([]string, len(statuses))
		for i, v := range statuses {
			s[i] = converter.ToORMNotifyTargetStatus(v)
		}
		u = u.Where(q.Status.In(s...))
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}

	return converter.ToBizNotifyTargetList(res), total, nil
}

func (n *NetzachRepo) GetNotifyTarget(ctx context.Context, id libmodel.InternalID) (*modelnetzach.NotifyTarget, error) {
	q := query.Use(n.data.db).NotifyTarget
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizNotifyTarget(res), nil
}

func (n *NetzachRepo) CreateNotifyFlow(
	ctx context.Context,
	userID libmodel.InternalID,
	f *modelnetzach.NotifyFlow,
) error {
	return n.data.WithTx(ctx, func(tx *query.Query) error {
		// Create Flow
		if err := tx.NotifyFlow.WithContext(ctx).Create(&model.NotifyFlow{
			ID:          f.ID,
			OwnerID:     userID,
			Name:        f.Name,
			Description: f.Description,
			Status:      converter.ToORMNotifySourceStatus(f.Status),
		}); err != nil {
			return err
		}

		// Create Sources
		if len(f.Sources) > 0 {
			flowSources := make([]*model.NotifyFlowSource, len(f.Sources))
			for i, source := range f.Sources {
				flowSources[i] = &model.NotifyFlowSource{
					NotifyFlowID:          f.ID,
					NotifySourceID:        source.SourceID,
					FilterExcludeKeywords: source.Filter.ExcludeKeywords,
					FilterIncludeKeywords: source.Filter.IncludeKeywords,
				}
			}
			if err := tx.NotifyFlowSource.WithContext(ctx).Create(flowSources...); err != nil {
				return err
			}
		}

		// Create Targets
		if len(f.Targets) > 0 {
			flowTargets := make([]*model.NotifyFlowTarget, len(f.Targets))
			for i, target := range f.Targets {
				flowTargets[i] = &model.NotifyFlowTarget{
					NotifyFlowID:          f.ID,
					NotifyTargetID:        target.TargetID,
					FilterExcludeKeywords: target.Filter.ExcludeKeywords,
					FilterIncludeKeywords: target.Filter.IncludeKeywords,
				}
			}
			if err := tx.NotifyFlowTarget.WithContext(ctx).Create(flowTargets...); err != nil {
				return err
			}
		}

		// Create FeedItemCollection
		if err := tx.FeedItemCollection.WithContext(ctx).Create(&model.FeedItemCollection{
			ID:          f.ID,
			UserID:      userID,
			Name:        f.Name,
			Description: f.Description,
			Category:    "",
		}); err != nil {
			return err
		}

		return nil
	})
}

func (n *NetzachRepo) UpdateNotifyFlow( //nolint:gocognit // complex logic
	ctx context.Context,
	userID libmodel.InternalID,
	f *modelnetzach.NotifyFlow,
) error {
	return n.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.NotifyFlow
		u := q.WithContext(ctx).Where(q.OwnerID.Eq(int64(userID)), q.ID.Eq(int64(f.ID)))

		updates := make(map[string]interface{})
		if len(f.Name) > 0 {
			updates["name"] = f.Name
		}
		if len(f.Description) > 0 {
			updates["description"] = f.Description
		}
		if f.Status != modelnetzach.NotifyFlowStatusUnspecified {
			updates["status"] = converter.ToORMNotifySourceStatus(f.Status)
		}

		if len(updates) > 0 {
			if _, err := u.Updates(updates); err != nil {
				return err
			}
		}

		if f.Sources != nil {
			qs := tx.NotifyFlowSource
			if _, err := qs.WithContext(ctx).Where(qs.NotifyFlowID.Eq(int64(f.ID))).Delete(); err != nil {
				return err
			}

			flowSources := make([]*model.NotifyFlowSource, len(f.Sources))
			for i, source := range f.Sources {
				flowSources[i] = &model.NotifyFlowSource{
					NotifyFlowID:          f.ID,
					NotifySourceID:        source.SourceID,
					FilterExcludeKeywords: source.Filter.ExcludeKeywords,
					FilterIncludeKeywords: source.Filter.IncludeKeywords,
				}
			}
			if len(flowSources) > 0 {
				if err := qs.WithContext(ctx).Create(flowSources...); err != nil {
					return err
				}
			}
		}

		if f.Targets != nil {
			qt := tx.NotifyFlowTarget
			if _, err := qt.WithContext(ctx).Where(qt.NotifyFlowID.Eq(int64(f.ID))).Delete(); err != nil {
				return err
			}

			flowTargets := make([]*model.NotifyFlowTarget, len(f.Targets))
			for i, target := range f.Targets {
				flowTargets[i] = &model.NotifyFlowTarget{
					NotifyFlowID:          f.ID,
					NotifyTargetID:        target.TargetID,
					FilterExcludeKeywords: target.Filter.ExcludeKeywords,
					FilterIncludeKeywords: target.Filter.IncludeKeywords,
				}
			}
			if len(flowTargets) > 0 {
				if err := qt.WithContext(ctx).Create(flowTargets...); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (n *NetzachRepo) ListNotifyFlows(
	ctx context.Context,
	paging libmodel.Paging,
	userID libmodel.InternalID,
	ids []libmodel.InternalID,
) ([]*modelnetzach.NotifyFlow, int64, error) {
	q := query.Use(n.data.db).NotifyFlow
	u := q.WithContext(ctx).Where(q.OwnerID.Eq(int64(userID)))

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		u = u.Where(q.ID.In(castIDs...))
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := u.Preload(q.NotifySources).
		Preload(q.NotifyTargets).
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	bizFlows := make([]*modelnetzach.NotifyFlow, len(res))
	for i, flow := range res {
		bizFlows[i] = &modelnetzach.NotifyFlow{
			ID:          flow.ID,
			Name:        flow.Name,
			Description: flow.Description,
			Status:      converter.ToBizNotifyFlowStatus(flow.Status),
			Sources:     nil,
			Targets:     nil,
		}
		n.fillFlowDetails(ctx, bizFlows[i])
	}

	return bizFlows, total, nil
}

func (n *NetzachRepo) fillFlowDetails(ctx context.Context, f *modelnetzach.NotifyFlow) {
	// Fetch sources
	q := query.Use(n.data.db)
	qs := q.NotifyFlowSource
	sources, _ := qs.WithContext(ctx).Where(qs.NotifyFlowID.Eq(int64(f.ID))).Find()
	f.Sources = make([]*modelnetzach.NotifyFlowSource, len(sources))
	for i, s := range sources {
		f.Sources[i] = &modelnetzach.NotifyFlowSource{
			SourceID: s.NotifySourceID,
			Filter: &modelnetzach.NotifyFilter{
				ExcludeKeywords: s.FilterExcludeKeywords,
				IncludeKeywords: s.FilterIncludeKeywords,
			},
		}
	}

	qt := q.NotifyFlowTarget
	targets, _ := qt.WithContext(ctx).Where(qt.NotifyFlowID.Eq(int64(f.ID))).Find()
	f.Targets = make([]*modelnetzach.NotifyFlowTarget, len(targets))
	for i, t := range targets {
		f.Targets[i] = &modelnetzach.NotifyFlowTarget{
			TargetID: t.NotifyTargetID,
			Filter: &modelnetzach.NotifyFilter{
				ExcludeKeywords: t.FilterExcludeKeywords,
				IncludeKeywords: t.FilterIncludeKeywords,
			},
		}
	}
}

func (n *NetzachRepo) GetNotifyFlow(ctx context.Context, id libmodel.InternalID) (*modelnetzach.NotifyFlow, error) {
	q := query.Use(n.data.db).NotifyFlow
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}

	flow := &modelnetzach.NotifyFlow{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Status:      converter.ToBizNotifyFlowStatus(res.Status),
		Sources:     nil,
		Targets:     nil,
	}
	n.fillFlowDetails(ctx, flow)
	return flow, nil
}

func (n *NetzachRepo) GetNotifyFlowIDsWithFeed(
	ctx context.Context,
	id libmodel.InternalID,
) ([]libmodel.InternalID, error) {
	q := query.Use(n.data.db).NotifyFlowSource

	var ids []int64
	err := q.WithContext(ctx).
		Where(q.NotifySourceID.Eq(int64(id))).
		Pluck(q.NotifyFlowID, &ids)

	res := make([]libmodel.InternalID, len(ids))
	for i, v := range ids {
		res[i] = libmodel.InternalID(v)
	}
	return res, err
}

func (n *NetzachRepo) UpsertSystemNotification(
	ctx context.Context,
	userID libmodel.InternalID,
	notification *modelnetzach.SystemNotification,
) error {
	q := query.Use(n.data.db).SystemNotification

	// Get old content if exists
	old, err := q.WithContext(ctx).Where(q.ID.Eq(int64(notification.ID))).First()
	if err == nil && len(old.Content) > 0 {
		notification.Content = fmt.Sprintf("%s\n%s", old.Content, notification.Content)
	}

	sysNotif := &model.SystemNotification{
		ID:      notification.ID,
		Type:    converter.ToORMSystemNotificationType(notification.Type),
		Level:   converter.ToORMSystemNotificationLevel(notification.Level),
		Status:  converter.ToORMSystemNotificationStatus(notification.Status),
		Title:   notification.Title,
		Content: notification.Content,
	}
	if notification.Type == modelnetzach.SystemNotificationTypeUser {
		sysNotif.UserID = userID
	}

	return q.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"type", "level", "status", "title", "content", "updated_at"}),
	}).Create(sysNotif)
}

func (n *NetzachRepo) ListSystemNotifications(
	ctx context.Context,
	paging libmodel.Paging,
	userID *libmodel.InternalID,
	types []modelnetzach.SystemNotificationType,
	levels []modelnetzach.SystemNotificationLevel,
	statuses []modelnetzach.SystemNotificationStatus,
) ([]*modelnetzach.SystemNotification, int64, error) {
	q := query.Use(n.data.db).SystemNotification
	u := q.WithContext(ctx).Order(q.UpdatedAt.Desc())

	if userID != nil {
		u = u.Where(q.UserID.Eq(int64(*userID)))
	}
	if len(types) > 0 {
		s := make([]string, len(types))
		for i, v := range types {
			s[i] = converter.ToORMSystemNotificationType(v)
		}
		u = u.Where(q.Type.In(s...))
	}
	if len(levels) > 0 {
		s := make([]string, len(levels))
		for i, v := range levels {
			s[i] = converter.ToORMSystemNotificationLevel(v)
		}
		u = u.Where(q.Level.In(s...))
	}
	if len(statuses) > 0 {
		s := make([]string, len(statuses))
		for i, v := range statuses {
			s[i] = converter.ToORMSystemNotificationStatus(v)
		}
		u = u.Where(q.Status.In(s...))
	}

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := u.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find()
	if err != nil {
		return nil, 0, err
	}

	return converter.ToBizSystemNotificationList(res), total, nil
}
