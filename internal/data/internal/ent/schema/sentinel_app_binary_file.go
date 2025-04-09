package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelAppBinaryFile struct {
	ent.Schema
}

func (SentinelAppBinaryFile) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("sentinel_app_binary_id").GoType(model.InternalID(0)),
		field.String("name"),
		field.Int64("size_bytes"),
		field.Bytes("sha256"),
		field.String("server_file_path"),
		field.String("chunks_info").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (SentinelAppBinaryFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sentinel_app_binary_id", "name"),
	}
}

func (SentinelAppBinaryFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sentinel_app_binary", SentinelAppBinary.Type).
			Required().
			Unique().
			Field("sentinel_app_binary_id"),
	}
}
