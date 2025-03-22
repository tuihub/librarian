// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/device"
	"github.com/tuihub/librarian/internal/data/internal/ent/session"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID model.InternalID `json:"user_id,omitempty"`
	// DeviceID holds the value of the "device_id" field.
	DeviceID model.InternalID `json:"device_id,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// ExpireAt holds the value of the "expire_at" field.
	ExpireAt time.Time `json:"expire_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges        SessionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SessionEdges) DeviceOrErr() (*Device, error) {
	if e.Device != nil {
		return e.Device, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "device"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case session.FieldID, session.FieldUserID, session.FieldDeviceID:
			values[i] = new(sql.NullInt64)
		case session.FieldRefreshToken:
			values[i] = new(sql.NullString)
		case session.FieldExpireAt, session.FieldUpdatedAt, session.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case session.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = model.InternalID(value.Int64)
			}
		case session.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = model.InternalID(value.Int64)
			}
		case session.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				s.DeviceID = model.InternalID(value.Int64)
			}
		case session.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				s.RefreshToken = value.String
			}
		case session.FieldExpireAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				s.ExpireAt = value.Time
			}
		case session.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case session.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Session.
// This includes values selected through modifiers, order, etc.
func (s *Session) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Session entity.
func (s *Session) QueryUser() *UserQuery {
	return NewSessionClient(s.config).QueryUser(s)
}

// QueryDevice queries the "device" edge of the Session entity.
func (s *Session) QueryDevice() *DeviceQuery {
	return NewSessionClient(s.config).QueryDevice(s)
}

// Update returns a builder for updating this Session.
// Note that you need to call Session.Unwrap() before calling this method if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return NewSessionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Session entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", s.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("refresh_token=")
	builder.WriteString(s.RefreshToken)
	builder.WriteString(", ")
	builder.WriteString("expire_at=")
	builder.WriteString(s.ExpireAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session
