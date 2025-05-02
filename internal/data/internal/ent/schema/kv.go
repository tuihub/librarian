package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type KV struct {
	ent.Schema
}

func (KV) Fields() []ent.Field {
	return []ent.Field{
		field.String("bucket"),
		field.String("key"),
		field.String("value"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (KV) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("bucket", "key").Unique(),
	}
}
