package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FeedConfig holds the schema definition for the FeedConfig entity.
type FeedConfig struct {
	ent.Schema
}

// Fields of the FeedConfig.
func (FeedConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.String("feed_url"),
		field.Int64("author_account"),
		field.Enum("source").Values("common"),
		field.Enum("status").
			Values("active", "suspend"),
		field.Time("pull_interval"),
		field.Time("last_pull_at"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the FeedConfig.
func (FeedConfig) Edges() []ent.Edge {
	return nil
}
