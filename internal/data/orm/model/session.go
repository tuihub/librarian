package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Session struct {
	ID           model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID       model.InternalID `gorm:"index:idx_session_user_id_device_id,priority:1"`
	DeviceID     model.InternalID `gorm:"index:idx_session_user_id_device_id,priority:2"`
	RefreshToken string           `gorm:"uniqueIndex"`
	ExpireAt     time.Time
	UpdatedAt    time.Time
	CreatedAt    time.Time
	User         *User   `gorm:"foreignKey:UserID"`
	Device       *Device `gorm:"foreignKey:DeviceID"`
}

func (Session) TableName() string {
	return "sessions"
}
