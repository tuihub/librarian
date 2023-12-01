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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/apppackageruntime"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppPackageRunTimeUpdate is the builder for updating AppPackageRunTime entities.
type AppPackageRunTimeUpdate struct {
	config
	hooks    []Hook
	mutation *AppPackageRunTimeMutation
}

// Where appends a list predicates to the AppPackageRunTimeUpdate builder.
func (aprtu *AppPackageRunTimeUpdate) Where(ps ...predicate.AppPackageRunTime) *AppPackageRunTimeUpdate {
	aprtu.mutation.Where(ps...)
	return aprtu
}

// SetUserID sets the "user_id" field.
func (aprtu *AppPackageRunTimeUpdate) SetUserID(mi model.InternalID) *AppPackageRunTimeUpdate {
	aprtu.mutation.ResetUserID()
	aprtu.mutation.SetUserID(mi)
	return aprtu
}

// AddUserID adds mi to the "user_id" field.
func (aprtu *AppPackageRunTimeUpdate) AddUserID(mi model.InternalID) *AppPackageRunTimeUpdate {
	aprtu.mutation.AddUserID(mi)
	return aprtu
}

// SetAppPackageID sets the "app_package_id" field.
func (aprtu *AppPackageRunTimeUpdate) SetAppPackageID(mi model.InternalID) *AppPackageRunTimeUpdate {
	aprtu.mutation.ResetAppPackageID()
	aprtu.mutation.SetAppPackageID(mi)
	return aprtu
}

// AddAppPackageID adds mi to the "app_package_id" field.
func (aprtu *AppPackageRunTimeUpdate) AddAppPackageID(mi model.InternalID) *AppPackageRunTimeUpdate {
	aprtu.mutation.AddAppPackageID(mi)
	return aprtu
}

// SetStartTime sets the "start_time" field.
func (aprtu *AppPackageRunTimeUpdate) SetStartTime(t time.Time) *AppPackageRunTimeUpdate {
	aprtu.mutation.SetStartTime(t)
	return aprtu
}

// SetRunDuration sets the "run_duration" field.
func (aprtu *AppPackageRunTimeUpdate) SetRunDuration(t time.Duration) *AppPackageRunTimeUpdate {
	aprtu.mutation.ResetRunDuration()
	aprtu.mutation.SetRunDuration(t)
	return aprtu
}

// AddRunDuration adds t to the "run_duration" field.
func (aprtu *AppPackageRunTimeUpdate) AddRunDuration(t time.Duration) *AppPackageRunTimeUpdate {
	aprtu.mutation.AddRunDuration(t)
	return aprtu
}

// SetUpdatedAt sets the "updated_at" field.
func (aprtu *AppPackageRunTimeUpdate) SetUpdatedAt(t time.Time) *AppPackageRunTimeUpdate {
	aprtu.mutation.SetUpdatedAt(t)
	return aprtu
}

// SetCreatedAt sets the "created_at" field.
func (aprtu *AppPackageRunTimeUpdate) SetCreatedAt(t time.Time) *AppPackageRunTimeUpdate {
	aprtu.mutation.SetCreatedAt(t)
	return aprtu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aprtu *AppPackageRunTimeUpdate) SetNillableCreatedAt(t *time.Time) *AppPackageRunTimeUpdate {
	if t != nil {
		aprtu.SetCreatedAt(*t)
	}
	return aprtu
}

// Mutation returns the AppPackageRunTimeMutation object of the builder.
func (aprtu *AppPackageRunTimeUpdate) Mutation() *AppPackageRunTimeMutation {
	return aprtu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aprtu *AppPackageRunTimeUpdate) Save(ctx context.Context) (int, error) {
	aprtu.defaults()
	return withHooks(ctx, aprtu.sqlSave, aprtu.mutation, aprtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aprtu *AppPackageRunTimeUpdate) SaveX(ctx context.Context) int {
	affected, err := aprtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aprtu *AppPackageRunTimeUpdate) Exec(ctx context.Context) error {
	_, err := aprtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aprtu *AppPackageRunTimeUpdate) ExecX(ctx context.Context) {
	if err := aprtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aprtu *AppPackageRunTimeUpdate) defaults() {
	if _, ok := aprtu.mutation.UpdatedAt(); !ok {
		v := apppackageruntime.UpdateDefaultUpdatedAt()
		aprtu.mutation.SetUpdatedAt(v)
	}
}

func (aprtu *AppPackageRunTimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apppackageruntime.Table, apppackageruntime.Columns, sqlgraph.NewFieldSpec(apppackageruntime.FieldID, field.TypeInt))
	if ps := aprtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aprtu.mutation.UserID(); ok {
		_spec.SetField(apppackageruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.AddedUserID(); ok {
		_spec.AddField(apppackageruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.AppPackageID(); ok {
		_spec.SetField(apppackageruntime.FieldAppPackageID, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.AddedAppPackageID(); ok {
		_spec.AddField(apppackageruntime.FieldAppPackageID, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.StartTime(); ok {
		_spec.SetField(apppackageruntime.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := aprtu.mutation.RunDuration(); ok {
		_spec.SetField(apppackageruntime.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.AddedRunDuration(); ok {
		_spec.AddField(apppackageruntime.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := aprtu.mutation.UpdatedAt(); ok {
		_spec.SetField(apppackageruntime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aprtu.mutation.CreatedAt(); ok {
		_spec.SetField(apppackageruntime.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aprtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apppackageruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aprtu.mutation.done = true
	return n, nil
}

// AppPackageRunTimeUpdateOne is the builder for updating a single AppPackageRunTime entity.
type AppPackageRunTimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppPackageRunTimeMutation
}

// SetUserID sets the "user_id" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetUserID(mi model.InternalID) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.ResetUserID()
	aprtuo.mutation.SetUserID(mi)
	return aprtuo
}

// AddUserID adds mi to the "user_id" field.
func (aprtuo *AppPackageRunTimeUpdateOne) AddUserID(mi model.InternalID) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.AddUserID(mi)
	return aprtuo
}

// SetAppPackageID sets the "app_package_id" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetAppPackageID(mi model.InternalID) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.ResetAppPackageID()
	aprtuo.mutation.SetAppPackageID(mi)
	return aprtuo
}

// AddAppPackageID adds mi to the "app_package_id" field.
func (aprtuo *AppPackageRunTimeUpdateOne) AddAppPackageID(mi model.InternalID) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.AddAppPackageID(mi)
	return aprtuo
}

// SetStartTime sets the "start_time" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetStartTime(t time.Time) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.SetStartTime(t)
	return aprtuo
}

// SetRunDuration sets the "run_duration" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetRunDuration(t time.Duration) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.ResetRunDuration()
	aprtuo.mutation.SetRunDuration(t)
	return aprtuo
}

// AddRunDuration adds t to the "run_duration" field.
func (aprtuo *AppPackageRunTimeUpdateOne) AddRunDuration(t time.Duration) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.AddRunDuration(t)
	return aprtuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetUpdatedAt(t time.Time) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.SetUpdatedAt(t)
	return aprtuo
}

// SetCreatedAt sets the "created_at" field.
func (aprtuo *AppPackageRunTimeUpdateOne) SetCreatedAt(t time.Time) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.SetCreatedAt(t)
	return aprtuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aprtuo *AppPackageRunTimeUpdateOne) SetNillableCreatedAt(t *time.Time) *AppPackageRunTimeUpdateOne {
	if t != nil {
		aprtuo.SetCreatedAt(*t)
	}
	return aprtuo
}

// Mutation returns the AppPackageRunTimeMutation object of the builder.
func (aprtuo *AppPackageRunTimeUpdateOne) Mutation() *AppPackageRunTimeMutation {
	return aprtuo.mutation
}

// Where appends a list predicates to the AppPackageRunTimeUpdate builder.
func (aprtuo *AppPackageRunTimeUpdateOne) Where(ps ...predicate.AppPackageRunTime) *AppPackageRunTimeUpdateOne {
	aprtuo.mutation.Where(ps...)
	return aprtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aprtuo *AppPackageRunTimeUpdateOne) Select(field string, fields ...string) *AppPackageRunTimeUpdateOne {
	aprtuo.fields = append([]string{field}, fields...)
	return aprtuo
}

// Save executes the query and returns the updated AppPackageRunTime entity.
func (aprtuo *AppPackageRunTimeUpdateOne) Save(ctx context.Context) (*AppPackageRunTime, error) {
	aprtuo.defaults()
	return withHooks(ctx, aprtuo.sqlSave, aprtuo.mutation, aprtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aprtuo *AppPackageRunTimeUpdateOne) SaveX(ctx context.Context) *AppPackageRunTime {
	node, err := aprtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aprtuo *AppPackageRunTimeUpdateOne) Exec(ctx context.Context) error {
	_, err := aprtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aprtuo *AppPackageRunTimeUpdateOne) ExecX(ctx context.Context) {
	if err := aprtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aprtuo *AppPackageRunTimeUpdateOne) defaults() {
	if _, ok := aprtuo.mutation.UpdatedAt(); !ok {
		v := apppackageruntime.UpdateDefaultUpdatedAt()
		aprtuo.mutation.SetUpdatedAt(v)
	}
}

func (aprtuo *AppPackageRunTimeUpdateOne) sqlSave(ctx context.Context) (_node *AppPackageRunTime, err error) {
	_spec := sqlgraph.NewUpdateSpec(apppackageruntime.Table, apppackageruntime.Columns, sqlgraph.NewFieldSpec(apppackageruntime.FieldID, field.TypeInt))
	id, ok := aprtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppPackageRunTime.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aprtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apppackageruntime.FieldID)
		for _, f := range fields {
			if !apppackageruntime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apppackageruntime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aprtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aprtuo.mutation.UserID(); ok {
		_spec.SetField(apppackageruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.AddedUserID(); ok {
		_spec.AddField(apppackageruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.AppPackageID(); ok {
		_spec.SetField(apppackageruntime.FieldAppPackageID, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.AddedAppPackageID(); ok {
		_spec.AddField(apppackageruntime.FieldAppPackageID, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.StartTime(); ok {
		_spec.SetField(apppackageruntime.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := aprtuo.mutation.RunDuration(); ok {
		_spec.SetField(apppackageruntime.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.AddedRunDuration(); ok {
		_spec.AddField(apppackageruntime.FieldRunDuration, field.TypeInt64, value)
	}
	if value, ok := aprtuo.mutation.UpdatedAt(); ok {
		_spec.SetField(apppackageruntime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aprtuo.mutation.CreatedAt(); ok {
		_spec.SetField(apppackageruntime.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &AppPackageRunTime{config: aprtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aprtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apppackageruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aprtuo.mutation.done = true
	return _node, nil
}