package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Tag struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserTag     model.InternalID `gorm:"index"` // UserID
	Name        string
	Description string
	Public      bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Owner       *User `gorm:"foreignKey:UserTag"`
}

func (Tag) TableName() string {
	return "tags"
}
