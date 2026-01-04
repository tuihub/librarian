package data

import (
	"context"
	"database/sql/driver"
	"fmt"

	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelnetzach"
	"github.com/tuihub/librarian/internal/model/modelyesod"

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
	t.OwnerID = id
	return q.WithContext(ctx).Create(t)
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
		updates["status"] = t.Status
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
		s := make([]driver.Valuer, len(statuses))
		for i, v := range statuses {
			s[i] = v
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

	return res, total, nil
}

func (n *NetzachRepo) GetNotifyTarget(ctx context.Context, id libmodel.InternalID) (*modelnetzach.NotifyTarget, error) {
	q := query.Use(n.data.db).NotifyTarget
	res, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (n *NetzachRepo) CreateNotifyFlow(
	ctx context.Context,
	userID libmodel.InternalID,
	f *modelnetzach.NotifyFlow,
) error {
	return n.data.WithTx(ctx, func(tx *query.Query) error {
		// Create Flow (includes Sources and Targets via GORM associations if set)
		f.OwnerID = userID
		if err := tx.NotifyFlow.WithContext(ctx).Create(f); err != nil {
			return err
		}

		// Create FeedItemCollection
		// NotifyFlow ID is same as FeedItemCollection ID?
		// Yes, logic implies it.
		if err := tx.FeedItemCollection.WithContext(ctx).Create(&modelyesod.FeedItemCollection{
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

//nolint:gocognit // complex logic
func (n *NetzachRepo) UpdateNotifyFlow(
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
			updates["status"] = f.Status
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
			if len(f.Sources) > 0 {
				for _, s := range f.Sources {
					s.NotifyFlowID = f.ID // Ensure ID is set
				}
				if err := qs.WithContext(ctx).Create(f.Sources...); err != nil {
					return err
				}
			}
		}

		if f.Targets != nil {
			qt := tx.NotifyFlowTarget
			if _, err := qt.WithContext(ctx).Where(qt.NotifyFlowID.Eq(int64(f.ID))).Delete(); err != nil {
				return err
			}
			if len(f.Targets) > 0 {
				for _, t := range f.Targets {
					t.NotifyFlowID = f.ID
				}
				if err := qt.WithContext(ctx).Create(f.Targets...); err != nil {
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

	res, err := u.Preload(q.Sources).
		Preload(q.Targets).
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

func (n *NetzachRepo) GetNotifyFlow(ctx context.Context, id libmodel.InternalID) (*modelnetzach.NotifyFlow, error) {
	q := query.Use(n.data.db).NotifyFlow
	res, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Preload(q.Sources).
		Preload(q.Targets).
		First()
	if err != nil {
		return nil, err
	}
	return res, nil
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

	if notification.Type == modelnetzach.SystemNotificationTypeUser {
		notification.UserID = userID
	}

	return q.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"type", "level", "status", "title", "content", "updated_at"}),
	}).Create(notification)
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
	u := q.WithContext(ctx).Order(q.UpdateTime.Desc())

	if userID != nil {
		u = u.Where(q.UserID.Eq(int64(*userID)))
	}
	if len(types) > 0 {
		t := make([]driver.Valuer, len(types))
		for i, v := range types {
			t[i] = v
		}
		u = u.Where(q.Type.In(t...))
	}
	if len(levels) > 0 {
		l := make([]driver.Valuer, len(levels))
		for i, v := range levels {
			l[i] = v
		}
		u = u.Where(q.Level.In(l...))
	}
	if len(statuses) > 0 {
		s := make([]driver.Valuer, len(statuses))
		for i, v := range statuses {
			s[i] = v
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

	return res, total, nil
}
