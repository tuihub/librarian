package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AppPackage holds the schema definition for the AppPackage entity.
type AppPackage struct {
	ent.Schema
}

// Fields of the AppPackage.
func (AppPackage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("internal_id").
			Unique(),
		field.Enum("source").
			Values("manual", "sentinel"),
		field.Int64("source_id"),
		field.String("source_package_id"),
		field.String("name"),
		field.Text("description"),
		field.String("binary_name"),
		field.String("binary_size"),
		field.Time("updated_at"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppPackage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source", "source_id", "source_package_id").
			Unique(),
	}
}

// Edges of the AppPackage.
func (AppPackage) Edges() []ent.Edge {
	return nil
}
