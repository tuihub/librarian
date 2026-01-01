package data

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
	"github.com/tuihub/librarian/internal/model"
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
	appInfo := gormschema.AppInfo{
		ID:                 a.ID,
		Source:             a.Source,
		SourceAppID:        a.SourceAppID,
		SourceURL:          a.SourceURL,
		Name:               a.Name,
		Type:               gormschema.ToSchemaAppType(a.Type),
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
		Tags:               gormschema.StringArrayVal(a.Tags),
		AlternativeNames:   gormschema.StringArrayVal(a.AlternativeNames),
		RawData:            a.RawData,
	}
	return g.data.db.WithContext(ctx).Create(&appInfo).Error
}

func (g *GeburaRepo) CreateAppInfoOrGet(ctx context.Context, a *modelgebura.AppInfo) (*modelgebura.AppInfo, error) {
	err := g.CreateAppInfo(ctx, a)
	if err == nil {
		return a, nil
	}
	// Check if it's a constraint error (duplicate)
	var ai gormschema.AppInfo
	if findErr := g.data.db.WithContext(ctx).
		Where("source = ? AND source_app_id = ?", a.Source, a.SourceAppID).
		First(&ai).Error; findErr == nil {
		return gormschema.ToBizAppInfo(&ai), nil
	}
	return nil, err
}

func (g *GeburaRepo) UpdateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	updates := map[string]any{
		"source_url":           a.SourceURL,
		"name":                 a.Name,
		"type":                 gormschema.ToSchemaAppType(a.Type),
		"short_description":    a.ShortDescription,
		"description":          a.Description,
		"icon_image_url":       a.IconImageURL,
		"icon_image_id":        a.IconImageID,
		"background_image_url": a.BackgroundImageURL,
		"background_image_id":  a.BackgroundImageID,
		"cover_image_url":      a.CoverImageURL,
		"cover_image_id":       a.CoverImageID,
		"release_date":         a.ReleaseDate,
		"developer":            a.Developer,
		"publisher":            a.Publisher,
		"tags":                 gormschema.StringArrayVal(a.Tags),
		"alternative_names":    gormschema.StringArrayVal(a.AlternativeNames),
		"raw_data":             a.RawData,
	}
	return g.data.db.WithContext(ctx).
		Model(&gormschema.AppInfo{}).
		Where("id = ? AND source = ? AND source_app_id = ?", a.ID, a.Source, a.SourceAppID).
		Updates(updates).Error
}

func (g *GeburaRepo) ListAppInfos(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
) ([]*modelgebura.AppInfo, int64, error) {
	var al []gormschema.AppInfo
	var total int64

	err := g.data.WithTx(ctx, func(tx *gorm.DB) error {
		query := tx.Model(&gormschema.AppInfo{})
		if len(sources) > 0 {
			query = query.Where("source IN ?", sources)
		}
		if len(types) > 0 {
			typeStrs := make([]string, len(types))
			for i, t := range types {
				typeStrs[i] = gormschema.ToSchemaAppType(t)
			}
			query = query.Where("type IN ?", typeStrs)
		}
		if len(ids) > 0 {
			query = query.Where("id IN ?", ids)
		}
		if err := query.Count(&total).Error; err != nil {
			return err
		}
		return query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&al).Error
	})
	if err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.AppInfo, len(al))
	for i := range al {
		res[i] = gormschema.ToBizAppInfo(&al[i])
	}
	return res, total, nil
}

func (g *GeburaRepo) GetAppInfo(ctx context.Context, id modelgebura.AppInfoID) (*modelgebura.AppInfo, error) {
	var ai gormschema.AppInfo
	err := g.data.db.WithContext(ctx).
		Where("source = ? AND source_app_id = ?", id.Source, id.SourceAppID).
		First(&ai).Error
	if err != nil {
		return nil, err
	}
	return gormschema.ToBizAppInfo(&ai), nil
}

func (g *GeburaRepo) CreateApp(ctx context.Context, userID model.InternalID, a *modelgebura.App) error {
	app := gormschema.App{
		ID:                 a.ID,
		UserID:             userID,
		VersionNumber:      a.VersionNumber,
		VersionDate:        a.VersionDate,
		CreatorDeviceID:    a.CreatorDeviceID,
		AppSources:         gormschema.StringMapVal(a.AppSources),
		Public:             a.Public,
		BoundStoreAppID:    a.BoundStoreAppID,
		StopStoreManage:    a.StopStoreManage,
		Name:               a.Name,
		Type:               gormschema.ToSchemaAppType(a.Type),
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
		Tags:               gormschema.StringArrayVal(a.Tags),
		AlternativeNames:   gormschema.StringArrayVal(a.AlternativeNames),
	}
	return g.data.db.WithContext(ctx).Create(&app).Error
}

func (g *GeburaRepo) UpdateApp(ctx context.Context, ownerID model.InternalID, a *modelgebura.App) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		var old gormschema.App
		if err := tx.First(&old, a.ID).Error; err != nil {
			return err
		}
		updates := map[string]any{
			"version_number":       old.VersionNumber + 1,
			"version_date":         a.VersionDate,
			"app_sources":          gormschema.StringMapVal(a.AppSources),
			"public":               a.Public,
			"name":                 a.Name,
			"type":                 gormschema.ToSchemaAppType(a.Type),
			"short_description":    a.ShortDescription,
			"description":          a.Description,
			"icon_image_url":       a.IconImageURL,
			"icon_image_id":        a.IconImageID,
			"background_image_url": a.BackgroundImageURL,
			"background_image_id":  a.BackgroundImageID,
			"cover_image_url":      a.CoverImageURL,
			"cover_image_id":       a.CoverImageID,
			"release_date":         a.ReleaseDate,
			"developer":            a.Developer,
			"publisher":            a.Publisher,
			"tags":                 gormschema.StringArrayVal(a.Tags),
			"alternative_names":    gormschema.StringArrayVal(a.AlternativeNames),
		}
		if a.StopStoreManage != nil {
			updates["stop_store_manage"] = *a.StopStoreManage
		}
		return tx.Model(&gormschema.App{}).
			Where("id = ? AND user_id = ?", a.ID, ownerID).
			Updates(updates).Error
	})
}

func (g *GeburaRepo) GetApp(ctx context.Context, id model.InternalID) (*modelgebura.App, error) {
	var a gormschema.App
	if err := g.data.db.WithContext(ctx).First(&a, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizApp(&a), nil
}

func (g *GeburaRepo) ListApps(
	ctx context.Context,
	paging model.Paging,
	ownerIDs []model.InternalID,
	ids []model.InternalID,
	publicOnly bool,
) ([]*modelgebura.App, int, error) {
	query := g.data.db.WithContext(ctx).Model(&gormschema.App{}).
		Where("user_id IN ?", ownerIDs)

	if len(ids) > 0 {
		query = query.Where("id IN ?", ids)
	}
	if publicOnly {
		query = query.Where("public = ?", true)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var apps []gormschema.App
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&apps).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.App, len(apps))
	for i := range apps {
		res[i] = gormschema.ToBizApp(&apps[i])
	}
	return res, int(total), nil
}

func (g *GeburaRepo) BatchCreateAppRunTime(
	ctx context.Context,
	userID model.InternalID,
	runTimes []*modelgebura.AppRunTime,
) error {
	rts := make([]gormschema.AppRunTime, 0, len(runTimes))
	for _, rt := range runTimes {
		rts = append(rts, gormschema.AppRunTime{
			ID:        rt.ID,
			UserID:    userID,
			AppID:     rt.AppID,
			StartTime: rt.RunTime.StartTime,
			Duration:  rt.RunTime.Duration,
		})
	}
	return g.data.db.WithContext(ctx).Create(&rts).Error
}

func (g *GeburaRepo) SumAppRunTime(
	ctx context.Context,
	userID model.InternalID,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	var result struct {
		Sum int64
	}
	query := g.data.db.WithContext(ctx).Model(&gormschema.AppRunTime{}).
		Select("COALESCE(SUM(duration), 0) as sum").
		Where("user_id = ?", userID)

	if len(appIDs) > 0 {
		query = query.Where("app_id IN ?", appIDs)
	}
	if len(deviceIDs) > 0 {
		query = query.Where("device_id IN ?", deviceIDs)
	}
	if timeRange != nil {
		query = query.Where("start_time >= ? AND start_time <= ?",
			timeRange.StartTime, timeRange.StartTime.Add(timeRange.Duration))
	}

	if err := query.Scan(&result).Error; err != nil {
		return 0, err
	}
	return time.Duration(result.Sum), nil
}

func (g *GeburaRepo) ListAppRunTimes(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) ([]*modelgebura.AppRunTime, int, error) {
	query := g.data.db.WithContext(ctx).Model(&gormschema.AppRunTime{}).
		Where("user_id = ?", userID)

	if len(appIDs) > 0 {
		query = query.Where("app_id IN ?", appIDs)
	}
	if len(deviceIDs) > 0 {
		query = query.Where("device_id IN ?", deviceIDs)
	}
	if timeRange != nil {
		query = query.Where("start_time >= ? AND start_time <= ?",
			timeRange.StartTime, timeRange.StartTime.Add(timeRange.Duration))
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var runtimes []gormschema.AppRunTime
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&runtimes).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.AppRunTime, len(runtimes))
	for i := range runtimes {
		res[i] = gormschema.ToBizAppRunTime(&runtimes[i])
	}
	return res, int(total), nil
}

func (g *GeburaRepo) DeleteAppRunTime(ctx context.Context, userID model.InternalID, id model.InternalID) error {
	return g.data.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&gormschema.AppRunTime{}).Error
}

func (g *GeburaRepo) CreateAppCategory(
	ctx context.Context,
	userID model.InternalID,
	ac *modelgebura.AppCategory,
) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		category := gormschema.AppCategory{
			ID:            ac.ID,
			UserID:        userID,
			VersionNumber: ac.VersionNumber,
			VersionDate:   ac.VersionDate,
			Name:          ac.Name,
		}
		if err := tx.Create(&category).Error; err != nil {
			return err
		}
		// Create associations
		for _, appID := range ac.AppIDs {
			aac := gormschema.AppAppCategory{
				AppID:         appID,
				AppCategoryID: ac.ID,
			}
			if err := tx.Create(&aac).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) ListAppCategories(
	ctx context.Context,
	userID model.InternalID,
) ([]*modelgebura.AppCategory, error) {
	var categories []gormschema.AppCategory
	if err := g.data.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&categories).Error; err != nil {
		return nil, err
	}

	res := make([]*modelgebura.AppCategory, len(categories))
	for i := range categories {
		res[i] = gormschema.ToBizAppCategory(&categories[i])
		// Get app IDs for this category
		var aacs []gormschema.AppAppCategory
		if err := g.data.db.WithContext(ctx).
			Where("app_category_id = ?", categories[i].ID).
			Find(&aacs).Error; err == nil {
			res[i].AppIDs = make([]model.InternalID, len(aacs))
			for j, aac := range aacs {
				res[i].AppIDs[j] = aac.AppID
			}
		}
	}
	return res, nil
}

func (g *GeburaRepo) UpdateAppCategory(
	ctx context.Context,
	userID model.InternalID,
	ac *modelgebura.AppCategory) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Get old version
		var old gormschema.AppCategory
		if err := tx.Where("id = ? AND user_id = ?", ac.ID, userID).First(&old).Error; err != nil {
			return err
		}
		// Remove existing associations
		if err := tx.Where("app_category_id = ?", ac.ID).Delete(&gormschema.AppAppCategory{}).Error; err != nil {
			return err
		}
		// Update category
		updates := map[string]any{
			"name":           ac.Name,
			"version_number": old.VersionNumber + 1,
			"version_date":   time.Now(),
		}
		if err := tx.Model(&gormschema.AppCategory{}).
			Where("id = ? AND user_id = ?", ac.ID, userID).
			Updates(updates).Error; err != nil {
			return err
		}
		// Create new associations
		for _, appID := range ac.AppIDs {
			aac := gormschema.AppAppCategory{
				AppID:         appID,
				AppCategoryID: ac.ID,
			}
			if err := tx.Create(&aac).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) DeleteAppCategory(
	ctx context.Context,
	userID model.InternalID,
	id model.InternalID,
) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Delete associations first
		if err := tx.Where("app_category_id = ?", id).Delete(&gormschema.AppAppCategory{}).Error; err != nil {
			return err
		}
		// Delete category
		return tx.Where("id = ? AND user_id = ?", id, userID).Delete(&gormschema.AppCategory{}).Error
	})
}

func (g *GeburaRepo) CreateSentinel(ctx context.Context, userID model.InternalID, s *modelgebura.Sentinel) error {
	sentinel := gormschema.Sentinel{
		ID:          s.ID,
		CreatorID:   userID,
		Name:        s.Name,
		Description: s.Description,
	}
	return g.data.db.WithContext(ctx).Create(&sentinel).Error
}

func (g *GeburaRepo) GetSentinel(ctx context.Context, id model.InternalID) (*modelgebura.Sentinel, error) {
	var s gormschema.Sentinel
	if err := g.data.db.WithContext(ctx).First(&s, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizSentinel(&s), nil
}

func (g *GeburaRepo) ListSentinels(ctx context.Context, page *model.Paging) ([]*modelgebura.Sentinel, int, error) {
	var sentinels []gormschema.Sentinel
	query := g.data.db.WithContext(ctx).Model(&gormschema.Sentinel{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(page.ToLimit()).Offset(page.ToOffset()).Find(&sentinels).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.Sentinel, len(sentinels))
	for i := range sentinels {
		res[i] = gormschema.ToBizSentinel(&sentinels[i])
	}
	return res, int(total), nil
}

func (g *GeburaRepo) UpdateSentinel(ctx context.Context, s *modelgebura.Sentinel) error {
	return g.data.db.WithContext(ctx).
		Model(&gormschema.Sentinel{}).
		Where("id = ?", s.ID).
		Updates(map[string]any{
			"name":        s.Name,
			"description": s.Description,
		}).Error
}

func (g *GeburaRepo) CreateSentinelSession(ctx context.Context, ss *modelgebura.SentinelSession) error {
	session := gormschema.SentinelSession{
		ID:           ss.ID,
		SentinelID:   ss.SentinelID,
		RefreshToken: ss.RefreshToken,
		Status:       gormschema.ToSchemaSentinelSessionStatus(ss.Status),
		CreatorID:    ss.CreatorID,
		ExpireAt:     ss.ExpireAt,
	}
	return g.data.db.WithContext(ctx).Create(&session).Error
}

func (g *GeburaRepo) GetSentinelSession(
	ctx context.Context,
	sentinelID model.InternalID,
	refreshToken string,
) (*modelgebura.SentinelSession, error) {
	var s gormschema.SentinelSession
	err := g.data.db.WithContext(ctx).
		Where("sentinel_id = ? AND refresh_token = ?", sentinelID, refreshToken).
		First(&s).Error
	if err != nil {
		return nil, err
	}
	return gormschema.ToBizSentinelSession(&s), nil
}

func (g *GeburaRepo) ListSentinelSessions(
	ctx context.Context,
	page *model.Paging,
	sentinelID model.InternalID,
) ([]*modelgebura.SentinelSession, int, error) {
	query := g.data.db.WithContext(ctx).Model(&gormschema.SentinelSession{}).
		Where("sentinel_id = ?", sentinelID)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var sessions []gormschema.SentinelSession
	if err := query.Limit(page.ToLimit()).Offset(page.ToOffset()).Find(&sessions).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.SentinelSession, len(sessions))
	for i := range sessions {
		res[i] = gormschema.ToBizSentinelSession(&sessions[i])
	}
	return res, int(total), nil
}

func (g *GeburaRepo) UpdateSentinelSessionStatus(
	ctx context.Context,
	id model.InternalID,
	status modelgebura.SentinelSessionStatus,
) error {
	return g.data.db.WithContext(ctx).
		Model(&gormschema.SentinelSession{}).
		Where("id = ?", id).
		Update("status", gormschema.ToSchemaSentinelSessionStatus(status)).Error
}

func (g *GeburaRepo) UpdateSentinelSessionToken(
	ctx context.Context,
	id model.InternalID,
	refreshToken string,
	expireAt time.Time,
	refreshedAt time.Time,
) error {
	return g.data.db.WithContext(ctx).
		Model(&gormschema.SentinelSession{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"refresh_token":     refreshToken,
			"expire_at":         expireAt,
			"last_refreshed_at": refreshedAt,
			"refresh_count":     gorm.Expr("refresh_count + 1"),
		}).Error
}

func (g *GeburaRepo) UpdateSentinelSessionLastUsed(
	ctx context.Context,
	id model.InternalID,
	usedAt time.Time,
) error {
	return g.data.db.WithContext(ctx).
		Model(&gormschema.SentinelSession{}).
		Where("id = ?", id).
		Update("last_used_at", usedAt).Error
}

func (g *GeburaRepo) DeleteSentinelSession(ctx context.Context, id model.InternalID) error {
	return g.data.db.WithContext(ctx).Delete(&gormschema.SentinelSession{}, id).Error
}

func (g *GeburaRepo) UpdateSentinelInfo(
	ctx context.Context,
	s *modelgebura.Sentinel,
) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Update sentinel info
		if err := tx.Model(&gormschema.Sentinel{}).
			Where("id = ?", s.ID).
			Updates(map[string]any{
				"url":                     s.URL,
				"alternative_urls":        gormschema.StringArrayVal(s.AlternativeUrls),
				"get_token_path":          s.GetTokenPath,
				"download_file_base_path": s.DownloadFileBasePath,
				"library_report_sequence": gorm.Expr("library_report_sequence + 1"),
			}).Error; err != nil {
			return err
		}

		// Get updated sentinel
		var sInfo gormschema.Sentinel
		if err := tx.First(&sInfo, s.ID).Error; err != nil {
			return err
		}

		// Upsert libraries
		for _, lib := range s.Libraries {
			newLib := gormschema.SentinelLibrary{
				ID:                    lib.ID,
				SentinelID:            sInfo.ID,
				ReportedID:            lib.ReportedID,
				DownloadBasePath:      lib.DownloadBasePath,
				LibraryReportSequence: sInfo.LibraryReportSequence,
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "sentinel_id"}, {Name: "reported_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"download_base_path", "library_report_sequence", "updated_at"}),
			}).Create(&newLib).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) UpsertAppBinaries( //nolint:gocognit,funlen // TODO
	ctx context.Context,
	sentinelID model.InternalID,
	abs []*modelgebura.SentinelAppBinary,
	snapshot *time.Time,
	commit bool,
) error {
	return g.data.WithTx(ctx, func(tx *gorm.DB) error {
		// Get sentinel with libraries
		var sInfo gormschema.Sentinel
		if err := tx.First(&sInfo, sentinelID).Error; err != nil {
			return err
		}

		var libraries []gormschema.SentinelLibrary
		if err := tx.Where("sentinel_id = ?", sentinelID).Find(&libraries).Error; err != nil {
			return err
		}

		libraryMap := make(map[int64]*gormschema.SentinelLibrary)
		for i := range libraries {
			libraryMap[libraries[i].ReportedID] = &libraries[i]
		}

		for _, ab := range abs {
			if _, ok := libraryMap[ab.SentinelLibraryID]; !ok {
				return errors.New("library not found")
			}
		}

		// Upsert binaries
		for _, ab := range abs {
			snapshotTime := lo.FromPtrOr(snapshot, libraryMap[ab.SentinelLibraryID].ActiveSnapshot)
			newAb := gormschema.SentinelAppBinary{
				ID:                        ab.ID,
				UnionID:                   ab.UnionID,
				SentinelID:                sentinelID,
				SentinelLibraryReportedID: ab.SentinelLibraryID,
				LibrarySnapshot:           snapshotTime,
				GeneratedID:               ab.GeneratedID,
				SizeBytes:                 ab.SizeBytes,
				NeedToken:                 ab.NeedToken,
				Name:                      ab.Name,
				Version:                   ab.Version,
				Developer:                 ab.Developer,
				Publisher:                 ab.Publisher,
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "sentinel_id"}, {Name: "sentinel_library_reported_id"}, {Name: "library_snapshot"}, {Name: "generated_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"union_id", "size_bytes", "need_token", "name", "version", "developer", "publisher", "updated_at"}),
			}).Create(&newAb).Error; err != nil {
				return err
			}

			// Upsert binary files
			for _, f := range ab.Files {
				newAbf := gormschema.SentinelAppBinaryFile{
					ID:                           f.ID,
					SentinelID:                   sentinelID,
					SentinelLibraryReportedID:    ab.SentinelLibraryID,
					LibrarySnapshot:              snapshotTime,
					SentinelAppBinaryGeneratedID: ab.GeneratedID,
					Name:                         f.Name,
					SizeBytes:                    f.SizeBytes,
					Sha256:                       f.Sha256,
					ServerFilePath:               f.ServerFilePath,
					ChunksInfo:                   f.ChunksInfo,
				}
				if err := tx.Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "sentinel_id"}, {Name: "sentinel_library_reported_id"}, {Name: "library_snapshot"}, {Name: "sentinel_app_binary_generated_id"}, {Name: "server_file_path"}},
					DoUpdates: clause.AssignmentColumns([]string{"name", "size_bytes", "sha256", "chunks_info", "updated_at"}),
				}).Create(&newAbf).Error; err != nil {
					return err
				}
			}
		}

		if snapshot != nil && commit {
			ids := make([]model.InternalID, 0, len(libraryMap))
			for _, lib := range libraryMap {
				ids = append(ids, lib.ID)
			}
			if err := tx.Model(&gormschema.SentinelLibrary{}).
				Where("id IN ?", ids).
				Update("active_snapshot", *snapshot).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *GeburaRepo) GetStoreApp(ctx context.Context, id model.InternalID) (*modelgebura.StoreApp, error) {
	var a gormschema.StoreApp
	if err := g.data.db.WithContext(ctx).First(&a, id).Error; err != nil {
		return nil, err
	}
	return gormschema.ToBizStoreApp(&a), nil
}

func (g *GeburaRepo) ListStoreApps(ctx context.Context, page *model.Paging) ([]*modelgebura.StoreApp, int, error) {
	query := g.data.db.WithContext(ctx).Model(&gormschema.StoreApp{})

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var storeApps []gormschema.StoreApp
	if err := query.Limit(page.ToLimit()).Offset(page.ToOffset()).Find(&storeApps).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.StoreApp, len(storeApps))
	for i := range storeApps {
		res[i] = gormschema.ToBizStoreApp(&storeApps[i])
	}
	return res, int(total), nil
}

func (g *GeburaRepo) ListStoreAppBinaries(
	ctx context.Context,
	page *model.Paging,
	appIDs []model.InternalID,
) ([]*modelgebura.StoreAppBinary, int, error) {
	// Join sentinel_app_binaries with sentinel_libraries to get active snapshot
	query := g.data.db.WithContext(ctx).Model(&gormschema.SentinelAppBinary{}).
		Joins("JOIN sentinel_libraries ON sentinel_app_binaries.sentinel_library_reported_id = sentinel_libraries.reported_id AND sentinel_app_binaries.library_snapshot = sentinel_libraries.active_snapshot")

	if len(appIDs) > 0 {
		query = query.Where("sentinel_app_binaries.store_app_id IN ?", appIDs)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var binaries []gormschema.SentinelAppBinary
	if err := query.Limit(page.ToLimit()).Offset(page.ToOffset()).Find(&binaries).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelgebura.StoreAppBinary, len(binaries))
	for i := range binaries {
		res[i] = gormschema.ToBizStoreAppBinary(&binaries[i])
	}
	return res, int(total), nil
}
