package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppInfo struct {
	ent.Schema
}

func (AppInfo) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("source"),
		field.String("source_app_id"),
		field.String("source_url").Optional(),
		field.String("name"),
		field.Enum("type").
			Values("unknown", "game"),
		field.String("short_description").Optional(),
		field.Text("description").Optional(),
		field.String("icon_image_url").Optional(),
		field.Int64("icon_image_id").GoType(model.InternalID(0)),
		field.String("background_image_url").Optional(),
		field.Int64("background_image_id").GoType(model.InternalID(0)),
		field.String("cover_image_url").Optional(),
		field.Int64("cover_image_id").GoType(model.InternalID(0)),
		field.String("release_date").Optional(),
		field.String("developer").Optional(),
		field.String("publisher").Optional(),
		field.Strings("tags"),
		field.Strings("alternative_names"),
		field.String("raw_data"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source", "source_app_id").
			Unique(),
	}
}

func (AppInfo) Edges() []ent.Edge {
	return []ent.Edge{}
}
