package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AppInst struct {
	ent.Schema
}

func (AppInst) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("device_id").GoType(model.InternalID(0)),
		field.Int64("app_id").GoType(model.InternalID(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppInst) Indexes() []ent.Index {
	return []ent.Index{}
}

func (AppInst) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("app_inst").
			Required().
			Unique(),
	}
}
