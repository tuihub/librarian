package schema

import (
	"github.com/tuihub/librarian/internal/model"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotifyTarget holds the schema definition for the NotifyTarget entity.
type NotifyTarget struct {
	ent.Schema
}

// Fields of the NotifyTarget.
func (NotifyTarget) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.JSON("destination", new(model.FeatureRequest)),
		field.Enum("status").
			Values("active", "suspend"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the NotifyTarget.
func (NotifyTarget) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("notify_target").
			Required().
			Unique(),
		edge.From("notify_flow", NotifyFlow.Type).
			Ref("notify_target").
			Through("notify_flow_target", NotifyFlowTarget.Type),
	}
}
