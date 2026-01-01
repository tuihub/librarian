package model

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"
)

type FeedItem struct {
	ID                  model.InternalID `gorm:"primaryKey;autoIncrement:false"`
	FeedID              model.InternalID `gorm:"index:idx_feed_item_feed_id_guid,priority:1"`
	Title               string
	Authors             []*modelfeed.Person `gorm:"serializer:json"`
	Description         string
	Content             string
	GUID                string `gorm:"column:guid;index:idx_feed_item_feed_id_guid,priority:2"`
	Link                string
	Image               *modelfeed.Image `gorm:"serializer:json"`
	Published           string
	PublishedParsed     time.Time
	Updated             string
	UpdatedParsed       *time.Time
	Enclosures          []*modelfeed.Enclosure `gorm:"serializer:json"`
	PublishPlatform     string                 `gorm:"index"`
	ReadCount           int64
	DigestDescription   string
	DigestImages        []*modelfeed.Image `gorm:"serializer:json"`
	UpdatedAt           time.Time
	CreatedAt           time.Time
	Feed                *Feed                `gorm:"foreignKey:FeedID"`
	FeedItemCollections []FeedItemCollection `gorm:"many2many:feed_item_collection_feed_items;"` // Ent usually creates a join table or adds FK
}

func (FeedItem) TableName() string {
	return "feed_items"
}
