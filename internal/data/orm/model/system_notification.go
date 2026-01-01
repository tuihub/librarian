package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type SystemNotification struct {
	ID        model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID    model.InternalID `gorm:"index"`
	Type      string
	Level     string
	Status    string
	Title     string
	Content   string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (SystemNotification) TableName() string {
	return "system_notifications"
}
