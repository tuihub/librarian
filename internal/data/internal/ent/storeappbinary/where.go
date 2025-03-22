// Code generated by ent, DO NOT EDIT.

package storeappbinary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// ID filters vertices based on their ID field.
func ID(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id model.InternalID) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldName, v))
}

// SizeBytes applies equality check predicate on the "size_bytes" field. It's identical to SizeBytesEQ.
func SizeBytes(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSizeBytes, v))
}

// PublicURL applies equality check predicate on the "public_url" field. It's identical to PublicURLEQ.
func PublicURL(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldPublicURL, v))
}

// Sha256 applies equality check predicate on the "sha256" field. It's identical to Sha256EQ.
func Sha256(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSha256, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldContainsFold(FieldName, v))
}

// SizeBytesEQ applies the EQ predicate on the "size_bytes" field.
func SizeBytesEQ(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSizeBytes, v))
}

// SizeBytesNEQ applies the NEQ predicate on the "size_bytes" field.
func SizeBytesNEQ(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldSizeBytes, v))
}

// SizeBytesIn applies the In predicate on the "size_bytes" field.
func SizeBytesIn(vs ...int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldSizeBytes, vs...))
}

// SizeBytesNotIn applies the NotIn predicate on the "size_bytes" field.
func SizeBytesNotIn(vs ...int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldSizeBytes, vs...))
}

// SizeBytesGT applies the GT predicate on the "size_bytes" field.
func SizeBytesGT(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldSizeBytes, v))
}

// SizeBytesGTE applies the GTE predicate on the "size_bytes" field.
func SizeBytesGTE(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldSizeBytes, v))
}

// SizeBytesLT applies the LT predicate on the "size_bytes" field.
func SizeBytesLT(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldSizeBytes, v))
}

// SizeBytesLTE applies the LTE predicate on the "size_bytes" field.
func SizeBytesLTE(v int64) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldSizeBytes, v))
}

// SizeBytesIsNil applies the IsNil predicate on the "size_bytes" field.
func SizeBytesIsNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIsNull(FieldSizeBytes))
}

// SizeBytesNotNil applies the NotNil predicate on the "size_bytes" field.
func SizeBytesNotNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotNull(FieldSizeBytes))
}

// PublicURLEQ applies the EQ predicate on the "public_url" field.
func PublicURLEQ(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldPublicURL, v))
}

// PublicURLNEQ applies the NEQ predicate on the "public_url" field.
func PublicURLNEQ(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldPublicURL, v))
}

// PublicURLIn applies the In predicate on the "public_url" field.
func PublicURLIn(vs ...string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldPublicURL, vs...))
}

// PublicURLNotIn applies the NotIn predicate on the "public_url" field.
func PublicURLNotIn(vs ...string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldPublicURL, vs...))
}

// PublicURLGT applies the GT predicate on the "public_url" field.
func PublicURLGT(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldPublicURL, v))
}

// PublicURLGTE applies the GTE predicate on the "public_url" field.
func PublicURLGTE(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldPublicURL, v))
}

// PublicURLLT applies the LT predicate on the "public_url" field.
func PublicURLLT(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldPublicURL, v))
}

// PublicURLLTE applies the LTE predicate on the "public_url" field.
func PublicURLLTE(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldPublicURL, v))
}

// PublicURLContains applies the Contains predicate on the "public_url" field.
func PublicURLContains(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldContains(FieldPublicURL, v))
}

// PublicURLHasPrefix applies the HasPrefix predicate on the "public_url" field.
func PublicURLHasPrefix(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldHasPrefix(FieldPublicURL, v))
}

// PublicURLHasSuffix applies the HasSuffix predicate on the "public_url" field.
func PublicURLHasSuffix(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldHasSuffix(FieldPublicURL, v))
}

// PublicURLIsNil applies the IsNil predicate on the "public_url" field.
func PublicURLIsNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIsNull(FieldPublicURL))
}

// PublicURLNotNil applies the NotNil predicate on the "public_url" field.
func PublicURLNotNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotNull(FieldPublicURL))
}

// PublicURLEqualFold applies the EqualFold predicate on the "public_url" field.
func PublicURLEqualFold(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEqualFold(FieldPublicURL, v))
}

// PublicURLContainsFold applies the ContainsFold predicate on the "public_url" field.
func PublicURLContainsFold(v string) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldContainsFold(FieldPublicURL, v))
}

// Sha256EQ applies the EQ predicate on the "sha256" field.
func Sha256EQ(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldEQ(FieldSha256, v))
}

// Sha256NEQ applies the NEQ predicate on the "sha256" field.
func Sha256NEQ(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNEQ(FieldSha256, v))
}

// Sha256In applies the In predicate on the "sha256" field.
func Sha256In(vs ...[]byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIn(FieldSha256, vs...))
}

// Sha256NotIn applies the NotIn predicate on the "sha256" field.
func Sha256NotIn(vs ...[]byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotIn(FieldSha256, vs...))
}

// Sha256GT applies the GT predicate on the "sha256" field.
func Sha256GT(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGT(FieldSha256, v))
}

// Sha256GTE applies the GTE predicate on the "sha256" field.
func Sha256GTE(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldGTE(FieldSha256, v))
}

// Sha256LT applies the LT predicate on the "sha256" field.
func Sha256LT(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLT(FieldSha256, v))
}

// Sha256LTE applies the LTE predicate on the "sha256" field.
func Sha256LTE(v []byte) predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldLTE(FieldSha256, v))
}

// Sha256IsNil applies the IsNil predicate on the "sha256" field.
func Sha256IsNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldIsNull(FieldSha256))
}

// Sha256NotNil applies the NotNil predicate on the "sha256" field.
func Sha256NotNil() predicate.StoreAppBinary {
	return predicate.StoreAppBinary(sql.FieldNotNull(FieldSha256))
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
