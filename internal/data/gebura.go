package data

import (
	"context"
	"database/sql/driver"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeburaRepo struct {
	data *Data
}

// NewGeburaRepo .
func NewGeburaRepo(data *Data) *GeburaRepo {
	return &GeburaRepo{
		data: data,
	}
}

func (g *GeburaRepo) CreateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	return query.Use(g.data.db).AppInfo.WithContext(ctx).Create(a)
}

func (g *GeburaRepo) CreateAppInfoOrGet(ctx context.Context, a *modelgebura.AppInfo) (*modelgebura.AppInfo, error) {
	err := g.CreateAppInfo(ctx, a)
	if err == nil {
		return a, nil
	}
	// Check for unique constraint violation
	q := query.Use(g.data.db).AppInfo
	ai, err2 := q.WithContext(ctx).Where(
		q.Source.Eq(a.Source),
		q.SourceAppID.Eq(a.SourceAppID),
	).First()
	if err2 == nil {
		return ai, nil
	}
	return nil, err
}

func (g *GeburaRepo) UpdateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	q := query.Use(g.data.db).AppInfo
	_, err := q.WithContext(ctx).
		Where(
			q.ID.Eq(int64(a.ID)),
			q.Source.Eq(a.Source),
			q.SourceAppID.Eq(a.SourceAppID),
		).
		Updates(a)
	return err
}

func (g *GeburaRepo) ListAppInfos(
	ctx context.Context,
	paging libmodel.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []libmodel.InternalID,
) ([]*modelgebura.AppInfo, int, error) {
	q := query.Use(g.data.db).AppInfo
	queryBuilder := q.WithContext(ctx)

	if len(sources) > 0 {
		queryBuilder = queryBuilder.Where(q.Source.In(sources...))
	}
	if len(types) > 0 {
		t := make([]driver.Valuer, len(types))
		for i, v := range types {
			t[i] = v
		}
		queryBuilder = queryBuilder.Where(q.Type.In(t...))
	}
	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.ID.In(castIDs...))
	}

	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}

	al, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	return al, int(total), nil
}

func (g *GeburaRepo) GetAppInfo(ctx context.Context, id modelgebura.AppInfoID) (*modelgebura.AppInfo, error) {
	q := query.Use(g.data.db).AppInfo
	res, err := q.WithContext(ctx).
		Where(
			q.Source.Eq(id.Source),
			q.SourceAppID.Eq(id.SourceAppID),
		).
		First()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (g *GeburaRepo) CreateApp(ctx context.Context, userID libmodel.InternalID, a *modelgebura.App) error {
	q := query.Use(g.data.db).App
	a.UserID = userID
	return q.WithContext(ctx).Create(a)
}

func (g *GeburaRepo) UpdateApp(ctx context.Context, ownerID libmodel.InternalID, a *modelgebura.App) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.App
		old, err := q.WithContext(ctx).Where(q.ID.Eq(int64(a.ID))).First()
		if err != nil {
			return err
		}

		a.VersionNumber = old.VersionNumber + 1

		_, err = q.WithContext(ctx).
			Where(
				q.ID.Eq(int64(a.ID)),
				q.UserID.Eq(int64(ownerID)),
			).
			Updates(a)
		return err
	})
}

func (g *GeburaRepo) GetApp(ctx context.Context, id libmodel.InternalID) (*modelgebura.App, error) {
	q := query.Use(g.data.db).App
	a, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (g *GeburaRepo) ListApps(
	ctx context.Context,
	paging libmodel.Paging,
	ownerIDs []libmodel.InternalID,
	ids []libmodel.InternalID,
	publicOnly bool,
) ([]*modelgebura.App, int, error) {
	q := query.Use(g.data.db).App
	queryBuilder := q.WithContext(ctx)

	if len(ownerIDs) > 0 {
		castOwnerIDs := make([]int64, len(ownerIDs))
		for i, v := range ownerIDs {
			castOwnerIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.UserID.In(castOwnerIDs...))
	}

	if len(ids) > 0 {
		castIDs := make([]int64, len(ids))
		for i, v := range ids {
			castIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.ID.In(castIDs...))
	}
	if publicOnly {
		queryBuilder = queryBuilder.Where(q.Public.Is(true))
	}

	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}

	ap, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	return ap, int(total), nil
}

func (g *GeburaRepo) BatchCreateAppRunTime(
	ctx context.Context,
	userID libmodel.InternalID,
	runTimes []*modelgebura.AppRunTime,
) error {
	for _, runTime := range runTimes {
		runTime.UserID = userID
	}
	return query.Use(g.data.db).AppRunTime.WithContext(ctx).Create(runTimes...)
}

func (g *GeburaRepo) SumAppRunTime(
	ctx context.Context,
	userID libmodel.InternalID,
	appIDs []libmodel.InternalID,
	deviceIDs []libmodel.InternalID,
	timeRange *libmodel.TimeRange,
) (time.Duration, error) {
	q := query.Use(g.data.db).AppRunTime
	queryBuilder := q.WithContext(ctx).Where(q.UserID.Eq(int64(userID)))

	if len(appIDs) > 0 {
		castAppIDs := make([]int64, len(appIDs))
		for i, v := range appIDs {
			castAppIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.AppID.In(castAppIDs...))
	}
	if len(deviceIDs) > 0 {
		castDeviceIDs := make([]int64, len(deviceIDs))
		for i, v := range deviceIDs {
			castDeviceIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.DeviceID.In(castDeviceIDs...))
	}

	queryBuilder = queryBuilder.Where(
		q.StartTime.Gte(timeRange.StartTime),
		q.StartTime.Lt(timeRange.StartTime.Add(timeRange.Duration)),
	)

	var sum int64
	err := queryBuilder.Select(q.Duration.Sum()).Scan(&sum)
	return time.Duration(sum), err
}

func (g *GeburaRepo) ListAppRunTimes(
	ctx context.Context,
	userID libmodel.InternalID,
	paging libmodel.Paging,
	appIDs []libmodel.InternalID,
	deviceIDs []libmodel.InternalID,
	timeRange *libmodel.TimeRange,
) ([]*modelgebura.AppRunTime, int, error) {
	q := query.Use(g.data.db).AppRunTime
	queryBuilder := q.WithContext(ctx).Where(q.UserID.Eq(int64(userID)))

	if len(appIDs) > 0 {
		castAppIDs := make([]int64, len(appIDs))
		for i, v := range appIDs {
			castAppIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.AppID.In(castAppIDs...))
	}
	if len(deviceIDs) > 0 {
		castDeviceIDs := make([]int64, len(deviceIDs))
		for i, v := range deviceIDs {
			castDeviceIDs[i] = int64(v)
		}
		queryBuilder = queryBuilder.Where(q.DeviceID.In(castDeviceIDs...))
	}

	queryBuilder = queryBuilder.Where(
		q.StartTime.Gte(timeRange.StartTime),
		q.StartTime.Lt(timeRange.StartTime.Add(timeRange.Duration)),
	)

	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}

	res, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}

	return res, int(total), nil
}

func (g *GeburaRepo) DeleteAppRunTime(ctx context.Context, userID libmodel.InternalID, id libmodel.InternalID) error {
	q := query.Use(g.data.db).AppRunTime
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id)), q.UserID.Eq(int64(userID))).
		Delete()
	return err
}

func (g *GeburaRepo) CreateAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	ac *modelgebura.AppCategory,
) error {
	ac.UserID = userID
	return query.Use(g.data.db).AppCategory.WithContext(ctx).Create(ac)
}

func (g *GeburaRepo) ListAppCategories(
	ctx context.Context,
	userID libmodel.InternalID,
) ([]*modelgebura.AppCategory, error) {
	q := query.Use(g.data.db).AppCategory
	return q.WithContext(ctx).Where(q.UserID.Eq(int64(userID))).Find()
}

func (g *GeburaRepo) UpdateAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	ac *modelgebura.AppCategory,
) error {
	q := query.Use(g.data.db).AppCategory
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(ac.ID)), q.UserID.Eq(int64(userID))).
		Updates(ac)
	return err
}

func (g *GeburaRepo) DeleteAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	id libmodel.InternalID,
) error {
	q := query.Use(g.data.db).AppCategory
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id)), q.UserID.Eq(int64(userID))).
		Delete()
	return err
}

func (g *GeburaRepo) CreateSentinel(
	ctx context.Context,
	userID libmodel.InternalID,
	s *modelgebura.Sentinel,
) error {
	s.CreatorID = userID
	return query.Use(g.data.db).Sentinel.WithContext(ctx).Create(s)
}

func (g *GeburaRepo) GetSentinel(
	ctx context.Context,
	id libmodel.InternalID,
) (*modelgebura.Sentinel, error) {
	q := query.Use(g.data.db).Sentinel
	return q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
}

func (g *GeburaRepo) ListSentinels(
	ctx context.Context,
	paging *libmodel.Paging,
) ([]*modelgebura.Sentinel, int, error) {
	q := query.Use(g.data.db).Sentinel
	queryBuilder := q.WithContext(ctx)
	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (g *GeburaRepo) UpdateSentinel(ctx context.Context, s *modelgebura.Sentinel) error {
	q := query.Use(g.data.db).Sentinel
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(s.ID))).
		Updates(s)
	return err
}

func (g *GeburaRepo) CreateSentinelSession(
	ctx context.Context,
	session *modelgebura.SentinelSession,
) error {
	return query.Use(g.data.db).SentinelSession.WithContext(ctx).Create(session)
}

func (g *GeburaRepo) ListSentinelSessions(
	ctx context.Context,
	paging *libmodel.Paging,
	sentinelID libmodel.InternalID,
) ([]*modelgebura.SentinelSession, int, error) {
	q := query.Use(g.data.db).SentinelSession
	queryBuilder := q.WithContext(ctx).Where(q.SentinelID.Eq(int64(sentinelID)))
	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (g *GeburaRepo) UpdateSentinelSessionStatus(
	ctx context.Context,
	id libmodel.InternalID,
	status modelgebura.SentinelSessionStatus,
) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Update(q.Status, status)
	return err
}

func (g *GeburaRepo) DeleteSentinelSession(ctx context.Context, id libmodel.InternalID) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Delete()
	return err
}

func (g *GeburaRepo) GetSentinelSession(
	ctx context.Context,
	userID libmodel.InternalID,
	refreshToken string,
) (*modelgebura.SentinelSession, error) {
	q := query.Use(g.data.db).SentinelSession
	return q.WithContext(ctx).
		Where(
			q.SentinelID.Eq(int64(userID)),
			q.RefreshToken.Eq(refreshToken),
		).
		First()
}

func (g *GeburaRepo) UpdateSentinelSessionToken(
	ctx context.Context,
	id libmodel.InternalID,
	refreshToken string,
	expireAt time.Time,
	refreshedAt time.Time,
) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Updates(map[string]interface{}{
			"refresh_token":     refreshToken,
			"expire_at":         expireAt,
			"last_refreshed_at": &refreshedAt,
			"refresh_count":     gorm.Expr("refresh_count + 1"),
		})
	return err
}

func (g *GeburaRepo) UpdateSentinelSessionLastUsed(
	ctx context.Context,
	id libmodel.InternalID,
	lastUsedAt time.Time,
) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Update(q.LastUsedAt, lastUsedAt)
	return err
}

func (g *GeburaRepo) UpdateSentinelInfo(ctx context.Context, s *modelgebura.Sentinel) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		if _, err := tx.Sentinel.WithContext(ctx).Where(tx.Sentinel.ID.Eq(int64(s.ID))).Updates(s); err != nil {
			return err
		}
		if _, err := tx.SentinelLibrary.WithContext(ctx).Where(tx.SentinelLibrary.SentinelID.Eq(int64(s.ID))).Delete(); err != nil {
			return err
		}
		if len(s.Libraries) > 0 {
			if err := tx.SentinelLibrary.WithContext(ctx).Create(s.Libraries...); err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) UpsertAppBinaries(
	ctx context.Context,
	sentinelID libmodel.InternalID,
	abs []*modelgebura.SentinelAppBinary,
	snapshot *time.Time,
	commit bool,
) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		if len(abs) > 0 {
			err := tx.SentinelAppBinary.WithContext(ctx).
				Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "union_id"}},
					UpdateAll: true,
				}).
				Create(abs...)
			if err != nil {
				return err
			}
		}

		if snapshot != nil {
			libraries, err := tx.SentinelLibrary.WithContext(ctx).
				Where(tx.SentinelLibrary.SentinelID.Eq(int64(sentinelID))).
				Find()
			if err != nil {
				return err
			}
			libIDs := make([]int64, len(libraries))
			for i, l := range libraries {
				libIDs[i] = int64(l.ID)
			}

			if len(libIDs) > 0 {
				_, err = tx.SentinelAppBinary.WithContext(ctx).
					Where(
						tx.SentinelAppBinary.SentinelLibraryID.In(libIDs...),
						tx.SentinelAppBinary.UpdatedAt.Lt(*snapshot),
					).
					Delete()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (g *GeburaRepo) GetStoreApp(
	ctx context.Context,
	id libmodel.InternalID,
) (*modelgebura.StoreApp, error) {
	q := query.Use(g.data.db).StoreApp
	return q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
}

func (g *GeburaRepo) ListStoreApps(
	ctx context.Context,
	paging *libmodel.Paging,
) ([]*modelgebura.StoreApp, int, error) {
	q := query.Use(g.data.db).StoreApp
	queryBuilder := q.WithContext(ctx)
	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}

func (g *GeburaRepo) ListStoreAppBinaries(
	ctx context.Context,
	paging *libmodel.Paging,
	appIDs []libmodel.InternalID,
) ([]*modelgebura.StoreAppBinary, int, error) {
	q := query.Use(g.data.db).StoreAppBinary
	queryBuilder := q.WithContext(ctx)
	if len(appIDs) > 0 {
		castAppIDs := make([]int64, len(appIDs))
		for i, v := range appIDs {
			castAppIDs[i] = int64(v)
		}
		// Assuming standard In check works for ID
		// Note: StoreAppBinary.AppID might be nullable pointer in struct.
		// If Gen handles it correctly, q.AppID.In(...) should accept int64s.
		// If not, we might need manual handling, but let's try standard way.
		queryBuilder = queryBuilder.Where(q.AppID.In(castAppIDs...))
	}
	total, err := queryBuilder.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := queryBuilder.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return res, int(total), nil
}
