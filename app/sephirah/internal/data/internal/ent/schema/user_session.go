package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserSession struct {
	ent.Schema
}

func (UserSession) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.String("refresh_token"),
		field.Time("expire_at"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (UserSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("refresh_token").
			Unique(),
	}
}

func (UserSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("device_info", DeviceInfo.Type).
			Ref("user_session").
			Unique(),
	}
}
