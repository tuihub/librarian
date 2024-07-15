package schema

import (
	"time"

	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PorterContext struct {
	ent.Schema
}

func (PorterContext) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("porter_id").GoType(model.InternalID(0)),
		field.JSON("context", new(modeltiphereth.PorterInstanceContext)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (PorterContext) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "porter_id").
			Unique(),
	}
}
