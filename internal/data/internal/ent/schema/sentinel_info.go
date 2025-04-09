package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SentinelInfo struct {
	ent.Schema
}

func (SentinelInfo) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("url"),
		field.Strings("alternative_urls").Optional(),
		field.String("get_token_path").Optional(),
		field.String("download_file_base_path"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (SentinelInfo) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SentinelInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sentinel_library", SentinelLibrary.Type).
			Ref("sentinel_info"),
	}
}
