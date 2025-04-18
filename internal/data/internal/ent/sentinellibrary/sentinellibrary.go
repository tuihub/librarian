// Code generated by ent, DO NOT EDIT.

package sentinellibrary

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sentinellibrary type in the database.
	Label = "sentinel_library"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSentinelInfoID holds the string denoting the sentinel_info_id field in the database.
	FieldSentinelInfoID = "sentinel_info_id"
	// FieldReportedID holds the string denoting the reported_id field in the database.
	FieldReportedID = "reported_id"
	// FieldDownloadBasePath holds the string denoting the download_base_path field in the database.
	FieldDownloadBasePath = "download_base_path"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLibraryReportSequence holds the string denoting the library_report_sequence field in the database.
	FieldLibraryReportSequence = "library_report_sequence"
	// EdgeSentinelInfo holds the string denoting the sentinel_info edge name in mutations.
	EdgeSentinelInfo = "sentinel_info"
	// Table holds the table name of the sentinellibrary in the database.
	Table = "sentinel_libraries"
	// SentinelInfoTable is the table that holds the sentinel_info relation/edge.
	SentinelInfoTable = "sentinel_libraries"
	// SentinelInfoInverseTable is the table name for the SentinelInfo entity.
	// It exists in this package in order to avoid circular dependency with the "sentinelinfo" package.
	SentinelInfoInverseTable = "sentinel_infos"
	// SentinelInfoColumn is the table column denoting the sentinel_info relation/edge.
	SentinelInfoColumn = "sentinel_info_id"
)

// Columns holds all SQL columns for sentinellibrary fields.
var Columns = []string{
	FieldID,
	FieldSentinelInfoID,
	FieldReportedID,
	FieldDownloadBasePath,
	FieldUpdatedAt,
	FieldCreatedAt,
	FieldLibraryReportSequence,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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

// OrderOption defines the ordering options for the SentinelLibrary queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySentinelInfoID orders the results by the sentinel_info_id field.
func BySentinelInfoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSentinelInfoID, opts...).ToFunc()
}

// ByReportedID orders the results by the reported_id field.
func ByReportedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReportedID, opts...).ToFunc()
}

// ByDownloadBasePath orders the results by the download_base_path field.
func ByDownloadBasePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadBasePath, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLibraryReportSequence orders the results by the library_report_sequence field.
func ByLibraryReportSequence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLibraryReportSequence, opts...).ToFunc()
}

// BySentinelInfoField orders the results by sentinel_info field.
func BySentinelInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSentinelInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newSentinelInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SentinelInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SentinelInfoTable, SentinelInfoColumn),
	)
}
