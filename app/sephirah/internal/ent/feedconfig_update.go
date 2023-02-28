// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/user"
)

// FeedConfigUpdate is the builder for updating FeedConfig entities.
type FeedConfigUpdate struct {
	config
	hooks    []Hook
	mutation *FeedConfigMutation
}

// Where appends a list predicates to the FeedConfigUpdate builder.
func (fcu *FeedConfigUpdate) Where(ps ...predicate.FeedConfig) *FeedConfigUpdate {
	fcu.mutation.Where(ps...)
	return fcu
}

// SetFeedURL sets the "feed_url" field.
func (fcu *FeedConfigUpdate) SetFeedURL(s string) *FeedConfigUpdate {
	fcu.mutation.SetFeedURL(s)
	return fcu
}

// SetAuthorAccount sets the "author_account" field.
func (fcu *FeedConfigUpdate) SetAuthorAccount(i int64) *FeedConfigUpdate {
	fcu.mutation.ResetAuthorAccount()
	fcu.mutation.SetAuthorAccount(i)
	return fcu
}

// AddAuthorAccount adds i to the "author_account" field.
func (fcu *FeedConfigUpdate) AddAuthorAccount(i int64) *FeedConfigUpdate {
	fcu.mutation.AddAuthorAccount(i)
	return fcu
}

// SetSource sets the "source" field.
func (fcu *FeedConfigUpdate) SetSource(f feedconfig.Source) *FeedConfigUpdate {
	fcu.mutation.SetSource(f)
	return fcu
}

// SetStatus sets the "status" field.
func (fcu *FeedConfigUpdate) SetStatus(f feedconfig.Status) *FeedConfigUpdate {
	fcu.mutation.SetStatus(f)
	return fcu
}

// SetPullInterval sets the "pull_interval" field.
func (fcu *FeedConfigUpdate) SetPullInterval(t time.Duration) *FeedConfigUpdate {
	fcu.mutation.ResetPullInterval()
	fcu.mutation.SetPullInterval(t)
	return fcu
}

// AddPullInterval adds t to the "pull_interval" field.
func (fcu *FeedConfigUpdate) AddPullInterval(t time.Duration) *FeedConfigUpdate {
	fcu.mutation.AddPullInterval(t)
	return fcu
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (fcu *FeedConfigUpdate) SetNextPullBeginAt(t time.Time) *FeedConfigUpdate {
	fcu.mutation.SetNextPullBeginAt(t)
	return fcu
}

// SetNillableNextPullBeginAt sets the "next_pull_begin_at" field if the given value is not nil.
func (fcu *FeedConfigUpdate) SetNillableNextPullBeginAt(t *time.Time) *FeedConfigUpdate {
	if t != nil {
		fcu.SetNextPullBeginAt(*t)
	}
	return fcu
}

// SetUpdatedAt sets the "updated_at" field.
func (fcu *FeedConfigUpdate) SetUpdatedAt(t time.Time) *FeedConfigUpdate {
	fcu.mutation.SetUpdatedAt(t)
	return fcu
}

// SetCreatedAt sets the "created_at" field.
func (fcu *FeedConfigUpdate) SetCreatedAt(t time.Time) *FeedConfigUpdate {
	fcu.mutation.SetCreatedAt(t)
	return fcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcu *FeedConfigUpdate) SetNillableCreatedAt(t *time.Time) *FeedConfigUpdate {
	if t != nil {
		fcu.SetCreatedAt(*t)
	}
	return fcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fcu *FeedConfigUpdate) SetUserID(id int64) *FeedConfigUpdate {
	fcu.mutation.SetUserID(id)
	return fcu
}

// SetUser sets the "user" edge to the User entity.
func (fcu *FeedConfigUpdate) SetUser(u *User) *FeedConfigUpdate {
	return fcu.SetUserID(u.ID)
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fcu *FeedConfigUpdate) SetFeedID(id int64) *FeedConfigUpdate {
	fcu.mutation.SetFeedID(id)
	return fcu
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (fcu *FeedConfigUpdate) SetNillableFeedID(id *int64) *FeedConfigUpdate {
	if id != nil {
		fcu = fcu.SetFeedID(*id)
	}
	return fcu
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fcu *FeedConfigUpdate) SetFeed(f *Feed) *FeedConfigUpdate {
	return fcu.SetFeedID(f.ID)
}

// Mutation returns the FeedConfigMutation object of the builder.
func (fcu *FeedConfigUpdate) Mutation() *FeedConfigMutation {
	return fcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fcu *FeedConfigUpdate) ClearUser() *FeedConfigUpdate {
	fcu.mutation.ClearUser()
	return fcu
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (fcu *FeedConfigUpdate) ClearFeed() *FeedConfigUpdate {
	fcu.mutation.ClearFeed()
	return fcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fcu *FeedConfigUpdate) Save(ctx context.Context) (int, error) {
	fcu.defaults()
	return withHooks[int, FeedConfigMutation](ctx, fcu.sqlSave, fcu.mutation, fcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcu *FeedConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := fcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fcu *FeedConfigUpdate) Exec(ctx context.Context) error {
	_, err := fcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcu *FeedConfigUpdate) ExecX(ctx context.Context) {
	if err := fcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcu *FeedConfigUpdate) defaults() {
	if _, ok := fcu.mutation.UpdatedAt(); !ok {
		v := feedconfig.UpdateDefaultUpdatedAt()
		fcu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcu *FeedConfigUpdate) check() error {
	if v, ok := fcu.mutation.Source(); ok {
		if err := feedconfig.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.source": %w`, err)}
		}
	}
	if v, ok := fcu.mutation.Status(); ok {
		if err := feedconfig.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.status": %w`, err)}
		}
	}
	if _, ok := fcu.mutation.UserID(); fcu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedConfig.user"`)
	}
	return nil
}

func (fcu *FeedConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feedconfig.Table, feedconfig.Columns, sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64))
	if ps := fcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcu.mutation.FeedURL(); ok {
		_spec.SetField(feedconfig.FieldFeedURL, field.TypeString, value)
	}
	if value, ok := fcu.mutation.AuthorAccount(); ok {
		_spec.SetField(feedconfig.FieldAuthorAccount, field.TypeInt64, value)
	}
	if value, ok := fcu.mutation.AddedAuthorAccount(); ok {
		_spec.AddField(feedconfig.FieldAuthorAccount, field.TypeInt64, value)
	}
	if value, ok := fcu.mutation.Source(); ok {
		_spec.SetField(feedconfig.FieldSource, field.TypeEnum, value)
	}
	if value, ok := fcu.mutation.Status(); ok {
		_spec.SetField(feedconfig.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fcu.mutation.PullInterval(); ok {
		_spec.SetField(feedconfig.FieldPullInterval, field.TypeInt64, value)
	}
	if value, ok := fcu.mutation.AddedPullInterval(); ok {
		_spec.AddField(feedconfig.FieldPullInterval, field.TypeInt64, value)
	}
	if value, ok := fcu.mutation.NextPullBeginAt(); ok {
		_spec.SetField(feedconfig.FieldNextPullBeginAt, field.TypeTime, value)
	}
	if value, ok := fcu.mutation.UpdatedAt(); ok {
		_spec.SetField(feedconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fcu.mutation.CreatedAt(); ok {
		_spec.SetField(feedconfig.FieldCreatedAt, field.TypeTime, value)
	}
	if fcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedconfig.UserTable,
			Columns: []string{feedconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedconfig.UserTable,
			Columns: []string{feedconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fcu.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   feedconfig.FeedTable,
			Columns: []string{feedconfig.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcu.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   feedconfig.FeedTable,
			Columns: []string{feedconfig.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feedconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fcu.mutation.done = true
	return n, nil
}

// FeedConfigUpdateOne is the builder for updating a single FeedConfig entity.
type FeedConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedConfigMutation
}

// SetFeedURL sets the "feed_url" field.
func (fcuo *FeedConfigUpdateOne) SetFeedURL(s string) *FeedConfigUpdateOne {
	fcuo.mutation.SetFeedURL(s)
	return fcuo
}

// SetAuthorAccount sets the "author_account" field.
func (fcuo *FeedConfigUpdateOne) SetAuthorAccount(i int64) *FeedConfigUpdateOne {
	fcuo.mutation.ResetAuthorAccount()
	fcuo.mutation.SetAuthorAccount(i)
	return fcuo
}

// AddAuthorAccount adds i to the "author_account" field.
func (fcuo *FeedConfigUpdateOne) AddAuthorAccount(i int64) *FeedConfigUpdateOne {
	fcuo.mutation.AddAuthorAccount(i)
	return fcuo
}

// SetSource sets the "source" field.
func (fcuo *FeedConfigUpdateOne) SetSource(f feedconfig.Source) *FeedConfigUpdateOne {
	fcuo.mutation.SetSource(f)
	return fcuo
}

// SetStatus sets the "status" field.
func (fcuo *FeedConfigUpdateOne) SetStatus(f feedconfig.Status) *FeedConfigUpdateOne {
	fcuo.mutation.SetStatus(f)
	return fcuo
}

// SetPullInterval sets the "pull_interval" field.
func (fcuo *FeedConfigUpdateOne) SetPullInterval(t time.Duration) *FeedConfigUpdateOne {
	fcuo.mutation.ResetPullInterval()
	fcuo.mutation.SetPullInterval(t)
	return fcuo
}

// AddPullInterval adds t to the "pull_interval" field.
func (fcuo *FeedConfigUpdateOne) AddPullInterval(t time.Duration) *FeedConfigUpdateOne {
	fcuo.mutation.AddPullInterval(t)
	return fcuo
}

// SetNextPullBeginAt sets the "next_pull_begin_at" field.
func (fcuo *FeedConfigUpdateOne) SetNextPullBeginAt(t time.Time) *FeedConfigUpdateOne {
	fcuo.mutation.SetNextPullBeginAt(t)
	return fcuo
}

// SetNillableNextPullBeginAt sets the "next_pull_begin_at" field if the given value is not nil.
func (fcuo *FeedConfigUpdateOne) SetNillableNextPullBeginAt(t *time.Time) *FeedConfigUpdateOne {
	if t != nil {
		fcuo.SetNextPullBeginAt(*t)
	}
	return fcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fcuo *FeedConfigUpdateOne) SetUpdatedAt(t time.Time) *FeedConfigUpdateOne {
	fcuo.mutation.SetUpdatedAt(t)
	return fcuo
}

// SetCreatedAt sets the "created_at" field.
func (fcuo *FeedConfigUpdateOne) SetCreatedAt(t time.Time) *FeedConfigUpdateOne {
	fcuo.mutation.SetCreatedAt(t)
	return fcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcuo *FeedConfigUpdateOne) SetNillableCreatedAt(t *time.Time) *FeedConfigUpdateOne {
	if t != nil {
		fcuo.SetCreatedAt(*t)
	}
	return fcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fcuo *FeedConfigUpdateOne) SetUserID(id int64) *FeedConfigUpdateOne {
	fcuo.mutation.SetUserID(id)
	return fcuo
}

// SetUser sets the "user" edge to the User entity.
func (fcuo *FeedConfigUpdateOne) SetUser(u *User) *FeedConfigUpdateOne {
	return fcuo.SetUserID(u.ID)
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fcuo *FeedConfigUpdateOne) SetFeedID(id int64) *FeedConfigUpdateOne {
	fcuo.mutation.SetFeedID(id)
	return fcuo
}

// SetNillableFeedID sets the "feed" edge to the Feed entity by ID if the given value is not nil.
func (fcuo *FeedConfigUpdateOne) SetNillableFeedID(id *int64) *FeedConfigUpdateOne {
	if id != nil {
		fcuo = fcuo.SetFeedID(*id)
	}
	return fcuo
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fcuo *FeedConfigUpdateOne) SetFeed(f *Feed) *FeedConfigUpdateOne {
	return fcuo.SetFeedID(f.ID)
}

// Mutation returns the FeedConfigMutation object of the builder.
func (fcuo *FeedConfigUpdateOne) Mutation() *FeedConfigMutation {
	return fcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fcuo *FeedConfigUpdateOne) ClearUser() *FeedConfigUpdateOne {
	fcuo.mutation.ClearUser()
	return fcuo
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (fcuo *FeedConfigUpdateOne) ClearFeed() *FeedConfigUpdateOne {
	fcuo.mutation.ClearFeed()
	return fcuo
}

// Where appends a list predicates to the FeedConfigUpdate builder.
func (fcuo *FeedConfigUpdateOne) Where(ps ...predicate.FeedConfig) *FeedConfigUpdateOne {
	fcuo.mutation.Where(ps...)
	return fcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fcuo *FeedConfigUpdateOne) Select(field string, fields ...string) *FeedConfigUpdateOne {
	fcuo.fields = append([]string{field}, fields...)
	return fcuo
}

// Save executes the query and returns the updated FeedConfig entity.
func (fcuo *FeedConfigUpdateOne) Save(ctx context.Context) (*FeedConfig, error) {
	fcuo.defaults()
	return withHooks[*FeedConfig, FeedConfigMutation](ctx, fcuo.sqlSave, fcuo.mutation, fcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fcuo *FeedConfigUpdateOne) SaveX(ctx context.Context) *FeedConfig {
	node, err := fcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fcuo *FeedConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := fcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcuo *FeedConfigUpdateOne) ExecX(ctx context.Context) {
	if err := fcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcuo *FeedConfigUpdateOne) defaults() {
	if _, ok := fcuo.mutation.UpdatedAt(); !ok {
		v := feedconfig.UpdateDefaultUpdatedAt()
		fcuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fcuo *FeedConfigUpdateOne) check() error {
	if v, ok := fcuo.mutation.Source(); ok {
		if err := feedconfig.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.source": %w`, err)}
		}
	}
	if v, ok := fcuo.mutation.Status(); ok {
		if err := feedconfig.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "FeedConfig.status": %w`, err)}
		}
	}
	if _, ok := fcuo.mutation.UserID(); fcuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedConfig.user"`)
	}
	return nil
}

func (fcuo *FeedConfigUpdateOne) sqlSave(ctx context.Context) (_node *FeedConfig, err error) {
	if err := fcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feedconfig.Table, feedconfig.Columns, sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64))
	id, ok := fcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FeedConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedconfig.FieldID)
		for _, f := range fields {
			if !feedconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feedconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcuo.mutation.FeedURL(); ok {
		_spec.SetField(feedconfig.FieldFeedURL, field.TypeString, value)
	}
	if value, ok := fcuo.mutation.AuthorAccount(); ok {
		_spec.SetField(feedconfig.FieldAuthorAccount, field.TypeInt64, value)
	}
	if value, ok := fcuo.mutation.AddedAuthorAccount(); ok {
		_spec.AddField(feedconfig.FieldAuthorAccount, field.TypeInt64, value)
	}
	if value, ok := fcuo.mutation.Source(); ok {
		_spec.SetField(feedconfig.FieldSource, field.TypeEnum, value)
	}
	if value, ok := fcuo.mutation.Status(); ok {
		_spec.SetField(feedconfig.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := fcuo.mutation.PullInterval(); ok {
		_spec.SetField(feedconfig.FieldPullInterval, field.TypeInt64, value)
	}
	if value, ok := fcuo.mutation.AddedPullInterval(); ok {
		_spec.AddField(feedconfig.FieldPullInterval, field.TypeInt64, value)
	}
	if value, ok := fcuo.mutation.NextPullBeginAt(); ok {
		_spec.SetField(feedconfig.FieldNextPullBeginAt, field.TypeTime, value)
	}
	if value, ok := fcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feedconfig.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fcuo.mutation.CreatedAt(); ok {
		_spec.SetField(feedconfig.FieldCreatedAt, field.TypeTime, value)
	}
	if fcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedconfig.UserTable,
			Columns: []string{feedconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedconfig.UserTable,
			Columns: []string{feedconfig.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fcuo.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   feedconfig.FeedTable,
			Columns: []string{feedconfig.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: feed.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fcuo.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   feedconfig.FeedTable,
			Columns: []string{feedconfig.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: feed.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FeedConfig{config: fcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feedconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fcuo.mutation.done = true
	return _node, nil
}