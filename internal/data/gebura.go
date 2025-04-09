package data

import (
	"context"
	"time"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/internal/ent"
	"github.com/tuihub/librarian/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/internal/data/internal/ent/appappcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/appcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/appruntime"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinellibrary"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelgebura"

	"entgo.io/ent/dialect/sql"
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
	q := g.data.db.AppInfo.Create().
		SetID(a.ID).
		SetSource(a.Source).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppInfoType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetDescription(a.Description).
		SetIconImageURL(a.IconImageURL).
		SetIconImageID(a.IconImageID).
		SetBackgroundImageURL(a.BackgroundImageURL).
		SetBackgroundImageID(a.BackgroundImageID).
		SetCoverImageURL(a.CoverImageURL).
		SetCoverImageID(a.CoverImageID).
		SetReleaseDate(a.ReleaseDate).
		SetDeveloper(a.Developer).
		SetPublisher(a.Publisher).
		SetTags(a.Tags).
		SetAlternativeNames(a.AlternativeNames).
		SetRawData(a.RawData)
	return q.Exec(ctx)
}

func (g *GeburaRepo) CreateAppInfoOrGet(ctx context.Context, a *modelgebura.AppInfo) (*modelgebura.AppInfo, error) {
	err := g.CreateAppInfo(ctx, a)
	if err == nil {
		return a, nil
	}
	if ent.IsConstraintError(err) {
		var ai *ent.AppInfo
		ai, err = g.data.db.AppInfo.Query().Where(
			appinfo.SourceEQ(a.Source),
			appinfo.SourceAppIDEQ(a.SourceAppID),
		).Only(ctx)
		if err == nil {
			return converter.ToBizAppInfo(ai), nil
		}
	}
	return nil, err
}

func (g *GeburaRepo) UpdateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	q := g.data.db.AppInfo.Update().
		Where(
			appinfo.IDEQ(a.ID),
			appinfo.SourceEQ(a.Source),
			appinfo.SourceAppIDEQ(a.SourceAppID),
		).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppInfoType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetDescription(a.Description).
		SetIconImageURL(a.IconImageURL).
		SetIconImageID(a.IconImageID).
		SetBackgroundImageURL(a.BackgroundImageURL).
		SetBackgroundImageID(a.BackgroundImageID).
		SetCoverImageURL(a.CoverImageURL).
		SetCoverImageID(a.CoverImageID).
		SetReleaseDate(a.ReleaseDate).
		SetDeveloper(a.Developer).
		SetPublisher(a.Publisher).
		SetTags(a.Tags).
		SetAlternativeNames(a.AlternativeNames).
		SetRawData(a.RawData)
	return q.Exec(ctx)
}

func (g *GeburaRepo) ListAppInfos(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
) ([]*modelgebura.AppInfo, int64, error) {
	var al []*ent.AppInfo
	var total int
	err := g.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.AppInfo.Query()
		if len(sources) > 0 {
			q.Where(appinfo.SourceIn(sources...))
		}
		if len(types) > 0 {
			typeFilter := make([]appinfo.Type, len(types))
			for i, appType := range types {
				typeFilter[i] = converter.ToEntAppInfoType(appType)
			}
			q.Where(appinfo.TypeIn(typeFilter...))
		}
		if len(ids) > 0 {
			q.Where(appinfo.IDIn(ids...))
		}
		var err error
		total, err = q.Count(ctx)
		if err != nil {
			return err
		}
		al, err = q.
			Limit(paging.ToLimit()).
			Offset(paging.ToOffset()).
			All(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizAppInfoList(al), int64(total), nil
}

func (g *GeburaRepo) GetAppInfo(ctx context.Context, id modelgebura.AppInfoID) (*modelgebura.AppInfo, error) {
	res, err := g.data.db.AppInfo.Query().
		Where(
			appinfo.SourceEQ(id.Source),
			appinfo.SourceAppIDEQ(id.SourceAppID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizAppInfo(res), nil
}

// func (g *GeburaRepo) GetBoundAppInfos(ctx context.Context, id model.InternalID) ([]*modelgebura.AppInfo, error) {
//	a, err := g.data.db.AppInfo.Get(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//	internalApp, err := a.QueryBindInternal().Only(ctx)
//	if err != nil {
//		return nil, err
//	}
//	externalApps, err := internalApp.QueryBindExternal().All(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return converter.ToBizAppInfoList(append(externalApps, internalApp)), nil
//}
//
// func (g *GeburaRepo) GetBatchBoundAppInfos(
//	ctx context.Context,
//	ids []model.InternalID,
// ) ([]*modelgebura.BoundAppInfos, error) {
//	infos, err := g.data.db.AppInfo.Query().
//		Where(
//			appinfo.IDIn(ids...),
//			appinfo.InternalEQ(true),
//		).
//		WithBindExternal().
//		All(ctx)
//	if err != nil {
//		return nil, err
//	}
//	res := make([]*modelgebura.BoundAppInfos, 0, len(infos))
//	for i := range infos {
//		res = append(res, new(modelgebura.BoundAppInfos))
//		res[i].Internal = converter.ToBizAppInfo(infos[i])
//		if externals, e := infos[i].Edges.BindExternalOrErr(); e == nil {
//			for _, external := range externals {
//				res[i].Others = append(res[i].Others, converter.ToBizAppInfo(external))
//			}
//		}
//		if res[i].Internal == nil {
//			res[i].Internal = new(modelgebura.AppInfo)
//		}
//	}
//	return res, nil
//}
//
// func (g *GeburaRepo) PurchaseAppInfo(
//	ctx context.Context,
//	userID model.InternalID,
//	appID *modelgebura.AppInfoID,
//	createFunc func(ctx2 context.Context) error,
// ) (model.InternalID, error) {
//	q := g.data.db.AppInfo.Query().WithBindInternal()
//	if appID.Internal {
//		q.Where(
//			appinfo.InternalEQ(true),
//			appinfo.SourceAppIDEQ(appID.SourceAppID),
//		)
//	} else {
//		q.Where(
//			appinfo.InternalEQ(false),
//			appinfo.SourceEQ(appID.Source),
//			appinfo.SourceAppIDEQ(appID.SourceAppID),
//		)
//	}
//	a, err := q.Only(ctx)
//	if ent.IsNotFound(err) && createFunc != nil {
//		err = createFunc(ctx)
//		if err != nil {
//			return 0, err
//		}
//		a, err = q.Only(ctx)
//		if err != nil {
//			return 0, err
//		}
//	}
//	if err != nil {
//		return 0, err
//	}
//	if a.Edges.BindInternal != nil {
//		err = g.data.db.User.UpdateOneID(userID).AddPurchasedAppIDs(a.Edges.BindInternal.ID).Exec(ctx)
//	} else {
//		err = errors.New("internal app not found")
//	}
//	return a.Edges.BindInternal.ID, err
//}
//
// func (g *GeburaRepo) GetPurchasedAppInfos(
//	ctx context.Context,
//	id model.InternalID,
//	source string,
// ) ([]*modelgebura.BoundAppInfos, error) {
//	q := g.data.db.AppInfo.Query().
//		Where(
//			appinfo.HasPurchasedByUserWith(user.IDEQ(id)),
//		)
//	if len(source) > 0 {
//		q.Where(appinfo.SourceEQ(source))
//	}
//	infos, err := q.
//		WithBindExternal().
//		All(ctx)
//	if err != nil {
//		return nil, err
//	}
//	res := make([]*modelgebura.BoundAppInfos, 0, len(infos))
//	for i := range infos {
//		res = append(res, new(modelgebura.BoundAppInfos))
//		res[i].Internal = converter.ToBizAppInfo(infos[i])
//		if externals, e := infos[i].Edges.BindExternalOrErr(); e == nil {
//			for _, external := range externals {
//				res[i].Others = append(res[i].Others, converter.ToBizAppInfo(external))
//			}
//		}
//	}
//	return res, nil
//}

func (g *GeburaRepo) CreateApp(ctx context.Context, userID model.InternalID, a *modelgebura.App) error {
	q := g.data.db.App.Create().
		SetUserID(userID).
		SetID(a.ID).
		SetVersionNumber(a.VersionNumber).
		SetVersionDate(a.VersionDate).
		SetCreatorDeviceID(a.CreatorDeviceID).
		SetAppSources(a.AppSources).
		SetPublic(a.Public).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetDescription(a.Description).
		SetIconImageURL(a.IconImageURL).
		SetIconImageID(a.IconImageID).
		SetBackgroundImageURL(a.BackgroundImageURL).
		SetBackgroundImageID(a.BackgroundImageID).
		SetCoverImageURL(a.CoverImageURL).
		SetCoverImageID(a.CoverImageID).
		SetReleaseDate(a.ReleaseDate).
		SetDeveloper(a.Developer).
		SetPublisher(a.Publisher).
		SetTags(a.Tags).
		SetAlternativeNames(a.AlternativeNames)
	if a.BoundStoreAppID != nil {
		q.SetBoundStoreAppID(*a.BoundStoreAppID)
	}
	if a.StopStoreManage != nil {
		q.SetStopStoreManage(*a.StopStoreManage)
	}
	return q.Exec(ctx)
}

func (g *GeburaRepo) UpdateApp(ctx context.Context, ownerID model.InternalID, a *modelgebura.App) error {
	return g.data.WithTx(ctx, func(tx *ent.Tx) error {
		old, err := tx.App.Get(ctx, a.ID)
		if err != nil {
			return err
		}
		q := tx.App.Update().
			Where(
				app.IDEQ(a.ID),
				app.HasUserWith(user.IDEQ(ownerID)),
			).
			SetVersionNumber(old.VersionNumber + 1).
			SetVersionDate(a.VersionDate).
			SetAppSources(a.AppSources).
			SetPublic(a.Public).
			SetName(a.Name).
			SetType(converter.ToEntAppType(a.Type)).
			SetShortDescription(a.ShortDescription).
			SetDescription(a.Description).
			SetIconImageURL(a.IconImageURL).
			SetIconImageID(a.IconImageID).
			SetBackgroundImageURL(a.BackgroundImageURL).
			SetBackgroundImageID(a.BackgroundImageID).
			SetCoverImageURL(a.CoverImageURL).
			SetCoverImageID(a.CoverImageID).
			SetReleaseDate(a.ReleaseDate).
			SetDeveloper(a.Developer).
			SetPublisher(a.Publisher).
			SetTags(a.Tags).
			SetAlternativeNames(a.AlternativeNames)
		if a.StopStoreManage != nil {
			q.SetStopStoreManage(*a.StopStoreManage)
		}
		return q.Exec(ctx)
	})
}

func (g *GeburaRepo) GetApp(ctx context.Context, id model.InternalID) (*modelgebura.App, error) {
	a, err := g.data.db.App.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return converter.ToBizApp(a), nil
}

func (g *GeburaRepo) ListApps(
	ctx context.Context,
	paging model.Paging,
	ownerIDs []model.InternalID,
	ids []model.InternalID,
	publicOnly bool,
) ([]*modelgebura.App, int, error) {
	q := g.data.db.App.Query().Where(
		app.HasUserWith(user.IDIn(ownerIDs...)),
	)
	if len(ids) > 0 {
		q.Where(app.IDIn(ids...))
	}
	if publicOnly {
		q.Where(app.PublicEQ(true))
	}
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ap, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*modelgebura.App, len(ap))
	for i := range ap {
		res[i] = converter.ToBizApp(ap[i])
	}
	return res, total, nil
}

// func (g *GeburaRepo) AssignApp(
//	ctx context.Context,
//	userID model.InternalID,
//	appID model.InternalID,
//	appInfoID model.InternalID,
// ) error {
//	err := g.data.db.App.Update().
//		Where(
//			app.HasOwnerWith(user.IDEQ(userID)),
//			app.IDEQ(appID),
//		).
//		SetAppInfoID(appInfoID).
//		Exec(ctx)
//	return err
//}
//
// func (g *GeburaRepo) UnAssignApp(
//	ctx context.Context,
//	userID model.InternalID,
//	appID model.InternalID,
// ) error {
//	err := g.data.db.App.Update().
//		Where(
//			app.HasOwnerWith(user.IDEQ(userID)),
//			app.IDEQ(appID),
//		).
//		ClearAppInfo().
//		Exec(ctx)
//	return err
//}

// func (g geburaRepo) ListAppPackageBinaryChecksumOfOneSource(
//	ctx context.Context,
//	source modelgebura.AppPackageSource,
//	sourceID model.InternalID,
// ) ([]string, error) {
//	return g.data.db.App.Query().
//		Where(
//			app.SourceEQ(converter.ToEntAppPackageSource(source)),
//			app.SourceIDEQ(sourceID),
//		).
//		Unique(true).
//		Select(app.FieldBinarySha256).
//		Strings(ctx)
//}

func (g *GeburaRepo) BatchCreateAppRunTime(
	ctx context.Context,
	userID model.InternalID,
	runTimes []*modelgebura.AppRunTime,
) error {
	rt := make([]*ent.AppRunTimeCreate, 0, len(runTimes))
	for _, runTime := range runTimes {
		rt = append(rt, g.data.db.AppRunTime.Create().
			SetID(runTime.ID).
			SetUserID(userID).
			SetAppID(runTime.AppID).
			SetStartTime(runTime.RunTime.StartTime).
			SetDuration(runTime.RunTime.Duration),
		)
	}
	return g.data.db.AppRunTime.CreateBulk(rt...).Exec(ctx)
}

func (g *GeburaRepo) SumAppRunTime(
	ctx context.Context,
	userID model.InternalID,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	var v []struct {
		Sum time.Duration
	}
	q := g.data.db.AppRunTime.Query().Where(
		appruntime.UserIDEQ(userID))
	if len(appIDs) > 0 {
		q.Where(appruntime.AppIDIn(appIDs...))
	}
	if len(deviceIDs) > 0 {
		q.Where(appruntime.DeviceIDIn(deviceIDs...))
	}
	err := q.Where(appruntime.And(
		appruntime.StartTimeGTE(timeRange.StartTime),
		appruntime.StartTimeLTE(timeRange.StartTime.Add(timeRange.Duration)),
	)).
		Aggregate(
			ent.Sum(appruntime.FieldDuration),
		).Scan(ctx, &v)
	if err != nil {
		return time.Duration(0), err
	}
	var res time.Duration
	for _, rt := range v {
		res += rt.Sum
	}
	return res, nil
}

func (g *GeburaRepo) ListAppRunTimes(
	ctx context.Context,
	userID model.InternalID,
	paging model.Paging,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
	timeRange *model.TimeRange,
) ([]*modelgebura.AppRunTime, int, error) {
	q := g.data.db.AppRunTime.Query().Where(
		appruntime.UserIDEQ(userID),
	)
	if len(appIDs) > 0 {
		q.Where(appruntime.AppIDIn(appIDs...))
	}
	if len(deviceIDs) > 0 {
		q.Where(appruntime.DeviceIDIn(deviceIDs...))
	}
	if timeRange != nil {
		q.Where(appruntime.And(
			appruntime.StartTimeGTE(timeRange.StartTime),
			appruntime.StartTimeLTE(timeRange.StartTime.Add(timeRange.Duration)),
		))
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
	return converter.ToBizAppRunTimeList(res), total, nil
}

func (g *GeburaRepo) DeleteAppRunTime(ctx context.Context, userID model.InternalID, id model.InternalID) error {
	return g.data.db.AppRunTime.DeleteOneID(id).Exec(ctx)
}

func (g *GeburaRepo) CreateAppCategory(ctx context.Context, userID model.InternalID, ac *modelgebura.AppCategory) error {
	q := g.data.db.AppCategory.Create().
		SetID(ac.ID).
		SetUserID(userID).
		SetVersionNumber(ac.VersionNumber).
		SetVersionDate(ac.VersionDate).
		SetName(ac.Name).
		AddAppIDs(ac.AppIDs...)
	return q.Exec(ctx)
}

func (g *GeburaRepo) ListAppCategories(ctx context.Context, userID model.InternalID) ([]*modelgebura.AppCategory, error) {
	acs, err := g.data.db.AppCategory.Query().
		WithAppAppCategory().
		Where(appcategory.UserIDEQ(userID)).
		All(ctx)
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
	userID model.InternalID,
	ac *modelgebura.AppCategory) error {
	return g.data.WithTx(ctx, func(tx *ent.Tx) error {
		// get old
		old, err := tx.AppCategory.Query().
			Where(
				appcategory.IDEQ(ac.ID),
				appcategory.UserIDEQ(userID),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		// remove existing
		_, err = tx.AppAppCategory.Delete().Where(
			appappcategory.HasAppCategoryWith(
				appcategory.IDEQ(ac.ID),
				appcategory.UserIDEQ(userID),
			),
		).Exec(ctx)
		if err != nil {
			return err
		}
		q := tx.AppCategory.Update().
			Where(
				appcategory.IDEQ(ac.ID),
				appcategory.UserIDEQ(userID),
			).
			SetName(ac.Name).
			SetVersionNumber(old.VersionNumber + 1).
			SetVersionDate(time.Now()).
			AddAppIDs(ac.AppIDs...)
		return q.Exec(ctx)
	})
}

func (g *GeburaRepo) DeleteAppCategory(
	ctx context.Context,
	userID model.InternalID,
	id model.InternalID,
) error {
	return g.data.WithTx(ctx, func(tx *ent.Tx) error {
		_, err := tx.AppAppCategory.Delete().Where(
			appappcategory.HasAppCategoryWith(
				appcategory.IDEQ(id),
				appcategory.UserIDEQ(userID),
			),
		).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = tx.AppCategory.Delete().Where(
			appcategory.IDEQ(id),
			appcategory.UserIDEQ(userID),
		).Exec(ctx)
		return err
	})
}

func (g *GeburaRepo) UpsertSentinelInfo(
	ctx context.Context,
	info *modelgebura.SentinelInfo,
) error {
	return g.data.WithTx(ctx, func(tx *ent.Tx) error {
		// upsert sentinel info
		q := tx.SentinelInfo.Create().
			SetID(info.ID).
			SetURL(info.Url).
			SetAlternativeUrls(info.AlternativeUrls).
			SetGetTokenPath(info.GetTokenPath).
			SetDownloadFileBasePath(info.DownloadFileBasePath)
		err := q.OnConflict(sql.ConflictColumns(sentinelinfo.FieldID)).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
		// upsert libraries
		sInfo, err := tx.SentinelInfo.Query().
			Where(sentinelinfo.IDEQ(info.ID)).
			Only(ctx)
		if err != nil {
			return err
		}
		oldLibs, err := tx.SentinelLibrary.Query().
			Where(sentinellibrary.SentinelInfoIDEQ(sInfo.ID)).
			All(ctx)
		if err != nil {
			return err
		}
		oldLibsMap := make(map[int64]*ent.SentinelLibrary, len(oldLibs))
		for _, lib := range oldLibs {
			oldLibsMap[lib.ReportedID] = lib
		}
		newLibs := make([]*ent.SentinelLibraryCreate, 0, len(info.Libraries))
		for _, lib := range info.Libraries {
			var reportSeq int64
			if oldLib, exists := oldLibsMap[lib.ReportedID]; exists {
				reportSeq = oldLib.ReportSequence + 1
			} else {
				reportSeq = 0
			}
			newLibs = append(newLibs, tx.SentinelLibrary.Create().
				SetID(lib.ID).
				SetSentinelInfoID(sInfo.ID).
				SetReportedID(lib.ReportedID).
				SetDownloadBasePath(lib.DownloadBasePath).
				SetReportSequence(reportSeq))
		}
		return tx.SentinelLibrary.CreateBulk(newLibs...).
			OnConflict(
				sql.ConflictColumns(
					sentinellibrary.FieldSentinelInfoID,
					sentinellibrary.FieldReportedID,
				),
				resolveWithIgnores([]string{
					sentinellibrary.FieldID,
				}),
			).
			UpdateNewValues().
			Exec(ctx)
	})
}
