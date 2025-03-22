package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppBinary struct {
	ent.Schema
}

func (AppBinary) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name").Optional(),
		field.Int64("size_bytes").Optional(),
		field.String("public_url").Optional(),
		field.Bytes("sha256").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppBinary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sha256").
			Unique(),
	}
}

func (AppBinary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app_info", AppInfo.Type).
			Ref("app_binary").
			Unique(),
	}
}
