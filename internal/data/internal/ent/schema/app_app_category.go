package schema

import (
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppAppCategory struct {
	ent.Schema
}

func (AppAppCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("app_category_id").GoType(model.InternalID(0)),
		field.Int64("app_id").GoType(model.InternalID(0)),
	}
}

func (AppAppCategory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_category_id", "app_id").
			Unique(),
	}
}

func (AppAppCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app_category", AppCategory.Type).
			Unique().
			Required().
			Field("app_category_id"),
		edge.To("app", App.Type).
			Unique().
			Required().
			Field("app_id"),
	}
}
