package data

import (
	"context"
	"errors"
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackageruntime"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/model"

	"entgo.io/ent/dialect/sql"
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

func (g geburaRepo) CreateApp(ctx context.Context, a *modelgebura.App) error {
	if a.Details == nil {
		a.Details = new(modelgebura.AppDetails)
	}
	q := g.data.db.App.Create().
		SetID(a.ID).
		SetSource(converter.ToEntAppSource(a.Source)).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetIconImageURL(a.IconImageURL).
		SetHeroImageURL(a.HeroImageURL).
		SetDescription(a.Details.Description).
		SetReleaseDate(a.Details.ReleaseDate).
		SetDeveloper(a.Details.Developer).
		SetPublisher(a.Details.Publisher).
		SetVersion(a.Details.Version).
		SetBindInternalID(a.BoundInternal)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateApp(ctx context.Context, a *modelgebura.App) error {
	q := g.data.db.App.Update().
		Where(
			app.IDEQ(a.ID),
			app.SourceEQ(converter.ToEntAppSource(a.Source)),
			app.SourceAppIDEQ(a.SourceAppID),
		).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetIconImageURL(a.IconImageURL).
		SetHeroImageURL(a.HeroImageURL)
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

func (g geburaRepo) ListApps(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppSource,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool) ([]*modelgebura.App, int64, error) {
	var al []*ent.App
	var total int
	err := g.data.WithTx(ctx, func(tx *ent.Tx) error {
		q := tx.App.Query()
		if len(sources) > 0 {
			sourceFilter := make([]app.Source, len(sources))
			for i, appSource := range sources {
				sourceFilter[i] = converter.ToEntAppSource(appSource)
			}
			q.Where(app.SourceIn(sourceFilter...))
		}
		if len(types) > 0 {
			typeFilter := make([]app.Type, len(types))
			for i, appType := range types {
				typeFilter[i] = converter.ToEntAppType(appType)
			}
			q.Where(app.TypeIn(typeFilter...))
		}
		if len(ids) > 0 {
			q.Where(app.IDIn(ids...))
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
	apps := make([]*modelgebura.App, len(al))
	for i, sa := range al {
		apps[i] = converter.ToBizApp(sa)
		if !containDetails {
			apps[i].Details = nil
		}
	}
	return apps, int64(total), nil
}

func (g geburaRepo) MergeApps(ctx context.Context, base modelgebura.App, merged model.InternalID) error {
	err := g.data.WithTx(ctx, func(tx *ent.Tx) error {
		baseApp := converter.ToEntApp(base)
		err := tx.App.UpdateOne(&baseApp).Exec(ctx)
		if err != nil {
			return err
		}
		mergedApp, err := tx.App.Get(ctx, merged)
		if err != nil {
			return err
		}
		if baseApp.Source != app.SourceInternal || mergedApp.Source != app.SourceInternal {
			return errors.New("source must be internal")
		}
		err = tx.User.Update().
			Where(user.HasPurchasedAppWith(app.IDEQ(mergedApp.ID))).
			RemovePurchasedAppIDs(mergedApp.ID).
			AddPurchasedAppIDs(baseApp.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.App.Update().
			Where(app.HasBindInternalWith(app.IDEQ(mergedApp.ID))).
			SetBindInternalID(baseApp.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.AppPackage.Update().
			Where(apppackage.HasAppWith(app.IDEQ(mergedApp.ID))).
			SetAppID(baseApp.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
		err = tx.App.DeleteOne(mergedApp).Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (g geburaRepo) GetBoundApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, error) {
	a, err := g.data.db.App.Get(ctx, id)
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
	return converter.ToBizAppList(append(externalApps, internalApp)), nil
}

func (g geburaRepo) GetBatchBoundApps(ctx context.Context, ids []model.InternalID) ([]*modelgebura.BoundApps, error) {
	apps, err := g.data.db.App.Query().
		Where(
			app.IDIn(ids...),
			app.SourceEQ(app.SourceInternal),
		).
		WithBindExternal().
		All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*modelgebura.BoundApps, 0, len(apps))
	for i := range apps {
		res = append(res, new(modelgebura.BoundApps))
		res[i].Internal = converter.ToBizApp(apps[i])
		if externals, e := apps[i].Edges.BindExternalOrErr(); e == nil {
			for _, external := range externals {
				if external.Source == app.SourceSteam {
					res[i].Steam = converter.ToBizApp(external)
				}
			}
		}
		if res[i].Internal == nil {
			res[i].Internal = new(modelgebura.App)
		}
		if res[i].Steam == nil {
			res[i].Steam = new(modelgebura.App)
		}
	}
	return res, nil
}

func (g geburaRepo) PurchaseApp(ctx context.Context, userID model.InternalID, appID model.InternalID) error {
	a, err := g.data.db.App.Get(ctx, appID)
	if err != nil {
		return err
	}
	if a.Source != app.SourceInternal {
		return errors.New("illegal app source")
	}
	err = g.data.db.User.UpdateOneID(userID).AddPurchasedAppIDs(appID).Exec(ctx)
	return err
}

func (g geburaRepo) GetPurchasedApps(ctx context.Context, id model.InternalID) ([]*modelgebura.BoundApps, error) {
	apps, err := g.data.db.App.Query().
		Where(
			app.HasPurchasedByUserWith(user.IDEQ(id)),
		).
		WithBindExternal().
		All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*modelgebura.BoundApps, 0, len(apps))
	for i := range apps {
		res = append(res, new(modelgebura.BoundApps))
		res[i].Internal = converter.ToBizApp(apps[i])
		if externals, e := apps[i].Edges.BindExternalOrErr(); e == nil {
			for _, external := range externals {
				if external.Source == app.SourceSteam {
					res[i].Steam = converter.ToBizApp(external)
				}
			}
		}
	}
	return res, nil
}

func (g geburaRepo) CreateAppPackage(ctx context.Context, userID model.InternalID, ap *modelgebura.AppPackage) error {
	q := g.data.db.AppPackage.Create().
		SetOwnerID(userID).
		SetID(ap.ID).
		SetSource(converter.ToEntAppPackageSource(ap.Source)).
		SetSourceID(ap.SourceID).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetPublic(ap.Public)
	if ap.Binary != nil {
		q.
			SetBinaryName(ap.Binary.Name).
			SetBinarySizeBytes(ap.Binary.SizeBytes).
			SetBinaryPublicURL(ap.Binary.PublicURL).
			SetBinarySha256(ap.Binary.Sha256)
	}
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateAppPackage(ctx context.Context, ownerID model.InternalID, ap *modelgebura.AppPackage) error {
	q := g.data.db.AppPackage.Update().
		Where(
			apppackage.IDEQ(ap.ID),
			apppackage.SourceEQ(converter.ToEntAppPackageSource(ap.Source)),
			apppackage.HasOwnerWith(user.IDEQ(ownerID)),
		).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetPublic(ap.Public)
	if ap.Binary != nil {
		q.
			SetBinaryName(ap.Binary.Name).
			SetBinarySizeBytes(ap.Binary.SizeBytes).
			SetBinaryPublicURL(ap.Binary.PublicURL).
			SetBinarySha256(ap.Binary.Sha256)
	}
	return q.Exec(ctx)
}

func (g geburaRepo) UpsertAppPackages(
	ctx context.Context,
	userID model.InternalID,
	apl []*modelgebura.AppPackage,
) error {
	appPackages := make([]*ent.AppPackageCreate, len(apl))
	for i, ap := range apl {
		appPackages[i] = g.data.db.AppPackage.Create().
			SetID(ap.ID).
			SetOwnerID(userID).
			SetSource(converter.ToEntAppPackageSource(ap.Source)).
			SetSourceID(ap.SourceID).
			SetName(ap.Name).
			SetDescription(ap.Description).
			SetPublic(ap.Public)
		if ap.Binary != nil {
			appPackages[i].
				SetBinaryName(ap.Binary.Name).
				SetBinarySizeBytes(ap.Binary.SizeBytes).
				SetBinaryPublicURL(ap.Binary.PublicURL).
				SetBinarySha256(ap.Binary.Sha256)
		}
	}
	return g.data.db.AppPackage.
		CreateBulk(appPackages...).
		OnConflict(
			sql.ConflictColumns(apppackage.FieldBinarySha256),
			resolveWithIgnores([]string{
				apppackage.FieldID,
				apppackage.FieldPublic,
			}),
		).
		Exec(ctx)
}

func (g geburaRepo) ListAppPackages(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppPackageSource,
	ids []model.InternalID,
) ([]*modelgebura.AppPackage, int, error) {
	q := g.data.db.AppPackage.Query()
	if len(sources) > 0 {
		q.Where(apppackage.SourceIn(converter.ToEntAppPackageSourceList(sources)...))
	}
	if len(ids) > 0 {
		q.Where(apppackage.IDIn(ids...))
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
	return converter.ToBizAppPackageList(ap), total, nil
}

func (g geburaRepo) AssignAppPackage(
	ctx context.Context,
	userID model.InternalID,
	appID model.InternalID,
	appPackageID model.InternalID,
) error {
	err := g.data.db.AppPackage.Update().
		Where(
			apppackage.HasOwnerWith(user.IDEQ(userID)),
			apppackage.IDEQ(appPackageID),
		).
		SetAppID(appID).
		Exec(ctx)
	return err
}

func (g geburaRepo) UnAssignAppPackage(
	ctx context.Context,
	userID model.InternalID,
	appPackageID model.InternalID,
) error {
	err := g.data.db.AppPackage.Update().
		Where(
			apppackage.HasOwnerWith(user.IDEQ(userID)),
			apppackage.IDEQ(appPackageID),
		).
		ClearApp().
		Exec(ctx)
	return err
}

func (g geburaRepo) ListAppPackageBinaryChecksumOfOneSource(
	ctx context.Context,
	source modelgebura.AppPackageSource,
	sourceID model.InternalID,
) ([]string, error) {
	return g.data.db.AppPackage.Query().
		Where(
			apppackage.SourceEQ(converter.ToEntAppPackageSource(source)),
			apppackage.SourceIDEQ(sourceID),
		).
		Unique(true).
		Select(apppackage.FieldBinarySha256).
		Strings(ctx)
}

func (g geburaRepo) AddAppPackageRunTime(
	ctx context.Context,
	userID model.InternalID,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) error {
	return g.data.db.AppPackageRunTime.Create().
		SetUserID(userID).
		SetAppPackageID(packageID).
		SetStartTime(timeRange.StartTime).
		SetRunDuration(timeRange.Duration).Exec(ctx)
}

func (g geburaRepo) SumAppPackageRunTime(
	ctx context.Context,
	userID model.InternalID,
	packageID model.InternalID,
	timeRange *model.TimeRange,
) (time.Duration, error) {
	res, err := g.data.db.AppPackageRunTime.Query().Where(
		apppackageruntime.UserIDEQ(userID),
		apppackageruntime.AppPackageIDEQ(packageID),
		apppackageruntime.StartTimeGTE(timeRange.StartTime),
		apppackageruntime.StartTimeLTE(timeRange.StartTime.Add(timeRange.Duration)),
	).Aggregate(
		ent.Sum(apppackageruntime.FieldRunDuration),
	).Only(ctx)
	return res.RunDuration, err
}
