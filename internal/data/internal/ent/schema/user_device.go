package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserDevice struct {
	ent.Schema
}

func (UserDevice) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("device_id").GoType(model.InternalID(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (UserDevice) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "device_id").
			Unique(),
	}
}

func (UserDevice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("device_info", DeviceInfo.Type).
			Unique().
			Required().
			Field("device_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
