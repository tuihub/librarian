package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FeedItem holds the schema definition for the FeedItem entity.
type FeedItem struct {
	ent.Schema
}

// Fields of the FeedItem.
func (FeedItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.String("title"),
		field.JSON("authors", []Person{}),
		field.String("description"),
		field.String("content"),
		field.String("guid"),
		field.String("link"),
		field.JSON("images", []Image{}),
		field.String("published"),
		field.Time("published_parsed"),
		field.String("updated"),
		field.Time("updated_parsed"),
		field.JSON("enclosure", []Enclosure{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

type Enclosure struct {
	URL    string
	Length string
	Type   string
}

// Edges of the FeedItem.
func (FeedItem) Edges() []ent.Edge {
	return nil
}
