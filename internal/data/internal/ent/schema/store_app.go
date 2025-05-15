package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type StoreApp struct {
	ent.Schema
}

func (StoreApp) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (StoreApp) Indexes() []ent.Index {
	return []ent.Index{}
}

func (StoreApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app_binary", SentinelAppBinary.Type).
			Through("store_app_binary", StoreAppBinary.Type),
	}
}
