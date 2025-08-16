package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model/modelsupervisor"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PorterInstance struct {
	ent.Schema
}

func (PorterInstance) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("version"),
		field.String("description"),
		field.String("source_code_address"),
		field.String("build_version"),
		field.String("build_date"),
		field.String("global_name"),
		field.String("address"),
		field.String("region"),
		field.JSON("feature_summary", new(modelsupervisor.PorterFeatureSummary)),
		field.String("context_json_schema"),
		field.Enum("status").
			Values("active", "blocked"),
		field.Enum("connection_status").
			Values("unspecified", "queueing", "connected", "disconnected", "active", "activation_failed", "downgraded"),
		field.String("connection_status_message"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (PorterInstance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("address").
			Unique(),
		index.Fields("global_name", "region"),
	}
}
