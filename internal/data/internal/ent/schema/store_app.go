package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type StoreApp struct {
	ent.Schema
}

func (StoreApp) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
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
	return []ent.Edge{}
}
