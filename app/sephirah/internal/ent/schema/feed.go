package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.String("title"),
		field.String("link"),
		field.String("description"),
		field.String("language"),
		field.JSON("authors", []Person{}),
		field.JSON("images", []Image{}),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

type Person struct {
	Name  string
	EMail string
}

type Image struct {
	URL   string
	Title string
}

// Edges of the Feed.
func (Feed) Edges() []ent.Edge {
	return nil
}
