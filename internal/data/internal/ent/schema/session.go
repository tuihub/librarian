package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("device_id").GoType(model.InternalID(0)).Optional(),
		field.String("refresh_token"),
		field.Time("expire_at"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "device_id").
			Unique(),
		index.Fields("refresh_token").
			Unique(),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("session").
			Field("user_id").
			Unique().
			Required(),
		edge.From("device", Device.Type).
			Ref("session").
			Field("device_id").
			Unique(),
	}
}
