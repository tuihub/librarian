package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Int64("internal_id").
			Unique(),
		field.Enum("platform").
			Values("steam"),
		field.String("platform_account_id"),
		field.String("name"),
		field.String("profile_url"),
		field.String("avatar_url"),
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
	return nil
}
