// Code generated by ent, DO NOT EDIT.

package notifyflowtarget

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLTE(FieldID, id))
}

// NotifyFlowID applies equality check predicate on the "notify_flow_id" field. It's identical to NotifyFlowIDEQ.
func NotifyFlowID(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldNotifyFlowID, vc))
}

// NotifyTargetID applies equality check predicate on the "notify_target_id" field. It's identical to NotifyTargetIDEQ.
func NotifyTargetID(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldNotifyTargetID, vc))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldCreatedAt, v))
}

// NotifyFlowIDEQ applies the EQ predicate on the "notify_flow_id" field.
func NotifyFlowIDEQ(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldNotifyFlowID, vc))
}

// NotifyFlowIDNEQ applies the NEQ predicate on the "notify_flow_id" field.
func NotifyFlowIDNEQ(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldNEQ(FieldNotifyFlowID, vc))
}

// NotifyFlowIDIn applies the In predicate on the "notify_flow_id" field.
func NotifyFlowIDIn(vs ...model.InternalID) predicate.NotifyFlowTarget {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowTarget(sql.FieldIn(FieldNotifyFlowID, v...))
}

// NotifyFlowIDNotIn applies the NotIn predicate on the "notify_flow_id" field.
func NotifyFlowIDNotIn(vs ...model.InternalID) predicate.NotifyFlowTarget {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowTarget(sql.FieldNotIn(FieldNotifyFlowID, v...))
}

// NotifyTargetIDEQ applies the EQ predicate on the "notify_target_id" field.
func NotifyTargetIDEQ(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldNotifyTargetID, vc))
}

// NotifyTargetIDNEQ applies the NEQ predicate on the "notify_target_id" field.
func NotifyTargetIDNEQ(v model.InternalID) predicate.NotifyFlowTarget {
	vc := int64(v)
	return predicate.NotifyFlowTarget(sql.FieldNEQ(FieldNotifyTargetID, vc))
}

// NotifyTargetIDIn applies the In predicate on the "notify_target_id" field.
func NotifyTargetIDIn(vs ...model.InternalID) predicate.NotifyFlowTarget {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowTarget(sql.FieldIn(FieldNotifyTargetID, v...))
}

// NotifyTargetIDNotIn applies the NotIn predicate on the "notify_target_id" field.
func NotifyTargetIDNotIn(vs ...model.InternalID) predicate.NotifyFlowTarget {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int64(vs[i])
	}
	return predicate.NotifyFlowTarget(sql.FieldNotIn(FieldNotifyTargetID, v...))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.FieldLTE(FieldCreatedAt, v))
}

// HasNotifyFlow applies the HasEdge predicate on the "notify_flow" edge.
func HasNotifyFlow() predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NotifyFlowTable, NotifyFlowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyFlowWith applies the HasEdge predicate on the "notify_flow" edge with a given conditions (other predicates).
func HasNotifyFlowWith(preds ...predicate.NotifyFlow) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(func(s *sql.Selector) {
		step := newNotifyFlowStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifyTarget applies the HasEdge predicate on the "notify_target" edge.
func HasNotifyTarget() predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NotifyTargetTable, NotifyTargetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyTargetWith applies the HasEdge predicate on the "notify_target" edge with a given conditions (other predicates).
func HasNotifyTargetWith(preds ...predicate.NotifyTarget) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(func(s *sql.Selector) {
		step := newNotifyTargetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotifyFlowTarget) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotifyFlowTarget) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotifyFlowTarget) predicate.NotifyFlowTarget {
	return predicate.NotifyFlowTarget(sql.NotPredicates(p))
}
