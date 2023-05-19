package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/converter"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelchesed"
	"github.com/tuihub/librarian/internal/model"
)

type chesedRepo struct {
	data *Data
}

// NewChesedRepo .
func NewChesedRepo(data *Data) bizchesed.ChesedRepo {
	return &chesedRepo{
		data: data,
	}
}

func (c chesedRepo) CreateImage(ctx context.Context, userID model.InternalID, image *modelchesed.Image) error {
	return c.data.db.Image.Create().
		SetID(image.ID).
		SetName(image.Name).
		SetDescription(image.Description).
		SetStatus(converter.ToEntImageStatus(image.Status)).
		SetFileID(image.ID).
		SetOwnerID(userID).
		Exec(ctx)
}

func (c chesedRepo) ListImageNeedScan(ctx context.Context) ([]*modelchesed.Image, error) {
	res, err := c.data.db.Image.Query().
		Where(image.StatusEQ(image.StatusUploaded)).
		Limit(10). //nolint:gomnd //TODO
		All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizImageList(res), nil
}

func (c chesedRepo) SetImageStatus(ctx context.Context, id model.InternalID, status modelchesed.ImageStatus) error {
	return c.data.db.Image.UpdateOneID(id).
		SetStatus(converter.ToEntImageStatus(status)).
		Exec(ctx)
}
