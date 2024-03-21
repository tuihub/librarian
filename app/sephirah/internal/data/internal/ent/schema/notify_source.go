package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type NotifySource struct {
	ent.Schema
}

func (NotifySource) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("feed_config_id").
			GoType(model.InternalID(0)).
			Optional(),
		field.Int64("feed_item_collection_id").
			GoType(model.InternalID(0)).
			Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (NotifySource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("notify_source").
			Required().
			Unique(),
		edge.From("notify_flow", NotifyFlow.Type).
			Ref("notify_source").
			Through("notify_flow_source", NotifyFlowSource.Type),
		edge.From("feed_config", FeedConfig.Type).
			Ref("notify_source").
			Unique().
			Field("feed_config_id"),
		edge.From("feed_item_collection", FeedItemCollection.Type).
			Ref("notify_source").
			Unique().
			Field("feed_item_collection_id"),
	}
}
