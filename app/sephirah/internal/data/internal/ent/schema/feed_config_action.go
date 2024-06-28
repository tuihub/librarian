package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FeedConfigAction struct {
	ent.Schema
}

func (FeedConfigAction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("feed_config_id").GoType(model.InternalID(0)),
		field.Int64("feed_action_set_id").GoType(model.InternalID(0)),
		field.Int64("index"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (FeedConfigAction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("feed_config_id", "feed_action_set_id").
			Unique(),
	}
}

func (FeedConfigAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feed_config", FeedConfig.Type).
			Unique().
			Required().
			Field("feed_config_id"),
		edge.To("feed_action_set", FeedActionSet.Type).
			Unique().
			Required().
			Field("feed_action_set_id"),
	}
}
