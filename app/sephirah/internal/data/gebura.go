package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizgebura"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
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

func (g geburaRepo) CreateApp(ctx context.Context, a *bizgebura.App) error {
	if a.Details == nil {
		a.Details = &bizgebura.AppDetails{}
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
