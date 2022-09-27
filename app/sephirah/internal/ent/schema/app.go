package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.Enum("source").
			Values("internal", "steam"),
		field.String("source_app_id"),
		field.String("source_url"),
		field.String("name"),
		field.Enum("type").
			Values("general", "game"),
		field.String("short_description"),
		field.Text("description"),
		field.String("image_url"),
		field.String("release_date"),
		field.String("developer"),
		field.String("publisher"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
