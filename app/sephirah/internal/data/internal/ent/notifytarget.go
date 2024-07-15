// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyTarget is the model entity for the NotifyTarget schema.
type NotifyTarget struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Destination holds the value of the "destination" field.
	Destination *modeltiphereth.FeatureRequest `json:"destination,omitempty"`
	// Status holds the value of the "status" field.
	Status notifytarget.Status `json:"status,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotifyTargetQuery when eager-loading is set.
	Edges              NotifyTargetEdges `json:"edges"`
	user_notify_target *model.InternalID
	selectValues       sql.SelectValues
}

// NotifyTargetEdges holds the relations/edges for other nodes in the graph.
type NotifyTargetEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// NotifyFlow holds the value of the notify_flow edge.
	NotifyFlow []*NotifyFlow `json:"notify_flow,omitempty"`
	// NotifyFlowTarget holds the value of the notify_flow_target edge.
	NotifyFlowTarget []*NotifyFlowTarget `json:"notify_flow_target,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotifyTargetEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// NotifyFlowOrErr returns the NotifyFlow value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyTargetEdges) NotifyFlowOrErr() ([]*NotifyFlow, error) {
	if e.loadedTypes[1] {
		return e.NotifyFlow, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow"}
}

// NotifyFlowTargetOrErr returns the NotifyFlowTarget value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyTargetEdges) NotifyFlowTargetOrErr() ([]*NotifyFlowTarget, error) {
	if e.loadedTypes[2] {
		return e.NotifyFlowTarget, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow_target"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotifyTarget) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notifytarget.FieldDestination:
			values[i] = new([]byte)
		case notifytarget.FieldID:
			values[i] = new(sql.NullInt64)
		case notifytarget.FieldName, notifytarget.FieldDescription, notifytarget.FieldStatus:
			values[i] = new(sql.NullString)
		case notifytarget.FieldUpdatedAt, notifytarget.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case notifytarget.ForeignKeys[0]: // user_notify_target
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotifyTarget fields.
func (nt *NotifyTarget) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notifytarget.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				nt.ID = model.InternalID(value.Int64)
			}
		case notifytarget.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				nt.Name = value.String
			}
		case notifytarget.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				nt.Description = value.String
			}
		case notifytarget.FieldDestination:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field destination", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &nt.Destination); err != nil {
					return fmt.Errorf("unmarshal field destination: %w", err)
				}
			}
		case notifytarget.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				nt.Status = notifytarget.Status(value.String)
			}
		case notifytarget.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				nt.UpdatedAt = value.Time
			}
		case notifytarget.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				nt.CreatedAt = value.Time
			}
		case notifytarget.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_notify_target", values[i])
			} else if value.Valid {
				nt.user_notify_target = new(model.InternalID)
				*nt.user_notify_target = model.InternalID(value.Int64)
			}
		default:
			nt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotifyTarget.
// This includes values selected through modifiers, order, etc.
func (nt *NotifyTarget) Value(name string) (ent.Value, error) {
	return nt.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the NotifyTarget entity.
func (nt *NotifyTarget) QueryOwner() *UserQuery {
	return NewNotifyTargetClient(nt.config).QueryOwner(nt)
}

// QueryNotifyFlow queries the "notify_flow" edge of the NotifyTarget entity.
func (nt *NotifyTarget) QueryNotifyFlow() *NotifyFlowQuery {
	return NewNotifyTargetClient(nt.config).QueryNotifyFlow(nt)
}

// QueryNotifyFlowTarget queries the "notify_flow_target" edge of the NotifyTarget entity.
func (nt *NotifyTarget) QueryNotifyFlowTarget() *NotifyFlowTargetQuery {
	return NewNotifyTargetClient(nt.config).QueryNotifyFlowTarget(nt)
}

// Update returns a builder for updating this NotifyTarget.
// Note that you need to call NotifyTarget.Unwrap() before calling this method if this NotifyTarget
// was returned from a transaction, and the transaction was committed or rolled back.
func (nt *NotifyTarget) Update() *NotifyTargetUpdateOne {
	return NewNotifyTargetClient(nt.config).UpdateOne(nt)
}

// Unwrap unwraps the NotifyTarget entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nt *NotifyTarget) Unwrap() *NotifyTarget {
	_tx, ok := nt.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotifyTarget is not a transactional entity")
	}
	nt.config.driver = _tx.drv
	return nt
}

// String implements the fmt.Stringer.
func (nt *NotifyTarget) String() string {
	var builder strings.Builder
	builder.WriteString("NotifyTarget(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nt.ID))
	builder.WriteString("name=")
	builder.WriteString(nt.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(nt.Description)
	builder.WriteString(", ")
	builder.WriteString("destination=")
	builder.WriteString(fmt.Sprintf("%v", nt.Destination))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", nt.Status))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(nt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(nt.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NotifyTargets is a parsable slice of NotifyTarget.
type NotifyTargets []*NotifyTarget
