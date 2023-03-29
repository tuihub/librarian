package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

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
			GoType(model.InternalID(0)).
			Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
				Incremental: &incrementalEnabled,
			}),
		field.String("username").
			Unique(),
		field.String("password"),
		field.Enum("status").
			Values("active", "blocked"),
		field.Enum("type").
			Values("admin", "normal", "sentinel"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bind_account", Account.Type),
		edge.To("purchased_app", App.Type),
		edge.To("app_package", AppPackage.Type),
		edge.To("feed_config", FeedConfig.Type),
		edge.To("created_user", User.Type).
			From("creator").
			Unique().
			Required(),
	}
}