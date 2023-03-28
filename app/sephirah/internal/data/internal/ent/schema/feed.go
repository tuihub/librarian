package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	incrementalEnabled := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			GoType(model.InternalID(0)).
			Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
				Incremental: &incrementalEnabled,
			}),
		field.String("title").Optional(),
		field.String("link").Optional(),
		field.String("description").Optional(),
		field.String("language").Optional(),
		field.JSON("authors", []*modelfeed.Person{}).Optional(),
		field.JSON("image", new(modelfeed.Image)).Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Feed.
func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", FeedItem.Type),
		edge.From("config", FeedConfig.Type).
			Ref("feed").
			Unique().
			Required(),
	}
}
