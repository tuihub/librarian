package schema

import (
	"time"

	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
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
	return []ent.Field{
		defaultPrimaryKey(),
		field.Enum("source").
			Values("manual", "sentinel"),
		field.Int64("source_id").
			GoType(model.InternalID(0)),
		field.String("name"),
		field.Text("description"),
		field.Bool("public"),
		field.String("binary_name"),
		field.Int64("binary_size_byte"),
		field.String("binary_public_url"),
		field.Bytes("binary_sha256"),
		field.Time("updated_at").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (AppPackage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("binary_sha256").
			Unique(),
	}
}

// Edges of the AppPackage.
func (AppPackage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("app_package").
			Required().
			Unique(),
		edge.From("app", App.Type).
			Ref("app_package").
			Unique(),
	}
}
