package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Account struct {
	ID                model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Platform          string           `gorm:"index:idx_account_platform_id,priority:1"`
	PlatformAccountID string           `gorm:"index:idx_account_platform_id,priority:2"`
	BoundUserID       model.InternalID `gorm:"index"`
	Name              string
	ProfileURL        string
	AvatarURL         string
	UpdatedAt         time.Time
	CreatedAt         time.Time
	BoundUser         *User `gorm:"foreignKey:BoundUserID"`
}

func (Account) TableName() string {
	return "accounts"
}
