package data

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/bizchesed"
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
		SetFileID(image.ID).
		SetOwnerID(userID).
		Exec(ctx)
}
