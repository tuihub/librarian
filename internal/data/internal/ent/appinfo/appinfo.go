// Code generated by ent, DO NOT EDIT.

package appinfo

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the appinfo type in the database.
	Label = "app_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldSourceAppID holds the string denoting the source_app_id field in the database.
	FieldSourceAppID = "source_app_id"
	// FieldSourceURL holds the string denoting the source_url field in the database.
	FieldSourceURL = "source_url"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldShortDescription holds the string denoting the short_description field in the database.
	FieldShortDescription = "short_description"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldIconImageURL holds the string denoting the icon_image_url field in the database.
	FieldIconImageURL = "icon_image_url"
	// FieldIconImageID holds the string denoting the icon_image_id field in the database.
	FieldIconImageID = "icon_image_id"
	// FieldBackgroundImageURL holds the string denoting the background_image_url field in the database.
	FieldBackgroundImageURL = "background_image_url"
	// FieldBackgroundImageID holds the string denoting the background_image_id field in the database.
	FieldBackgroundImageID = "background_image_id"
	// FieldCoverImageURL holds the string denoting the cover_image_url field in the database.
	FieldCoverImageURL = "cover_image_url"
	// FieldCoverImageID holds the string denoting the cover_image_id field in the database.
	FieldCoverImageID = "cover_image_id"
	// FieldReleaseDate holds the string denoting the release_date field in the database.
	FieldReleaseDate = "release_date"
	// FieldDeveloper holds the string denoting the developer field in the database.
	FieldDeveloper = "developer"
	// FieldPublisher holds the string denoting the publisher field in the database.
	FieldPublisher = "publisher"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldAlternativeNames holds the string denoting the alternative_names field in the database.
	FieldAlternativeNames = "alternative_names"
	// FieldRawData holds the string denoting the raw_data field in the database.
	FieldRawData = "raw_data"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the appinfo in the database.
	Table = "app_infos"
)

// Columns holds all SQL columns for appinfo fields.
var Columns = []string{
	FieldID,
	FieldSource,
	FieldSourceAppID,
	FieldSourceURL,
	FieldName,
	FieldType,
	FieldShortDescription,
	FieldDescription,
	FieldIconImageURL,
	FieldIconImageID,
	FieldBackgroundImageURL,
	FieldBackgroundImageID,
	FieldCoverImageURL,
	FieldCoverImageID,
	FieldReleaseDate,
	FieldDeveloper,
	FieldPublisher,
	FieldTags,
	FieldAlternativeNames,
	FieldRawData,
	FieldUpdatedAt,
	FieldCreatedAt,
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeUnknown Type = "unknown"
	TypeGame    Type = "game"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeUnknown, TypeGame:
		return nil
	default:
		return fmt.Errorf("appinfo: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the AppInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// BySourceAppID orders the results by the source_app_id field.
func BySourceAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceAppID, opts...).ToFunc()
}

// BySourceURL orders the results by the source_url field.
func BySourceURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceURL, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByShortDescription orders the results by the short_description field.
func ByShortDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShortDescription, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByIconImageURL orders the results by the icon_image_url field.
func ByIconImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIconImageURL, opts...).ToFunc()
}

// ByIconImageID orders the results by the icon_image_id field.
func ByIconImageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIconImageID, opts...).ToFunc()
}

// ByBackgroundImageURL orders the results by the background_image_url field.
func ByBackgroundImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackgroundImageURL, opts...).ToFunc()
}

// ByBackgroundImageID orders the results by the background_image_id field.
func ByBackgroundImageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackgroundImageID, opts...).ToFunc()
}

// ByCoverImageURL orders the results by the cover_image_url field.
func ByCoverImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverImageURL, opts...).ToFunc()
}

// ByCoverImageID orders the results by the cover_image_id field.
func ByCoverImageID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoverImageID, opts...).ToFunc()
}

// ByReleaseDate orders the results by the release_date field.
func ByReleaseDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleaseDate, opts...).ToFunc()
}

// ByDeveloper orders the results by the developer field.
func ByDeveloper(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeveloper, opts...).ToFunc()
}

// ByPublisher orders the results by the publisher field.
func ByPublisher(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublisher, opts...).ToFunc()
}

// ByRawData orders the results by the raw_data field.
func ByRawData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRawData, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
