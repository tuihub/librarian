package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("username").
			Unique(),
		field.String("password"),
		field.Enum("status").
			Values("active", "blocked"),
		field.Enum("type").
			Values("admin", "normal"),
		field.Int64("creator_id").
			GoType(model.InternalID(0)).Immutable(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session", Session.Type),
		edge.To("account", Account.Type),
		edge.To("app", App.Type),
		edge.To("feed_config", FeedConfig.Type),
		edge.To("feed_action_set", FeedActionSet.Type),
		edge.To("feed_item_collection", FeedItemCollection.Type),
		edge.To("notify_source", NotifySource.Type),
		edge.To("notify_target", NotifyTarget.Type),
		edge.To("notify_flow", NotifyFlow.Type),
		edge.To("image", Image.Type),
		edge.To("file", File.Type),
		edge.To("tag", Tag.Type),
		edge.To("porter_context", PorterContext.Type),
		edge.To("created_user", User.Type).
			From("creator").
			Field("creator_id").
			Unique().
			Required().
			Immutable(),
	}
}
