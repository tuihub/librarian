package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/data/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelgebura"
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/exp/slices"
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

func (g geburaRepo) IsApp(ctx context.Context, id model.InternalID) error {
	a, err := g.ListApp(ctx, model.Paging{
		PageSize: 1,
		PageNum:  0,
	}, nil, nil, []model.InternalID{id}, false)
	if err != nil {
		return err
	}
	if len(a) != 1 {
		return errors.New("no such app")
	}
	return nil
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
		SetVersion(a.Details.Version)
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

func (g geburaRepo) UpsertApp(ctx context.Context, al []*modelgebura.App) error {
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
			SetImageURL(a.ImageURL)
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
			sql.ResolveWithIgnore(),
			sql.ResolveWith(func(u *sql.UpdateSet) {
				ignores := []string{
					app.FieldID,
					app.FieldSource,
					app.FieldSourceAppID,
				}
				for _, c := range u.Columns() {
					if slices.Contains(ignores, c) {
						u.SetIgnore(c)
					}
					u.SetExcluded(c)
				}
			}),
		).
		Exec(ctx)
}

func (g geburaRepo) ListApp(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppSource,
	types []modelgebura.AppType,
	ids []model.InternalID,
	containDetails bool) ([]*modelgebura.App, error) {
	q := g.data.db.App.Query()
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
	a, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	apps := make([]*modelgebura.App, len(a))
	for i, sa := range a {
		apps[i] = g.data.converter.ToBizApp(sa)
		if !containDetails {
			apps[i].Details = nil
		}
	}
	return apps, nil
}

func (g geburaRepo) IsAppPackage(ctx context.Context, id model.InternalID) error {
	a, err := g.ListAppPackage(ctx, model.Paging{
		PageSize: 1,
		PageNum:  0,
	}, nil, []model.InternalID{id})
	if err != nil {
		return err
	}
	if len(a) != 1 {
		return errors.New("no such app package")
	}
	return nil
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
		SetBinarySize(ap.Binary.Size)
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
		SetBinarySize(ap.Binary.Size)
	return q.Exec(ctx)
}

func (g geburaRepo) UpsertAppPackage(ctx context.Context, apl []*modelgebura.AppPackage) error {
	appPackages := make([]*ent.AppPackageCreate, len(apl))
	for i, ap := range apl {
		appPackages[i] = g.data.db.AppPackage.Create().
			SetID(ap.ID).
			SetSource(converter.ToEntAppPackageSource(ap.Source)).
			SetSourceID(ap.SourceID).
			SetSourcePackageID(ap.SourcePackageID)
		if len(ap.Name) > 0 {
			appPackages[i].SetName(ap.Name)
		}
		if len(ap.Description) > 0 {
			appPackages[i].SetDescription(ap.Description)
		}
		if len(ap.Binary.Name) > 0 {
			appPackages[i].SetBinaryName(ap.Binary.Name)
		}
		if ap.Binary.Size > 0 {
			appPackages[i].SetBinarySize(ap.Binary.Size)
		}
	}
	return g.data.db.AppPackage.
		CreateBulk(appPackages...).
		OnConflict(
			sql.ConflictColumns(apppackage.FieldSource, apppackage.FieldSourceID, apppackage.FieldSourcePackageID),
			sql.ResolveWith(func(u *sql.UpdateSet) {
				ignores := []string{
					apppackage.FieldID,
					apppackage.FieldSource,
					apppackage.FieldSourceID,
					apppackage.FieldSourcePackageID,
				}
				for _, c := range u.Columns() {
					if slices.Contains(ignores, c) {
						u.SetIgnore(c)
					}
					u.SetExcluded(c)
				}
			}),
		).
		Exec(ctx)
}

func (g geburaRepo) ListAppPackage(
	ctx context.Context,
	paging model.Paging,
	sources []modelgebura.AppPackageSource,
	ids []model.InternalID,
) ([]*modelgebura.AppPackage, error) {
	q := g.data.db.AppPackage.Query()
	if len(sources) > 0 {
		sourceFilter := make([]apppackage.Source, len(sources))
		for i, apSource := range sources {
			sourceFilter[i] = converter.ToEntAppPackageSource(apSource)
		}
		q.Where(apppackage.SourceIn(sourceFilter...))
	}
	if len(ids) > 0 {
		q.Where(apppackage.IDIn(ids...))
	}
	ap, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return g.data.converter.ToBizAppPackageList(ap), nil
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
