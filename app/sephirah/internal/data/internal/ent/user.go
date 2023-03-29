// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// Type holds the value of the "type" field.
	Type user.Type `json:"type,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges             UserEdges `json:"edges"`
	user_created_user *model.InternalID
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// BindAccount holds the value of the bind_account edge.
	BindAccount []*Account `json:"bind_account,omitempty"`
	// PurchasedApp holds the value of the purchased_app edge.
	PurchasedApp []*App `json:"purchased_app,omitempty"`
	// AppPackage holds the value of the app_package edge.
	AppPackage []*AppPackage `json:"app_package,omitempty"`
	// FeedConfig holds the value of the feed_config edge.
	FeedConfig []*FeedConfig `json:"feed_config,omitempty"`
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// CreatedUser holds the value of the created_user edge.
	CreatedUser []*User `json:"created_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// BindAccountOrErr returns the BindAccount value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) BindAccountOrErr() ([]*Account, error) {
	if e.loadedTypes[0] {
		return e.BindAccount, nil
	}
	return nil, &NotLoadedError{edge: "bind_account"}
}

// PurchasedAppOrErr returns the PurchasedApp value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PurchasedAppOrErr() ([]*App, error) {
	if e.loadedTypes[1] {
		return e.PurchasedApp, nil
	}
	return nil, &NotLoadedError{edge: "purchased_app"}
}

// AppPackageOrErr returns the AppPackage value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AppPackageOrErr() ([]*AppPackage, error) {
	if e.loadedTypes[2] {
		return e.AppPackage, nil
	}
	return nil, &NotLoadedError{edge: "app_package"}
}

// FeedConfigOrErr returns the FeedConfig value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FeedConfigOrErr() ([]*FeedConfig, error) {
	if e.loadedTypes[3] {
		return e.FeedConfig, nil
	}
	return nil, &NotLoadedError{edge: "feed_config"}
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// CreatedUserOrErr returns the CreatedUser value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreatedUserOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.CreatedUser, nil
	}
	return nil, &NotLoadedError{edge: "created_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldPassword, user.FieldStatus, user.FieldType:
			values[i] = new(sql.NullString)
		case user.FieldUpdatedAt, user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // user_created_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = model.InternalID(value.Int64)
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				u.Type = user.Type(value.String)
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_created_user", values[i])
			} else if value.Valid {
				u.user_created_user = new(model.InternalID)
				*u.user_created_user = model.InternalID(value.Int64)
			}
		}
	}
	return nil
}

// QueryBindAccount queries the "bind_account" edge of the User entity.
func (u *User) QueryBindAccount() *AccountQuery {
	return NewUserClient(u.config).QueryBindAccount(u)
}

// QueryPurchasedApp queries the "purchased_app" edge of the User entity.
func (u *User) QueryPurchasedApp() *AppQuery {
	return NewUserClient(u.config).QueryPurchasedApp(u)
}

// QueryAppPackage queries the "app_package" edge of the User entity.
func (u *User) QueryAppPackage() *AppPackageQuery {
	return NewUserClient(u.config).QueryAppPackage(u)
}

// QueryFeedConfig queries the "feed_config" edge of the User entity.
func (u *User) QueryFeedConfig() *FeedConfigQuery {
	return NewUserClient(u.config).QueryFeedConfig(u)
}

// QueryCreator queries the "creator" edge of the User entity.
func (u *User) QueryCreator() *UserQuery {
	return NewUserClient(u.config).QueryCreator(u)
}

// QueryCreatedUser queries the "created_user" edge of the User entity.
func (u *User) QueryCreatedUser() *UserQuery {
	return NewUserClient(u.config).QueryCreatedUser(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", u.Type))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User