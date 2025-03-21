package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppInfo struct {
	ent.Schema
}

func (AppInfo) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Bool("internal"),
		field.String("source"),
		field.String("source_app_id"),
		field.String("source_url").Optional(),
		field.String("name"),
		field.Enum("type").
			Values("unknown", "game"),
		field.String("short_description").Optional(),
		field.Text("description").Optional(),
		field.String("icon_image_url").Optional(),
		field.String("background_image_url").Optional(),
		field.String("cover_image_url").Optional(),
		field.String("release_date").Optional(),
		field.String("developer").Optional(),
		field.String("publisher").Optional(),
		field.String("version").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source", "source_app_id").
			Unique(),
	}
}

func (AppInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("purchased_by_account", Account.Type).
			Ref("purchased_app"),
		edge.From("purchased_by_user", User.Type).
			Ref("purchased_app"),
		edge.To("app", App.Type),
		edge.To("app_binary", AppBinary.Type),
		edge.To("bind_external", AppInfo.Type).
			From("bind_internal").
			Unique(),
	}
}
