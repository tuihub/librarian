package schema

import (
	"time"

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
		edge.To("purchased_app", AppInfo.Type),
		edge.To("app", App.Type),
		edge.To("app_inst", AppInst.Type),
		edge.To("feed_config", FeedConfig.Type),
		edge.To("feed_action_set", FeedActionSet.Type),
		edge.To("feed_item_collection", FeedItemCollection.Type),
		edge.To("notify_source", NotifySource.Type),
		edge.To("notify_target", NotifyTarget.Type),
		edge.To("notify_flow", NotifyFlow.Type),
		edge.To("image", Image.Type),
		edge.To("file", File.Type),
		edge.To("device_info", DeviceInfo.Type).
			Through("user_device", UserDevice.Type),
		edge.To("tag", Tag.Type),
		edge.To("porter_context", PorterContext.Type),
		edge.To("created_user", User.Type).
			From("creator").
			Unique().
			Required(),
	}
}
