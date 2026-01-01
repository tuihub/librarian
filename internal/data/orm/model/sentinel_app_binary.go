package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type SentinelAppBinary struct {
	ID                        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UnionID                   string           `gorm:"index"`
	SentinelID                model.InternalID
	SentinelLibraryReportedID int64
	LibrarySnapshot           time.Time
	GeneratedID               string
	SizeBytes                 int64
	NeedToken                 bool
	Name                      string
	Version                   string
	Developer                 string
	Publisher                 string
	UpdatedAt                 time.Time
	CreatedAt                 time.Time
	StoreApps                 []StoreApp `gorm:"many2many:store_app_binaries;joinForeignKey:SentinelAppBinaryUnionID;joinReferences:StoreAppID"`
}

func (SentinelAppBinary) TableName() string {
	return "sentinel_app_binaries"
}
