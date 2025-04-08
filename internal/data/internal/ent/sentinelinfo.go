// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelinfo"
	"github.com/tuihub/librarian/internal/model"
)

// SentinelInfo is the model entity for the SentinelInfo schema.
type SentinelInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID model.InternalID `json:"user_id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// AlternativeUrls holds the value of the "alternative_urls" field.
	AlternativeUrls []string `json:"alternative_urls,omitempty"`
	// GetTokenPath holds the value of the "get_token_path" field.
	GetTokenPath string `json:"get_token_path,omitempty"`
	// DownloadFileBasePath holds the value of the "download_file_base_path" field.
	DownloadFileBasePath string `json:"download_file_base_path,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SentinelInfoQuery when eager-loading is set.
	Edges        SentinelInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SentinelInfoEdges holds the relations/edges for other nodes in the graph.
type SentinelInfoEdges struct {
	// SentinelLibrary holds the value of the sentinel_library edge.
	SentinelLibrary []*SentinelLibrary `json:"sentinel_library,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SentinelLibraryOrErr returns the SentinelLibrary value or an error if the edge
// was not loaded in eager-loading.
func (e SentinelInfoEdges) SentinelLibraryOrErr() ([]*SentinelLibrary, error) {
	if e.loadedTypes[0] {
		return e.SentinelLibrary, nil
	}
	return nil, &NotLoadedError{edge: "sentinel_library"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SentinelInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sentinelinfo.FieldAlternativeUrls:
			values[i] = new([]byte)
		case sentinelinfo.FieldID, sentinelinfo.FieldUserID:
			values[i] = new(sql.NullInt64)
		case sentinelinfo.FieldURL, sentinelinfo.FieldGetTokenPath, sentinelinfo.FieldDownloadFileBasePath:
			values[i] = new(sql.NullString)
		case sentinelinfo.FieldUpdatedAt, sentinelinfo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SentinelInfo fields.
func (si *SentinelInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sentinelinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			si.ID = int(value.Int64)
		case sentinelinfo.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				si.UserID = model.InternalID(value.Int64)
			}
		case sentinelinfo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				si.URL = value.String
			}
		case sentinelinfo.FieldAlternativeUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field alternative_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &si.AlternativeUrls); err != nil {
					return fmt.Errorf("unmarshal field alternative_urls: %w", err)
				}
			}
		case sentinelinfo.FieldGetTokenPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field get_token_path", values[i])
			} else if value.Valid {
				si.GetTokenPath = value.String
			}
		case sentinelinfo.FieldDownloadFileBasePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field download_file_base_path", values[i])
			} else if value.Valid {
				si.DownloadFileBasePath = value.String
			}
		case sentinelinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				si.UpdatedAt = value.Time
			}
		case sentinelinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				si.CreatedAt = value.Time
			}
		default:
			si.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SentinelInfo.
// This includes values selected through modifiers, order, etc.
func (si *SentinelInfo) Value(name string) (ent.Value, error) {
	return si.selectValues.Get(name)
}

// QuerySentinelLibrary queries the "sentinel_library" edge of the SentinelInfo entity.
func (si *SentinelInfo) QuerySentinelLibrary() *SentinelLibraryQuery {
	return NewSentinelInfoClient(si.config).QuerySentinelLibrary(si)
}

// Update returns a builder for updating this SentinelInfo.
// Note that you need to call SentinelInfo.Unwrap() before calling this method if this SentinelInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (si *SentinelInfo) Update() *SentinelInfoUpdateOne {
	return NewSentinelInfoClient(si.config).UpdateOne(si)
}

// Unwrap unwraps the SentinelInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (si *SentinelInfo) Unwrap() *SentinelInfo {
	_tx, ok := si.config.driver.(*txDriver)
	if !ok {
		panic("ent: SentinelInfo is not a transactional entity")
	}
	si.config.driver = _tx.drv
	return si
}

// String implements the fmt.Stringer.
func (si *SentinelInfo) String() string {
	var builder strings.Builder
	builder.WriteString("SentinelInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", si.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", si.UserID))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(si.URL)
	builder.WriteString(", ")
	builder.WriteString("alternative_urls=")
	builder.WriteString(fmt.Sprintf("%v", si.AlternativeUrls))
	builder.WriteString(", ")
	builder.WriteString("get_token_path=")
	builder.WriteString(si.GetTokenPath)
	builder.WriteString(", ")
	builder.WriteString("download_file_base_path=")
	builder.WriteString(si.DownloadFileBasePath)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(si.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(si.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SentinelInfos is a parsable slice of SentinelInfo.
type SentinelInfos []*SentinelInfo
