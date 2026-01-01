package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type Feed struct {
	ID          model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	Title       string
	Link        string
	Description string
	Language    string
	Authors     []*modelfeed.Person `gorm:"serializer:json"`
	Image       *modelfeed.Image    `gorm:"serializer:json"`
	UpdatedAt   time.Time
	CreatedAt   time.Time
	Item        []FeedItem  `gorm:"foreignKey:FeedID"`
	Config      *FeedConfig `gorm:"foreignKey:ID;references:ID"` // Relation check needed
}

func (Feed) TableName() string {
	return "feeds"
}
