// Code generated by ent, DO NOT EDIT.

package notifyflowsource

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/model"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLTE(FieldID, id))
}

// NotifyFlowID applies equality check predicate on the "notify_flow_id" field. It's identical to NotifyFlowIDEQ.
func NotifyFlowID(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldNotifyFlowID, vc))
}

// NotifySourceID applies equality check predicate on the "notify_source_id" field. It's identical to NotifySourceIDEQ.
func NotifySourceID(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldNotifySourceID, vc))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldCreatedAt, v))
}

// NotifyFlowIDEQ applies the EQ predicate on the "notify_flow_id" field.
func NotifyFlowIDEQ(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldNotifyFlowID, vc))
}

// NotifyFlowIDNEQ applies the NEQ predicate on the "notify_flow_id" field.
func NotifyFlowIDNEQ(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldNEQ(FieldNotifyFlowID, vc))
}

// NotifyFlowIDIn applies the In predicate on the "notify_flow_id" field.
func NotifyFlowIDIn(vs ...model.InternalID) predicate.NotifyFlowSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowSource(sql.FieldIn(FieldNotifyFlowID, v...))
}

// NotifyFlowIDNotIn applies the NotIn predicate on the "notify_flow_id" field.
func NotifyFlowIDNotIn(vs ...model.InternalID) predicate.NotifyFlowSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowSource(sql.FieldNotIn(FieldNotifyFlowID, v...))
}

// NotifySourceIDEQ applies the EQ predicate on the "notify_source_id" field.
func NotifySourceIDEQ(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldNotifySourceID, vc))
}

// NotifySourceIDNEQ applies the NEQ predicate on the "notify_source_id" field.
func NotifySourceIDNEQ(v model.InternalID) predicate.NotifyFlowSource {
	vc := int64(v)
	return predicate.NotifyFlowSource(sql.FieldNEQ(FieldNotifySourceID, vc))
}

// NotifySourceIDIn applies the In predicate on the "notify_source_id" field.
func NotifySourceIDIn(vs ...model.InternalID) predicate.NotifyFlowSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowSource(sql.FieldIn(FieldNotifySourceID, v...))
}

// NotifySourceIDNotIn applies the NotIn predicate on the "notify_source_id" field.
func NotifySourceIDNotIn(vs ...model.InternalID) predicate.NotifyFlowSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowSource(sql.FieldNotIn(FieldNotifySourceID, v...))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.FieldLTE(FieldCreatedAt, v))
}

// HasNotifyFlow applies the HasEdge predicate on the "notify_flow" edge.
func HasNotifyFlow() predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NotifyFlowTable, NotifyFlowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyFlowWith applies the HasEdge predicate on the "notify_flow" edge with a given conditions (other predicates).
func HasNotifyFlowWith(preds ...predicate.NotifyFlow) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(func(s *sql.Selector) {
		step := newNotifyFlowStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifySource applies the HasEdge predicate on the "notify_source" edge.
func HasNotifySource() predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NotifySourceTable, NotifySourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifySourceWith applies the HasEdge predicate on the "notify_source" edge with a given conditions (other predicates).
func HasNotifySourceWith(preds ...predicate.FeedConfig) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(func(s *sql.Selector) {
		step := newNotifySourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotifyFlowSource) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotifyFlowSource) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotifyFlowSource) predicate.NotifyFlowSource {
	return predicate.NotifyFlowSource(sql.NotPredicates(p))
}
