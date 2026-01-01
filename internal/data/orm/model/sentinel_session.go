package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type SentinelSession struct {
	ID              model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	SentinelID      model.InternalID
	RefreshToken    string
	ExpireAt        time.Time
	Status          string
	CreatorID       model.InternalID
	LastUsedAt      *time.Time
	LastRefreshedAt *time.Time
	RefreshCount    int64
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Sentinel        *Sentinel `gorm:"foreignKey:SentinelID"`
}

func (SentinelSession) TableName() string {
	return "sentinel_sessions"
}
