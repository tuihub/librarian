package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelInfo struct {
	ent.Schema
}

func (SentinelInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").GoType(model.InternalID(0)),
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
	return []ent.Index{
		index.Fields("user_id"),
	}
}

func (SentinelInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sentinel_library", SentinelLibrary.Type).
			Ref("sentinel_info"),
	}
}
