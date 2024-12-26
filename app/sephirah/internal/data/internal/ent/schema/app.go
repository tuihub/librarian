package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type App struct {
	ent.Schema
}

func (App) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("name"),
		field.Text("description"),
		field.Int64("device_id").GoType(model.InternalID(0)),
		field.Bool("public"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (App) Indexes() []ent.Index {
	return []ent.Index{}
}

func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("app").
			Required().
			Unique(),
		edge.From("app_info", AppInfo.Type).
			Ref("app").
			Unique(),
	}
}
