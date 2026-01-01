package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type StoreAppBinary struct {
	StoreAppID               model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelAppBinaryUnionID model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UpdatedAt                time.Time
	CreatedAt                time.Time
}

func (StoreAppBinary) TableName() string {
	return "store_app_binaries"
}
