package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("device_name"),
		field.Enum("system_type").Values("ios", "android", "web", "windows", "macos", "linux", "unknown"),
		field.String("system_version"),
		field.String("client_name"),
		field.String("client_source_code_address"),
		field.String("client_version"),
		field.String("client_local_id").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session", Session.Type),
		edge.To("app", App.Type),
	}
}
