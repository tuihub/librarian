package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model/modelfeed"

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
		field.JSON("authors", []modelfeed.Person{}),
		field.String("description"),
		field.String("content"),
		field.String("guid"),
		field.String("link"),
		field.JSON("image", new(modelfeed.Image)),
		field.String("published"),
		field.Time("published_parsed"),
		field.String("updated"),
		field.Time("updated_parsed"),
		field.JSON("enclosure", []modelfeed.Enclosure{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the FeedItem.
func (FeedItem) Edges() []ent.Edge {
	return nil
}
