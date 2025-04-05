// Code generated by ent, DO NOT EDIT.

package appcategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the appcategory type in the database.
	Label = "app_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldVersionNumber holds the string denoting the version_number field in the database.
	FieldVersionNumber = "version_number"
	// FieldVersionDate holds the string denoting the version_date field in the database.
	FieldVersionDate = "version_date"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeAppAppCategory holds the string denoting the app_app_category edge name in mutations.
	EdgeAppAppCategory = "app_app_category"
	// Table holds the table name of the appcategory in the database.
	Table = "app_categories"
	// AppTable is the table that holds the app relation/edge. The primary key declared below.
	AppTable = "app_app_categories"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// AppAppCategoryTable is the table that holds the app_app_category relation/edge.
	AppAppCategoryTable = "app_app_categories"
	// AppAppCategoryInverseTable is the table name for the AppAppCategory entity.
	// It exists in this package in order to avoid circular dependency with the "appappcategory" package.
	AppAppCategoryInverseTable = "app_app_categories"
	// AppAppCategoryColumn is the table column denoting the app_app_category relation/edge.
	AppAppCategoryColumn = "app_category_id"
)

// Columns holds all SQL columns for appcategory fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldVersionNumber,
	FieldVersionDate,
	FieldName,
	FieldUpdatedAt,
	FieldCreatedAt,
}

var (
	// AppPrimaryKey and AppColumn2 are the table columns denoting the
	// primary key for the app relation (M2M).
	AppPrimaryKey = []string{"app_category_id", "app_id"}
)

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

// OrderOption defines the ordering options for the AppCategory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByVersionNumber orders the results by the version_number field.
func ByVersionNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersionNumber, opts...).ToFunc()
}

// ByVersionDate orders the results by the version_date field.
func ByVersionDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersionDate, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAppCount orders the results by app count.
func ByAppCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppStep(), opts...)
	}
}

// ByApp orders the results by app terms.
func ByApp(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAppAppCategoryCount orders the results by app_app_category count.
func ByAppAppCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppAppCategoryStep(), opts...)
	}
}

// ByAppAppCategory orders the results by app_app_category terms.
func ByAppAppCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppAppCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAppStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AppTable, AppPrimaryKey...),
	)
}
func newAppAppCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppAppCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AppAppCategoryTable, AppAppCategoryColumn),
	)
}
