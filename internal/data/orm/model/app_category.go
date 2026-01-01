package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type AppCategory struct {
	ID               model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID           model.InternalID `gorm:"index"`
	VersionNumber    uint64
	VersionDate      time.Time
	Name             string
	UpdatedAt        time.Time
	CreatedAt        time.Time
	Apps             []App            `gorm:"many2many:app_app_categories;"`
	AppAppCategories []AppAppCategory `gorm:"foreignKey:AppCategoryID"`
}

func (AppCategory) TableName() string {
	return "app_categories"
}
