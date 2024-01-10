package schema

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PorterPrivilege struct {
	ent.Schema
}

func (PorterPrivilege) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("porter_id").GoType(model.InternalID(0)),
		field.JSON("privilege", new(modeltiphereth.PorterInstancePrivilege)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (PorterPrivilege) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "porter_id").
			Unique(),
	}
}
