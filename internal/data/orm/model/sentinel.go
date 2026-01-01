package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Sentinel struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Name                  string
	Description           string
	URL                   string
	AlternativeURLs       []string `gorm:"serializer:json"`
	GetTokenPath          string
	DownloadFileBasePath  string
	CreatorID             model.InternalID
	LibraryReportSequence int64
	UpdatedAt             time.Time
	CreatedAt             time.Time
	SentinelSessions      []SentinelSession `gorm:"foreignKey:SentinelID"`
	SentinelLibraries     []SentinelLibrary `gorm:"foreignKey:SentinelID"`
}

func (Sentinel) TableName() string {
	return "sentinels"
}
