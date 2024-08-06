// Code generated by ent, DO NOT EDIT.

package portercontext

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the portercontext type in the database.
	Label = "porter_context"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGlobalName holds the string denoting the global_name field in the database.
	FieldGlobalName = "global_name"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldContextJSON holds the string denoting the context_json field in the database.
	FieldContextJSON = "context_json"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the portercontext in the database.
	Table = "porter_contexts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "porter_contexts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_porter_context"
)

// Columns holds all SQL columns for portercontext fields.
var Columns = []string{
	FieldID,
	FieldGlobalName,
	FieldRegion,
	FieldContextJSON,
	FieldName,
	FieldDescription,
	FieldStatus,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "porter_contexts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_porter_context",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusActive   Status = "active"
	StatusDisabled Status = "disabled"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusActive, StatusDisabled:
		return nil
	default:
		return fmt.Errorf("portercontext: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the PorterContext queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGlobalName orders the results by the global_name field.
func ByGlobalName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGlobalName, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByContextJSON orders the results by the context_json field.
func ByContextJSON(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContextJSON, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
