// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the deviceinfo type in the database.
	Label = "device_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeviceName holds the string denoting the device_name field in the database.
	FieldDeviceName = "device_name"
	// FieldSystemType holds the string denoting the system_type field in the database.
	FieldSystemType = "system_type"
	// FieldSystemVersion holds the string denoting the system_version field in the database.
	FieldSystemVersion = "system_version"
	// FieldClientName holds the string denoting the client_name field in the database.
	FieldClientName = "client_name"
	// FieldClientSourceCodeAddress holds the string denoting the client_source_code_address field in the database.
	FieldClientSourceCodeAddress = "client_source_code_address"
	// FieldClientVersion holds the string denoting the client_version field in the database.
	FieldClientVersion = "client_version"
	// FieldClientLocalID holds the string denoting the client_local_id field in the database.
	FieldClientLocalID = "client_local_id"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeUserSession holds the string denoting the user_session edge name in mutations.
	EdgeUserSession = "user_session"
	// EdgeUserDevice holds the string denoting the user_device edge name in mutations.
	EdgeUserDevice = "user_device"
	// Table holds the table name of the deviceinfo in the database.
	Table = "device_infos"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_devices"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserSessionTable is the table that holds the user_session relation/edge.
	UserSessionTable = "user_sessions"
	// UserSessionInverseTable is the table name for the UserSession entity.
	// It exists in this package in order to avoid circular dependency with the "usersession" package.
	UserSessionInverseTable = "user_sessions"
	// UserSessionColumn is the table column denoting the user_session relation/edge.
	UserSessionColumn = "device_info_user_session"
	// UserDeviceTable is the table that holds the user_device relation/edge.
	UserDeviceTable = "user_devices"
	// UserDeviceInverseTable is the table name for the UserDevice entity.
	// It exists in this package in order to avoid circular dependency with the "userdevice" package.
	UserDeviceInverseTable = "user_devices"
	// UserDeviceColumn is the table column denoting the user_device relation/edge.
	UserDeviceColumn = "device_id"
)

// Columns holds all SQL columns for deviceinfo fields.
var Columns = []string{
	FieldID,
	FieldDeviceName,
	FieldSystemType,
	FieldSystemVersion,
	FieldClientName,
	FieldClientSourceCodeAddress,
	FieldClientVersion,
	FieldClientLocalID,
	FieldUpdatedAt,
	FieldCreatedAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "device_id"}
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

// SystemType defines the type for the "system_type" enum field.
type SystemType string

// SystemType values.
const (
	SystemTypeIos     SystemType = "ios"
	SystemTypeAndroid SystemType = "android"
	SystemTypeWeb     SystemType = "web"
	SystemTypeWindows SystemType = "windows"
	SystemTypeMacos   SystemType = "macos"
	SystemTypeLinux   SystemType = "linux"
	SystemTypeUnknown SystemType = "unknown"
)

func (st SystemType) String() string {
	return string(st)
}

// SystemTypeValidator is a validator for the "system_type" field enum values. It is called by the builders before save.
func SystemTypeValidator(st SystemType) error {
	switch st {
	case SystemTypeIos, SystemTypeAndroid, SystemTypeWeb, SystemTypeWindows, SystemTypeMacos, SystemTypeLinux, SystemTypeUnknown:
		return nil
	default:
		return fmt.Errorf("deviceinfo: invalid enum value for system_type field: %q", st)
	}
}

// OrderOption defines the ordering options for the DeviceInfo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeviceName orders the results by the device_name field.
func ByDeviceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceName, opts...).ToFunc()
}

// BySystemType orders the results by the system_type field.
func BySystemType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemType, opts...).ToFunc()
}

// BySystemVersion orders the results by the system_version field.
func BySystemVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemVersion, opts...).ToFunc()
}

// ByClientName orders the results by the client_name field.
func ByClientName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientName, opts...).ToFunc()
}

// ByClientSourceCodeAddress orders the results by the client_source_code_address field.
func ByClientSourceCodeAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientSourceCodeAddress, opts...).ToFunc()
}

// ByClientVersion orders the results by the client_version field.
func ByClientVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientVersion, opts...).ToFunc()
}

// ByClientLocalID orders the results by the client_local_id field.
func ByClientLocalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientLocalID, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserSessionCount orders the results by user_session count.
func ByUserSessionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserSessionStep(), opts...)
	}
}

// ByUserSession orders the results by user_session terms.
func ByUserSession(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserSessionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserDeviceCount orders the results by user_device count.
func ByUserDeviceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserDeviceStep(), opts...)
	}
}

// ByUserDevice orders the results by user_device terms.
func ByUserDevice(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserDeviceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
func newUserSessionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserSessionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserSessionTable, UserSessionColumn),
	)
}
func newUserDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserDeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UserDeviceTable, UserDeviceColumn),
	)
}
