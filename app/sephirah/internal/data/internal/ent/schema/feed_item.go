package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"
	"github.com/tuihub/librarian/internal/model/modelfeed"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// FeedItem holds the schema definition for the FeedItem entity.
type FeedItem struct {
	ent.Schema
}

// Fields of the FeedItem.
func (FeedItem) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("feed_id").
			Immutable().
			GoType(model.InternalID(0)),
		field.String("title").Optional(),
		field.JSON("authors", []*modelfeed.Person{}).Optional(),
		field.String("description").Optional(),
		field.String("content").Optional(),
		field.String("guid").Immutable(),
		field.String("link").Optional(),
		field.JSON("image", new(modelfeed.Image)).Optional(),
		field.String("published").Optional(),
		field.Time("published_parsed"),
		field.String("updated").Optional(),
		field.Time("updated_parsed").Optional().Nillable(),
		field.JSON("enclosures", []*modelfeed.Enclosure{}).Optional(),
		field.String("publish_platform").Optional(),
		field.Int64("read_count").Default(0),
		field.String("digest_description").Optional(),
		field.JSON("digest_images", []*modelfeed.Image{}).Optional(),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (FeedItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("feed_id", "guid").
			Unique(),
		index.Fields("publish_platform"),
	}
}

// Edges of the FeedItem.
func (FeedItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feed", Feed.Type).
			Ref("item").
			Field("feed_id").
			Unique().
			Immutable().
			Required(),
	}
}
