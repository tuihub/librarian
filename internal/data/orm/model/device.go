package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Device struct {
	ID                      model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	DeviceName              string
	SystemType              string
	SystemVersion           string
	ClientName              string
	ClientSourceCodeAddress string
	ClientVersion           string
	ClientLocalID           string
	UpdatedAt               time.Time
	CreatedAt               time.Time
	Sessions                []Session `gorm:"foreignKey:DeviceID"`
	App                     []App     `gorm:"foreignKey:CreatorDeviceID"`
}

func (Device) TableName() string {
	return "devices"
}
