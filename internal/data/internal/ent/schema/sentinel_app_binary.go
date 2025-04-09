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
		field.Int64("sentinel_library_id").GoType(model.InternalID(0)),
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
		field.Int64("report_sequence"),
	}
}

func (SentinelAppBinary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sentinel_library_id", "generated_id").
			Unique(),
		index.Fields("generated_id"),
		index.Fields("report_sequence"),
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
