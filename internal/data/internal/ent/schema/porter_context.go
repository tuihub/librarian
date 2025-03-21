package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PorterContext struct {
	ent.Schema
}

func (PorterContext) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("global_name"),
		field.String("region"),
		field.String("context_json"),
		field.String("name"),
		field.String("description"),
		field.Enum("status").
			Values("active", "disabled"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (PorterContext) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("global_name", "region").
			Unique(),
	}
}

func (PorterContext) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("porter_context").
			Required().
			Unique(),
	}
}
