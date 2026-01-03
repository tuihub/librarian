package modelfeed

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
)

type Feed struct {
	ID          model.InternalID `json:"internal_id,omitempty" gorm:"primaryKey;autoIncrement:false"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Link        string           `json:"link,omitempty"`
	Authors     []*Person        `json:"authors,omitempty"     gorm:"serializer:json"`
	Language    string           `json:"language,omitempty"`
	Image       *Image           `json:"image,omitempty"       gorm:"serializer:json"`
	Items       []*Item          `json:"items"                 gorm:"foreignKey:FeedID"`
	FeedType    string           `json:"feedType"              gorm:"-"` // Not stored in DB
	FeedVersion string           `json:"feedVersion"           gorm:"-"` // Not stored in DB
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

func (Feed) TableName() string {
	return "feeds"
}

type Item struct {
	ID                model.InternalID `json:"internal_id,omitempty"        gorm:"primaryKey;autoIncrement:false"`
	FeedID            model.InternalID `json:"feed_id,omitempty"            gorm:"index:idx_feed_item_feed_id_guid,priority:1"`
	Title             string           `json:"title,omitempty"`
	Description       string           `json:"description,omitempty"`
	Content           string           `json:"content,omitempty"`
	Link              string           `json:"link,omitempty"`
	Updated           string           `json:"updated,omitempty"`
	UpdatedParsed     *time.Time       `json:"updatedParsed,omitempty"`
	Published         string           `json:"published,omitempty"`
	PublishedParsed   *time.Time       `json:"publishedParsed,omitempty"`
	Authors           []*Person        `json:"authors,omitempty"            gorm:"serializer:json"`
	GUID              string           `json:"guid,omitempty"               gorm:"column:guid;index:idx_feed_item_feed_id_guid,priority:2"`
	Image             *Image           `json:"image,omitempty"              gorm:"serializer:json"`
	Enclosures        []*Enclosure     `json:"enclosures,omitempty"         gorm:"serializer:json"`
	PublishPlatform   string           `json:"publish_platform,omitempty"   gorm:"index"`
	ReadCount         int64            `json:"read_count,omitempty"`
	DigestDescription string           `json:"digest_description,omitempty"`
	DigestImages      []*Image         `json:"digest_images,omitempty"      gorm:"serializer:json"`
	UpdatedAt         time.Time
	CreatedAt         time.Time
	Feed              *Feed `json:"-"                            gorm:"foreignKey:FeedID"`
	// FeedItemCollections relation will be defined in the join table or implicit if needed
}

func (Item) TableName() string {
	return "feed_items"
}

// Person is an individual specified in a feed
// (e.g. an author).
type Person struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Image is an image that is the artwork for a given
// feed or item.
type Image struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

// Enclosure is a file associated with a given Item.
type Enclosure struct {
	URL    string `json:"url,omitempty"`
	Length string `json:"length,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Len returns the length of Items.
func (f Feed) Len() int {
	return len(f.Items)
}

// Less compares PublishedParsed of Items[i], Items[k]
// and returns true if Items[i] is less than Items[k].
func (f Feed) Less(i, k int) bool {
	if f.Items[i].PublishedParsed == nil || f.Items[k].PublishedParsed == nil {
		return false
	}
	return f.Items[i].PublishedParsed.Before(
		*f.Items[k].PublishedParsed,
	)
}

// Swap swaps Items[i] and Items[k].
func (f Feed) Swap(i, k int) {
	f.Items[i], f.Items[k] = f.Items[k], f.Items[i]
}
