package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.String("title"),
		field.String("link"),
		field.String("description"),
		field.String("language"),
		field.JSON("authors", []*modelfeed.Person{}),
		field.JSON("image", new(modelfeed.Image)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Feed.
func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("config", FeedConfig.Type).
			Ref("feed").
			Unique().
			Required(),
	}
}
