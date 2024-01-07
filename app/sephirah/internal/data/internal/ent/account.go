// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/model"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Platform holds the value of the "platform" field.
	Platform account.Platform `json:"platform,omitempty"`
	// PlatformAccountID holds the value of the "platform_account_id" field.
	PlatformAccountID string `json:"platform_account_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ProfileURL holds the value of the "profile_url" field.
	ProfileURL string `json:"profile_url,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges             AccountEdges `json:"edges"`
	user_bind_account *model.InternalID
	selectValues      sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// PurchasedApp holds the value of the purchased_app edge.
	PurchasedApp []*App `json:"purchased_app,omitempty"`
	// BindUser holds the value of the bind_user edge.
	BindUser *User `json:"bind_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PurchasedAppOrErr returns the PurchasedApp value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) PurchasedAppOrErr() ([]*App, error) {
	if e.loadedTypes[0] {
		return e.PurchasedApp, nil
	}
	return nil, &NotLoadedError{edge: "purchased_app"}
}

// BindUserOrErr returns the BindUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) BindUserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.BindUser == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BindUser, nil
	}
	return nil, &NotLoadedError{edge: "bind_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			values[i] = new(sql.NullInt64)
		case account.FieldPlatform, account.FieldPlatformAccountID, account.FieldName, account.FieldProfileURL, account.FieldAvatarURL:
			values[i] = new(sql.NullString)
		case account.FieldUpdatedAt, account.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case account.ForeignKeys[0]: // user_bind_account
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = model.InternalID(value.Int64)
			}
		case account.FieldPlatform:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform", values[i])
			} else if value.Valid {
				a.Platform = account.Platform(value.String)
			}
		case account.FieldPlatformAccountID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field platform_account_id", values[i])
			} else if value.Valid {
				a.PlatformAccountID = value.String
			}
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldProfileURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_url", values[i])
			} else if value.Valid {
				a.ProfileURL = value.String
			}
		case account.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				a.AvatarURL = value.String
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_bind_account", values[i])
			} else if value.Valid {
				a.user_bind_account = new(model.InternalID)
				*a.user_bind_account = model.InternalID(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryPurchasedApp queries the "purchased_app" edge of the Account entity.
func (a *Account) QueryPurchasedApp() *AppQuery {
	return NewAccountClient(a.config).QueryPurchasedApp(a)
}

// QueryBindUser queries the "bind_user" edge of the Account entity.
func (a *Account) QueryBindUser() *UserQuery {
	return NewAccountClient(a.config).QueryBindUser(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("platform=")
	builder.WriteString(fmt.Sprintf("%v", a.Platform))
	builder.WriteString(", ")
	builder.WriteString("platform_account_id=")
	builder.WriteString(a.PlatformAccountID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("profile_url=")
	builder.WriteString(a.ProfileURL)
	builder.WriteString(", ")
	builder.WriteString("avatar_url=")
	builder.WriteString(a.AvatarURL)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
