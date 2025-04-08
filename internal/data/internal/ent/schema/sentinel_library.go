package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelLibrary struct {
	ent.Schema
}

func (SentinelLibrary) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sentinel_info_id"),
		field.Int64("reported_id"),
		field.String("download_base_path"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}
func (SentinelLibrary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sentinel_info_id", "reported_id"),
	}
}

func (SentinelLibrary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sentinel_info", SentinelInfo.Type).
			Required().
			Unique().
			Field("sentinel_info_id"),
		edge.From("sentinel_app_binary", SentinelAppBinary.Type).
			Ref("sentinel_library"),
	}
}
