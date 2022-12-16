// Code generated by ent, DO NOT EDIT.

package apppackage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// InternalID applies equality check predicate on the "internal_id" field. It's identical to InternalIDEQ.
func InternalID(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInternalID), v))
	})
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceID), v))
	})
}

// SourcePackageID applies equality check predicate on the "source_package_id" field. It's identical to SourcePackageIDEQ.
func SourcePackageID(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourcePackageID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// BinaryName applies equality check predicate on the "binary_name" field. It's identical to BinaryNameEQ.
func BinaryName(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBinaryName), v))
	})
}

// BinarySize applies equality check predicate on the "binary_size" field. It's identical to BinarySizeEQ.
func BinarySize(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBinarySize), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// InternalIDEQ applies the EQ predicate on the "internal_id" field.
func InternalIDEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInternalID), v))
	})
}

// InternalIDNEQ applies the NEQ predicate on the "internal_id" field.
func InternalIDNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInternalID), v))
	})
}

// InternalIDIn applies the In predicate on the "internal_id" field.
func InternalIDIn(vs ...int64) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInternalID), v...))
	})
}

// InternalIDNotIn applies the NotIn predicate on the "internal_id" field.
func InternalIDNotIn(vs ...int64) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInternalID), v...))
	})
}

// InternalIDGT applies the GT predicate on the "internal_id" field.
func InternalIDGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInternalID), v))
	})
}

// InternalIDGTE applies the GTE predicate on the "internal_id" field.
func InternalIDGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInternalID), v))
	})
}

// InternalIDLT applies the LT predicate on the "internal_id" field.
func InternalIDLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInternalID), v))
	})
}

// InternalIDLTE applies the LTE predicate on the "internal_id" field.
func InternalIDLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInternalID), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v Source) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v Source) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...Source) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...Source) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceID), v))
	})
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSourceID), v))
	})
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...int64) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSourceID), v...))
	})
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...int64) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSourceID), v...))
	})
}

// SourceIDGT applies the GT predicate on the "source_id" field.
func SourceIDGT(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSourceID), v))
	})
}

// SourceIDGTE applies the GTE predicate on the "source_id" field.
func SourceIDGTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSourceID), v))
	})
}

// SourceIDLT applies the LT predicate on the "source_id" field.
func SourceIDLT(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSourceID), v))
	})
}

// SourceIDLTE applies the LTE predicate on the "source_id" field.
func SourceIDLTE(v int64) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSourceID), v))
	})
}

// SourcePackageIDEQ applies the EQ predicate on the "source_package_id" field.
func SourcePackageIDEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDNEQ applies the NEQ predicate on the "source_package_id" field.
func SourcePackageIDNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDIn applies the In predicate on the "source_package_id" field.
func SourcePackageIDIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSourcePackageID), v...))
	})
}

// SourcePackageIDNotIn applies the NotIn predicate on the "source_package_id" field.
func SourcePackageIDNotIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSourcePackageID), v...))
	})
}

// SourcePackageIDGT applies the GT predicate on the "source_package_id" field.
func SourcePackageIDGT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDGTE applies the GTE predicate on the "source_package_id" field.
func SourcePackageIDGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDLT applies the LT predicate on the "source_package_id" field.
func SourcePackageIDLT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDLTE applies the LTE predicate on the "source_package_id" field.
func SourcePackageIDLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDContains applies the Contains predicate on the "source_package_id" field.
func SourcePackageIDContains(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDHasPrefix applies the HasPrefix predicate on the "source_package_id" field.
func SourcePackageIDHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDHasSuffix applies the HasSuffix predicate on the "source_package_id" field.
func SourcePackageIDHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDEqualFold applies the EqualFold predicate on the "source_package_id" field.
func SourcePackageIDEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSourcePackageID), v))
	})
}

// SourcePackageIDContainsFold applies the ContainsFold predicate on the "source_package_id" field.
func SourcePackageIDContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSourcePackageID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// BinaryNameEQ applies the EQ predicate on the "binary_name" field.
func BinaryNameEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBinaryName), v))
	})
}

// BinaryNameNEQ applies the NEQ predicate on the "binary_name" field.
func BinaryNameNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBinaryName), v))
	})
}

// BinaryNameIn applies the In predicate on the "binary_name" field.
func BinaryNameIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBinaryName), v...))
	})
}

// BinaryNameNotIn applies the NotIn predicate on the "binary_name" field.
func BinaryNameNotIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBinaryName), v...))
	})
}

// BinaryNameGT applies the GT predicate on the "binary_name" field.
func BinaryNameGT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBinaryName), v))
	})
}

// BinaryNameGTE applies the GTE predicate on the "binary_name" field.
func BinaryNameGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBinaryName), v))
	})
}

// BinaryNameLT applies the LT predicate on the "binary_name" field.
func BinaryNameLT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBinaryName), v))
	})
}

// BinaryNameLTE applies the LTE predicate on the "binary_name" field.
func BinaryNameLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBinaryName), v))
	})
}

// BinaryNameContains applies the Contains predicate on the "binary_name" field.
func BinaryNameContains(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBinaryName), v))
	})
}

// BinaryNameHasPrefix applies the HasPrefix predicate on the "binary_name" field.
func BinaryNameHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBinaryName), v))
	})
}

// BinaryNameHasSuffix applies the HasSuffix predicate on the "binary_name" field.
func BinaryNameHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBinaryName), v))
	})
}

// BinaryNameEqualFold applies the EqualFold predicate on the "binary_name" field.
func BinaryNameEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBinaryName), v))
	})
}

// BinaryNameContainsFold applies the ContainsFold predicate on the "binary_name" field.
func BinaryNameContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBinaryName), v))
	})
}

// BinarySizeEQ applies the EQ predicate on the "binary_size" field.
func BinarySizeEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBinarySize), v))
	})
}

// BinarySizeNEQ applies the NEQ predicate on the "binary_size" field.
func BinarySizeNEQ(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBinarySize), v))
	})
}

// BinarySizeIn applies the In predicate on the "binary_size" field.
func BinarySizeIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBinarySize), v...))
	})
}

// BinarySizeNotIn applies the NotIn predicate on the "binary_size" field.
func BinarySizeNotIn(vs ...string) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBinarySize), v...))
	})
}

// BinarySizeGT applies the GT predicate on the "binary_size" field.
func BinarySizeGT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBinarySize), v))
	})
}

// BinarySizeGTE applies the GTE predicate on the "binary_size" field.
func BinarySizeGTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBinarySize), v))
	})
}

// BinarySizeLT applies the LT predicate on the "binary_size" field.
func BinarySizeLT(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBinarySize), v))
	})
}

// BinarySizeLTE applies the LTE predicate on the "binary_size" field.
func BinarySizeLTE(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBinarySize), v))
	})
}

// BinarySizeContains applies the Contains predicate on the "binary_size" field.
func BinarySizeContains(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBinarySize), v))
	})
}

// BinarySizeHasPrefix applies the HasPrefix predicate on the "binary_size" field.
func BinarySizeHasPrefix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBinarySize), v))
	})
}

// BinarySizeHasSuffix applies the HasSuffix predicate on the "binary_size" field.
func BinarySizeHasSuffix(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBinarySize), v))
	})
}

// BinarySizeEqualFold applies the EqualFold predicate on the "binary_size" field.
func BinarySizeEqualFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBinarySize), v))
	})
}

// BinarySizeContainsFold applies the ContainsFold predicate on the "binary_size" field.
func BinarySizeContainsFold(v string) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBinarySize), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppPackage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppPackage) predicate.AppPackage {
	return predicate.AppPackage(func(s *sql.Selector) {
		p(s.Not())
	})
}
