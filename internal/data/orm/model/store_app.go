package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type StoreApp struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Name                string
	Description         string
	UpdatedAt           time.Time
	CreatedAt           time.Time
	SentinelAppBinaries []SentinelAppBinary `gorm:"many2many:store_app_binaries;joinForeignKey:StoreAppID;joinReferences:SentinelAppBinaryUnionID"`
}

func (StoreApp) TableName() string {
	return "store_apps"
}
