// Code generated by ent, DO NOT EDIT.

package kv

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldID, id))
}

// Bucket applies equality check predicate on the "bucket" field. It's identical to BucketEQ.
func Bucket(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldBucket, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldValue, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldCreatedAt, v))
}

// BucketEQ applies the EQ predicate on the "bucket" field.
func BucketEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldBucket, v))
}

// BucketNEQ applies the NEQ predicate on the "bucket" field.
func BucketNEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldBucket, v))
}

// BucketIn applies the In predicate on the "bucket" field.
func BucketIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldBucket, vs...))
}

// BucketNotIn applies the NotIn predicate on the "bucket" field.
func BucketNotIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldBucket, vs...))
}

// BucketGT applies the GT predicate on the "bucket" field.
func BucketGT(v string) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldBucket, v))
}

// BucketGTE applies the GTE predicate on the "bucket" field.
func BucketGTE(v string) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldBucket, v))
}

// BucketLT applies the LT predicate on the "bucket" field.
func BucketLT(v string) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldBucket, v))
}

// BucketLTE applies the LTE predicate on the "bucket" field.
func BucketLTE(v string) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldBucket, v))
}

// BucketContains applies the Contains predicate on the "bucket" field.
func BucketContains(v string) predicate.KV {
	return predicate.KV(sql.FieldContains(FieldBucket, v))
}

// BucketHasPrefix applies the HasPrefix predicate on the "bucket" field.
func BucketHasPrefix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasPrefix(FieldBucket, v))
}

// BucketHasSuffix applies the HasSuffix predicate on the "bucket" field.
func BucketHasSuffix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasSuffix(FieldBucket, v))
}

// BucketEqualFold applies the EqualFold predicate on the "bucket" field.
func BucketEqualFold(v string) predicate.KV {
	return predicate.KV(sql.FieldEqualFold(FieldBucket, v))
}

// BucketContainsFold applies the ContainsFold predicate on the "bucket" field.
func BucketContainsFold(v string) predicate.KV {
	return predicate.KV(sql.FieldContainsFold(FieldBucket, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.KV {
	return predicate.KV(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.KV {
	return predicate.KV(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.KV {
	return predicate.KV(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.KV {
	return predicate.KV(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.KV {
	return predicate.KV(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.KV {
	return predicate.KV(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.KV {
	return predicate.KV(sql.FieldContainsFold(FieldValue, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.KV {
	return predicate.KV(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.KV {
	return predicate.KV(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.KV {
	return predicate.KV(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KV) predicate.KV {
	return predicate.KV(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KV) predicate.KV {
	return predicate.KV(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.KV) predicate.KV {
	return predicate.KV(sql.NotPredicates(p))
}
