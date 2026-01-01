package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type App struct {
	ID                 model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	VersionNumber      uint64
	VersionDate        time.Time
	UserID             model.InternalID `gorm:"index"`
	CreatorDeviceID    model.InternalID
	AppSources         map[string]string `gorm:"serializer:json"`
	Public             bool
	BoundStoreAppID    model.InternalID
	StopStoreManage    bool
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
	UpdatedAt          time.Time
	CreatedAt          time.Time
	User               *User         `gorm:"foreignKey:UserID"`
	Device             *Device       `gorm:"foreignKey:CreatorDeviceID"`
	AppRunTime         []AppRunTime  `gorm:"foreignKey:AppID"`
	AppCategories      []AppCategory `gorm:"many2many:app_app_categories;"`
}

func (App) TableName() string {
	return "apps"
}
