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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/deviceinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/usersession"
	"github.com/tuihub/librarian/internal/model"
)

// DeviceInfoUpdate is the builder for updating DeviceInfo entities.
type DeviceInfoUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceInfoMutation
}

// Where appends a list predicates to the DeviceInfoUpdate builder.
func (diu *DeviceInfoUpdate) Where(ps ...predicate.DeviceInfo) *DeviceInfoUpdate {
	diu.mutation.Where(ps...)
	return diu
}

// SetDeviceName sets the "device_name" field.
func (diu *DeviceInfoUpdate) SetDeviceName(s string) *DeviceInfoUpdate {
	diu.mutation.SetDeviceName(s)
	return diu
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableDeviceName(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetDeviceName(*s)
	}
	return diu
}

// SetSystemType sets the "system_type" field.
func (diu *DeviceInfoUpdate) SetSystemType(dt deviceinfo.SystemType) *DeviceInfoUpdate {
	diu.mutation.SetSystemType(dt)
	return diu
}

// SetNillableSystemType sets the "system_type" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableSystemType(dt *deviceinfo.SystemType) *DeviceInfoUpdate {
	if dt != nil {
		diu.SetSystemType(*dt)
	}
	return diu
}

// SetSystemVersion sets the "system_version" field.
func (diu *DeviceInfoUpdate) SetSystemVersion(s string) *DeviceInfoUpdate {
	diu.mutation.SetSystemVersion(s)
	return diu
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableSystemVersion(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetSystemVersion(*s)
	}
	return diu
}

// SetClientName sets the "client_name" field.
func (diu *DeviceInfoUpdate) SetClientName(s string) *DeviceInfoUpdate {
	diu.mutation.SetClientName(s)
	return diu
}

// SetNillableClientName sets the "client_name" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableClientName(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetClientName(*s)
	}
	return diu
}

// SetClientSourceCodeAddress sets the "client_source_code_address" field.
func (diu *DeviceInfoUpdate) SetClientSourceCodeAddress(s string) *DeviceInfoUpdate {
	diu.mutation.SetClientSourceCodeAddress(s)
	return diu
}

// SetNillableClientSourceCodeAddress sets the "client_source_code_address" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableClientSourceCodeAddress(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetClientSourceCodeAddress(*s)
	}
	return diu
}

// SetClientVersion sets the "client_version" field.
func (diu *DeviceInfoUpdate) SetClientVersion(s string) *DeviceInfoUpdate {
	diu.mutation.SetClientVersion(s)
	return diu
}

// SetNillableClientVersion sets the "client_version" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableClientVersion(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetClientVersion(*s)
	}
	return diu
}

// SetUpdatedAt sets the "updated_at" field.
func (diu *DeviceInfoUpdate) SetUpdatedAt(t time.Time) *DeviceInfoUpdate {
	diu.mutation.SetUpdatedAt(t)
	return diu
}

// SetCreatedAt sets the "created_at" field.
func (diu *DeviceInfoUpdate) SetCreatedAt(t time.Time) *DeviceInfoUpdate {
	diu.mutation.SetCreatedAt(t)
	return diu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableCreatedAt(t *time.Time) *DeviceInfoUpdate {
	if t != nil {
		diu.SetCreatedAt(*t)
	}
	return diu
}

// AddUserSessionIDs adds the "user_session" edge to the UserSession entity by IDs.
func (diu *DeviceInfoUpdate) AddUserSessionIDs(ids ...model.InternalID) *DeviceInfoUpdate {
	diu.mutation.AddUserSessionIDs(ids...)
	return diu
}

// AddUserSession adds the "user_session" edges to the UserSession entity.
func (diu *DeviceInfoUpdate) AddUserSession(u ...*UserSession) *DeviceInfoUpdate {
	ids := make([]model.InternalID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return diu.AddUserSessionIDs(ids...)
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (diu *DeviceInfoUpdate) Mutation() *DeviceInfoMutation {
	return diu.mutation
}

// ClearUserSession clears all "user_session" edges to the UserSession entity.
func (diu *DeviceInfoUpdate) ClearUserSession() *DeviceInfoUpdate {
	diu.mutation.ClearUserSession()
	return diu
}

// RemoveUserSessionIDs removes the "user_session" edge to UserSession entities by IDs.
func (diu *DeviceInfoUpdate) RemoveUserSessionIDs(ids ...model.InternalID) *DeviceInfoUpdate {
	diu.mutation.RemoveUserSessionIDs(ids...)
	return diu
}

// RemoveUserSession removes "user_session" edges to UserSession entities.
func (diu *DeviceInfoUpdate) RemoveUserSession(u ...*UserSession) *DeviceInfoUpdate {
	ids := make([]model.InternalID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return diu.RemoveUserSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (diu *DeviceInfoUpdate) Save(ctx context.Context) (int, error) {
	diu.defaults()
	return withHooks(ctx, diu.sqlSave, diu.mutation, diu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diu *DeviceInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := diu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (diu *DeviceInfoUpdate) Exec(ctx context.Context) error {
	_, err := diu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diu *DeviceInfoUpdate) ExecX(ctx context.Context) {
	if err := diu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diu *DeviceInfoUpdate) defaults() {
	if _, ok := diu.mutation.UpdatedAt(); !ok {
		v := deviceinfo.UpdateDefaultUpdatedAt()
		diu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (diu *DeviceInfoUpdate) check() error {
	if v, ok := diu.mutation.SystemType(); ok {
		if err := deviceinfo.SystemTypeValidator(v); err != nil {
			return &ValidationError{Name: "system_type", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.system_type": %w`, err)}
		}
	}
	return nil
}

func (diu *DeviceInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := diu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deviceinfo.Table, deviceinfo.Columns, sqlgraph.NewFieldSpec(deviceinfo.FieldID, field.TypeInt64))
	if ps := diu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diu.mutation.DeviceName(); ok {
		_spec.SetField(deviceinfo.FieldDeviceName, field.TypeString, value)
	}
	if value, ok := diu.mutation.SystemType(); ok {
		_spec.SetField(deviceinfo.FieldSystemType, field.TypeEnum, value)
	}
	if value, ok := diu.mutation.SystemVersion(); ok {
		_spec.SetField(deviceinfo.FieldSystemVersion, field.TypeString, value)
	}
	if value, ok := diu.mutation.ClientName(); ok {
		_spec.SetField(deviceinfo.FieldClientName, field.TypeString, value)
	}
	if value, ok := diu.mutation.ClientSourceCodeAddress(); ok {
		_spec.SetField(deviceinfo.FieldClientSourceCodeAddress, field.TypeString, value)
	}
	if value, ok := diu.mutation.ClientVersion(); ok {
		_spec.SetField(deviceinfo.FieldClientVersion, field.TypeString, value)
	}
	if value, ok := diu.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diu.mutation.CreatedAt(); ok {
		_spec.SetField(deviceinfo.FieldCreatedAt, field.TypeTime, value)
	}
	if diu.mutation.UserSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diu.mutation.RemovedUserSessionIDs(); len(nodes) > 0 && !diu.mutation.UserSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diu.mutation.UserSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, diu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	diu.mutation.done = true
	return n, nil
}

// DeviceInfoUpdateOne is the builder for updating a single DeviceInfo entity.
type DeviceInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceInfoMutation
}

// SetDeviceName sets the "device_name" field.
func (diuo *DeviceInfoUpdateOne) SetDeviceName(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetDeviceName(s)
	return diuo
}

// SetNillableDeviceName sets the "device_name" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableDeviceName(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetDeviceName(*s)
	}
	return diuo
}

// SetSystemType sets the "system_type" field.
func (diuo *DeviceInfoUpdateOne) SetSystemType(dt deviceinfo.SystemType) *DeviceInfoUpdateOne {
	diuo.mutation.SetSystemType(dt)
	return diuo
}

// SetNillableSystemType sets the "system_type" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableSystemType(dt *deviceinfo.SystemType) *DeviceInfoUpdateOne {
	if dt != nil {
		diuo.SetSystemType(*dt)
	}
	return diuo
}

// SetSystemVersion sets the "system_version" field.
func (diuo *DeviceInfoUpdateOne) SetSystemVersion(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetSystemVersion(s)
	return diuo
}

// SetNillableSystemVersion sets the "system_version" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableSystemVersion(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetSystemVersion(*s)
	}
	return diuo
}

// SetClientName sets the "client_name" field.
func (diuo *DeviceInfoUpdateOne) SetClientName(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetClientName(s)
	return diuo
}

// SetNillableClientName sets the "client_name" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableClientName(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetClientName(*s)
	}
	return diuo
}

// SetClientSourceCodeAddress sets the "client_source_code_address" field.
func (diuo *DeviceInfoUpdateOne) SetClientSourceCodeAddress(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetClientSourceCodeAddress(s)
	return diuo
}

// SetNillableClientSourceCodeAddress sets the "client_source_code_address" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableClientSourceCodeAddress(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetClientSourceCodeAddress(*s)
	}
	return diuo
}

// SetClientVersion sets the "client_version" field.
func (diuo *DeviceInfoUpdateOne) SetClientVersion(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetClientVersion(s)
	return diuo
}

// SetNillableClientVersion sets the "client_version" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableClientVersion(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetClientVersion(*s)
	}
	return diuo
}

// SetUpdatedAt sets the "updated_at" field.
func (diuo *DeviceInfoUpdateOne) SetUpdatedAt(t time.Time) *DeviceInfoUpdateOne {
	diuo.mutation.SetUpdatedAt(t)
	return diuo
}

// SetCreatedAt sets the "created_at" field.
func (diuo *DeviceInfoUpdateOne) SetCreatedAt(t time.Time) *DeviceInfoUpdateOne {
	diuo.mutation.SetCreatedAt(t)
	return diuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableCreatedAt(t *time.Time) *DeviceInfoUpdateOne {
	if t != nil {
		diuo.SetCreatedAt(*t)
	}
	return diuo
}

// AddUserSessionIDs adds the "user_session" edge to the UserSession entity by IDs.
func (diuo *DeviceInfoUpdateOne) AddUserSessionIDs(ids ...model.InternalID) *DeviceInfoUpdateOne {
	diuo.mutation.AddUserSessionIDs(ids...)
	return diuo
}

// AddUserSession adds the "user_session" edges to the UserSession entity.
func (diuo *DeviceInfoUpdateOne) AddUserSession(u ...*UserSession) *DeviceInfoUpdateOne {
	ids := make([]model.InternalID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return diuo.AddUserSessionIDs(ids...)
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (diuo *DeviceInfoUpdateOne) Mutation() *DeviceInfoMutation {
	return diuo.mutation
}

// ClearUserSession clears all "user_session" edges to the UserSession entity.
func (diuo *DeviceInfoUpdateOne) ClearUserSession() *DeviceInfoUpdateOne {
	diuo.mutation.ClearUserSession()
	return diuo
}

// RemoveUserSessionIDs removes the "user_session" edge to UserSession entities by IDs.
func (diuo *DeviceInfoUpdateOne) RemoveUserSessionIDs(ids ...model.InternalID) *DeviceInfoUpdateOne {
	diuo.mutation.RemoveUserSessionIDs(ids...)
	return diuo
}

// RemoveUserSession removes "user_session" edges to UserSession entities.
func (diuo *DeviceInfoUpdateOne) RemoveUserSession(u ...*UserSession) *DeviceInfoUpdateOne {
	ids := make([]model.InternalID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return diuo.RemoveUserSessionIDs(ids...)
}

// Where appends a list predicates to the DeviceInfoUpdate builder.
func (diuo *DeviceInfoUpdateOne) Where(ps ...predicate.DeviceInfo) *DeviceInfoUpdateOne {
	diuo.mutation.Where(ps...)
	return diuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (diuo *DeviceInfoUpdateOne) Select(field string, fields ...string) *DeviceInfoUpdateOne {
	diuo.fields = append([]string{field}, fields...)
	return diuo
}

// Save executes the query and returns the updated DeviceInfo entity.
func (diuo *DeviceInfoUpdateOne) Save(ctx context.Context) (*DeviceInfo, error) {
	diuo.defaults()
	return withHooks(ctx, diuo.sqlSave, diuo.mutation, diuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diuo *DeviceInfoUpdateOne) SaveX(ctx context.Context) *DeviceInfo {
	node, err := diuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (diuo *DeviceInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := diuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diuo *DeviceInfoUpdateOne) ExecX(ctx context.Context) {
	if err := diuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diuo *DeviceInfoUpdateOne) defaults() {
	if _, ok := diuo.mutation.UpdatedAt(); !ok {
		v := deviceinfo.UpdateDefaultUpdatedAt()
		diuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (diuo *DeviceInfoUpdateOne) check() error {
	if v, ok := diuo.mutation.SystemType(); ok {
		if err := deviceinfo.SystemTypeValidator(v); err != nil {
			return &ValidationError{Name: "system_type", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.system_type": %w`, err)}
		}
	}
	return nil
}

func (diuo *DeviceInfoUpdateOne) sqlSave(ctx context.Context) (_node *DeviceInfo, err error) {
	if err := diuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deviceinfo.Table, deviceinfo.Columns, sqlgraph.NewFieldSpec(deviceinfo.FieldID, field.TypeInt64))
	id, ok := diuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeviceInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := diuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinfo.FieldID)
		for _, f := range fields {
			if !deviceinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deviceinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := diuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diuo.mutation.DeviceName(); ok {
		_spec.SetField(deviceinfo.FieldDeviceName, field.TypeString, value)
	}
	if value, ok := diuo.mutation.SystemType(); ok {
		_spec.SetField(deviceinfo.FieldSystemType, field.TypeEnum, value)
	}
	if value, ok := diuo.mutation.SystemVersion(); ok {
		_spec.SetField(deviceinfo.FieldSystemVersion, field.TypeString, value)
	}
	if value, ok := diuo.mutation.ClientName(); ok {
		_spec.SetField(deviceinfo.FieldClientName, field.TypeString, value)
	}
	if value, ok := diuo.mutation.ClientSourceCodeAddress(); ok {
		_spec.SetField(deviceinfo.FieldClientSourceCodeAddress, field.TypeString, value)
	}
	if value, ok := diuo.mutation.ClientVersion(); ok {
		_spec.SetField(deviceinfo.FieldClientVersion, field.TypeString, value)
	}
	if value, ok := diuo.mutation.UpdatedAt(); ok {
		_spec.SetField(deviceinfo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diuo.mutation.CreatedAt(); ok {
		_spec.SetField(deviceinfo.FieldCreatedAt, field.TypeTime, value)
	}
	if diuo.mutation.UserSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diuo.mutation.RemovedUserSessionIDs(); len(nodes) > 0 && !diuo.mutation.UserSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := diuo.mutation.UserSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   deviceinfo.UserSessionTable,
			Columns: []string{deviceinfo.UserSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DeviceInfo{config: diuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, diuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	diuo.mutation.done = true
	return _node, nil
}