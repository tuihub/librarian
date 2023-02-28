// Code generated by ent, DO NOT EDIT.

package feeditem

import (
	"time"
)

const (
	// Label holds the string label denoting the feeditem type in the database.
	Label = "feed_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthors holds the string denoting the authors field in the database.
	FieldAuthors = "authors"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldGUID holds the string denoting the guid field in the database.
	FieldGUID = "guid"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldPublished holds the string denoting the published field in the database.
	FieldPublished = "published"
	// FieldPublishedParsed holds the string denoting the published_parsed field in the database.
	FieldPublishedParsed = "published_parsed"
	// FieldUpdated holds the string denoting the updated field in the database.
	FieldUpdated = "updated"
	// FieldUpdatedParsed holds the string denoting the updated_parsed field in the database.
	FieldUpdatedParsed = "updated_parsed"
	// FieldEnclosure holds the string denoting the enclosure field in the database.
	FieldEnclosure = "enclosure"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeFeed holds the string denoting the feed edge name in mutations.
	EdgeFeed = "feed"
	// Table holds the table name of the feeditem in the database.
	Table = "feed_items"
	// FeedTable is the table that holds the feed relation/edge.
	FeedTable = "feed_items"
	// FeedInverseTable is the table name for the Feed entity.
	// It exists in this package in order to avoid circular dependency with the "feed" package.
	FeedInverseTable = "feeds"
	// FeedColumn is the table column denoting the feed relation/edge.
	FeedColumn = "feed_item"
)

// Columns holds all SQL columns for feeditem fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldAuthors,
	FieldDescription,
	FieldContent,
	FieldGUID,
	FieldLink,
	FieldImage,
	FieldPublished,
	FieldPublishedParsed,
	FieldUpdated,
	FieldUpdatedParsed,
	FieldEnclosure,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "feed_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"feed_item",
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