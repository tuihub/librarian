package model

import (
	"github.com/tuihub/librarian/internal/model"
)

type AppAppCategory struct {
	AppCategoryID model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	AppID         model.InternalID `gorm:"primaryKey;autoIncrement:false"`
}

func (AppAppCategory) TableName() string {
	return "app_app_categories"
}
