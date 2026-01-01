package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type User struct {
	ID            model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Username      string           `gorm:"uniqueIndex"`
	Password      string
	Status        string
	Type          string
	CreatorID     model.InternalID
	UpdatedAt     time.Time
	CreatedAt     time.Time
	Sessions      []Session       `gorm:"foreignKey:UserID"`
	Account       []Account       `gorm:"foreignKey:BoundUserID"`
	App           []App           `gorm:"foreignKey:UserID"`
	FeedConfig    []FeedConfig    `gorm:"foreignKey:UserFeedConfig"`
	NotifySource  []NotifySource  `gorm:"foreignKey:OwnerID"`
	NotifyTarget  []NotifyTarget  `gorm:"foreignKey:OwnerID"`
	NotifyFlow    []NotifyFlow    `gorm:"foreignKey:OwnerID"`
	Image         []Image         `gorm:"foreignKey:OwnerID"`
	File          []File          `gorm:"foreignKey:OwnerID"`
	Tag           []Tag           `gorm:"foreignKey:UserTag"`
	PorterContext []PorterContext `gorm:"foreignKey:OwnerID"` // Check relationship later
	CreatedUser   []User          `gorm:"foreignKey:CreatorID"`
	Creator       *User           `gorm:"foreignKey:CreatorID"`
}

func (User) TableName() string {
	return "users"
}
