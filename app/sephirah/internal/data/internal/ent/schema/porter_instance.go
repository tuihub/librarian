package schema

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"

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
		field.String("global_name"),
		field.String("address"),
		field.JSON("feature_summary", new(modeltiphereth.PorterFeatureSummary)),
		field.Enum("status").
			Values("active", "blocked"),
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
	}
}
