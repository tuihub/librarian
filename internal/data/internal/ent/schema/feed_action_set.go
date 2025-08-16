package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type FeedActionSet struct {
	ent.Schema
}

func (FeedActionSet) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.JSON("actions", []*model.FeatureRequest{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (FeedActionSet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("feed_action_set").
			Required().
			Unique(),
		edge.From("feed_config", FeedConfig.Type).
			Ref("feed_action_set"),
	}
}
