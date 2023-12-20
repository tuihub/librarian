// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedConfig is the model entity for the FeedConfig schema.
type FeedConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID model.InternalID `json:"id,omitempty"`
	// UserFeedConfig holds the value of the "user_feed_config" field.
	UserFeedConfig model.InternalID `json:"user_feed_config,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// FeedURL holds the value of the "feed_url" field.
	FeedURL string `json:"feed_url,omitempty"`
	// AuthorAccount holds the value of the "author_account" field.
	AuthorAccount model.InternalID `json:"author_account,omitempty"`
	// Source holds the value of the "source" field.
	Source feedconfig.Source `json:"source,omitempty"`
	// Status holds the value of the "status" field.
	Status feedconfig.Status `json:"status,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// PullInterval holds the value of the "pull_interval" field.
	PullInterval time.Duration `json:"pull_interval,omitempty"`
	// HideItems holds the value of the "hide_items" field.
	HideItems bool `json:"hide_items,omitempty"`
	// LatestPullAt holds the value of the "latest_pull_at" field.
	LatestPullAt time.Time `json:"latest_pull_at,omitempty"`
	// NextPullBeginAt holds the value of the "next_pull_begin_at" field.
	NextPullBeginAt time.Time `json:"next_pull_begin_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FeedConfigQuery when eager-loading is set.
	Edges        FeedConfigEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FeedConfigEdges holds the relations/edges for other nodes in the graph.
type FeedConfigEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Feed holds the value of the feed edge.
	Feed *Feed `json:"feed,omitempty"`
	// NotifyFlow holds the value of the notify_flow edge.
	NotifyFlow []*NotifyFlow `json:"notify_flow,omitempty"`
	// NotifyFlowSource holds the value of the notify_flow_source edge.
	NotifyFlowSource []*NotifyFlowSource `json:"notify_flow_source,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FeedConfigEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// FeedOrErr returns the Feed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FeedConfigEdges) FeedOrErr() (*Feed, error) {
	if e.loadedTypes[1] {
		if e.Feed == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: feed.Label}
		}
		return e.Feed, nil
	}
	return nil, &NotLoadedError{edge: "feed"}
}

// NotifyFlowOrErr returns the NotifyFlow value or an error if the edge
// was not loaded in eager-loading.
func (e FeedConfigEdges) NotifyFlowOrErr() ([]*NotifyFlow, error) {
	if e.loadedTypes[2] {
		return e.NotifyFlow, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow"}
}

// NotifyFlowSourceOrErr returns the NotifyFlowSource value or an error if the edge
// was not loaded in eager-loading.
func (e FeedConfigEdges) NotifyFlowSourceOrErr() ([]*NotifyFlowSource, error) {
	if e.loadedTypes[3] {
		return e.NotifyFlowSource, nil
	}
	return nil, &NotLoadedError{edge: "notify_flow_source"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FeedConfig) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feedconfig.FieldHideItems:
			values[i] = new(sql.NullBool)
		case feedconfig.FieldID, feedconfig.FieldUserFeedConfig, feedconfig.FieldAuthorAccount, feedconfig.FieldPullInterval:
			values[i] = new(sql.NullInt64)
		case feedconfig.FieldName, feedconfig.FieldFeedURL, feedconfig.FieldSource, feedconfig.FieldStatus, feedconfig.FieldCategory:
			values[i] = new(sql.NullString)
		case feedconfig.FieldLatestPullAt, feedconfig.FieldNextPullBeginAt, feedconfig.FieldUpdatedAt, feedconfig.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FeedConfig fields.
func (fc *FeedConfig) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feedconfig.FieldID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				fc.ID = model.InternalID(value.Int64)
			}
		case feedconfig.FieldUserFeedConfig:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_feed_config", values[i])
			} else if value.Valid {
				fc.UserFeedConfig = model.InternalID(value.Int64)
			}
		case feedconfig.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				fc.Name = value.String
			}
		case feedconfig.FieldFeedURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_url", values[i])
			} else if value.Valid {
				fc.FeedURL = value.String
			}
		case feedconfig.FieldAuthorAccount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author_account", values[i])
			} else if value.Valid {
				fc.AuthorAccount = model.InternalID(value.Int64)
			}
		case feedconfig.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				fc.Source = feedconfig.Source(value.String)
			}
		case feedconfig.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				fc.Status = feedconfig.Status(value.String)
			}
		case feedconfig.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				fc.Category = value.String
			}
		case feedconfig.FieldPullInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pull_interval", values[i])
			} else if value.Valid {
				fc.PullInterval = time.Duration(value.Int64)
			}
		case feedconfig.FieldHideItems:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field hide_items", values[i])
			} else if value.Valid {
				fc.HideItems = value.Bool
			}
		case feedconfig.FieldLatestPullAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field latest_pull_at", values[i])
			} else if value.Valid {
				fc.LatestPullAt = value.Time
			}
		case feedconfig.FieldNextPullBeginAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field next_pull_begin_at", values[i])
			} else if value.Valid {
				fc.NextPullBeginAt = value.Time
			}
		case feedconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fc.UpdatedAt = value.Time
			}
		case feedconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fc.CreatedAt = value.Time
			}
		default:
			fc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FeedConfig.
// This includes values selected through modifiers, order, etc.
func (fc *FeedConfig) Value(name string) (ent.Value, error) {
	return fc.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the FeedConfig entity.
func (fc *FeedConfig) QueryOwner() *UserQuery {
	return NewFeedConfigClient(fc.config).QueryOwner(fc)
}

// QueryFeed queries the "feed" edge of the FeedConfig entity.
func (fc *FeedConfig) QueryFeed() *FeedQuery {
	return NewFeedConfigClient(fc.config).QueryFeed(fc)
}

// QueryNotifyFlow queries the "notify_flow" edge of the FeedConfig entity.
func (fc *FeedConfig) QueryNotifyFlow() *NotifyFlowQuery {
	return NewFeedConfigClient(fc.config).QueryNotifyFlow(fc)
}

// QueryNotifyFlowSource queries the "notify_flow_source" edge of the FeedConfig entity.
func (fc *FeedConfig) QueryNotifyFlowSource() *NotifyFlowSourceQuery {
	return NewFeedConfigClient(fc.config).QueryNotifyFlowSource(fc)
}

// Update returns a builder for updating this FeedConfig.
// Note that you need to call FeedConfig.Unwrap() before calling this method if this FeedConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (fc *FeedConfig) Update() *FeedConfigUpdateOne {
	return NewFeedConfigClient(fc.config).UpdateOne(fc)
}

// Unwrap unwraps the FeedConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fc *FeedConfig) Unwrap() *FeedConfig {
	_tx, ok := fc.config.driver.(*txDriver)
	if !ok {
		panic("ent: FeedConfig is not a transactional entity")
	}
	fc.config.driver = _tx.drv
	return fc
}

// String implements the fmt.Stringer.
func (fc *FeedConfig) String() string {
	var builder strings.Builder
	builder.WriteString("FeedConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fc.ID))
	builder.WriteString("user_feed_config=")
	builder.WriteString(fmt.Sprintf("%v", fc.UserFeedConfig))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(fc.Name)
	builder.WriteString(", ")
	builder.WriteString("feed_url=")
	builder.WriteString(fc.FeedURL)
	builder.WriteString(", ")
	builder.WriteString("author_account=")
	builder.WriteString(fmt.Sprintf("%v", fc.AuthorAccount))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", fc.Source))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", fc.Status))
	builder.WriteString(", ")
	builder.WriteString("category=")
	builder.WriteString(fc.Category)
	builder.WriteString(", ")
	builder.WriteString("pull_interval=")
	builder.WriteString(fmt.Sprintf("%v", fc.PullInterval))
	builder.WriteString(", ")
	builder.WriteString("hide_items=")
	builder.WriteString(fmt.Sprintf("%v", fc.HideItems))
	builder.WriteString(", ")
	builder.WriteString("latest_pull_at=")
	builder.WriteString(fc.LatestPullAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("next_pull_begin_at=")
	builder.WriteString(fc.NextPullBeginAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fc.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FeedConfigs is a parsable slice of FeedConfig.
type FeedConfigs []*FeedConfig
