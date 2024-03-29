// Code generated by ent, DO NOT EDIT.

package appbinary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the appbinary type in the database.
	Label = "app_binary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSizeBytes holds the string denoting the size_bytes field in the database.
	FieldSizeBytes = "size_bytes"
	// FieldPublicURL holds the string denoting the public_url field in the database.
	FieldPublicURL = "public_url"
	// FieldSha256 holds the string denoting the sha256 field in the database.
	FieldSha256 = "sha256"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAppInfo holds the string denoting the app_info edge name in mutations.
	EdgeAppInfo = "app_info"
	// Table holds the table name of the appbinary in the database.
	Table = "app_binaries"
	// AppInfoTable is the table that holds the app_info relation/edge.
	AppInfoTable = "app_binaries"
	// AppInfoInverseTable is the table name for the AppInfo entity.
	// It exists in this package in order to avoid circular dependency with the "appinfo" package.
	AppInfoInverseTable = "app_infos"
	// AppInfoColumn is the table column denoting the app_info relation/edge.
	AppInfoColumn = "app_info_app_binary"
)

// Columns holds all SQL columns for appbinary fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSizeBytes,
	FieldPublicURL,
	FieldSha256,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "app_binaries"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"app_info_app_binary",
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

// OrderOption defines the ordering options for the AppBinary queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySizeBytes orders the results by the size_bytes field.
func BySizeBytes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSizeBytes, opts...).ToFunc()
}

// ByPublicURL orders the results by the public_url field.
func ByPublicURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicURL, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAppInfoField orders the results by app_info field.
func ByAppInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newAppInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AppInfoTable, AppInfoColumn),
	)
}
