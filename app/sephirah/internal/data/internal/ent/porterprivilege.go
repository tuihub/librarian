// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterprivilege"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"
)

// PorterPrivilege is the model entity for the PorterPrivilege schema.
type PorterPrivilege struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID model.InternalID `json:"user_id,omitempty"`
	// PorterID holds the value of the "porter_id" field.
	PorterID model.InternalID `json:"porter_id,omitempty"`
	// Privilege holds the value of the "privilege" field.
	Privilege *modeltiphereth.PorterInstancePrivilege `json:"privilege,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PorterPrivilege) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case porterprivilege.FieldPrivilege:
			values[i] = new([]byte)
		case porterprivilege.FieldID, porterprivilege.FieldUserID, porterprivilege.FieldPorterID:
			values[i] = new(sql.NullInt64)
		case porterprivilege.FieldUpdatedAt, porterprivilege.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PorterPrivilege fields.
func (pp *PorterPrivilege) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case porterprivilege.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pp.ID = int(value.Int64)
		case porterprivilege.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pp.UserID = model.InternalID(value.Int64)
			}
		case porterprivilege.FieldPorterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field porter_id", values[i])
			} else if value.Valid {
				pp.PorterID = model.InternalID(value.Int64)
			}
		case porterprivilege.FieldPrivilege:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field privilege", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pp.Privilege); err != nil {
					return fmt.Errorf("unmarshal field privilege: %w", err)
				}
			}
		case porterprivilege.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pp.UpdatedAt = value.Time
			}
		case porterprivilege.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pp.CreatedAt = value.Time
			}
		default:
			pp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PorterPrivilege.
// This includes values selected through modifiers, order, etc.
func (pp *PorterPrivilege) Value(name string) (ent.Value, error) {
	return pp.selectValues.Get(name)
}

// Update returns a builder for updating this PorterPrivilege.
// Note that you need to call PorterPrivilege.Unwrap() before calling this method if this PorterPrivilege
// was returned from a transaction, and the transaction was committed or rolled back.
func (pp *PorterPrivilege) Update() *PorterPrivilegeUpdateOne {
	return NewPorterPrivilegeClient(pp.config).UpdateOne(pp)
}

// Unwrap unwraps the PorterPrivilege entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pp *PorterPrivilege) Unwrap() *PorterPrivilege {
	_tx, ok := pp.config.driver.(*txDriver)
	if !ok {
		panic("ent: PorterPrivilege is not a transactional entity")
	}
	pp.config.driver = _tx.drv
	return pp
}

// String implements the fmt.Stringer.
func (pp *PorterPrivilege) String() string {
	var builder strings.Builder
	builder.WriteString("PorterPrivilege(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pp.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.UserID))
	builder.WriteString(", ")
	builder.WriteString("porter_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.PorterID))
	builder.WriteString(", ")
	builder.WriteString("privilege=")
	builder.WriteString(fmt.Sprintf("%v", pp.Privilege))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pp.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PorterPrivileges is a parsable slice of PorterPrivilege.
type PorterPrivileges []*PorterPrivilege
