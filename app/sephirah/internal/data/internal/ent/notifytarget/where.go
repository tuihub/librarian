// Code generated by ent, DO NOT EDIT.

package notifytarget

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldID, id))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldToken, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldDescription, v))
}

// Destination applies equality check predicate on the "destination" field. It's identical to DestinationEQ.
func Destination(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldDestination, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldCreatedAt, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContainsFold(FieldToken, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContainsFold(FieldDescription, v))
}

// DestinationEQ applies the EQ predicate on the "destination" field.
func DestinationEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldDestination, v))
}

// DestinationNEQ applies the NEQ predicate on the "destination" field.
func DestinationNEQ(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldDestination, v))
}

// DestinationIn applies the In predicate on the "destination" field.
func DestinationIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldDestination, vs...))
}

// DestinationNotIn applies the NotIn predicate on the "destination" field.
func DestinationNotIn(vs ...string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldDestination, vs...))
}

// DestinationGT applies the GT predicate on the "destination" field.
func DestinationGT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldDestination, v))
}

// DestinationGTE applies the GTE predicate on the "destination" field.
func DestinationGTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldDestination, v))
}

// DestinationLT applies the LT predicate on the "destination" field.
func DestinationLT(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldDestination, v))
}

// DestinationLTE applies the LTE predicate on the "destination" field.
func DestinationLTE(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldDestination, v))
}

// DestinationContains applies the Contains predicate on the "destination" field.
func DestinationContains(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContains(FieldDestination, v))
}

// DestinationHasPrefix applies the HasPrefix predicate on the "destination" field.
func DestinationHasPrefix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasPrefix(FieldDestination, v))
}

// DestinationHasSuffix applies the HasSuffix predicate on the "destination" field.
func DestinationHasSuffix(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldHasSuffix(FieldDestination, v))
}

// DestinationEqualFold applies the EqualFold predicate on the "destination" field.
func DestinationEqualFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEqualFold(FieldDestination, v))
}

// DestinationContainsFold applies the ContainsFold predicate on the "destination" field.
func DestinationContainsFold(v string) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldContainsFold(FieldDestination, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldStatus, vs...))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifyFlow applies the HasEdge predicate on the "notify_flow" edge.
func HasNotifyFlow() predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, NotifyFlowTable, NotifyFlowPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyFlowWith applies the HasEdge predicate on the "notify_flow" edge with a given conditions (other predicates).
func HasNotifyFlowWith(preds ...predicate.NotifyFlow) predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := newNotifyFlowStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotifyFlowTarget applies the HasEdge predicate on the "notify_flow_target" edge.
func HasNotifyFlowTarget() predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, NotifyFlowTargetTable, NotifyFlowTargetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotifyFlowTargetWith applies the HasEdge predicate on the "notify_flow_target" edge with a given conditions (other predicates).
func HasNotifyFlowTargetWith(preds ...predicate.NotifyFlowTarget) predicate.NotifyTarget {
	return predicate.NotifyTarget(func(s *sql.Selector) {
		step := newNotifyFlowTargetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotifyTarget) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotifyTarget) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotifyTarget) predicate.NotifyTarget {
	return predicate.NotifyTarget(sql.NotPredicates(p))
}
