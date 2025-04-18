// Code generated by ent, DO NOT EDIT.

package appcategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldEQ(FieldUserID, vc))
}

// VersionNumber applies equality check predicate on the "version_number" field. It's identical to VersionNumberEQ.
func VersionNumber(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldVersionNumber, v))
}

// VersionDate applies equality check predicate on the "version_date" field. It's identical to VersionDateEQ.
func VersionDate(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldVersionDate, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldName, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldEQ(FieldUserID, vc))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldNEQ(FieldUserID, vc))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...model.InternalID) predicate.AppCategory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppCategory(sql.FieldIn(FieldUserID, v...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...model.InternalID) predicate.AppCategory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.AppCategory(sql.FieldNotIn(FieldUserID, v...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldGT(FieldUserID, vc))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldGTE(FieldUserID, vc))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldLT(FieldUserID, vc))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v model.InternalID) predicate.AppCategory {
	vc := int64(v)
	return predicate.AppCategory(sql.FieldLTE(FieldUserID, vc))
}

// VersionNumberEQ applies the EQ predicate on the "version_number" field.
func VersionNumberEQ(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldVersionNumber, v))
}

// VersionNumberNEQ applies the NEQ predicate on the "version_number" field.
func VersionNumberNEQ(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldVersionNumber, v))
}

// VersionNumberIn applies the In predicate on the "version_number" field.
func VersionNumberIn(vs ...uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldVersionNumber, vs...))
}

// VersionNumberNotIn applies the NotIn predicate on the "version_number" field.
func VersionNumberNotIn(vs ...uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldVersionNumber, vs...))
}

// VersionNumberGT applies the GT predicate on the "version_number" field.
func VersionNumberGT(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldVersionNumber, v))
}

// VersionNumberGTE applies the GTE predicate on the "version_number" field.
func VersionNumberGTE(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldVersionNumber, v))
}

// VersionNumberLT applies the LT predicate on the "version_number" field.
func VersionNumberLT(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldVersionNumber, v))
}

// VersionNumberLTE applies the LTE predicate on the "version_number" field.
func VersionNumberLTE(v uint64) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldVersionNumber, v))
}

// VersionDateEQ applies the EQ predicate on the "version_date" field.
func VersionDateEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldVersionDate, v))
}

// VersionDateNEQ applies the NEQ predicate on the "version_date" field.
func VersionDateNEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldVersionDate, v))
}

// VersionDateIn applies the In predicate on the "version_date" field.
func VersionDateIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldVersionDate, vs...))
}

// VersionDateNotIn applies the NotIn predicate on the "version_date" field.
func VersionDateNotIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldVersionDate, vs...))
}

// VersionDateGT applies the GT predicate on the "version_date" field.
func VersionDateGT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldVersionDate, v))
}

// VersionDateGTE applies the GTE predicate on the "version_date" field.
func VersionDateGTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldVersionDate, v))
}

// VersionDateLT applies the LT predicate on the "version_date" field.
func VersionDateLT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldVersionDate, v))
}

// VersionDateLTE applies the LTE predicate on the "version_date" field.
func VersionDateLTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldVersionDate, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldContainsFold(FieldName, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppCategory {
	return predicate.AppCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.AppCategory {
	return predicate.AppCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AppTable, AppPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.AppCategory {
	return predicate.AppCategory(func(s *sql.Selector) {
		step := newAppStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAppAppCategory applies the HasEdge predicate on the "app_app_category" edge.
func HasAppAppCategory() predicate.AppCategory {
	return predicate.AppCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AppAppCategoryTable, AppAppCategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppAppCategoryWith applies the HasEdge predicate on the "app_app_category" edge with a given conditions (other predicates).
func HasAppAppCategoryWith(preds ...predicate.AppAppCategory) predicate.AppCategory {
	return predicate.AppCategory(func(s *sql.Selector) {
		step := newAppAppCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppCategory) predicate.AppCategory {
	return predicate.AppCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppCategory) predicate.AppCategory {
	return predicate.AppCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppCategory) predicate.AppCategory {
	return predicate.AppCategory(sql.NotPredicates(p))
}
