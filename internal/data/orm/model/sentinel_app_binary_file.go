package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type SentinelAppBinaryFile struct {
	ID                           model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID                   model.InternalID
	SentinelLibraryReportedID    int64
	LibrarySnapshot              time.Time
	SentinelAppBinaryGeneratedID string
	Name                         string
	SizeBytes                    int64
	Sha256                       []byte
	ServerFilePath               string
	ChunksInfo                   string
	UpdatedAt                    time.Time
	CreatedAt                    time.Time
}

func (SentinelAppBinaryFile) TableName() string {
	return "sentinel_app_binary_files"
}
