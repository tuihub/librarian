// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinellibrary"
)

// SentinelAppBinary is the model entity for the SentinelAppBinary schema.
type SentinelAppBinary struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SentinelLibraryID holds the value of the "sentinel_library_id" field.
	SentinelLibraryID int `json:"sentinel_library_id,omitempty"`
	// GeneratedID holds the value of the "generated_id" field.
	GeneratedID string `json:"generated_id,omitempty"`
	// SizeBytes holds the value of the "size_bytes" field.
	SizeBytes int64 `json:"size_bytes,omitempty"`
	// NeedToken holds the value of the "need_token" field.
	NeedToken bool `json:"need_token,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Developer holds the value of the "developer" field.
	Developer string `json:"developer,omitempty"`
	// Publisher holds the value of the "publisher" field.
	Publisher string `json:"publisher,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SentinelAppBinaryQuery when eager-loading is set.
	Edges        SentinelAppBinaryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SentinelAppBinaryEdges holds the relations/edges for other nodes in the graph.
type SentinelAppBinaryEdges struct {
	// SentinelLibrary holds the value of the sentinel_library edge.
	SentinelLibrary *SentinelLibrary `json:"sentinel_library,omitempty"`
	// SentinelAppBinaryFile holds the value of the sentinel_app_binary_file edge.
	SentinelAppBinaryFile []*SentinelAppBinaryFile `json:"sentinel_app_binary_file,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SentinelLibraryOrErr returns the SentinelLibrary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SentinelAppBinaryEdges) SentinelLibraryOrErr() (*SentinelLibrary, error) {
	if e.SentinelLibrary != nil {
		return e.SentinelLibrary, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: sentinellibrary.Label}
	}
	return nil, &NotLoadedError{edge: "sentinel_library"}
}

// SentinelAppBinaryFileOrErr returns the SentinelAppBinaryFile value or an error if the edge
// was not loaded in eager-loading.
func (e SentinelAppBinaryEdges) SentinelAppBinaryFileOrErr() ([]*SentinelAppBinaryFile, error) {
	if e.loadedTypes[1] {
		return e.SentinelAppBinaryFile, nil
	}
	return nil, &NotLoadedError{edge: "sentinel_app_binary_file"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SentinelAppBinary) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sentinelappbinary.FieldNeedToken:
			values[i] = new(sql.NullBool)
		case sentinelappbinary.FieldID, sentinelappbinary.FieldSentinelLibraryID, sentinelappbinary.FieldSizeBytes:
			values[i] = new(sql.NullInt64)
		case sentinelappbinary.FieldGeneratedID, sentinelappbinary.FieldName, sentinelappbinary.FieldVersion, sentinelappbinary.FieldDeveloper, sentinelappbinary.FieldPublisher:
			values[i] = new(sql.NullString)
		case sentinelappbinary.FieldUpdatedAt, sentinelappbinary.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SentinelAppBinary fields.
func (sab *SentinelAppBinary) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sentinelappbinary.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sab.ID = int(value.Int64)
		case sentinelappbinary.FieldSentinelLibraryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sentinel_library_id", values[i])
			} else if value.Valid {
				sab.SentinelLibraryID = int(value.Int64)
			}
		case sentinelappbinary.FieldGeneratedID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field generated_id", values[i])
			} else if value.Valid {
				sab.GeneratedID = value.String
			}
		case sentinelappbinary.FieldSizeBytes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size_bytes", values[i])
			} else if value.Valid {
				sab.SizeBytes = value.Int64
			}
		case sentinelappbinary.FieldNeedToken:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field need_token", values[i])
			} else if value.Valid {
				sab.NeedToken = value.Bool
			}
		case sentinelappbinary.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sab.Name = value.String
			}
		case sentinelappbinary.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				sab.Version = value.String
			}
		case sentinelappbinary.FieldDeveloper:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field developer", values[i])
			} else if value.Valid {
				sab.Developer = value.String
			}
		case sentinelappbinary.FieldPublisher:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher", values[i])
			} else if value.Valid {
				sab.Publisher = value.String
			}
		case sentinelappbinary.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sab.UpdatedAt = value.Time
			}
		case sentinelappbinary.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sab.CreatedAt = value.Time
			}
		default:
			sab.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SentinelAppBinary.
// This includes values selected through modifiers, order, etc.
func (sab *SentinelAppBinary) Value(name string) (ent.Value, error) {
	return sab.selectValues.Get(name)
}

// QuerySentinelLibrary queries the "sentinel_library" edge of the SentinelAppBinary entity.
func (sab *SentinelAppBinary) QuerySentinelLibrary() *SentinelLibraryQuery {
	return NewSentinelAppBinaryClient(sab.config).QuerySentinelLibrary(sab)
}

// QuerySentinelAppBinaryFile queries the "sentinel_app_binary_file" edge of the SentinelAppBinary entity.
func (sab *SentinelAppBinary) QuerySentinelAppBinaryFile() *SentinelAppBinaryFileQuery {
	return NewSentinelAppBinaryClient(sab.config).QuerySentinelAppBinaryFile(sab)
}

// Update returns a builder for updating this SentinelAppBinary.
// Note that you need to call SentinelAppBinary.Unwrap() before calling this method if this SentinelAppBinary
// was returned from a transaction, and the transaction was committed or rolled back.
func (sab *SentinelAppBinary) Update() *SentinelAppBinaryUpdateOne {
	return NewSentinelAppBinaryClient(sab.config).UpdateOne(sab)
}

// Unwrap unwraps the SentinelAppBinary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sab *SentinelAppBinary) Unwrap() *SentinelAppBinary {
	_tx, ok := sab.config.driver.(*txDriver)
	if !ok {
		panic("ent: SentinelAppBinary is not a transactional entity")
	}
	sab.config.driver = _tx.drv
	return sab
}

// String implements the fmt.Stringer.
func (sab *SentinelAppBinary) String() string {
	var builder strings.Builder
	builder.WriteString("SentinelAppBinary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sab.ID))
	builder.WriteString("sentinel_library_id=")
	builder.WriteString(fmt.Sprintf("%v", sab.SentinelLibraryID))
	builder.WriteString(", ")
	builder.WriteString("generated_id=")
	builder.WriteString(sab.GeneratedID)
	builder.WriteString(", ")
	builder.WriteString("size_bytes=")
	builder.WriteString(fmt.Sprintf("%v", sab.SizeBytes))
	builder.WriteString(", ")
	builder.WriteString("need_token=")
	builder.WriteString(fmt.Sprintf("%v", sab.NeedToken))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sab.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(sab.Version)
	builder.WriteString(", ")
	builder.WriteString("developer=")
	builder.WriteString(sab.Developer)
	builder.WriteString(", ")
	builder.WriteString("publisher=")
	builder.WriteString(sab.Publisher)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sab.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sab.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SentinelAppBinaries is a parsable slice of SentinelAppBinary.
type SentinelAppBinaries []*SentinelAppBinary
