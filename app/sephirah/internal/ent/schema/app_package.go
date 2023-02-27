package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AppPackage holds the schema definition for the AppPackage entity.
type AppPackage struct {
	ent.Schema
}

// Fields of the AppPackage.
func (AppPackage) Fields() []ent.Field {
	incrementalEnabled := false
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable().
			Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
				Incremental: &incrementalEnabled,
			}),
		field.Enum("source").
			Values("manual", "sentinel"),
		field.Int64("source_id"),
		field.String("source_package_id"),
		field.String("name"),
		field.Text("description"),
		field.String("binary_name"),
		field.Int64("binary_size"),
		field.String("binary_public_url"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
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
	return []ent.Edge{
		edge.From("app", App.Type).
			Ref("app_package").
			Unique(),
	}
}
