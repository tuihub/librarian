package schema

import (
	"time"

	"github.com/tuihub/librarian/model/modelfeed"

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
		defaultPrimaryKey(),
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
