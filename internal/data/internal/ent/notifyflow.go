// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyFlow is the model entity for the NotifyFlow schema.
type NotifyFlow struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status notifyflow.Status `json:"status,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotifyFlowQuery when eager-loading is set.
	Edges            NotifyFlowEdges `json:"edges"`
	user_notify_flow *model.InternalID
	selectValues     sql.SelectValues
}

// NotifyFlowEdges holds the relations/edges for other nodes in the graph.
type NotifyFlowEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// NotifyTarget holds the value of the notify_target edge.
	NotifyTarget []*NotifyTarget `json:"notify_target,omitempty"`
	// NotifySource holds the value of the notify_source edge.
	NotifySource []*NotifySource `json:"notify_source,omitempty"`
	// NotifyFlowTarget holds the value of the notify_flow_target edge.
	NotifyFlowTarget []*NotifyFlowTarget `json:"notify_flow_target,omitempty"`
	// NotifyFlowSource holds the value of the notify_flow_source edge.
	NotifyFlowSource []*NotifyFlowSource `json:"notify_flow_source,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotifyFlowEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// NotifyTargetOrErr returns the NotifyTarget value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyFlowEdges) NotifyTargetOrErr() ([]*NotifyTarget, error) {
	if e.loadedTypes[1] {
		return e.NotifyTarget, nil
	}
	return nil, &NotLoadedError{edge: "notify_target"}
}

// NotifySourceOrErr returns the NotifySource value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyFlowEdges) NotifySourceOrErr() ([]*NotifySource, error) {
	if e.loadedTypes[2] {
		return e.NotifySource, nil
	}
	return nil, &NotLoadedError{edge: "notify_source"}
}

// NotifyFlowTargetOrErr returns the NotifyFlowTarget value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyFlowEdges) NotifyFlowTargetOrErr() ([]*NotifyFlowTarget, error) {
	if e.loadedTypes[3] {
		return e.NotifyFlowTarget, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow_target"}
}

// NotifyFlowSourceOrErr returns the NotifyFlowSource value or an error if the edge
// was not loaded in eager-loading.
func (e NotifyFlowEdges) NotifyFlowSourceOrErr() ([]*NotifyFlowSource, error) {
	if e.loadedTypes[4] {
		return e.NotifyFlowSource, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow_source"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotifyFlow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notifyflow.FieldID:
			values[i] = new(sql.NullInt64)
		case notifyflow.FieldName, notifyflow.FieldDescription, notifyflow.FieldStatus:
			values[i] = new(sql.NullString)
		case notifyflow.FieldUpdatedAt, notifyflow.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case notifyflow.ForeignKeys[0]: // user_notify_flow
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotifyFlow fields.
func (nf *NotifyFlow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notifyflow.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				nf.ID = model.InternalID(value.Int64)
			}
		case notifyflow.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				nf.Name = value.String
			}
		case notifyflow.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				nf.Description = value.String
			}
		case notifyflow.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				nf.Status = notifyflow.Status(value.String)
			}
		case notifyflow.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				nf.UpdatedAt = value.Time
			}
		case notifyflow.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				nf.CreatedAt = value.Time
			}
		case notifyflow.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_notify_flow", values[i])
			} else if value.Valid {
				nf.user_notify_flow = new(model.InternalID)
				*nf.user_notify_flow = model.InternalID(value.Int64)
			}
		default:
			nf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotifyFlow.
// This includes values selected through modifiers, order, etc.
func (nf *NotifyFlow) Value(name string) (ent.Value, error) {
	return nf.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the NotifyFlow entity.
func (nf *NotifyFlow) QueryOwner() *UserQuery {
	return NewNotifyFlowClient(nf.config).QueryOwner(nf)
}

// QueryNotifyTarget queries the "notify_target" edge of the NotifyFlow entity.
func (nf *NotifyFlow) QueryNotifyTarget() *NotifyTargetQuery {
	return NewNotifyFlowClient(nf.config).QueryNotifyTarget(nf)
}

// QueryNotifySource queries the "notify_source" edge of the NotifyFlow entity.
func (nf *NotifyFlow) QueryNotifySource() *NotifySourceQuery {
	return NewNotifyFlowClient(nf.config).QueryNotifySource(nf)
}

// QueryNotifyFlowTarget queries the "notify_flow_target" edge of the NotifyFlow entity.
func (nf *NotifyFlow) QueryNotifyFlowTarget() *NotifyFlowTargetQuery {
	return NewNotifyFlowClient(nf.config).QueryNotifyFlowTarget(nf)
}

// QueryNotifyFlowSource queries the "notify_flow_source" edge of the NotifyFlow entity.
func (nf *NotifyFlow) QueryNotifyFlowSource() *NotifyFlowSourceQuery {
	return NewNotifyFlowClient(nf.config).QueryNotifyFlowSource(nf)
}

// Update returns a builder for updating this NotifyFlow.
// Note that you need to call NotifyFlow.Unwrap() before calling this method if this NotifyFlow
// was returned from a transaction, and the transaction was committed or rolled back.
func (nf *NotifyFlow) Update() *NotifyFlowUpdateOne {
	return NewNotifyFlowClient(nf.config).UpdateOne(nf)
}

// Unwrap unwraps the NotifyFlow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nf *NotifyFlow) Unwrap() *NotifyFlow {
	_tx, ok := nf.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotifyFlow is not a transactional entity")
	}
	nf.config.driver = _tx.drv
	return nf
}

// String implements the fmt.Stringer.
func (nf *NotifyFlow) String() string {
	var builder strings.Builder
	builder.WriteString("NotifyFlow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nf.ID))
	builder.WriteString("name=")
	builder.WriteString(nf.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(nf.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", nf.Status))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(nf.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(nf.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NotifyFlows is a parsable slice of NotifyFlow.
type NotifyFlows []*NotifyFlow
