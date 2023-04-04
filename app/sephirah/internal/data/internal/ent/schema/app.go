package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Enum("source").
			Values("internal", "steam"),
		field.String("source_app_id"),
		field.String("source_url"),
		field.String("name"),
		field.Enum("type").
			Values("game"),
		field.String("short_description"),
		field.Text("description"),
		field.String("image_url"),
		field.String("release_date"),
		field.String("developer"),
		field.String("publisher"),
		field.String("version"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source", "source_app_id").
			Unique(),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("purchased_by", User.Type).
			Ref("purchased_app"),
		edge.To("app_package", AppPackage.Type),
		edge.To("bind_external", App.Type).
			From("bind_internal").
			Unique().
			Required(),
	}
}
