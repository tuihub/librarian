// Code generated by ent, DO NOT EDIT.

package feeditemcollection

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the feeditemcollection type in the database.
	Label = "feed_item_collection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeFeedItem holds the string denoting the feed_item edge name in mutations.
	EdgeFeedItem = "feed_item"
	// EdgeNotifySource holds the string denoting the notify_source edge name in mutations.
	EdgeNotifySource = "notify_source"
	// Table holds the table name of the feeditemcollection in the database.
	Table = "feed_item_collections"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "feed_item_collections"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_feed_item_collection"
	// FeedItemTable is the table that holds the feed_item relation/edge. The primary key declared below.
	FeedItemTable = "feed_item_feed_item_collection"
	// FeedItemInverseTable is the table name for the FeedItem entity.
	// It exists in this package in order to avoid circular dependency with the "feeditem" package.
	FeedItemInverseTable = "feed_items"
	// NotifySourceTable is the table that holds the notify_source relation/edge.
	NotifySourceTable = "notify_sources"
	// NotifySourceInverseTable is the table name for the NotifySource entity.
	// It exists in this package in order to avoid circular dependency with the "notifysource" package.
	NotifySourceInverseTable = "notify_sources"
	// NotifySourceColumn is the table column denoting the notify_source relation/edge.
	NotifySourceColumn = "feed_item_collection_id"
)

// Columns holds all SQL columns for feeditemcollection fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCategory,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "feed_item_collections"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_feed_item_collection",
}

var (
	// FeedItemPrimaryKey and FeedItemColumn2 are the table columns denoting the
	// primary key for the feed_item relation (M2M).
	FeedItemPrimaryKey = []string{"feed_item_id", "feed_item_collection_id"}
)

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

// OrderOption defines the ordering options for the FeedItemCollection queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCategory orders the results by the category field.
func ByCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategory, opts...).ToFunc()
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

// ByFeedItemCount orders the results by feed_item count.
func ByFeedItemCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFeedItemStep(), opts...)
	}
}

// ByFeedItem orders the results by feed_item terms.
func ByFeedItem(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeedItemStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotifySourceCount orders the results by notify_source count.
func ByNotifySourceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotifySourceStep(), opts...)
	}
}

// ByNotifySource orders the results by notify_source terms.
func ByNotifySource(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifySourceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newFeedItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeedItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, FeedItemTable, FeedItemPrimaryKey...),
	)
}
func newNotifySourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifySourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotifySourceTable, NotifySourceColumn),
	)
}
