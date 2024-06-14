package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type SystemNotification struct {
	ent.Schema
}

func (SystemNotification) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("user_id").GoType(model.InternalID(0)).Optional(),
		field.Enum("type").
			Values("system", "user"),
		field.Enum("level").
			Values("info", "warn", "error", "ongoing"),
		field.Enum("status").
			Values("unread", "read", "dismissed"),
		field.String("title"),
		field.String("content"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}
