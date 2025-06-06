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
	"github.com/tuihub/librarian/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/internal/data/internal/ent/appruntime"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppRunTimeUpdate is the builder for updating AppRunTime entities.
type AppRunTimeUpdate struct {
	config
	hooks    []Hook
	mutation *AppRunTimeMutation
}

// Where appends a list predicates to the AppRunTimeUpdate builder.
func (artu *AppRunTimeUpdate) Where(ps ...predicate.AppRunTime) *AppRunTimeUpdate {
	artu.mutation.Where(ps...)
	return artu
}

// SetUserID sets the "user_id" field.
func (artu *AppRunTimeUpdate) SetUserID(mi model.InternalID) *AppRunTimeUpdate {
	artu.mutation.ResetUserID()
	artu.mutation.SetUserID(mi)
	return artu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableUserID(mi *model.InternalID) *AppRunTimeUpdate {
	if mi != nil {
		artu.SetUserID(*mi)
	}
	return artu
}

// AddUserID adds mi to the "user_id" field.
func (artu *AppRunTimeUpdate) AddUserID(mi model.InternalID) *AppRunTimeUpdate {
	artu.mutation.AddUserID(mi)
	return artu
}

// SetAppID sets the "app_id" field.
func (artu *AppRunTimeUpdate) SetAppID(mi model.InternalID) *AppRunTimeUpdate {
	artu.mutation.SetAppID(mi)
	return artu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableAppID(mi *model.InternalID) *AppRunTimeUpdate {
	if mi != nil {
		artu.SetAppID(*mi)
	}
	return artu
}

// SetDeviceID sets the "device_id" field.
func (artu *AppRunTimeUpdate) SetDeviceID(mi model.InternalID) *AppRunTimeUpdate {
	artu.mutation.ResetDeviceID()
	artu.mutation.SetDeviceID(mi)
	return artu
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableDeviceID(mi *model.InternalID) *AppRunTimeUpdate {
	if mi != nil {
		artu.SetDeviceID(*mi)
	}
	return artu
}

// AddDeviceID adds mi to the "device_id" field.
func (artu *AppRunTimeUpdate) AddDeviceID(mi model.InternalID) *AppRunTimeUpdate {
	artu.mutation.AddDeviceID(mi)
	return artu
}

// SetStartTime sets the "start_time" field.
func (artu *AppRunTimeUpdate) SetStartTime(t time.Time) *AppRunTimeUpdate {
	artu.mutation.SetStartTime(t)
	return artu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableStartTime(t *time.Time) *AppRunTimeUpdate {
	if t != nil {
		artu.SetStartTime(*t)
	}
	return artu
}

// SetDuration sets the "duration" field.
func (artu *AppRunTimeUpdate) SetDuration(t time.Duration) *AppRunTimeUpdate {
	artu.mutation.ResetDuration()
	artu.mutation.SetDuration(t)
	return artu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableDuration(t *time.Duration) *AppRunTimeUpdate {
	if t != nil {
		artu.SetDuration(*t)
	}
	return artu
}

// AddDuration adds t to the "duration" field.
func (artu *AppRunTimeUpdate) AddDuration(t time.Duration) *AppRunTimeUpdate {
	artu.mutation.AddDuration(t)
	return artu
}

// SetUpdatedAt sets the "updated_at" field.
func (artu *AppRunTimeUpdate) SetUpdatedAt(t time.Time) *AppRunTimeUpdate {
	artu.mutation.SetUpdatedAt(t)
	return artu
}

// SetCreatedAt sets the "created_at" field.
func (artu *AppRunTimeUpdate) SetCreatedAt(t time.Time) *AppRunTimeUpdate {
	artu.mutation.SetCreatedAt(t)
	return artu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (artu *AppRunTimeUpdate) SetNillableCreatedAt(t *time.Time) *AppRunTimeUpdate {
	if t != nil {
		artu.SetCreatedAt(*t)
	}
	return artu
}

// SetApp sets the "app" edge to the App entity.
func (artu *AppRunTimeUpdate) SetApp(a *App) *AppRunTimeUpdate {
	return artu.SetAppID(a.ID)
}

// Mutation returns the AppRunTimeMutation object of the builder.
func (artu *AppRunTimeUpdate) Mutation() *AppRunTimeMutation {
	return artu.mutation
}

// ClearApp clears the "app" edge to the App entity.
func (artu *AppRunTimeUpdate) ClearApp() *AppRunTimeUpdate {
	artu.mutation.ClearApp()
	return artu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (artu *AppRunTimeUpdate) Save(ctx context.Context) (int, error) {
	artu.defaults()
	return withHooks(ctx, artu.sqlSave, artu.mutation, artu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (artu *AppRunTimeUpdate) SaveX(ctx context.Context) int {
	affected, err := artu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (artu *AppRunTimeUpdate) Exec(ctx context.Context) error {
	_, err := artu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (artu *AppRunTimeUpdate) ExecX(ctx context.Context) {
	if err := artu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (artu *AppRunTimeUpdate) defaults() {
	if _, ok := artu.mutation.UpdatedAt(); !ok {
		v := appruntime.UpdateDefaultUpdatedAt()
		artu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (artu *AppRunTimeUpdate) check() error {
	if artu.mutation.AppCleared() && len(artu.mutation.AppIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "AppRunTime.app"`)
	}
	return nil
}

func (artu *AppRunTimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := artu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(appruntime.Table, appruntime.Columns, sqlgraph.NewFieldSpec(appruntime.FieldID, field.TypeInt64))
	if ps := artu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := artu.mutation.UserID(); ok {
		_spec.SetField(appruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.AddedUserID(); ok {
		_spec.AddField(appruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.DeviceID(); ok {
		_spec.SetField(appruntime.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.AddedDeviceID(); ok {
		_spec.AddField(appruntime.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.StartTime(); ok {
		_spec.SetField(appruntime.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := artu.mutation.Duration(); ok {
		_spec.SetField(appruntime.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.AddedDuration(); ok {
		_spec.AddField(appruntime.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := artu.mutation.UpdatedAt(); ok {
		_spec.SetField(appruntime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := artu.mutation.CreatedAt(); ok {
		_spec.SetField(appruntime.FieldCreatedAt, field.TypeTime, value)
	}
	if artu.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appruntime.AppTable,
			Columns: []string{appruntime.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := artu.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appruntime.AppTable,
			Columns: []string{appruntime.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, artu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	artu.mutation.done = true
	return n, nil
}

// AppRunTimeUpdateOne is the builder for updating a single AppRunTime entity.
type AppRunTimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppRunTimeMutation
}

// SetUserID sets the "user_id" field.
func (artuo *AppRunTimeUpdateOne) SetUserID(mi model.InternalID) *AppRunTimeUpdateOne {
	artuo.mutation.ResetUserID()
	artuo.mutation.SetUserID(mi)
	return artuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableUserID(mi *model.InternalID) *AppRunTimeUpdateOne {
	if mi != nil {
		artuo.SetUserID(*mi)
	}
	return artuo
}

// AddUserID adds mi to the "user_id" field.
func (artuo *AppRunTimeUpdateOne) AddUserID(mi model.InternalID) *AppRunTimeUpdateOne {
	artuo.mutation.AddUserID(mi)
	return artuo
}

// SetAppID sets the "app_id" field.
func (artuo *AppRunTimeUpdateOne) SetAppID(mi model.InternalID) *AppRunTimeUpdateOne {
	artuo.mutation.SetAppID(mi)
	return artuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableAppID(mi *model.InternalID) *AppRunTimeUpdateOne {
	if mi != nil {
		artuo.SetAppID(*mi)
	}
	return artuo
}

// SetDeviceID sets the "device_id" field.
func (artuo *AppRunTimeUpdateOne) SetDeviceID(mi model.InternalID) *AppRunTimeUpdateOne {
	artuo.mutation.ResetDeviceID()
	artuo.mutation.SetDeviceID(mi)
	return artuo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableDeviceID(mi *model.InternalID) *AppRunTimeUpdateOne {
	if mi != nil {
		artuo.SetDeviceID(*mi)
	}
	return artuo
}

// AddDeviceID adds mi to the "device_id" field.
func (artuo *AppRunTimeUpdateOne) AddDeviceID(mi model.InternalID) *AppRunTimeUpdateOne {
	artuo.mutation.AddDeviceID(mi)
	return artuo
}

// SetStartTime sets the "start_time" field.
func (artuo *AppRunTimeUpdateOne) SetStartTime(t time.Time) *AppRunTimeUpdateOne {
	artuo.mutation.SetStartTime(t)
	return artuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableStartTime(t *time.Time) *AppRunTimeUpdateOne {
	if t != nil {
		artuo.SetStartTime(*t)
	}
	return artuo
}

// SetDuration sets the "duration" field.
func (artuo *AppRunTimeUpdateOne) SetDuration(t time.Duration) *AppRunTimeUpdateOne {
	artuo.mutation.ResetDuration()
	artuo.mutation.SetDuration(t)
	return artuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableDuration(t *time.Duration) *AppRunTimeUpdateOne {
	if t != nil {
		artuo.SetDuration(*t)
	}
	return artuo
}

// AddDuration adds t to the "duration" field.
func (artuo *AppRunTimeUpdateOne) AddDuration(t time.Duration) *AppRunTimeUpdateOne {
	artuo.mutation.AddDuration(t)
	return artuo
}

// SetUpdatedAt sets the "updated_at" field.
func (artuo *AppRunTimeUpdateOne) SetUpdatedAt(t time.Time) *AppRunTimeUpdateOne {
	artuo.mutation.SetUpdatedAt(t)
	return artuo
}

// SetCreatedAt sets the "created_at" field.
func (artuo *AppRunTimeUpdateOne) SetCreatedAt(t time.Time) *AppRunTimeUpdateOne {
	artuo.mutation.SetCreatedAt(t)
	return artuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (artuo *AppRunTimeUpdateOne) SetNillableCreatedAt(t *time.Time) *AppRunTimeUpdateOne {
	if t != nil {
		artuo.SetCreatedAt(*t)
	}
	return artuo
}

// SetApp sets the "app" edge to the App entity.
func (artuo *AppRunTimeUpdateOne) SetApp(a *App) *AppRunTimeUpdateOne {
	return artuo.SetAppID(a.ID)
}

// Mutation returns the AppRunTimeMutation object of the builder.
func (artuo *AppRunTimeUpdateOne) Mutation() *AppRunTimeMutation {
	return artuo.mutation
}

// ClearApp clears the "app" edge to the App entity.
func (artuo *AppRunTimeUpdateOne) ClearApp() *AppRunTimeUpdateOne {
	artuo.mutation.ClearApp()
	return artuo
}

// Where appends a list predicates to the AppRunTimeUpdate builder.
func (artuo *AppRunTimeUpdateOne) Where(ps ...predicate.AppRunTime) *AppRunTimeUpdateOne {
	artuo.mutation.Where(ps...)
	return artuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (artuo *AppRunTimeUpdateOne) Select(field string, fields ...string) *AppRunTimeUpdateOne {
	artuo.fields = append([]string{field}, fields...)
	return artuo
}

// Save executes the query and returns the updated AppRunTime entity.
func (artuo *AppRunTimeUpdateOne) Save(ctx context.Context) (*AppRunTime, error) {
	artuo.defaults()
	return withHooks(ctx, artuo.sqlSave, artuo.mutation, artuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (artuo *AppRunTimeUpdateOne) SaveX(ctx context.Context) *AppRunTime {
	node, err := artuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (artuo *AppRunTimeUpdateOne) Exec(ctx context.Context) error {
	_, err := artuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (artuo *AppRunTimeUpdateOne) ExecX(ctx context.Context) {
	if err := artuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (artuo *AppRunTimeUpdateOne) defaults() {
	if _, ok := artuo.mutation.UpdatedAt(); !ok {
		v := appruntime.UpdateDefaultUpdatedAt()
		artuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (artuo *AppRunTimeUpdateOne) check() error {
	if artuo.mutation.AppCleared() && len(artuo.mutation.AppIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "AppRunTime.app"`)
	}
	return nil
}

func (artuo *AppRunTimeUpdateOne) sqlSave(ctx context.Context) (_node *AppRunTime, err error) {
	if err := artuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(appruntime.Table, appruntime.Columns, sqlgraph.NewFieldSpec(appruntime.FieldID, field.TypeInt64))
	id, ok := artuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppRunTime.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := artuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appruntime.FieldID)
		for _, f := range fields {
			if !appruntime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appruntime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := artuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := artuo.mutation.UserID(); ok {
		_spec.SetField(appruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.AddedUserID(); ok {
		_spec.AddField(appruntime.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.DeviceID(); ok {
		_spec.SetField(appruntime.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.AddedDeviceID(); ok {
		_spec.AddField(appruntime.FieldDeviceID, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.StartTime(); ok {
		_spec.SetField(appruntime.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := artuo.mutation.Duration(); ok {
		_spec.SetField(appruntime.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.AddedDuration(); ok {
		_spec.AddField(appruntime.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := artuo.mutation.UpdatedAt(); ok {
		_spec.SetField(appruntime.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := artuo.mutation.CreatedAt(); ok {
		_spec.SetField(appruntime.FieldCreatedAt, field.TypeTime, value)
	}
	if artuo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appruntime.AppTable,
			Columns: []string{appruntime.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := artuo.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appruntime.AppTable,
			Columns: []string{appruntime.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppRunTime{config: artuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, artuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appruntime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	artuo.mutation.done = true
	return _node, nil
}
