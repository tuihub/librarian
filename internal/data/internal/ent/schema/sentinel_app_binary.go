package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelAppBinary struct {
	ent.Schema
}

func (SentinelAppBinary) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("union_id"),
		field.Int64("sentinel_id").GoType(model.InternalID(0)),
		field.Int64("sentinel_library_reported_id"),
		field.Time("library_snapshot"),
		field.String("generated_id"),
		field.Int64("size_bytes"),
		field.Bool("need_token"),
		field.String("name"),
		field.String("version").Optional(),
		field.String("developer").Optional(),
		field.String("publisher").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (SentinelAppBinary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("union_id"),
		index.Fields("sentinel_id", "sentinel_library_reported_id", "library_snapshot", "generated_id").
			Unique(),
	}
}

func (SentinelAppBinary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("store_app", StoreApp.Type).
			Ref("app_binary").
			Through("store_app_binary", StoreAppBinary.Type),
	}
}
