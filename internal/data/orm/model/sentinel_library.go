package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type SentinelLibrary struct {
	ID                    model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID            model.InternalID `gorm:"index:idx_sentinel_library_sentinel_id_reported_id,priority:1"`
	ReportedID            int64            `gorm:"index:idx_sentinel_library_sentinel_id_reported_id,priority:2"`
	DownloadBasePath      string
	ActiveSnapshot        *time.Time
	LibraryReportSequence int64 `gorm:"index"`
	UpdatedAt             time.Time
	CreatedAt             time.Time
	Sentinel              *Sentinel `gorm:"foreignKey:SentinelID"`
}

func (SentinelLibrary) TableName() string {
	return "sentinel_libraries"
}
