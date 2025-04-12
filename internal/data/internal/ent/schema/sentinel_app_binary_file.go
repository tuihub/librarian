package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelAppBinaryFile struct {
	ent.Schema
}

func (SentinelAppBinaryFile) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("sentinel_app_binary_id"),
		field.Int64("sentinel_info_id").GoType(model.InternalID(0)),
		field.Int64("sentinel_library_reported_id"),
		field.String("sentinel_app_binary_generated_id"),
		field.String("name"),
		field.Int64("size_bytes"),
		field.Bytes("sha256"),
		field.String("server_file_path"),
		field.String("chunks_info").Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
		field.Int64("app_binary_report_sequence"),
	}
}

func (SentinelAppBinaryFile) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("sentinel_app_binary_id", "server_file_path").
		//	Unique(),
		index.Fields("sentinel_info_id", "sentinel_library_reported_id",
			"sentinel_app_binary_generated_id", "server_file_path").
			Unique(),
		index.Fields("app_binary_report_sequence"),
	}
}

func (SentinelAppBinaryFile) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("sentinel_app_binary", SentinelAppBinary.Type).
		//	Required().
		//	Unique().
		//	Field("sentinel_app_binary_id"),
	}
}
