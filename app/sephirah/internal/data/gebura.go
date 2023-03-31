package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model"

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
		SetImageURL(a.ImageURL).
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
		).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(converter.ToEntAppType(a.Type)).
		SetShortDescription(a.ShortDescription).
		SetImageURL(a.ImageURL)
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

func (g geburaRepo) UpsertApps(ctx context.Context, al []*modelgebura.App) error {
	apps := make([]*ent.AppCreate, len(al))
	for i, a := range al {
		if a.Details == nil {
			a.Details = new(modelgebura.AppDetails)
		}
		apps[i] = g.data.db.App.Create().
			SetID(a.ID).
			SetSource(converter.ToEntAppSource(a.Source)).
			SetSourceAppID(a.SourceAppID).
			SetSourceURL(a.SourceURL).
			SetName(a.Name).
			SetType(converter.ToEntAppType(a.Type)).
			SetShortDescription(a.ShortDescription).
			SetImageURL(a.ImageURL).
			SetBindInternalID(a.BoundInternal)
		if a.Details != nil {
			apps[i].
				SetDescription(a.Details.Description).
				SetReleaseDate(a.Details.ReleaseDate).
				SetDeveloper(a.Details.Developer).
				SetPublisher(a.Details.Publisher).
				SetVersion(a.Details.Version)
		}
	}
	return g.data.db.App.
		CreateBulk(apps...).
		OnConflict(
			sql.ConflictColumns(app.FieldSource, app.FieldSourceAppID),
			resolveWithIgnores([]string{
				app.FieldID,
			}),
		).
		Exec(ctx)
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
			Limit(paging.PageSize).
			Offset((paging.PageNum - 1) * paging.PageSize).
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

func (g geburaRepo) SearchApps(ctx context.Context, paging model.Paging, keyword string) (
	[]*modelgebura.App, int, error) {
	q := g.data.db.App.Query().
		Where(
			app.Or(
				app.NameContains(keyword),
				app.ShortDescriptionContains(keyword),
				app.DescriptionContains(keyword),
			),
		)
	total, err := q.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	apps, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizAppList(apps), total, nil
}

func (g geburaRepo) GetBindApps(ctx context.Context, id model.InternalID) ([]*modelgebura.App, error) {
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

func (g geburaRepo) PurchaseApp(ctx context.Context, userID model.InternalID, appID model.InternalID) error {
	err := g.data.db.User.UpdateOneID(userID).AddPurchasedAppIDs(appID).Exec(ctx)
	return err
}

func (g geburaRepo) GetPurchasedApps(ctx context.Context, id model.InternalID) ([]model.InternalID, error) {
	appIDs, err := g.data.db.App.Query().
		Where(
			app.HasPurchasedByWith(user.IDEQ(id)),
		).
		IDs(ctx)
	if err != nil {
		return nil, err
	}
	return appIDs, nil
}

func (g geburaRepo) CreateAppPackage(ctx context.Context, ap *modelgebura.AppPackage) error {
	q := g.data.db.AppPackage.Create().
		SetID(ap.ID).
		SetSource(converter.ToEntAppPackageSource(ap.Source)).
		SetSourceID(ap.SourceID).
		SetSourcePackageID(ap.SourcePackageID).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetBinaryName(ap.Binary.Name).
		SetBinarySizeByte(ap.Binary.SizeByte)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateAppPackage(ctx context.Context, ap *modelgebura.AppPackage) error {
	q := g.data.db.AppPackage.Update().
		Where(
			apppackage.IDEQ(ap.ID),
			apppackage.SourceEQ(converter.ToEntAppPackageSource(ap.Source)),
		).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetBinaryName(ap.Binary.Name).
		SetBinarySizeByte(ap.Binary.SizeByte)
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
			SetSourcePackageID(ap.SourcePackageID).
			SetName(ap.Name).
			SetDescription(ap.Description).
			SetPublic(ap.Public).
			SetBinaryName(ap.Binary.Name).
			SetBinarySizeByte(ap.Binary.SizeByte).
			SetBinaryPublicURL(ap.Binary.PublicURL)
	}
	return g.data.db.AppPackage.
		CreateBulk(appPackages...).
		OnConflict(
			sql.ConflictColumns(apppackage.FieldSource, apppackage.FieldSourceID, apppackage.FieldSourcePackageID),
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
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
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

func (g geburaRepo) ListAllAppPackageIDOfOneSource(
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
		Select(apppackage.FieldSourcePackageID).
		Strings(ctx)
}
