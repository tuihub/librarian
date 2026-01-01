package data

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/orm/model"
	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"

	"github.com/samber/lo"
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
	return query.Use(g.data.db).AppInfo.WithContext(ctx).Create(&model.AppInfo{
		ID:                 a.ID,
		Source:             a.Source,
		SourceAppID:        a.SourceAppID,
		SourceURL:          a.SourceURL,
		Name:               a.Name,
		Type:               converter.ToORMAppInfoTypeManual(a.Type),
		ShortDescription:   a.ShortDescription,
		Description:        a.Description,
		IconImageURL:       a.IconImageURL,
		IconImageID:        a.IconImageID,
		BackgroundImageURL: a.BackgroundImageURL,
		BackgroundImageID:  a.BackgroundImageID,
		CoverImageURL:      a.CoverImageURL,
		CoverImageID:       a.CoverImageID,
		ReleaseDate:        a.ReleaseDate,
		Developer:          a.Developer,
		Publisher:          a.Publisher,
		Tags:               a.Tags,
		AlternativeNames:   a.AlternativeNames,
		RawData:            a.RawData,
	})
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
		return converter.ToBizAppInfo(ai), nil
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
		Updates(&model.AppInfo{
			SourceURL:          a.SourceURL,
			Name:               a.Name,
			Type:               converter.ToORMAppInfoTypeManual(a.Type),
			ShortDescription:   a.ShortDescription,
			Description:        a.Description,
			IconImageURL:       a.IconImageURL,
			IconImageID:        a.IconImageID,
			BackgroundImageURL: a.BackgroundImageURL,
			BackgroundImageID:  a.BackgroundImageID,
			CoverImageURL:      a.CoverImageURL,
			CoverImageID:       a.CoverImageID,
			ReleaseDate:        a.ReleaseDate,
			Developer:          a.Developer,
			Publisher:          a.Publisher,
			Tags:               a.Tags,
			AlternativeNames:   a.AlternativeNames,
			RawData:            a.RawData,
		})
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
		typeFilter := make([]string, len(types))
		for i, appType := range types {
			typeFilter[i] = converter.ToORMAppInfoTypeManual(appType)
		}
		queryBuilder = queryBuilder.Where(q.Type.In(typeFilter...))
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

	return converter.ToBizAppInfoList(al), int(total), nil
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
	return converter.ToBizAppInfo(res), nil
}

func (g *GeburaRepo) CreateApp(ctx context.Context, userID libmodel.InternalID, a *modelgebura.App) error {
	q := query.Use(g.data.db).App
	app := &model.App{
		ID:                 a.ID,
		UserID:             userID,
		VersionNumber:      a.VersionNumber,
		VersionDate:        a.VersionDate,
		CreatorDeviceID:    a.CreatorDeviceID,
		AppSources:         a.AppSources,
		Public:             a.Public,
		Name:               a.Name,
		Type:               converter.ToORMAppInfoTypeManual(a.Type),
		ShortDescription:   a.ShortDescription,
		Description:        a.Description,
		IconImageURL:       a.IconImageURL,
		IconImageID:        a.IconImageID,
		BackgroundImageURL: a.BackgroundImageURL,
		BackgroundImageID:  a.BackgroundImageID,
		CoverImageURL:      a.CoverImageURL,
		CoverImageID:       a.CoverImageID,
		ReleaseDate:        a.ReleaseDate,
		Developer:          a.Developer,
		Publisher:          a.Publisher,
		Tags:               a.Tags,
		AlternativeNames:   a.AlternativeNames,
	}
	if a.BoundStoreAppID != nil {
		app.BoundStoreAppID = *a.BoundStoreAppID
	}
	if a.StopStoreManage != nil {
		app.StopStoreManage = *a.StopStoreManage
	}
	return q.WithContext(ctx).Create(app)
}

func (g *GeburaRepo) UpdateApp(ctx context.Context, ownerID libmodel.InternalID, a *modelgebura.App) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.App
		old, err := q.WithContext(ctx).Where(q.ID.Eq(int64(a.ID))).First()
		if err != nil {
			return err
		}

		updates := &model.App{
			VersionNumber:      old.VersionNumber + 1,
			VersionDate:        a.VersionDate,
			AppSources:         a.AppSources,
			Public:             a.Public,
			Name:               a.Name,
			Type:               converter.ToORMAppInfoTypeManual(a.Type),
			ShortDescription:   a.ShortDescription,
			Description:        a.Description,
			IconImageURL:       a.IconImageURL,
			IconImageID:        a.IconImageID,
			BackgroundImageURL: a.BackgroundImageURL,
			BackgroundImageID:  a.BackgroundImageID,
			CoverImageURL:      a.CoverImageURL,
			CoverImageID:       a.CoverImageID,
			ReleaseDate:        a.ReleaseDate,
			Developer:          a.Developer,
			Publisher:          a.Publisher,
			Tags:               a.Tags,
			AlternativeNames:   a.AlternativeNames,
		}
		if a.StopStoreManage != nil {
			updates.StopStoreManage = *a.StopStoreManage
		}

		_, err = q.WithContext(ctx).
			Where(
				q.ID.Eq(int64(a.ID)),
				q.UserID.Eq(int64(ownerID)),
			).
			Updates(updates)
		return err
	})
}

func (g *GeburaRepo) GetApp(ctx context.Context, id libmodel.InternalID) (*modelgebura.App, error) {
	q := query.Use(g.data.db).App
	a, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizApp(a), nil
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

	return converter.ToBizAppList(ap), int(total), nil
}

func (g *GeburaRepo) BatchCreateAppRunTime(
	ctx context.Context,
	userID libmodel.InternalID,
	runTimes []*modelgebura.AppRunTime,
) error {
	rt := make([]*model.AppRunTime, 0, len(runTimes))
	for _, runTime := range runTimes {
		rt = append(rt, &model.AppRunTime{
			ID:        runTime.ID,
			UserID:    userID,
			AppID:     runTime.AppID,
			StartTime: runTime.RunTime.StartTime,
			Duration:  runTime.RunTime.Duration,
		})
	}
	return query.Use(g.data.db).AppRunTime.WithContext(ctx).Create(rt...)
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
		q.StartTime.Lte(timeRange.StartTime.Add(timeRange.Duration)),
	)

	var sum int64
	err := queryBuilder.Select(q.Duration.Sum()).Scan(&sum)
	if err != nil {
		return time.Duration(0), err
	}
	return time.Duration(sum), nil
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
	if timeRange != nil {
		queryBuilder = queryBuilder.Where(
			q.StartTime.Gte(timeRange.StartTime),
			q.StartTime.Lte(timeRange.StartTime.Add(timeRange.Duration)),
		)
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

	return converter.ToBizAppRunTimeList(res), int(total), nil
}

func (g *GeburaRepo) DeleteAppRunTime(ctx context.Context, userID libmodel.InternalID, id libmodel.InternalID) error {
	q := query.Use(g.data.db).AppRunTime
	_, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).Delete()
	return err
}

func (g *GeburaRepo) CreateAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	ac *modelgebura.AppCategory,
) error {
	cat := &model.AppCategory{
		ID:            ac.ID,
		UserID:        userID,
		VersionNumber: ac.VersionNumber,
		VersionDate:   ac.VersionDate,
		Name:          ac.Name,
	}
	joinEntries := make([]model.AppAppCategory, len(ac.AppIDs))
	for i, appID := range ac.AppIDs {
		joinEntries[i] = model.AppAppCategory{
			AppCategoryID: ac.ID,
			AppID:         appID,
		}
	}
	cat.AppAppCategories = joinEntries

	return query.Use(g.data.db).AppCategory.WithContext(ctx).Create(cat)
}

func (g *GeburaRepo) ListAppCategories(
	ctx context.Context,
	userID libmodel.InternalID,
) ([]*modelgebura.AppCategory, error) {
	q := query.Use(g.data.db).AppCategory
	acs, err := q.WithContext(ctx).
		Preload(q.AppAppCategories).
		Where(q.UserID.Eq(int64(userID))).
		Find()
	if err != nil {
		return nil, err
	}

	res := make([]*modelgebura.AppCategory, len(acs))
	for i := range acs {
		res[i] = converter.ToBizAppCategoryExtend(acs[i])
	}
	return res, nil
}

func (g *GeburaRepo) UpdateAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	ac *modelgebura.AppCategory) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.AppCategory
		old, err := q.WithContext(ctx).
			Where(
				q.ID.Eq(int64(ac.ID)),
				q.UserID.Eq(int64(userID)),
			).
			First()
		if err != nil {
			return err
		}

		// Update fields
		_, err = q.WithContext(ctx).
			Where(q.ID.Eq(int64(ac.ID))).
			Updates(&model.AppCategory{
				Name:          ac.Name,
				VersionNumber: old.VersionNumber + 1,
				VersionDate:   time.Now(),
			})
		if err != nil {
			return err
		}

		// Update Associations
		qa := tx.AppAppCategory
		_, err = qa.WithContext(ctx).Where(qa.AppCategoryID.Eq(int64(ac.ID))).Delete()
		if err != nil {
			return err
		}

		// Insert new
		if len(ac.AppIDs) > 0 {
			newEntries := make([]*model.AppAppCategory, len(ac.AppIDs))
			for i, appID := range ac.AppIDs {
				newEntries[i] = &model.AppAppCategory{
					AppCategoryID: ac.ID,
					AppID:         appID,
				}
			}
			err = qa.WithContext(ctx).Create(newEntries...)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) DeleteAppCategory(
	ctx context.Context,
	userID libmodel.InternalID,
	id libmodel.InternalID,
) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		q := tx.AppCategory
		qa := tx.AppAppCategory

		count, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id)), q.UserID.Eq(int64(userID))).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			return nil
		}

		_, err = qa.WithContext(ctx).Where(qa.AppCategoryID.Eq(int64(id))).Delete()
		if err != nil {
			return err
		}

		_, err = q.WithContext(ctx).Where(q.ID.Eq(int64(id))).Delete()
		return err
	})
}

func (g *GeburaRepo) CreateSentinel(ctx context.Context, userID libmodel.InternalID, s *modelgebura.Sentinel) error {
	return query.Use(g.data.db).Sentinel.WithContext(ctx).Create(&model.Sentinel{
		ID:          s.ID,
		CreatorID:   userID,
		Name:        s.Name,
		Description: s.Description,
	})
}

func (g *GeburaRepo) GetSentinel(ctx context.Context, id libmodel.InternalID) (*modelgebura.Sentinel, error) {
	s, err := query.Use(g.data.db).Sentinel.WithContext(ctx).Where(query.Sentinel.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizSentinel(s), nil
}

func (g *GeburaRepo) ListSentinels(ctx context.Context, page *libmodel.Paging) ([]*modelgebura.Sentinel, int, error) {
	q := query.Use(g.data.db).Sentinel
	sentinels, err := q.WithContext(ctx).
		Limit(page.ToLimit()).
		Offset(page.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	total, err := q.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizSentinelList(sentinels), int(total), nil
}

func (g *GeburaRepo) UpdateSentinel(ctx context.Context, s *modelgebura.Sentinel) error {
	q := query.Use(g.data.db).Sentinel
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(s.ID))).
		Updates(&model.Sentinel{
			Name:        s.Name,
			Description: s.Description,
		})
	return err
}

func (g *GeburaRepo) CreateSentinelSession(ctx context.Context, ss *modelgebura.SentinelSession) error {
	return query.Use(g.data.db).SentinelSession.WithContext(ctx).Create(&model.SentinelSession{
		ID:           ss.ID,
		SentinelID:   ss.SentinelID,
		RefreshToken: ss.RefreshToken,
		Status:       converter.ToORMSentinelSessionStatus(ss.Status),
		CreatorID:    ss.CreatorID,
		ExpireAt:     ss.ExpireAt,
	})
}

func (g *GeburaRepo) GetSentinelSession(
	ctx context.Context,
	sentinelID libmodel.InternalID,
	refreshToken string,
) (*modelgebura.SentinelSession, error) {
	q := query.Use(g.data.db).SentinelSession
	s, err := q.WithContext(ctx).
		Where(
			q.SentinelID.Eq(int64(sentinelID)),
			q.RefreshToken.Eq(refreshToken),
		).
		First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizSentinelSession(s), nil
}

func (g *GeburaRepo) ListSentinelSessions(
	ctx context.Context,
	page *libmodel.Paging,
	sentinelID libmodel.InternalID,
) ([]*modelgebura.SentinelSession, int, error) {
	q := query.Use(g.data.db).SentinelSession
	sessions, err := q.WithContext(ctx).
		Where(q.SentinelID.Eq(int64(sentinelID))).
		Limit(page.ToLimit()).
		Offset(page.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	total, err := q.WithContext(ctx).
		Where(q.SentinelID.Eq(int64(sentinelID))).
		Count()
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizSentinelSessionList(sessions), int(total), nil
}

func (g *GeburaRepo) UpdateSentinelSessionStatus(
	ctx context.Context,
	id libmodel.InternalID,
	status modelgebura.SentinelSessionStatus,
) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Update(q.Status, converter.ToORMSentinelSessionStatus(status))
	return err
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
			"last_refreshed_at": refreshedAt,
			"refresh_count":     gorm.Expr("refresh_count + ?", 1),
		})
	return err
}

func (g *GeburaRepo) UpdateSentinelSessionLastUsed(
	ctx context.Context,
	id libmodel.InternalID,
	usedAt time.Time,
) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Update(q.LastUsedAt, usedAt)
	return err
}

func (g *GeburaRepo) DeleteSentinelSession(ctx context.Context, id libmodel.InternalID) error {
	q := query.Use(g.data.db).SentinelSession
	_, err := q.WithContext(ctx).Where(q.ID.Eq(int64(id))).Delete()
	return err
}

func (g *GeburaRepo) UpdateSentinelInfo(
	ctx context.Context,
	s *modelgebura.Sentinel,
) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		// update sentinel info
		_, err := tx.Sentinel.WithContext(ctx).
			Where(tx.Sentinel.ID.Eq(int64(s.ID))).
			Updates(map[string]interface{}{
				"url":                     s.URL,
				"alternative_urls":        s.AlternativeUrls,
				"get_token_path":          s.GetTokenPath,
				"download_file_base_path": s.DownloadFileBasePath,
				"library_report_sequence": gorm.Expr("library_report_sequence + ?", 1),
			})
		if err != nil {
			return err
		}

		// upsert libraries
		sInfo, err := tx.Sentinel.WithContext(ctx).Where(tx.Sentinel.ID.Eq(int64(s.ID))).First()
		if err != nil {
			return err
		}

		newLibs := make([]*model.SentinelLibrary, 0, len(s.Libraries))
		for _, lib := range s.Libraries {
			newLibs = append(newLibs, &model.SentinelLibrary{
				ID:                    lib.ID,
				SentinelID:            sInfo.ID,
				ReportedID:            lib.ReportedID,
				DownloadBasePath:      lib.DownloadBasePath,
				LibraryReportSequence: sInfo.LibraryReportSequence,
			})
		}

		return tx.SentinelLibrary.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "sentinel_id"}, {Name: "reported_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"download_base_path", "library_report_sequence",
			}),
		}).Create(newLibs...)
	})
}

func (g *GeburaRepo) UpsertAppBinaries( //nolint:gocognit,funlen // complex logic
	ctx context.Context,
	sentinelID libmodel.InternalID,
	abs []*modelgebura.SentinelAppBinary,
	snapshot *time.Time,
	commit bool,
) error {
	return g.data.WithTx(ctx, func(tx *query.Query) error {
		sInfo, err := tx.Sentinel.WithContext(ctx).
			Where(tx.Sentinel.ID.Eq(int64(sentinelID))).
			Preload(tx.Sentinel.SentinelLibraries).
			First()
		if err != nil {
			return err
		}

		libraryMap := make(map[int64]*model.SentinelLibrary)
		for _, lib := range sInfo.SentinelLibraries {
			libraryMap[lib.ReportedID] = &lib
		}
		for _, ab := range abs {
			if _, ok := libraryMap[ab.SentinelLibraryID]; !ok {
				return errors.New("library not found")
			}
		}

		// upsert binaries
		newAbs := make([]*model.SentinelAppBinary, 0, len(abs))
		for _, ab := range abs {
			newAbs = append(newAbs, &model.SentinelAppBinary{
				ID:                        ab.ID,
				UnionID:                   ab.UnionID,
				SentinelID:                sentinelID,
				SentinelLibraryReportedID: ab.SentinelLibraryID,
				LibrarySnapshot: lo.FromPtrOr(
					snapshot,
					lo.FromPtr(libraryMap[ab.SentinelLibraryID].ActiveSnapshot),
				),
				GeneratedID: ab.GeneratedID,
				SizeBytes:   ab.SizeBytes,
				NeedToken:   ab.NeedToken,
				Name:        ab.Name,
				Version:     ab.Version,
				Developer:   ab.Developer,
				Publisher:   ab.Publisher,
			})
		}

		err = tx.SentinelAppBinary.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "sentinel_id"},
				{Name: "sentinel_library_reported_id"},
				{Name: "library_snapshot"},
				{Name: "generated_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"union_id", "size_bytes", "need_token", "name", "version", "developer", "publisher",
			}),
		}).Create(newAbs...)
		if err != nil {
			return err
		}

		// upsert binary files
		abfCount := lo.Sum(lo.Map(abs, func(ab *modelgebura.SentinelAppBinary, _ int) int {
			return len(ab.Files)
		}))
		newAbfs := make([]*model.SentinelAppBinaryFile, 0, abfCount)
		for _, ab := range abs {
			for _, f := range ab.Files {
				newAbfs = append(newAbfs, &model.SentinelAppBinaryFile{
					ID:                        f.ID,
					SentinelID:                sentinelID,
					SentinelLibraryReportedID: ab.SentinelLibraryID,
					LibrarySnapshot: lo.FromPtrOr(
						snapshot,
						lo.FromPtr(libraryMap[ab.SentinelLibraryID].ActiveSnapshot),
					),
					SentinelAppBinaryGeneratedID: ab.GeneratedID,
					Name:                         f.Name,
					SizeBytes:                    f.SizeBytes,
					Sha256:                       f.Sha256,
					ServerFilePath:               f.ServerFilePath,
					ChunksInfo:                   f.ChunksInfo,
				})
			}
		}

		err = tx.SentinelAppBinaryFile.WithContext(ctx).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "sentinel_id"},
				{Name: "sentinel_library_reported_id"},
				{Name: "library_snapshot"},
				{Name: "sentinel_app_binary_generated_id"},
				{Name: "server_file_path"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "size_bytes", "sha256", "chunks_info",
			}),
		}).Create(newAbfs...)
		if err != nil {
			return err
		}

		if snapshot != nil && commit {
			ids := make([]int64, 0, len(libraryMap))
			for _, lib := range libraryMap {
				ids = append(ids, int64(lib.ID))
			}
			_, err = tx.SentinelLibrary.WithContext(ctx).
				Where(tx.SentinelLibrary.ID.In(ids...)).
				Updates(&model.SentinelLibrary{
					ActiveSnapshot: snapshot,
				})
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) GetStoreApp(ctx context.Context, id libmodel.InternalID) (*modelgebura.StoreApp, error) {
	a, err := query.Use(g.data.db).StoreApp.WithContext(ctx).Where(query.StoreApp.ID.Eq(int64(id))).First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizStoreApp(a), nil
}

func (g *GeburaRepo) ListStoreApps(ctx context.Context, page *libmodel.Paging) ([]*modelgebura.StoreApp, int, error) {
	q := query.Use(g.data.db).StoreApp
	total, err := q.WithContext(ctx).Count()
	if err != nil {
		return nil, 0, err
	}
	storeApps, err := q.WithContext(ctx).
		Limit(page.ToLimit()).
		Offset(page.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizStoreAppList(storeApps), int(total), nil
}

func (g *GeburaRepo) ListStoreAppBinaries(
	ctx context.Context,
	page *libmodel.Paging,
	appIDs []libmodel.InternalID,
) ([]*modelgebura.StoreAppBinary, int, error) {
	sab := query.Use(g.data.db).SentinelAppBinary
	var storeAppBinaries []*model.SentinelAppBinary

	db := sab.WithContext(ctx).UnderlyingDB()

	// Join with SentinelLibrary to check active snapshot
	db = db.Joins(
		"JOIN sentinel_libraries ON sentinel_app_binaries.sentinel_library_reported_id = sentinel_libraries.reported_id AND sentinel_app_binaries.library_snapshot = sentinel_libraries.active_snapshot",
	)

	if len(appIDs) > 0 {
		db = db.Joins("StoreApps").Where("store_apps.id IN ?", appIDs)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset(page.ToOffset()).Limit(page.ToLimit()).Find(&storeAppBinaries).Error
	if err != nil {
		return nil, 0, err
	}

	return converter.ToBizStoreAppBinaryList(storeAppBinaries), int(total), nil
}
