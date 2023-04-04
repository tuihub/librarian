package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotifyFlow holds the schema definition for the NotifyFlow entity.
type NotifyFlow struct {
	ent.Schema
}

// Fields of the NotifyFlow.
func (NotifyFlow) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.Enum("status").
			Values("active", "suspend"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the NotifyFlow.
func (NotifyFlow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("notify_flow").
			Required().
			Unique(),
		edge.To("notify_target", NotifyTarget.Type).
			Through("notify_flow_target", NotifyFlowTarget.Type),
		edge.From("feed_config", FeedConfig.Type).
			Ref("notify_flow"),
	}
}
