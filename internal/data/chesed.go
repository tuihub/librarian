package data

import (
	"context"

	"github.com/tuihub/librarian/internal/data/internal/gormschema"
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
	img := gormschema.Image{
		ID:          image.ID,
		OwnerID:     userID,
		FileID:      image.ID,
		Name:        image.Name,
		Description: image.Description,
		Status:      gormschema.ToSchemaImageStatus(image.Status),
	}
	return c.data.db.WithContext(ctx).Create(&img).Error
}

func (c *ChesedRepo) ListImages(ctx context.Context, userID model.InternalID, paging model.Paging) (
	[]*modelchesed.Image, int64, error) {
	query := c.data.db.WithContext(ctx).Model(&gormschema.Image{}).
		Where("owner_id = ?", userID)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var images []gormschema.Image
	if err := query.Limit(paging.ToLimit()).Offset(paging.ToOffset()).Find(&images).Error; err != nil {
		return nil, 0, err
	}

	res := make([]*modelchesed.Image, len(images))
	for i := range images {
		res[i] = gormschema.ToBizImage(&images[i])
	}
	return res, total, nil
}

func (c *ChesedRepo) ListImageNeedScan(ctx context.Context) ([]*modelchesed.Image, error) {
	var images []gormschema.Image
	err := c.data.db.WithContext(ctx).
		Where("status = ?", "uploaded").
		Limit(10). //nolint:mnd //TODO
		Find(&images).Error
	if err != nil {
		return nil, err
	}

	res := make([]*modelchesed.Image, len(images))
	for i := range images {
		res[i] = gormschema.ToBizImage(&images[i])
	}
	return res, nil
}

func (c *ChesedRepo) SetImageStatus(ctx context.Context, id model.InternalID, status modelchesed.ImageStatus) error {
	return c.data.db.WithContext(ctx).
		Model(&gormschema.Image{}).
		Where("id = ?", id).
		Update("status", gormschema.ToSchemaImageStatus(status)).Error
}

func (c *ChesedRepo) GetImage(ctx context.Context, userID model.InternalID, id model.InternalID) (
	*modelchesed.Image, error) {
	var img gormschema.Image
	err := c.data.db.WithContext(ctx).
		Where("id = ? AND owner_id = ?", id, userID).
		First(&img).Error
	if err != nil {
		return nil, err
	}
	return gormschema.ToBizImage(&img), nil
}
