package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppRunTime struct {
	ent.Schema
}

func (AppRunTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("app_id").GoType(model.InternalID(0)),
		field.Time("start_time"),
		field.Int64("run_duration").GoType(time.Duration(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppRunTime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_id"),
		index.Fields("start_time", "run_duration").Unique(),
	}
}

func (AppRunTime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).
			Ref("app_run_time").
			Field("app_id").
			Required().
			Unique(),
	}
}
