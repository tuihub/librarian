package data

import (
	"context"
	"errors"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/ent"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/apppackage"

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

func (g geburaRepo) IsApp(ctx context.Context, id int64) error {
	a, err := g.ListApp(ctx, bizgebura.Paging{
		PageSize: 1,
		PageNum:  0,
	}, nil, nil, []int64{id}, false)
	if err != nil {
		return err
	}
	if len(a) != 1 {
		return errors.New("no such app")
	}
	return nil
}

func (g geburaRepo) CreateApp(ctx context.Context, a *bizgebura.App) error {
	if a.Details == nil {
		a.Details = new(bizgebura.AppDetails)
	}
	q := g.data.db.App.Create().
		SetInternalID(a.InternalID).
		SetSource(toEntAppSource(a.Source)).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(toEntAppType(a.Type)).
		SetShortDescription(a.ShorDescription).
		SetImageURL(a.ImageURL).
		SetDescription(a.Details.Description).
		SetReleaseDate(a.Details.ReleaseDate).
		SetDeveloper(a.Details.Developer).
		SetPublisher(a.Details.Publisher)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateApp(ctx context.Context, a *bizgebura.App) error {
	q := g.data.db.App.Update().
		Where(
			app.InternalIDEQ(a.InternalID),
			app.SourceEQ(toEntAppSource(a.Source)),
		).
		SetSourceAppID(a.SourceAppID).
		SetSourceURL(a.SourceURL).
		SetName(a.Name).
		SetType(toEntAppType(a.Type)).
		SetShortDescription(a.ShorDescription).
		SetImageURL(a.ImageURL)
	if a.Details != nil {
		q.
			SetDescription(a.Details.Description).
			SetReleaseDate(a.Details.ReleaseDate).
			SetDeveloper(a.Details.Developer).
			SetPublisher(a.Details.Publisher)
	}
	return q.Exec(ctx)
}

func (g geburaRepo) UpsertApp(ctx context.Context, al []*bizgebura.App) error {
	apps := make([]*ent.AppCreate, len(al))
	for i, a := range al {
		if a.Details == nil {
			a.Details = new(bizgebura.AppDetails)
		}
		apps[i] = g.data.db.App.Create().
			SetInternalID(a.InternalID).
			SetSource(toEntAppSource(a.Source)).
			SetSourceAppID(a.SourceAppID).
			SetSourceURL(a.SourceURL).
			SetName(a.Name).
			SetType(toEntAppType(a.Type)).
			SetShortDescription(a.ShorDescription).
			SetImageURL(a.ImageURL)
		if a.Details != nil {
			apps[i].
				SetDescription(a.Details.Description).
				SetReleaseDate(a.Details.ReleaseDate).
				SetDeveloper(a.Details.Developer).
				SetPublisher(a.Details.Publisher)
		}
	}
	return g.data.db.App.
		CreateBulk(apps...).
		OnConflict(
			sql.ConflictColumns(app.FieldSource, app.FieldSourceAppID),
			sql.ResolveWithIgnore(),
			sql.ResolveWith(func(u *sql.UpdateSet) {
				ignores := []string{
					app.FieldInternalID,
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
	paging bizgebura.Paging,
	sources []bizgebura.AppSource,
	types []bizgebura.AppType,
	ids []int64,
	containDetails bool) ([]*bizgebura.App, error) {
	q := g.data.db.App.Query()
	if len(sources) > 0 {
		sourceFilter := make([]app.Source, len(sources))
		for i, appSource := range sources {
			sourceFilter[i] = toEntAppSource(appSource)
		}
		q.Where(app.SourceIn(sourceFilter...))
	}
	if len(types) > 0 {
		typeFilter := make([]app.Type, len(types))
		for i, appType := range types {
			typeFilter[i] = toEntAppType(appType)
		}
		q.Where(app.TypeIn(typeFilter...))
	}
	if len(ids) > 0 {
		q.Where(app.InternalIDIn(ids...))
	}
	a, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	apps := make([]*bizgebura.App, len(a))
	for i, sa := range a {
		apps[i] = toBizApp(sa)
		if containDetails {
			apps[i].Details = toBizAppDetails(sa)
		}
	}
	return apps, nil
}

func (g geburaRepo) IsAppPackage(ctx context.Context, id int64) error {
	a, err := g.ListAppPackage(ctx, bizgebura.Paging{
		PageSize: 1,
		PageNum:  0,
	}, nil, []int64{id})
	if err != nil {
		return err
	}
	if len(a) != 1 {
		return errors.New("no such app package")
	}
	return nil
}

func (g geburaRepo) CreateAppPackage(ctx context.Context, ap *bizgebura.AppPackage) error {
	q := g.data.db.AppPackage.Create().
		SetInternalID(ap.InternalID).
		SetSource(toEntAppPackageSource(ap.Source)).
		SetSourceID(ap.SourceID).
		SetSourcePackageID(ap.SourcePackageID).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetBinaryName(ap.Binary.Name).
		SetBinarySize(ap.Binary.Size)
	return q.Exec(ctx)
}

func (g geburaRepo) UpdateAppPackage(ctx context.Context, ap *bizgebura.AppPackage) error {
	q := g.data.db.AppPackage.Update().
		Where(
			apppackage.InternalIDEQ(ap.InternalID),
			apppackage.SourceEQ(toEntAppPackageSource(ap.Source)),
		).
		SetName(ap.Name).
		SetDescription(ap.Description).
		SetBinaryName(ap.Binary.Name).
		SetBinarySize(ap.Binary.Size)
	return q.Exec(ctx)
}

func (g geburaRepo) UpsertAppPackage(ctx context.Context, apl []*bizgebura.AppPackage) error {
	appPackages := make([]*ent.AppPackageCreate, len(apl))
	for i, ap := range apl {
		appPackages[i] = g.data.db.AppPackage.Create().
			SetInternalID(ap.InternalID).
			SetSource(toEntAppPackageSource(ap.Source)).
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
					apppackage.FieldInternalID,
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
	paging bizgebura.Paging,
	sources []bizgebura.AppPackageSource,
	ids []int64,
) ([]*bizgebura.AppPackage, error) {
	q := g.data.db.AppPackage.Query()
	if len(sources) > 0 {
		sourceFilter := make([]apppackage.Source, len(sources))
		for i, apSource := range sources {
			sourceFilter[i] = toEntAppPackageSource(apSource)
		}
		q.Where(apppackage.SourceIn(sourceFilter...))
	}
	if len(ids) > 0 {
		q.Where(apppackage.InternalIDIn(ids...))
	}
	ap, err := q.
		Limit(paging.PageSize).
		Offset((paging.PageNum - 1) * paging.PageSize).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return toBizAppPackages(ap), nil
}

func (g geburaRepo) ListAllAppPackageIDOfOneSource(
	ctx context.Context,
	source bizgebura.AppPackageSource,
	sourceID int64,
) ([]string, error) {
	return g.data.db.AppPackage.Query().
		Where(
			apppackage.SourceEQ(toEntAppPackageSource(source)),
			apppackage.SourceIDEQ(sourceID),
		).
		Unique(true).
		Select(apppackage.FieldSourcePackageID).
		Strings(ctx)
}
