package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type StoreAppBinary struct {
	ent.Schema
}

func (StoreAppBinary) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("store_app_id").GoType(model.InternalID(0)),
		field.Int64("sentinel_app_binary_union_id").GoType(model.InternalID(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (StoreAppBinary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("store_app_id", "sentinel_app_binary_union_id").
			Unique(),
		index.Fields("sentinel_app_binary_union_id").
			Unique(),
	}
}

func (StoreAppBinary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("store_app", StoreApp.Type).
			Unique().
			Required().
			Field("store_app_id"),
		edge.To("sentinel_app_binary", SentinelAppBinary.Type).
			Unique().
			Required().
			Field("sentinel_app_binary_union_id"),
	}
}
