package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FeedItemCollection struct {
	ent.Schema
}

func (FeedItemCollection) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.String("description"),
		field.String("category"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (FeedItemCollection) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("category"),
	}
}

func (FeedItemCollection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("feed_item_collection").
			Required().
			Unique(),
		edge.From("feed_item", FeedItem.Type).
			Ref("feed_item_collection"),
		edge.To("notify_source", NotifySource.Type),
	}
}
