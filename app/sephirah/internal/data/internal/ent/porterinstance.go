// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/model"
)

// PorterInstance is the model entity for the PorterInstance schema.
type PorterInstance struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// GlobalName holds the value of the "global_name" field.
	GlobalName string `json:"global_name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// FeatureSummary holds the value of the "feature_summary" field.
	FeatureSummary *modeltiphereth.PorterFeatureSummary `json:"feature_summary,omitempty"`
	// Status holds the value of the "status" field.
	Status porterinstance.Status `json:"status,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PorterInstance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case porterinstance.FieldFeatureSummary:
			values[i] = new([]byte)
		case porterinstance.FieldID:
			values[i] = new(sql.NullInt64)
		case porterinstance.FieldName, porterinstance.FieldVersion, porterinstance.FieldGlobalName, porterinstance.FieldAddress, porterinstance.FieldStatus:
			values[i] = new(sql.NullString)
		case porterinstance.FieldUpdatedAt, porterinstance.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PorterInstance fields.
func (pi *PorterInstance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case porterinstance.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pi.ID = model.InternalID(value.Int64)
			}
		case porterinstance.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pi.Name = value.String
			}
		case porterinstance.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pi.Version = value.String
			}
		case porterinstance.FieldGlobalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field global_name", values[i])
			} else if value.Valid {
				pi.GlobalName = value.String
			}
		case porterinstance.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				pi.Address = value.String
			}
		case porterinstance.FieldFeatureSummary:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field feature_summary", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pi.FeatureSummary); err != nil {
					return fmt.Errorf("unmarshal field feature_summary: %w", err)
				}
			}
		case porterinstance.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pi.Status = porterinstance.Status(value.String)
			}
		case porterinstance.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		case porterinstance.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PorterInstance.
// This includes values selected through modifiers, order, etc.
func (pi *PorterInstance) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// Update returns a builder for updating this PorterInstance.
// Note that you need to call PorterInstance.Unwrap() before calling this method if this PorterInstance
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PorterInstance) Update() *PorterInstanceUpdateOne {
	return NewPorterInstanceClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the PorterInstance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PorterInstance) Unwrap() *PorterInstance {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PorterInstance is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PorterInstance) String() string {
	var builder strings.Builder
	builder.WriteString("PorterInstance(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("name=")
	builder.WriteString(pi.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pi.Version)
	builder.WriteString(", ")
	builder.WriteString("global_name=")
	builder.WriteString(pi.GlobalName)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(pi.Address)
	builder.WriteString(", ")
	builder.WriteString("feature_summary=")
	builder.WriteString(fmt.Sprintf("%v", pi.FeatureSummary))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pi.Status))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PorterInstances is a parsable slice of PorterInstance.
type PorterInstances []*PorterInstance
