package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type FeedItemCollection struct {
	ID           model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	UserID       model.InternalID `gorm:"column:user_feed_item_collection"` // Inferred
	Name         string
	Description  string
	Category     string `gorm:"index"`
	UpdatedAt    time.Time
	CreatedAt    time.Time
	Owner        *User          `gorm:"foreignKey:UserID"`
	FeedItems    []FeedItem     `gorm:"many2many:feed_item_collection_feed_items;"`
	NotifySource []NotifySource `gorm:"foreignKey:ID;references:ID"` // Check
}

func (FeedItemCollection) TableName() string {
	return "feed_item_collections"
}
