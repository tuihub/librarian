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

// NotifyFlowTarget holds the schema definition for the NotifyFlowTarget entity.
type NotifyFlowTarget struct {
	ent.Schema
}

// Fields of the NotifyFlowTarget.
func (NotifyFlowTarget) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("notify_flow_id").GoType(model.InternalID(0)),
		field.Int64("notify_target_id").GoType(model.InternalID(0)),
		field.JSON("filter_include_keywords", []string{}),
		field.JSON("filter_exclude_keywords", []string{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (NotifyFlowTarget) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("notify_flow_id", "notify_target_id").
			Unique(),
	}
}

// Edges of the NotifyFlowTarget.
func (NotifyFlowTarget) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notify_flow", NotifyFlow.Type).
			Unique().
			Required().
			Field("notify_flow_id"),
		edge.To("notify_target", NotifyTarget.Type).
			Unique().
			Required().
			Field("notify_target_id"),
	}
}
