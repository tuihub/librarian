// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// NotifySource is the model entity for the NotifySource schema.
type NotifySource struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// FeedConfigID holds the value of the "feed_config_id" field.
	FeedConfigID model.InternalID `json:"feed_config_id,omitempty"`
	// FeedItemCollectionID holds the value of the "feed_item_collection_id" field.
	FeedItemCollectionID model.InternalID `json:"feed_item_collection_id,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotifySourceQuery when eager-loading is set.
	Edges              NotifySourceEdges `json:"edges"`
	user_notify_source *model.InternalID
	selectValues       sql.SelectValues
}

// NotifySourceEdges holds the relations/edges for other nodes in the graph.
type NotifySourceEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// NotifyFlow holds the value of the notify_flow edge.
	NotifyFlow []*NotifyFlow `json:"notify_flow,omitempty"`
	// FeedConfig holds the value of the feed_config edge.
	FeedConfig *FeedConfig `json:"feed_config,omitempty"`
	// FeedItemCollection holds the value of the feed_item_collection edge.
	FeedItemCollection *FeedItemCollection `json:"feed_item_collection,omitempty"`
	// NotifyFlowSource holds the value of the notify_flow_source edge.
	NotifyFlowSource []*NotifyFlowSource `json:"notify_flow_source,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotifySourceEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// NotifyFlowOrErr returns the NotifyFlow value or an error if the edge
// was not loaded in eager-loading.
func (e NotifySourceEdges) NotifyFlowOrErr() ([]*NotifyFlow, error) {
	if e.loadedTypes[1] {
		return e.NotifyFlow, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow"}
}

// FeedConfigOrErr returns the FeedConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotifySourceEdges) FeedConfigOrErr() (*FeedConfig, error) {
	if e.FeedConfig != nil {
		return e.FeedConfig, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: feedconfig.Label}
	}
	return nil, &NotLoadedError{edge: "feed_config"}
}

// FeedItemCollectionOrErr returns the FeedItemCollection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NotifySourceEdges) FeedItemCollectionOrErr() (*FeedItemCollection, error) {
	if e.FeedItemCollection != nil {
		return e.FeedItemCollection, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: feeditemcollection.Label}
	}
	return nil, &NotLoadedError{edge: "feed_item_collection"}
}

// NotifyFlowSourceOrErr returns the NotifyFlowSource value or an error if the edge
// was not loaded in eager-loading.
func (e NotifySourceEdges) NotifyFlowSourceOrErr() ([]*NotifyFlowSource, error) {
	if e.loadedTypes[4] {
		return e.NotifyFlowSource, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow_source"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotifySource) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notifysource.FieldID, notifysource.FieldFeedConfigID, notifysource.FieldFeedItemCollectionID:
			values[i] = new(sql.NullInt64)
		case notifysource.FieldUpdatedAt, notifysource.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case notifysource.ForeignKeys[0]: // user_notify_source
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotifySource fields.
func (ns *NotifySource) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notifysource.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ns.ID = model.InternalID(value.Int64)
			}
		case notifysource.FieldFeedConfigID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field feed_config_id", values[i])
			} else if value.Valid {
				ns.FeedConfigID = model.InternalID(value.Int64)
			}
		case notifysource.FieldFeedItemCollectionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field feed_item_collection_id", values[i])
			} else if value.Valid {
				ns.FeedItemCollectionID = model.InternalID(value.Int64)
			}
		case notifysource.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ns.UpdatedAt = value.Time
			}
		case notifysource.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ns.CreatedAt = value.Time
			}
		case notifysource.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_notify_source", values[i])
			} else if value.Valid {
				ns.user_notify_source = new(model.InternalID)
				*ns.user_notify_source = model.InternalID(value.Int64)
			}
		default:
			ns.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotifySource.
// This includes values selected through modifiers, order, etc.
func (ns *NotifySource) Value(name string) (ent.Value, error) {
	return ns.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the NotifySource entity.
func (ns *NotifySource) QueryOwner() *UserQuery {
	return NewNotifySourceClient(ns.config).QueryOwner(ns)
}

// QueryNotifyFlow queries the "notify_flow" edge of the NotifySource entity.
func (ns *NotifySource) QueryNotifyFlow() *NotifyFlowQuery {
	return NewNotifySourceClient(ns.config).QueryNotifyFlow(ns)
}

// QueryFeedConfig queries the "feed_config" edge of the NotifySource entity.
func (ns *NotifySource) QueryFeedConfig() *FeedConfigQuery {
	return NewNotifySourceClient(ns.config).QueryFeedConfig(ns)
}

// QueryFeedItemCollection queries the "feed_item_collection" edge of the NotifySource entity.
func (ns *NotifySource) QueryFeedItemCollection() *FeedItemCollectionQuery {
	return NewNotifySourceClient(ns.config).QueryFeedItemCollection(ns)
}

// QueryNotifyFlowSource queries the "notify_flow_source" edge of the NotifySource entity.
func (ns *NotifySource) QueryNotifyFlowSource() *NotifyFlowSourceQuery {
	return NewNotifySourceClient(ns.config).QueryNotifyFlowSource(ns)
}

// Update returns a builder for updating this NotifySource.
// Note that you need to call NotifySource.Unwrap() before calling this method if this NotifySource
// was returned from a transaction, and the transaction was committed or rolled back.
func (ns *NotifySource) Update() *NotifySourceUpdateOne {
	return NewNotifySourceClient(ns.config).UpdateOne(ns)
}

// Unwrap unwraps the NotifySource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ns *NotifySource) Unwrap() *NotifySource {
	_tx, ok := ns.config.driver.(*txDriver)
	if !ok {
		panic("ent: NotifySource is not a transactional entity")
	}
	ns.config.driver = _tx.drv
	return ns
}

// String implements the fmt.Stringer.
func (ns *NotifySource) String() string {
	var builder strings.Builder
	builder.WriteString("NotifySource(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ns.ID))
	builder.WriteString("feed_config_id=")
	builder.WriteString(fmt.Sprintf("%v", ns.FeedConfigID))
	builder.WriteString(", ")
	builder.WriteString("feed_item_collection_id=")
	builder.WriteString(fmt.Sprintf("%v", ns.FeedItemCollectionID))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ns.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ns.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// NotifySources is a parsable slice of NotifySource.
type NotifySources []*NotifySource