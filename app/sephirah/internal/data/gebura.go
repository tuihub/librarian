package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz"
)

type geburaRepo struct {
	data *Data
}

// NewGeburaRepo .
func NewGeburaRepo(data *Data) biz.GeburaRepo {
	return &geburaRepo{
		data: data,
	}
}

func (g geburaRepo) CreateApp(ctx context.Context, app *biz.App) error {
	_, err := g.data.db.App.Create().
		SetInternalID(app.InternalID).
		SetSource(toEntAppSource(app.Source)).
		SetSourceAppID(app.SourceAppID).
		SetSourceURL(app.SourceURL).
		SetName(app.Name).
		SetType(toEntAppType(app.Type)).
		SetShortDescription(app.ShorDescription).
		SetImageURL(app.ImageURL).
		SetDescription(app.Details.Description).
		SetReleaseDate(app.Details.ReleaseDate).
		SetDeveloper(app.Details.Developer).
		SetPublisher(app.Details.Publisher).
		Save(ctx)
	return err
}
