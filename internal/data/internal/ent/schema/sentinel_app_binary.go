package schema

import (
	"time"

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
		field.Int("sentinel_library_id"),
		field.String("generated_id"),
		field.Int64("size_bytes"),
		field.Bool("need_token"),
		field.String("name").Optional(),
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
		index.Fields("sentinel_library_id", "generated_id"),
	}
}

func (SentinelAppBinary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sentinel_library", SentinelLibrary.Type).
			Required().
			Unique().
			Field("sentinel_library_id"),
		edge.From("sentinel_app_binary_file", SentinelAppBinaryFile.Type).
			Ref("sentinel_app_binary"),
	}
}
