package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

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
		defaultPrimaryKey(),
		field.Int64("sentinel_info_id").GoType(model.InternalID(0)),
		field.Int64("reported_id"),
		field.String("download_base_path"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
		field.Int64("report_sequence"),
	}
}

func (SentinelLibrary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sentinel_info_id", "reported_id").
			Unique(),
		index.Fields("report_sequence"),
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
