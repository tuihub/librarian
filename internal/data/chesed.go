package data

import (
	"context"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/internal/ent/image"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelchesed"
)

type ChesedRepo struct {
	data *Data
}

// NewChesedRepo .
func NewChesedRepo(data *Data) *ChesedRepo {
	return &ChesedRepo{
		data: data,
	}
}

func (c *ChesedRepo) CreateImage(ctx context.Context, userID model.InternalID, image *modelchesed.Image) error {
	return c.data.db.Image.Create().
		SetID(image.ID).
		SetName(image.Name).
		SetDescription(image.Description).
		SetStatus(converter.ToEntImageStatus(image.Status)).
		SetFileID(image.ID).
		SetOwnerID(userID).
		Exec(ctx)
}

func (c *ChesedRepo) ListImages(ctx context.Context, userID model.InternalID, paging model.Paging) (
	[]*modelchesed.Image, int64, error) {
	q := c.data.db.Image.Query().
		Where(
			image.HasOwnerWith(user.IDEQ(userID)),
		)
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
	return converter.ToBizImageList(res), int64(total), nil
}

func (c *ChesedRepo) ListImageNeedScan(ctx context.Context) ([]*modelchesed.Image, error) {
	res, err := c.data.db.Image.Query().
		Where(image.StatusEQ(image.StatusUploaded)).
		Limit(10). //nolint:mnd //TODO
		All(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizImageList(res), nil
}

func (c *ChesedRepo) SetImageStatus(ctx context.Context, id model.InternalID, status modelchesed.ImageStatus) error {
	return c.data.db.Image.UpdateOneID(id).
		SetStatus(converter.ToEntImageStatus(status)).
		Exec(ctx)
}

func (c *ChesedRepo) GetImage(ctx context.Context, userID model.InternalID, id model.InternalID) (
	*modelchesed.Image, error) {
	res, err := c.data.db.Image.Query().
		Where(
			image.IDEQ(id),
			image.HasOwnerWith(user.IDEQ(userID)),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return converter.ToBizImage(res), nil
}
