//nolint:dupl // TODO
package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotifyFlowSource holds the schema definition for the NotifyFlowSource entity.
type NotifyFlowSource struct {
	ent.Schema
}

// Fields of the NotifyFlowSource.
func (NotifyFlowSource) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("notify_flow_id").GoType(model.InternalID(0)),
		field.Int64("notify_source_id").GoType(model.InternalID(0)),
		field.JSON("filter_include_keywords", []string{}),
		field.JSON("filter_exclude_keywords", []string{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (NotifyFlowSource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("notify_flow_id", "notify_source_id").
			Unique(),
	}
}

// Edges of the NotifyFlowSource.
func (NotifyFlowSource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notify_flow", NotifyFlow.Type).
			Unique().
			Required().
			Field("notify_flow_id"),
		edge.To("notify_source", NotifySource.Type).
			Unique().
			Required().
			Field("notify_source_id"),
	}
}
