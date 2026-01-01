package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type AppInfo struct {
	ID                 model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Source             string           `gorm:"index:idx_app_info_source_source_app_id,priority:1"`
	SourceAppID        string           `gorm:"index:idx_app_info_source_source_app_id,priority:2"`
	SourceURL          string
	Name               string
	Type               string
	ShortDescription   string
	Description        string
	IconImageURL       string
	IconImageID        model.InternalID
	BackgroundImageURL string
	BackgroundImageID  model.InternalID
	CoverImageURL      string
	CoverImageID       model.InternalID
	ReleaseDate        string
	Developer          string
	Publisher          string
	Tags               []string `gorm:"serializer:json"`
	AlternativeNames   []string `gorm:"serializer:json"`
	RawData            string
	UpdatedAt          time.Time
	CreatedAt          time.Time
}

func (AppInfo) TableName() string {
	return "app_infos"
}
