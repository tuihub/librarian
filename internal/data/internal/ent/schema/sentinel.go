package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Sentinel struct {
	ent.Schema
}

func (Sentinel) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.String("url").Default(""),
		field.Strings("alternative_urls").Optional(),
		field.String("get_token_path").Optional(),
		field.String("download_file_base_path").Default(""),
		field.Int64("creator_id").Immutable().GoType(model.InternalID(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
		field.Int64("library_report_sequence").Default(0),
		field.Int64("app_binary_report_sequence").Default(0),
	}
}

func (Sentinel) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Sentinel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sentinel_session", SentinelSession.Type),
		edge.To("sentinel_library", SentinelLibrary.Type),
	}
}
