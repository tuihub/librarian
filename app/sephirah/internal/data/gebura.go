package data

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinst"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinstruntime"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model"
)

type geburaRepo struct {
	data *Data
}

// NewGeburaRepo .
func NewGeburaRepo(data *Data) bizgebura.GeburaRepo {
	return &geburaRepo{
		data: data,
	}
}

func (g geburaRepo) CreateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	if a.Details == nil {
		a.Details = new(modelgebura.AppInfoDetails)
	}
	q := g.data.db.AppInfo.Create().
		SetID(a.ID).
		SetInternal(a.Internal).
		SetSource(a.Source).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetIconImageURL(a.IconImageURL).
		SetBackgroundImageURL(a.BackgroundImageURL).
		SetCoverImageURL(a.CoverImageURL).
		SetDescription(a.Details.Description).
		SetReleaseDate(a.Details.ReleaseDate).
		SetDeveloper(a.Details.Developer).
		SetPublisher(a.Details.Publisher).
		SetVersion(a.Details.Version).
		SetBindInternalID(a.BoundInternal)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateAppInfo(ctx context.Context, a *modelgebura.AppInfo) error {
	q := g.data.db.AppInfo.Update().
		Where(
			appinfo.IDEQ(a.ID),
			appinfo.SourceEQ(a.Source),
			appinfo.SourceAppIDEQ(a.SourceAppID),
		).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetIconImageURL(a.IconImageURL).
		SetBackgroundImageURL(a.BackgroundImageURL).
		SetCoverImageURL(a.CoverImageURL)
	if a.Details != nil {
		q.
			SetDescription(a.Details.Description).
			SetReleaseDate(a.Details.ReleaseDate).
			SetDeveloper(a.Details.Developer).
			SetPublisher(a.Details.Publisher).
			SetVersion(a.Details.Version)
	}
	return q.Exec(ctx)
}

func (g geburaRepo) ListAppInfos(
	ctx context.Context,
	paging model.Paging,
	sources []string,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool,
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
				typeFilter[i] = converter.ToEntAppType(appType)
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
	infos := make([]*modelgebura.AppInfo, len(al))
	for i, sa := range al {
		infos[i] = converter.ToBizAppInfo(sa)
		if !containDetails {
			infos[i].Details = nil
		}
	}
	return infos, int64(total), nil
}

func (g geburaRepo) MergeAppInfos(ctx context.Context, base modelgebura.AppInfo, merged model.InternalID) error {
	err := g.data.WithTx(ctx, func(tx *ent.Tx) error {
		baseAppInfo := converter.ToEntAppInfo(base)
		err := tx.AppInfo.UpdateOne(&baseAppInfo).Exec(ctx)
		if err != nil {
			return err
		}
		mergedAppInfo, err := tx.AppInfo.Get(ctx, merged)
		if err != nil {
			return err
		}
		if !baseAppInfo.Internal || !mergedAppInfo.Internal {
			return errors.New("source must be internal")
		}
		err = tx.User.Update().
			Where(user.HasPurchasedAppWith(appinfo.IDEQ(mergedAppInfo.ID))).
			RemovePurchasedAppIDs(mergedAppInfo.ID).
			AddPurchasedAppIDs(baseAppInfo.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.AppInfo.Update().
			Where(appinfo.HasBindInternalWith(appinfo.IDEQ(mergedAppInfo.ID))).
			SetBindInternalID(baseAppInfo.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.App.Update().
			Where(app.HasAppInfoWith(appinfo.IDEQ(mergedAppInfo.ID))).
			SetAppInfoID(baseAppInfo.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.AppInfo.DeleteOne(mergedAppInfo).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (g geburaRepo) GetAppInfo(ctx context.Context, id modelgebura.AppInfoID) (*modelgebura.AppInfo, error) {
	res, err := g.data.db.AppInfo.Query().
		Where(
			appinfo.InternalEQ(id.Internal),
			appinfo.SourceEQ(id.Source),
			appinfo.SourceAppIDEQ(id.SourceAppID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizAppInfo(res), nil
}

func (g geburaRepo) GetBoundAppInfos(ctx context.Context, id model.InternalID) ([]*modelgebura.AppInfo, error) {
	a, err := g.data.db.AppInfo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	internalApp, err := a.QueryBindInternal().Only(ctx)
	if err != nil {
		return nil, err
	}
	externalApps, err := internalApp.QueryBindExternal().All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizAppInfoList(append(externalApps, internalApp)), nil
}

func (g geburaRepo) GetBatchBoundAppInfos(
	ctx context.Context,
	ids []model.InternalID,
) ([]*modelgebura.BoundAppInfos, error) {
	infos, err := g.data.db.AppInfo.Query().
		Where(
			appinfo.IDIn(ids...),
			appinfo.InternalEQ(true),
		).
		WithBindExternal().
		All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*modelgebura.BoundAppInfos, 0, len(infos))
	for i := range infos {
		res = append(res, new(modelgebura.BoundAppInfos))
		res[i].Internal = converter.ToBizAppInfo(infos[i])
		if externals, e := infos[i].Edges.BindExternalOrErr(); e == nil {
			for _, external := range externals {
				res[i].Others = append(res[i].Others, converter.ToBizAppInfo(external))
			}
		}
		if res[i].Internal == nil {
			res[i].Internal = new(modelgebura.AppInfo)
		}
	}
	return res, nil
}

func (g geburaRepo) PurchaseAppInfo(
	ctx context.Context,
	userID model.InternalID,
	appID *modelgebura.AppInfoID,
	createFunc func(ctx2 context.Context) error,
) (model.InternalID, error) {
	q := g.data.db.AppInfo.Query().WithBindInternal()
	if appID.Internal {
		q.Where(
			appinfo.InternalEQ(true),
			appinfo.SourceAppIDEQ(appID.SourceAppID),
		)
	} else {
		q.Where(
			appinfo.InternalEQ(false),
			appinfo.SourceEQ(appID.Source),
			appinfo.SourceAppIDEQ(appID.SourceAppID),
		)
	}
	a, err := q.Only(ctx)
	if ent.IsNotFound(err) && createFunc != nil {
		err = createFunc(ctx)
		if err != nil {
			return 0, err
		}
		a, err = q.Only(ctx)
		if err != nil {
			return 0, err
		}
	}
	if err != nil {
		return 0, err
	}
	if a.Edges.BindInternal != nil {
		err = g.data.db.User.UpdateOneID(userID).AddPurchasedAppIDs(a.Edges.BindInternal.ID).Exec(ctx)
	} else {
		err = errors.New("internal app not found")
	}
	return a.Edges.BindInternal.ID, err
}

func (g geburaRepo) GetPurchasedAppInfos(
	ctx context.Context,
	id model.InternalID,
	source string,
) ([]*modelgebura.BoundAppInfos, error) {
	q := g.data.db.AppInfo.Query().
		Where(
			appinfo.HasPurchasedByUserWith(user.IDEQ(id)),
		)
	if len(source) > 0 {
		q.Where(appinfo.SourceEQ(source))
	}
	infos, err := q.
		WithBindExternal().
		All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*modelgebura.BoundAppInfos, 0, len(infos))
	for i := range infos {
		res = append(res, new(modelgebura.BoundAppInfos))
		res[i].Internal = converter.ToBizAppInfo(infos[i])
		if externals, e := infos[i].Edges.BindExternalOrErr(); e == nil {
			for _, external := range externals {
				res[i].Others = append(res[i].Others, converter.ToBizAppInfo(external))
			}
		}
	}
	return res, nil
}

func (g geburaRepo) CreateApp(ctx context.Context, userID model.InternalID, ap *modelgebura.App) error {
	q := g.data.db.App.Create().
		SetOwnerID(userID).
		SetID(ap.ID).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetPublic(ap.Public)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateApp(ctx context.Context, ownerID model.InternalID, ap *modelgebura.App) error {
	q := g.data.db.App.Update().
		Where(
			app.IDEQ(ap.ID),
			app.HasOwnerWith(user.IDEQ(ownerID)),
		).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetPublic(ap.Public)
	return q.Exec(ctx)
}

func (g geburaRepo) ListApps(
	ctx context.Context,
	ownerID model.InternalID,
	paging model.Paging,
	appInfoIDs []model.InternalID,
	ids []model.InternalID,
) ([]*modelgebura.App, int, error) {
	q := g.data.db.App.Query().Where(app.HasOwnerWith(user.IDEQ(ownerID)))
	if len(ids) > 0 {
		q.Where(app.IDIn(ids...))
	}
	if len(appInfoIDs) > 0 {
		q.Where(app.HasAppInfoWith(appinfo.IDIn(appInfoIDs...)))
	}
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	ap, err := q.
		WithAppInfo().
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	res := make([]*modelgebura.App, len(ap))
	for i := range ap {
		res[i] = converter.ToBizApp(ap[i])
		if ap[i].Edges.AppInfo != nil {
			res[i].AssignedAppInfoID = ap[i].Edges.AppInfo.ID
		}
	}
	return res, total, nil
}

func (g geburaRepo) AssignApp(
	ctx context.Context,
	userID model.InternalID,
	appID model.InternalID,
	appInfoID model.InternalID,
) error {
	err := g.data.db.App.Update().
		Where(
			app.HasOwnerWith(user.IDEQ(userID)),
			app.IDEQ(appID),
		).
		SetAppInfoID(appInfoID).
		Exec(ctx)
	return err
}

func (g geburaRepo) UnAssignApp(
	ctx context.Context,
	userID model.InternalID,
	appID model.InternalID,
) error {
	err := g.data.db.App.Update().
		Where(
			app.HasOwnerWith(user.IDEQ(userID)),
			app.IDEQ(appID),
		).
		ClearAppInfo().
		Exec(ctx)
	return err
}

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

func (g geburaRepo) CreateAppInst(ctx context.Context, ownerID model.InternalID, inst *modelgebura.AppInst) error {
	_, err := g.data.db.App.Query().Where(
		app.IDEQ(inst.AppID),
		app.HasOwnerWith(user.IDEQ(ownerID)),
	).Only(ctx)
	if err != nil {
		return err
	}
	q := g.data.db.AppInst.Create().
		SetID(inst.ID).
		SetOwnerID(ownerID).
		SetAppID(inst.AppID).
		SetDeviceID(inst.DeviceID)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateAppInst(ctx context.Context, ownerID model.InternalID, inst *modelgebura.AppInst) error {
	_, err := g.data.db.App.Query().Where(
		app.IDEQ(inst.AppID),
		app.HasOwnerWith(user.IDEQ(ownerID)),
	).Only(ctx)
	if err != nil {
		return err
	}
	q := g.data.db.AppInst.Update().
		Where(
			appinst.IDEQ(inst.ID),
			appinst.HasOwnerWith(user.IDEQ(ownerID)),
		).
		SetAppID(inst.AppID)
	return q.Exec(ctx)
}

func (g geburaRepo) ListAppInsts(
	ctx context.Context,
	ownerID model.InternalID,
	paging model.Paging,
	ids []model.InternalID,
	appIDs []model.InternalID,
	deviceIDs []model.InternalID,
) ([]*modelgebura.AppInst, int, error) {
	q := g.data.db.AppInst.Query().Where(appinst.HasOwnerWith(user.IDEQ(ownerID)))
	if len(ids) > 0 {
		q.Where(appinst.IDIn(ids...))
	}
	if len(appIDs) > 0 {
		q.Where(appinst.AppIDIn(appIDs...))
	}
	if len(deviceIDs) > 0 {
		q.Where(appinst.DeviceIDIn(deviceIDs...))
	}
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	insts, err := q.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizAppInstList(insts), total, nil
}

func (g geburaRepo) AddAppInstRunTime(
	ctx context.Context,
	userID model.InternalID,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) error {
	return g.data.db.AppInstRunTime.Create().
		SetUserID(userID).
		SetAppID(packageID).
		SetStartTime(timeRange.StartTime).
		SetRunDuration(timeRange.Duration).Exec(ctx)
}

func (g geburaRepo) SumAppInstRunTime(
	ctx context.Context,
	userID model.InternalID,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	res, err := g.data.db.AppInstRunTime.Query().Where(
		appinstruntime.UserIDEQ(userID),
		appinstruntime.AppIDEQ(packageID),
		appinstruntime.StartTimeGTE(timeRange.StartTime),
		appinstruntime.StartTimeLTE(timeRange.StartTime.Add(timeRange.Duration)),
	).Aggregate(
		ent.Sum(appinstruntime.FieldRunDuration),
	).Only(ctx)
	return res.RunDuration, err
}
