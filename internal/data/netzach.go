package data

import (
	"context"
	"fmt"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelnetzach"

	"gorm.io/gorm"
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

func (n *NetzachRepo) CreateNotifyTarget(ctx context.Context, id model.InternalID, t *modelnetzach.NotifyTarget) error {
	var destVal *gormschema.FeatureRequestVal
	if t.Destination != nil {
		v := gormschema.FeatureRequestVal(*t.Destination)
		destVal = &v
	}
	target := gormschema.NotifyTarget{
		ID:          t.ID,
		OwnerID:     id,
		Name:        t.Name,
		Description: t.Description,
		Destination: destVal,
		Status:      gormschema.ToSchemaNotifyTargetStatus(t.Status),
	}
	return n.data.db.WithContext(ctx).Create(&target).Error
}

func (n *NetzachRepo) UpdateNotifyTarget(
	ctx context.Context,
	userID model.InternalID,
	t *modelnetzach.NotifyTarget,
) error {
	updates := make(map[string]any)
	if len(t.Name) > 0 {
		updates["name"] = t.Name
	}
	if len(t.Description) > 0 {
		updates["description"] = t.Description
	}
	if t.Destination != nil {
		v := gormschema.FeatureRequestVal(*t.Destination)
		updates["destination"] = &v
	}
	if t.Status != modelnetzach.NotifyTargetStatusUnspecified {
		updates["status"] = gormschema.ToSchemaNotifyTargetStatus(t.Status)
	}
	return n.data.db.WithContext(ctx).
		Model(&gormschema.NotifyTarget{}).
		Where("id = ? AND owner_id = ?", t.ID, userID).
		Updates(updates).Error
}

func (n *NetzachRepo) ListNotifyTargets(
	ctx context.Context,
	paging model.Paging,
	userID model.InternalID,
	ids []model.InternalID,
	statuses []modelnetzach.NotifyTargetStatus,
) ([]*modelnetzach.NotifyTarget, int64, error) {
	query := n.data.db.WithContext(ctx).Model(&gormschema.NotifyTarget{}).
		Where("owner_id = ?", userID)

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}
	if len(statuses) > 0 {
		statusStrs := make([]string, len(statuses))
		for i, s := range statuses {
			statusStrs[i] = gormschema.ToSchemaNotifyTargetStatus(s)
		}
		query = query.Where("status IN ?", statusStrs)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var targets []gormschema.NotifyTarget
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&targets).Error; err != nil {
		return nil, 0, err
	}

	return gormschema.ToBizNotifyTargetList(ptrSlice(targets)), total, nil
}

func (n *NetzachRepo) GetNotifyTarget(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyTarget, error) {
	var target gormschema.NotifyTarget
	if err := n.data.db.WithContext(ctx).First(&target, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizNotifyTarget(&target), nil
}

func (n *NetzachRepo) CreateNotifyFlow(ctx context.Context, userID model.InternalID, f *modelnetzach.NotifyFlow) error {
	return n.data.WithTx(ctx, func(tx *gorm.DB) error {
		flow := gormschema.NotifyFlow{
			ID:          f.ID,
			OwnerID:     userID,
			Name:        f.Name,
			Description: f.Description,
			Status:      gormschema.ToSchemaNotifyFlowStatus(f.Status),
		}
		if err := tx.Create(&flow).Error; err != nil {
			return err
		}

		for _, source := range f.Sources {
			fs := gormschema.NotifyFlowSource{
				NotifyFlowID:          f.ID,
				NotifySourceID:        source.SourceID,
				FilterExcludeKeywords: gormschema.StringArrayVal(source.Filter.ExcludeKeywords),
				FilterIncludeKeywords: gormschema.StringArrayVal(source.Filter.IncludeKeywords),
			}
			if err := tx.Create(&fs).Error; err != nil {
				return err
			}
		}

		for _, target := range f.Targets {
			ft := gormschema.NotifyFlowTarget{
				NotifyFlowID:          f.ID,
				NotifyTargetID:        target.TargetID,
				FilterExcludeKeywords: gormschema.StringArrayVal(target.Filter.ExcludeKeywords),
				FilterIncludeKeywords: gormschema.StringArrayVal(target.Filter.IncludeKeywords),
			}
			if err := tx.Create(&ft).Error; err != nil {
				return err
			}
		}

		// For save flow items
		fic := gormschema.FeedItemCollection{
			ID:          f.ID,
			OwnerID:     userID,
			Name:        f.Name,
			Description: f.Description,
			Category:    "",
		}
		return tx.Create(&fic).Error
	})
}

func (n *NetzachRepo) UpdateNotifyFlow(
	ctx context.Context,
	userID model.InternalID,
	f *modelnetzach.NotifyFlow,
) error {
	return n.data.WithTx(ctx, func(tx *gorm.DB) error {
		updates := make(map[string]any)
		if len(f.Name) > 0 {
			updates["name"] = f.Name
		}
		if len(f.Description) > 0 {
			updates["description"] = f.Description
		}
		if f.Status != modelnetzach.NotifyFlowStatusUnspecified {
			updates["status"] = gormschema.ToSchemaNotifyFlowStatus(f.Status)
		}

		if f.Sources != nil {
			if err := tx.Where("notify_flow_id = ?", f.ID).Delete(&gormschema.NotifyFlowSource{}).Error; err != nil {
				return err
			}
			for _, source := range f.Sources {
				fs := gormschema.NotifyFlowSource{
					NotifyFlowID:          f.ID,
					NotifySourceID:        source.SourceID,
					FilterExcludeKeywords: gormschema.StringArrayVal(source.Filter.ExcludeKeywords),
					FilterIncludeKeywords: gormschema.StringArrayVal(source.Filter.IncludeKeywords),
				}
				if err := tx.Create(&fs).Error; err != nil {
					return err
				}
			}
		}

		if f.Targets != nil {
			if err := tx.Where("notify_flow_id = ?", f.ID).Delete(&gormschema.NotifyFlowTarget{}).Error; err != nil {
				return err
			}
			for _, target := range f.Targets {
				ft := gormschema.NotifyFlowTarget{
					NotifyFlowID:          f.ID,
					NotifyTargetID:        target.TargetID,
					FilterExcludeKeywords: gormschema.StringArrayVal(target.Filter.ExcludeKeywords),
					FilterIncludeKeywords: gormschema.StringArrayVal(target.Filter.IncludeKeywords),
				}
				if err := tx.Create(&ft).Error; err != nil {
					return err
				}
			}
		}

		return tx.Model(&gormschema.NotifyFlow{}).
			Where("id = ? AND owner_id = ?", f.ID, userID).
			Updates(updates).Error
	})
}

func (n *NetzachRepo) ListNotifyFlows(
	ctx context.Context,
	paging model.Paging,
	userID model.InternalID,
	ids []model.InternalID,
) ([]*modelnetzach.NotifyFlow, int64, error) {
	query := n.data.db.WithContext(ctx).Model(&gormschema.NotifyFlow{}).
		Where("owner_id = ?", userID)

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var flows []gormschema.NotifyFlow
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&flows).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelnetzach.NotifyFlow, len(flows))
	for i := range flows {
		res[i] = gormschema.ToBizNotifyFlow(&flows[i])
		// Get sources
		var sources []gormschema.NotifyFlowSource
		if err := n.data.db.WithContext(ctx).Where("notify_flow_id = ?", flows[i].ID).Find(&sources).Error; err == nil {
			res[i].Sources = make([]*modelnetzach.NotifyFlowSource, len(sources))
			for j, s := range sources {
				res[i].Sources[j] = &modelnetzach.NotifyFlowSource{
					SourceID: s.NotifySourceID,
					Filter: &modelnetzach.NotifyFilter{
						ExcludeKeywords: []string(s.FilterExcludeKeywords),
						IncludeKeywords: []string(s.FilterIncludeKeywords),
					},
				}
			}
		}
		// Get targets
		var targets []gormschema.NotifyFlowTarget
		if err := n.data.db.WithContext(ctx).Where("notify_flow_id = ?", flows[i].ID).Find(&targets).Error; err == nil {
			res[i].Targets = make([]*modelnetzach.NotifyFlowTarget, len(targets))
			for j, t := range targets {
				res[i].Targets[j] = &modelnetzach.NotifyFlowTarget{
					TargetID: t.NotifyTargetID,
					Filter: &modelnetzach.NotifyFilter{
						ExcludeKeywords: []string(t.FilterExcludeKeywords),
						IncludeKeywords: []string(t.FilterIncludeKeywords),
					},
				}
			}
		}
	}
	return res, total, nil
}

func (n *NetzachRepo) GetNotifyFlow(ctx context.Context, id model.InternalID) (*modelnetzach.NotifyFlow, error) {
	var flow gormschema.NotifyFlow
	if err := n.data.db.WithContext(ctx).First(&flow, id).Error; err != nil {
		return nil, err
	}

	res := gormschema.ToBizNotifyFlow(&flow)
	// Get sources
	var sources []gormschema.NotifyFlowSource
	if err := n.data.db.WithContext(ctx).Where("notify_flow_id = ?", flow.ID).Find(&sources).Error; err == nil {
		res.Sources = make([]*modelnetzach.NotifyFlowSource, len(sources))
		for i, s := range sources {
			res.Sources[i] = &modelnetzach.NotifyFlowSource{
				SourceID: s.NotifySourceID,
				Filter: &modelnetzach.NotifyFilter{
					ExcludeKeywords: []string(s.FilterExcludeKeywords),
					IncludeKeywords: []string(s.FilterIncludeKeywords),
				},
			}
		}
	}
	// Get targets
	var targets []gormschema.NotifyFlowTarget
	if err := n.data.db.WithContext(ctx).Where("notify_flow_id = ?", flow.ID).Find(&targets).Error; err == nil {
		res.Targets = make([]*modelnetzach.NotifyFlowTarget, len(targets))
		for i, t := range targets {
			res.Targets[i] = &modelnetzach.NotifyFlowTarget{
				TargetID: t.NotifyTargetID,
				Filter: &modelnetzach.NotifyFilter{
					ExcludeKeywords: []string(t.FilterExcludeKeywords),
					IncludeKeywords: []string(t.FilterIncludeKeywords),
				},
			}
		}
	}
	return res, nil
}

func (n *NetzachRepo) GetNotifyFlowIDsWithFeed(ctx context.Context, id model.InternalID) ([]model.InternalID, error) {
	var sources []gormschema.NotifyFlowSource
	if err := n.data.db.WithContext(ctx).
		Where("notify_source_id = ?", id).
		Find(&sources).Error; err != nil {
		return nil, err
	}
	ids := make([]model.InternalID, len(sources))
	for i, s := range sources {
		ids[i] = s.NotifyFlowID
	}
	return ids, nil
}

func (n *NetzachRepo) UpsertSystemNotification(
	ctx context.Context,
	userID model.InternalID,
	notification *modelnetzach.SystemNotification,
) error {
	// Check if exists and append content
	var old gormschema.SystemNotification
	if err := n.data.db.WithContext(ctx).First(&old, notification.ID).Error; err == nil && len(old.Content) > 0 {
		notification.Content = fmt.Sprintf("%s\n%s", old.Content, notification.Content)
	}

	sn := gormschema.SystemNotification{
		ID:      notification.ID,
		Type:    gormschema.ToSchemaSystemNotificationType(notification.Type),
		Level:   gormschema.ToSchemaSystemNotificationLevel(notification.Level),
		Status:  gormschema.ToSchemaSystemNotificationStatus(notification.Status),
		Title:   notification.Title,
		Content: notification.Content,
	}
	if notification.Type == modelnetzach.SystemNotificationTypeUser {
		sn.UserID = &userID
	}

	return n.data.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"level", "status", "title", "content", "updated_at"}),
	}).Create(&sn).Error
}

func (n *NetzachRepo) ListSystemNotifications(
	ctx context.Context,
	paging model.Paging,
	userID *model.InternalID,
	types []modelnetzach.SystemNotificationType,
	levels []modelnetzach.SystemNotificationLevel,
	statuses []modelnetzach.SystemNotificationStatus,
) ([]*modelnetzach.SystemNotification, int64, error) {
	query := n.data.db.WithContext(ctx).Model(&gormschema.SystemNotification{}).
		Order("updated_at DESC")

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if len(types) > 0 {
		typeStrs := make([]string, len(types))
		for i, t := range types {
			typeStrs[i] = gormschema.ToSchemaSystemNotificationType(t)
		}
		query = query.Where("type IN ?", typeStrs)
	}
	if len(levels) > 0 {
		levelStrs := make([]string, len(levels))
		for i, l := range levels {
			levelStrs[i] = gormschema.ToSchemaSystemNotificationLevel(l)
		}
		query = query.Where("level IN ?", levelStrs)
	}
	if len(statuses) > 0 {
		statusStrs := make([]string, len(statuses))
		for i, s := range statuses {
			statusStrs[i] = gormschema.ToSchemaSystemNotificationStatus(s)
		}
		query = query.Where("status IN ?", statusStrs)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var notifications []gormschema.SystemNotification
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&notifications).Error; err != nil {
		return nil, 0, err
	}

	return gormschema.ToBizSystemNotificationList(ptrSlice(notifications)), total, nil
}
