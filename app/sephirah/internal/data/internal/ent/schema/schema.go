package schema

import (
	"github.com/tuihub/librarian/internal/model"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

func defaultPrimaryKey() ent.Field {
	incrementalEnabled := false
	return field.Int64("id").
		Unique().
		Immutable().
		GoType(model.InternalID(0)).
		Annotations(entsql.Annotation{ //nolint:exhaustruct // no need
			Incremental: &incrementalEnabled,
		})
}
