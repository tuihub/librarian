package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type App struct {
	ent.Schema
}

func (App) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Uint64("version_number"),
		field.Time("version_date"),
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("creator_device_id").GoType(model.InternalID(0)).Immutable(),
		field.JSON("app_sources", map[string]string{}),
		field.Bool("public"),
		field.Int64("bound_store_app_id").GoType(model.InternalID(0)).
			Optional().Immutable(),
		field.Bool("stop_store_manage").Optional(),
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
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("app").
			Field("user_id").
			Required().
			Unique(),
		edge.From("device", Device.Type).
			Ref("app").
			Field("creator_device_id").
			Required().
			Unique().
			Immutable(),
		edge.To("app_run_time", AppRunTime.Type),
	}
}
