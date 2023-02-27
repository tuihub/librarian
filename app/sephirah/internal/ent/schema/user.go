package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	incrementalEnabled := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
				Incremental: &incrementalEnabled,
			}),
		field.String("username").
			Unique(),
		field.String("password"),
		field.Enum("status").
			Values("active", "blocked"),
		field.Enum("type").
			Values("admin"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type),
		edge.To("app", App.Type),
		edge.To("feed_config", FeedConfig.Type),
		edge.To("create", User.Type).
			From("creator").
			Unique(),
	}
}
