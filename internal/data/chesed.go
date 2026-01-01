package data

import (
	"context"

	"github.com/tuihub/librarian/internal/data/internal/converter"
	"github.com/tuihub/librarian/internal/data/orm/model"
	"github.com/tuihub/librarian/internal/data/orm/query"
	libmodel "github.com/tuihub/librarian/internal/model"
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

func (c *ChesedRepo) CreateImage(ctx context.Context, userID libmodel.InternalID, image *modelchesed.Image) error {
	return query.Use(c.data.db).Image.WithContext(ctx).Create(&model.Image{
		ID:          image.ID,
		OwnerID:     userID,
		Name:        image.Name,
		Description: image.Description,
		Status:      converter.ToORMImageStatus(image.Status),
		FileID:      image.ID,
	})
}

func (c *ChesedRepo) ListImages(ctx context.Context, userID libmodel.InternalID, paging libmodel.Paging) (
	[]*modelchesed.Image, int64, error) {
	q := query.Use(c.data.db).Image
	u := q.WithContext(ctx).Where(q.OwnerID.Eq(int64(userID)))

	total, err := u.Count()
	if err != nil {
		return nil, 0, err
	}
	res, err := u.
		Limit(paging.ToLimit()).
		Offset(paging.ToOffset()).
		Find()
	if err != nil {
		return nil, 0, err
	}
	return converter.ToBizImageList(res), total, nil
}

func (c *ChesedRepo) ListImageNeedScan(ctx context.Context) ([]*modelchesed.Image, error) {
	q := query.Use(c.data.db).Image
	res, err := q.WithContext(ctx).
		Where(q.Status.Eq(converter.ToORMImageStatus(modelchesed.ImageStatusUploaded))).
		Limit(10). //nolint:mnd //TODO
		Find()
	if err != nil {
		return nil, err
	}
	return converter.ToBizImageList(res), nil
}

func (c *ChesedRepo) SetImageStatus(ctx context.Context, id libmodel.InternalID, status modelchesed.ImageStatus) error {
	q := query.Use(c.data.db).Image
	_, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id))).
		Update(q.Status, converter.ToORMImageStatus(status))
	return err
}

func (c *ChesedRepo) GetImage(ctx context.Context, userID libmodel.InternalID, id libmodel.InternalID) (
	*modelchesed.Image, error) {
	q := query.Use(c.data.db).Image
	res, err := q.WithContext(ctx).
		Where(q.ID.Eq(int64(id)), q.OwnerID.Eq(int64(userID))).
		First()
	if err != nil {
		return nil, err
	}
	return converter.ToBizImage(res), nil
}
