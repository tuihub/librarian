package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		defaultPrimaryKey(),
		field.String("platform"),
		field.String("platform_account_id"),
		field.Int64("bound_user_id").GoType(model.InternalID(0)).Optional(),
		field.String("name"),
		field.String("profile_url"),
		field.String("avatar_url"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("platform", "platform_account_id").
			Unique(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bound_user", User.Type).
			Ref("account").
			Field("bound_user_id").
			Unique(),
	}
}
