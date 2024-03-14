package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AppInstRunTime struct {
	ent.Schema
}

func (AppInstRunTime) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").GoType(model.InternalID(0)),
		field.Int64("app_inst_id").GoType(model.InternalID(0)),
		field.Time("start_time"),
		field.Int64("run_duration").GoType(time.Duration(0)),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppInstRunTime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_inst_id"),
	}
}
