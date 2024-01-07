package schema

import (
	"time"

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
		edge.To("purchased_app", App.Type),
		edge.From("bind_user", User.Type).
			Ref("bind_account").
			Unique(),
	}
}
