package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SentinelSession struct {
	ent.Schema
}

func (SentinelSession) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.Int64("sentinel_id").GoType(model.InternalID(0)),
		field.String("refresh_token"),
		field.Time("expire_at"),
		field.Enum("status").Values("active", "suspend"),
		field.Int64("creator_id").Immutable().GoType(model.InternalID(0)),
		field.Time("last_used_at").Optional(),
		field.Time("last_refreshed_at").Optional(),
		field.Int64("refresh_count").Default(0),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (SentinelSession) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sentinel_id").
			Unique(),
	}
}

func (SentinelSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sentinel", Sentinel.Type).
			Ref("sentinel_session").
			Field("sentinel_id").
			Unique().
			Required(),
	}
}
