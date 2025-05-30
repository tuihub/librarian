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
	"github.com/tuihub/librarian/internal/data/internal/ent/device"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/session"
	"github.com/tuihub/librarian/internal/model"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetDeviceName sets the "device_name" field.
func (du *DeviceUpdate) SetDeviceName(s string) *DeviceUpdate {
	du.mutation.SetDeviceName(s)
	return du
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableDeviceName(s *string) *DeviceUpdate {
	if s != nil {
		du.SetDeviceName(*s)
	}
	return du
}

// SetSystemType sets the "system_type" field.
func (du *DeviceUpdate) SetSystemType(dt device.SystemType) *DeviceUpdate {
	du.mutation.SetSystemType(dt)
	return du
}

// SetNillableSystemType sets the "system_type" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableSystemType(dt *device.SystemType) *DeviceUpdate {
	if dt != nil {
		du.SetSystemType(*dt)
	}
	return du
}

// SetSystemVersion sets the "system_version" field.
func (du *DeviceUpdate) SetSystemVersion(s string) *DeviceUpdate {
	du.mutation.SetSystemVersion(s)
	return du
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableSystemVersion(s *string) *DeviceUpdate {
	if s != nil {
		du.SetSystemVersion(*s)
	}
	return du
}

// SetClientName sets the "client_name" field.
func (du *DeviceUpdate) SetClientName(s string) *DeviceUpdate {
	du.mutation.SetClientName(s)
	return du
}

// SetNillableClientName sets the "client_name" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableClientName(s *string) *DeviceUpdate {
	if s != nil {
		du.SetClientName(*s)
	}
	return du
}

// SetClientSourceCodeAddress sets the "client_source_code_address" field.
func (du *DeviceUpdate) SetClientSourceCodeAddress(s string) *DeviceUpdate {
	du.mutation.SetClientSourceCodeAddress(s)
	return du
}

// SetNillableClientSourceCodeAddress sets the "client_source_code_address" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableClientSourceCodeAddress(s *string) *DeviceUpdate {
	if s != nil {
		du.SetClientSourceCodeAddress(*s)
	}
	return du
}

// SetClientVersion sets the "client_version" field.
func (du *DeviceUpdate) SetClientVersion(s string) *DeviceUpdate {
	du.mutation.SetClientVersion(s)
	return du
}

// SetNillableClientVersion sets the "client_version" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableClientVersion(s *string) *DeviceUpdate {
	if s != nil {
		du.SetClientVersion(*s)
	}
	return du
}

// SetClientLocalID sets the "client_local_id" field.
func (du *DeviceUpdate) SetClientLocalID(s string) *DeviceUpdate {
	du.mutation.SetClientLocalID(s)
	return du
}

// SetNillableClientLocalID sets the "client_local_id" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableClientLocalID(s *string) *DeviceUpdate {
	if s != nil {
		du.SetClientLocalID(*s)
	}
	return du
}

// ClearClientLocalID clears the value of the "client_local_id" field.
func (du *DeviceUpdate) ClearClientLocalID() *DeviceUpdate {
	du.mutation.ClearClientLocalID()
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeviceUpdate) SetUpdatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DeviceUpdate) SetCreatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableCreatedAt(t *time.Time) *DeviceUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// AddSessionIDs adds the "session" edge to the Session entity by IDs.
func (du *DeviceUpdate) AddSessionIDs(ids ...model.InternalID) *DeviceUpdate {
	du.mutation.AddSessionIDs(ids...)
	return du
}

// AddSession adds the "session" edges to the Session entity.
func (du *DeviceUpdate) AddSession(s ...*Session) *DeviceUpdate {
	ids := make([]model.InternalID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddSessionIDs(ids...)
}

// AddAppIDs adds the "app" edge to the App entity by IDs.
func (du *DeviceUpdate) AddAppIDs(ids ...model.InternalID) *DeviceUpdate {
	du.mutation.AddAppIDs(ids...)
	return du
}

// AddApp adds the "app" edges to the App entity.
func (du *DeviceUpdate) AddApp(a ...*App) *DeviceUpdate {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.AddAppIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearSession clears all "session" edges to the Session entity.
func (du *DeviceUpdate) ClearSession() *DeviceUpdate {
	du.mutation.ClearSession()
	return du
}

// RemoveSessionIDs removes the "session" edge to Session entities by IDs.
func (du *DeviceUpdate) RemoveSessionIDs(ids ...model.InternalID) *DeviceUpdate {
	du.mutation.RemoveSessionIDs(ids...)
	return du
}

// RemoveSession removes "session" edges to Session entities.
func (du *DeviceUpdate) RemoveSession(s ...*Session) *DeviceUpdate {
	ids := make([]model.InternalID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveSessionIDs(ids...)
}

// ClearApp clears all "app" edges to the App entity.
func (du *DeviceUpdate) ClearApp() *DeviceUpdate {
	du.mutation.ClearApp()
	return du
}

// RemoveAppIDs removes the "app" edge to App entities by IDs.
func (du *DeviceUpdate) RemoveAppIDs(ids ...model.InternalID) *DeviceUpdate {
	du.mutation.RemoveAppIDs(ids...)
	return du
}

// RemoveApp removes "app" edges to App entities.
func (du *DeviceUpdate) RemoveApp(a ...*App) *DeviceUpdate {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.RemoveAppIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeviceUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeviceUpdate) check() error {
	if v, ok := du.mutation.SystemType(); ok {
		if err := device.SystemTypeValidator(v); err != nil {
			return &ValidationError{Name: "system_type", err: fmt.Errorf(`ent: validator failed for field "Device.system_type": %w`, err)}
		}
	}
	return nil
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.DeviceName(); ok {
		_spec.SetField(device.FieldDeviceName, field.TypeString, value)
	}
	if value, ok := du.mutation.SystemType(); ok {
		_spec.SetField(device.FieldSystemType, field.TypeEnum, value)
	}
	if value, ok := du.mutation.SystemVersion(); ok {
		_spec.SetField(device.FieldSystemVersion, field.TypeString, value)
	}
	if value, ok := du.mutation.ClientName(); ok {
		_spec.SetField(device.FieldClientName, field.TypeString, value)
	}
	if value, ok := du.mutation.ClientSourceCodeAddress(); ok {
		_spec.SetField(device.FieldClientSourceCodeAddress, field.TypeString, value)
	}
	if value, ok := du.mutation.ClientVersion(); ok {
		_spec.SetField(device.FieldClientVersion, field.TypeString, value)
	}
	if value, ok := du.mutation.ClientLocalID(); ok {
		_spec.SetField(device.FieldClientLocalID, field.TypeString, value)
	}
	if du.mutation.ClientLocalIDCleared() {
		_spec.ClearField(device.FieldClientLocalID, field.TypeString)
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
	}
	if du.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSessionIDs(); len(nodes) > 0 && !du.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAppIDs(); len(nodes) > 0 && !du.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetDeviceName sets the "device_name" field.
func (duo *DeviceUpdateOne) SetDeviceName(s string) *DeviceUpdateOne {
	duo.mutation.SetDeviceName(s)
	return duo
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableDeviceName(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetDeviceName(*s)
	}
	return duo
}

// SetSystemType sets the "system_type" field.
func (duo *DeviceUpdateOne) SetSystemType(dt device.SystemType) *DeviceUpdateOne {
	duo.mutation.SetSystemType(dt)
	return duo
}

// SetNillableSystemType sets the "system_type" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableSystemType(dt *device.SystemType) *DeviceUpdateOne {
	if dt != nil {
		duo.SetSystemType(*dt)
	}
	return duo
}

// SetSystemVersion sets the "system_version" field.
func (duo *DeviceUpdateOne) SetSystemVersion(s string) *DeviceUpdateOne {
	duo.mutation.SetSystemVersion(s)
	return duo
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableSystemVersion(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetSystemVersion(*s)
	}
	return duo
}

// SetClientName sets the "client_name" field.
func (duo *DeviceUpdateOne) SetClientName(s string) *DeviceUpdateOne {
	duo.mutation.SetClientName(s)
	return duo
}

// SetNillableClientName sets the "client_name" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableClientName(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetClientName(*s)
	}
	return duo
}

// SetClientSourceCodeAddress sets the "client_source_code_address" field.
func (duo *DeviceUpdateOne) SetClientSourceCodeAddress(s string) *DeviceUpdateOne {
	duo.mutation.SetClientSourceCodeAddress(s)
	return duo
}

// SetNillableClientSourceCodeAddress sets the "client_source_code_address" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableClientSourceCodeAddress(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetClientSourceCodeAddress(*s)
	}
	return duo
}

// SetClientVersion sets the "client_version" field.
func (duo *DeviceUpdateOne) SetClientVersion(s string) *DeviceUpdateOne {
	duo.mutation.SetClientVersion(s)
	return duo
}

// SetNillableClientVersion sets the "client_version" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableClientVersion(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetClientVersion(*s)
	}
	return duo
}

// SetClientLocalID sets the "client_local_id" field.
func (duo *DeviceUpdateOne) SetClientLocalID(s string) *DeviceUpdateOne {
	duo.mutation.SetClientLocalID(s)
	return duo
}

// SetNillableClientLocalID sets the "client_local_id" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableClientLocalID(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetClientLocalID(*s)
	}
	return duo
}

// ClearClientLocalID clears the value of the "client_local_id" field.
func (duo *DeviceUpdateOne) ClearClientLocalID() *DeviceUpdateOne {
	duo.mutation.ClearClientLocalID()
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeviceUpdateOne) SetUpdatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DeviceUpdateOne) SetCreatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableCreatedAt(t *time.Time) *DeviceUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// AddSessionIDs adds the "session" edge to the Session entity by IDs.
func (duo *DeviceUpdateOne) AddSessionIDs(ids ...model.InternalID) *DeviceUpdateOne {
	duo.mutation.AddSessionIDs(ids...)
	return duo
}

// AddSession adds the "session" edges to the Session entity.
func (duo *DeviceUpdateOne) AddSession(s ...*Session) *DeviceUpdateOne {
	ids := make([]model.InternalID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddSessionIDs(ids...)
}

// AddAppIDs adds the "app" edge to the App entity by IDs.
func (duo *DeviceUpdateOne) AddAppIDs(ids ...model.InternalID) *DeviceUpdateOne {
	duo.mutation.AddAppIDs(ids...)
	return duo
}

// AddApp adds the "app" edges to the App entity.
func (duo *DeviceUpdateOne) AddApp(a ...*App) *DeviceUpdateOne {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.AddAppIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearSession clears all "session" edges to the Session entity.
func (duo *DeviceUpdateOne) ClearSession() *DeviceUpdateOne {
	duo.mutation.ClearSession()
	return duo
}

// RemoveSessionIDs removes the "session" edge to Session entities by IDs.
func (duo *DeviceUpdateOne) RemoveSessionIDs(ids ...model.InternalID) *DeviceUpdateOne {
	duo.mutation.RemoveSessionIDs(ids...)
	return duo
}

// RemoveSession removes "session" edges to Session entities.
func (duo *DeviceUpdateOne) RemoveSession(s ...*Session) *DeviceUpdateOne {
	ids := make([]model.InternalID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveSessionIDs(ids...)
}

// ClearApp clears all "app" edges to the App entity.
func (duo *DeviceUpdateOne) ClearApp() *DeviceUpdateOne {
	duo.mutation.ClearApp()
	return duo
}

// RemoveAppIDs removes the "app" edge to App entities by IDs.
func (duo *DeviceUpdateOne) RemoveAppIDs(ids ...model.InternalID) *DeviceUpdateOne {
	duo.mutation.RemoveAppIDs(ids...)
	return duo
}

// RemoveApp removes "app" edges to App entities.
func (duo *DeviceUpdateOne) RemoveApp(a ...*App) *DeviceUpdateOne {
	ids := make([]model.InternalID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.RemoveAppIDs(ids...)
}

// Where appends a list predicates to the DeviceUpdate builder.
func (duo *DeviceUpdateOne) Where(ps ...predicate.Device) *DeviceUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeviceUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeviceUpdateOne) check() error {
	if v, ok := duo.mutation.SystemType(); ok {
		if err := device.SystemTypeValidator(v); err != nil {
			return &ValidationError{Name: "system_type", err: fmt.Errorf(`ent: validator failed for field "Device.system_type": %w`, err)}
		}
	}
	return nil
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeInt64))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.DeviceName(); ok {
		_spec.SetField(device.FieldDeviceName, field.TypeString, value)
	}
	if value, ok := duo.mutation.SystemType(); ok {
		_spec.SetField(device.FieldSystemType, field.TypeEnum, value)
	}
	if value, ok := duo.mutation.SystemVersion(); ok {
		_spec.SetField(device.FieldSystemVersion, field.TypeString, value)
	}
	if value, ok := duo.mutation.ClientName(); ok {
		_spec.SetField(device.FieldClientName, field.TypeString, value)
	}
	if value, ok := duo.mutation.ClientSourceCodeAddress(); ok {
		_spec.SetField(device.FieldClientSourceCodeAddress, field.TypeString, value)
	}
	if value, ok := duo.mutation.ClientVersion(); ok {
		_spec.SetField(device.FieldClientVersion, field.TypeString, value)
	}
	if value, ok := duo.mutation.ClientLocalID(); ok {
		_spec.SetField(device.FieldClientLocalID, field.TypeString, value)
	}
	if duo.mutation.ClientLocalIDCleared() {
		_spec.ClearField(device.FieldClientLocalID, field.TypeString)
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
	}
	if duo.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSessionIDs(); len(nodes) > 0 && !duo.mutation.SessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionTable,
			Columns: []string{device.SessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAppIDs(); len(nodes) > 0 && !duo.mutation.AppCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.AppTable,
			Columns: []string{device.AppColumn},
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
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
