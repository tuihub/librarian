package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DeviceInfo struct {
	ent.Schema
}

func (DeviceInfo) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("device_model"),
		field.String("system_version"),
		field.String("client_name"),
		field.String("client_source_code_address"),
		field.String("client_version"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (DeviceInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_session", UserSession.Type),
	}
}
