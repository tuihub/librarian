package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FeedConfig holds the schema definition for the FeedConfig entity.
type FeedConfig struct {
	ent.Schema
}

// Fields of the FeedConfig.
func (FeedConfig) Fields() []ent.Field {
	incrementalEnabled := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			GoType(model.InternalID(0)).
			Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
				Incremental: &incrementalEnabled,
			}),
		field.String("feed_url"),
		field.Int64("author_account").
			GoType(model.InternalID(0)),
		field.Enum("source").Values("common"),
		field.Enum("status").
			Values("active", "suspend"),
		field.Int64("pull_interval").
			GoType(time.Duration(0)),
		field.Time("next_pull_begin_at").
			Default(time.UnixMicro(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the FeedConfig.
func (FeedConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("feed_config").
			Required().
			Unique(),
		edge.To("feed", Feed.Type).
			Unique(),
	}
}
