package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Image struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID     model.InternalID `gorm:"column:user_image"` // Inferred
	FileID      model.InternalID `gorm:"column:file_image"` // Inferred
	Name        string
	Description string
	Status      string
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *User `gorm:"foreignKey:OwnerID"`
	File        *File `gorm:"foreignKey:FileID"`
}

func (Image) TableName() string {
	return "images"
}
