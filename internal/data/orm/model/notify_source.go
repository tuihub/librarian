package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type NotifySource struct {
	ID                   model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	OwnerID              model.InternalID `gorm:"column:user_notify_source"` // Inferred
	FeedConfigID         *model.InternalID
	FeedItemCollectionID *model.InternalID
	UpdatedAt            time.Time
	CreatedAt            time.Time
	Owner                *User               `gorm:"foreignKey:OwnerID"`
	FeedConfig           *FeedConfig         `gorm:"foreignKey:FeedConfigID"`
	FeedItemCollection   *FeedItemCollection `gorm:"foreignKey:FeedItemCollectionID"`
	NotifyFlows          []NotifyFlow        `gorm:"many2many:notify_flow_sources;"`
}

func (NotifySource) TableName() string {
	return "notify_sources"
}
