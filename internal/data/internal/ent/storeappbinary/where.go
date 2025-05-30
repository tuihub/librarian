// Code generated by ent, DO NOT EDIT.

package storeappbinary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldID, id))
}

// StoreAppID applies equality check predicate on the "store_app_id" field. It's identical to StoreAppIDEQ.
func StoreAppID(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldEQ(FieldStoreAppID, vc))
}

// SentinelAppBinaryUnionID applies equality check predicate on the "sentinel_app_binary_union_id" field. It's identical to SentinelAppBinaryUnionIDEQ.
func SentinelAppBinaryUnionID(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSentinelAppBinaryUnionID, vc))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldCreatedAt, v))
}

// StoreAppIDEQ applies the EQ predicate on the "store_app_id" field.
func StoreAppIDEQ(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldEQ(FieldStoreAppID, vc))
}

// StoreAppIDNEQ applies the NEQ predicate on the "store_app_id" field.
func StoreAppIDNEQ(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldStoreAppID, vc))
}

// StoreAppIDIn applies the In predicate on the "store_app_id" field.
func StoreAppIDIn(vs ...model.InternalID) predicate.StoreAppBinary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.StoreAppBinary(sql.FieldIn(FieldStoreAppID, v...))
}

// StoreAppIDNotIn applies the NotIn predicate on the "store_app_id" field.
func StoreAppIDNotIn(vs ...model.InternalID) predicate.StoreAppBinary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldStoreAppID, v...))
}

// SentinelAppBinaryUnionIDEQ applies the EQ predicate on the "sentinel_app_binary_union_id" field.
func SentinelAppBinaryUnionIDEQ(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSentinelAppBinaryUnionID, vc))
}

// SentinelAppBinaryUnionIDNEQ applies the NEQ predicate on the "sentinel_app_binary_union_id" field.
func SentinelAppBinaryUnionIDNEQ(v model.InternalID) predicate.StoreAppBinary {
	vc := int64(v)
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldSentinelAppBinaryUnionID, vc))
}

// SentinelAppBinaryUnionIDIn applies the In predicate on the "sentinel_app_binary_union_id" field.
func SentinelAppBinaryUnionIDIn(vs ...model.InternalID) predicate.StoreAppBinary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.StoreAppBinary(sql.FieldIn(FieldSentinelAppBinaryUnionID, v...))
}

// SentinelAppBinaryUnionIDNotIn applies the NotIn predicate on the "sentinel_app_binary_union_id" field.
func SentinelAppBinaryUnionIDNotIn(vs ...model.InternalID) predicate.StoreAppBinary {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldSentinelAppBinaryUnionID, v...))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldCreatedAt, v))
}

// HasStoreApp applies the HasEdge predicate on the "store_app" edge.
func HasStoreApp() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StoreAppTable, StoreAppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStoreAppWith applies the HasEdge predicate on the "store_app" edge with a given conditions (other predicates).
func HasStoreAppWith(preds ...predicate.StoreApp) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(func(s *sql.Selector) {
		step := newStoreAppStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSentinelAppBinary applies the HasEdge predicate on the "sentinel_app_binary" edge.
func HasSentinelAppBinary() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SentinelAppBinaryTable, SentinelAppBinaryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSentinelAppBinaryWith applies the HasEdge predicate on the "sentinel_app_binary" edge with a given conditions (other predicates).
func HasSentinelAppBinaryWith(preds ...predicate.SentinelAppBinary) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(func(s *sql.Selector) {
		step := newSentinelAppBinaryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StoreAppBinary) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StoreAppBinary) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.StoreAppBinary) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.NotPredicates(p))
}
