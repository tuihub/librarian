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
	"github.com/tuihub/librarian/internal/data/internal/ent/appcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppCategoryUpdate is the builder for updating AppCategory entities.
type AppCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *AppCategoryMutation
}

// Where appends a list predicates to the AppCategoryUpdate builder.
func (acu *AppCategoryUpdate) Where(ps ...predicate.AppCategory) *AppCategoryUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetUserID sets the "user_id" field.
func (acu *AppCategoryUpdate) SetUserID(mi model.InternalID) *AppCategoryUpdate {
	acu.mutation.ResetUserID()
	acu.mutation.SetUserID(mi)
	return acu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (acu *AppCategoryUpdate) SetNillableUserID(mi *model.InternalID) *AppCategoryUpdate {
	if mi != nil {
		acu.SetUserID(*mi)
	}
	return acu
}

// AddUserID adds mi to the "user_id" field.
func (acu *AppCategoryUpdate) AddUserID(mi model.InternalID) *AppCategoryUpdate {
	acu.mutation.AddUserID(mi)
	return acu
}

// SetAppID sets the "app_id" field.
func (acu *AppCategoryUpdate) SetAppID(mi model.InternalID) *AppCategoryUpdate {
	acu.mutation.ResetAppID()
	acu.mutation.SetAppID(mi)
	return acu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (acu *AppCategoryUpdate) SetNillableAppID(mi *model.InternalID) *AppCategoryUpdate {
	if mi != nil {
		acu.SetAppID(*mi)
	}
	return acu
}

// AddAppID adds mi to the "app_id" field.
func (acu *AppCategoryUpdate) AddAppID(mi model.InternalID) *AppCategoryUpdate {
	acu.mutation.AddAppID(mi)
	return acu
}

// SetStartTime sets the "start_time" field.
func (acu *AppCategoryUpdate) SetStartTime(t time.Time) *AppCategoryUpdate {
	acu.mutation.SetStartTime(t)
	return acu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (acu *AppCategoryUpdate) SetNillableStartTime(t *time.Time) *AppCategoryUpdate {
	if t != nil {
		acu.SetStartTime(*t)
	}
	return acu
}

// SetRunDuration sets the "run_duration" field.
func (acu *AppCategoryUpdate) SetRunDuration(t time.Duration) *AppCategoryUpdate {
	acu.mutation.ResetRunDuration()
	acu.mutation.SetRunDuration(t)
	return acu
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (acu *AppCategoryUpdate) SetNillableRunDuration(t *time.Duration) *AppCategoryUpdate {
	if t != nil {
		acu.SetRunDuration(*t)
	}
	return acu
}

// AddRunDuration adds t to the "run_duration" field.
func (acu *AppCategoryUpdate) AddRunDuration(t time.Duration) *AppCategoryUpdate {
	acu.mutation.AddRunDuration(t)
	return acu
}

// SetUpdatedAt sets the "updated_at" field.
func (acu *AppCategoryUpdate) SetUpdatedAt(t time.Time) *AppCategoryUpdate {
	acu.mutation.SetUpdatedAt(t)
	return acu
}

// SetCreatedAt sets the "created_at" field.
func (acu *AppCategoryUpdate) SetCreatedAt(t time.Time) *AppCategoryUpdate {
	acu.mutation.SetCreatedAt(t)
	return acu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acu *AppCategoryUpdate) SetNillableCreatedAt(t *time.Time) *AppCategoryUpdate {
	if t != nil {
		acu.SetCreatedAt(*t)
	}
	return acu
}

// Mutation returns the AppCategoryMutation object of the builder.
func (acu *AppCategoryUpdate) Mutation() *AppCategoryMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AppCategoryUpdate) Save(ctx context.Context) (int, error) {
	acu.defaults()
	return withHooks(ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AppCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AppCategoryUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AppCategoryUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acu *AppCategoryUpdate) defaults() {
	if _, ok := acu.mutation.UpdatedAt(); !ok {
		v := appcategory.UpdateDefaultUpdatedAt()
		acu.mutation.SetUpdatedAt(v)
	}
}

func (acu *AppCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appcategory.Table, appcategory.Columns, sqlgraph.NewFieldSpec(appcategory.FieldID, field.TypeInt))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.UserID(); ok {
		_spec.SetField(appcategory.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.AddedUserID(); ok {
		_spec.AddField(appcategory.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.AppID(); ok {
		_spec.SetField(appcategory.FieldAppID, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.AddedAppID(); ok {
		_spec.AddField(appcategory.FieldAppID, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.StartTime(); ok {
		_spec.SetField(appcategory.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := acu.mutation.RunDuration(); ok {
		_spec.SetField(appcategory.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.AddedRunDuration(); ok {
		_spec.AddField(appcategory.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := acu.mutation.UpdatedAt(); ok {
		_spec.SetField(appcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := acu.mutation.CreatedAt(); ok {
		_spec.SetField(appcategory.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// AppCategoryUpdateOne is the builder for updating a single AppCategory entity.
type AppCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppCategoryMutation
}

// SetUserID sets the "user_id" field.
func (acuo *AppCategoryUpdateOne) SetUserID(mi model.InternalID) *AppCategoryUpdateOne {
	acuo.mutation.ResetUserID()
	acuo.mutation.SetUserID(mi)
	return acuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (acuo *AppCategoryUpdateOne) SetNillableUserID(mi *model.InternalID) *AppCategoryUpdateOne {
	if mi != nil {
		acuo.SetUserID(*mi)
	}
	return acuo
}

// AddUserID adds mi to the "user_id" field.
func (acuo *AppCategoryUpdateOne) AddUserID(mi model.InternalID) *AppCategoryUpdateOne {
	acuo.mutation.AddUserID(mi)
	return acuo
}

// SetAppID sets the "app_id" field.
func (acuo *AppCategoryUpdateOne) SetAppID(mi model.InternalID) *AppCategoryUpdateOne {
	acuo.mutation.ResetAppID()
	acuo.mutation.SetAppID(mi)
	return acuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (acuo *AppCategoryUpdateOne) SetNillableAppID(mi *model.InternalID) *AppCategoryUpdateOne {
	if mi != nil {
		acuo.SetAppID(*mi)
	}
	return acuo
}

// AddAppID adds mi to the "app_id" field.
func (acuo *AppCategoryUpdateOne) AddAppID(mi model.InternalID) *AppCategoryUpdateOne {
	acuo.mutation.AddAppID(mi)
	return acuo
}

// SetStartTime sets the "start_time" field.
func (acuo *AppCategoryUpdateOne) SetStartTime(t time.Time) *AppCategoryUpdateOne {
	acuo.mutation.SetStartTime(t)
	return acuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (acuo *AppCategoryUpdateOne) SetNillableStartTime(t *time.Time) *AppCategoryUpdateOne {
	if t != nil {
		acuo.SetStartTime(*t)
	}
	return acuo
}

// SetRunDuration sets the "run_duration" field.
func (acuo *AppCategoryUpdateOne) SetRunDuration(t time.Duration) *AppCategoryUpdateOne {
	acuo.mutation.ResetRunDuration()
	acuo.mutation.SetRunDuration(t)
	return acuo
}

// SetNillableRunDuration sets the "run_duration" field if the given value is not nil.
func (acuo *AppCategoryUpdateOne) SetNillableRunDuration(t *time.Duration) *AppCategoryUpdateOne {
	if t != nil {
		acuo.SetRunDuration(*t)
	}
	return acuo
}

// AddRunDuration adds t to the "run_duration" field.
func (acuo *AppCategoryUpdateOne) AddRunDuration(t time.Duration) *AppCategoryUpdateOne {
	acuo.mutation.AddRunDuration(t)
	return acuo
}

// SetUpdatedAt sets the "updated_at" field.
func (acuo *AppCategoryUpdateOne) SetUpdatedAt(t time.Time) *AppCategoryUpdateOne {
	acuo.mutation.SetUpdatedAt(t)
	return acuo
}

// SetCreatedAt sets the "created_at" field.
func (acuo *AppCategoryUpdateOne) SetCreatedAt(t time.Time) *AppCategoryUpdateOne {
	acuo.mutation.SetCreatedAt(t)
	return acuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acuo *AppCategoryUpdateOne) SetNillableCreatedAt(t *time.Time) *AppCategoryUpdateOne {
	if t != nil {
		acuo.SetCreatedAt(*t)
	}
	return acuo
}

// Mutation returns the AppCategoryMutation object of the builder.
func (acuo *AppCategoryUpdateOne) Mutation() *AppCategoryMutation {
	return acuo.mutation
}

// Where appends a list predicates to the AppCategoryUpdate builder.
func (acuo *AppCategoryUpdateOne) Where(ps ...predicate.AppCategory) *AppCategoryUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AppCategoryUpdateOne) Select(field string, fields ...string) *AppCategoryUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AppCategory entity.
func (acuo *AppCategoryUpdateOne) Save(ctx context.Context) (*AppCategory, error) {
	acuo.defaults()
	return withHooks(ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AppCategoryUpdateOne) SaveX(ctx context.Context) *AppCategory {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AppCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AppCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acuo *AppCategoryUpdateOne) defaults() {
	if _, ok := acuo.mutation.UpdatedAt(); !ok {
		v := appcategory.UpdateDefaultUpdatedAt()
		acuo.mutation.SetUpdatedAt(v)
	}
}

func (acuo *AppCategoryUpdateOne) sqlSave(ctx context.Context) (_node *AppCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(appcategory.Table, appcategory.Columns, sqlgraph.NewFieldSpec(appcategory.FieldID, field.TypeInt))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcategory.FieldID)
		for _, f := range fields {
			if !appcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.UserID(); ok {
		_spec.SetField(appcategory.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.AddedUserID(); ok {
		_spec.AddField(appcategory.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.AppID(); ok {
		_spec.SetField(appcategory.FieldAppID, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.AddedAppID(); ok {
		_spec.AddField(appcategory.FieldAppID, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.StartTime(); ok {
		_spec.SetField(appcategory.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := acuo.mutation.RunDuration(); ok {
		_spec.SetField(appcategory.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.AddedRunDuration(); ok {
		_spec.AddField(appcategory.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := acuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := acuo.mutation.CreatedAt(); ok {
		_spec.SetField(appcategory.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &AppCategory{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}
