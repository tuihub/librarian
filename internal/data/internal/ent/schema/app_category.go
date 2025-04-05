package schema

import (
	"entgo.io/ent/schema/edge"
	"github.com/tuihub/librarian/internal/model"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppCategory struct {
	ent.Schema
}

func (AppCategory) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Uint64("version_number"),
		field.Time("version_date"),
		field.String("name"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

func (AppCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).
			Through("app_app_category", AppAppCategory.Type),
	}
}
