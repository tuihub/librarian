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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/userdevice"
	"github.com/tuihub/librarian/internal/model"
)

// UserDeviceCreate is the builder for creating a UserDevice entity.
type UserDeviceCreate struct {
	config
	mutation *UserDeviceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (udc *UserDeviceCreate) SetUserID(mi model.InternalID) *UserDeviceCreate {
	udc.mutation.SetUserID(mi)
	return udc
}

// SetDeviceID sets the "device_id" field.
func (udc *UserDeviceCreate) SetDeviceID(mi model.InternalID) *UserDeviceCreate {
	udc.mutation.SetDeviceID(mi)
	return udc
}

// SetUpdatedAt sets the "updated_at" field.
func (udc *UserDeviceCreate) SetUpdatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetUpdatedAt(t)
	return udc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableUpdatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetUpdatedAt(*t)
	}
	return udc
}

// SetCreatedAt sets the "created_at" field.
func (udc *UserDeviceCreate) SetCreatedAt(t time.Time) *UserDeviceCreate {
	udc.mutation.SetCreatedAt(t)
	return udc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (udc *UserDeviceCreate) SetNillableCreatedAt(t *time.Time) *UserDeviceCreate {
	if t != nil {
		udc.SetCreatedAt(*t)
	}
	return udc
}

// SetDeviceInfoID sets the "device_info" edge to the DeviceInfo entity by ID.
func (udc *UserDeviceCreate) SetDeviceInfoID(id model.InternalID) *UserDeviceCreate {
	udc.mutation.SetDeviceInfoID(id)
	return udc
}

// SetDeviceInfo sets the "device_info" edge to the DeviceInfo entity.
func (udc *UserDeviceCreate) SetDeviceInfo(d *DeviceInfo) *UserDeviceCreate {
	return udc.SetDeviceInfoID(d.ID)
}

// SetUser sets the "user" edge to the User entity.
func (udc *UserDeviceCreate) SetUser(u *User) *UserDeviceCreate {
	return udc.SetUserID(u.ID)
}

// Mutation returns the UserDeviceMutation object of the builder.
func (udc *UserDeviceCreate) Mutation() *UserDeviceMutation {
	return udc.mutation
}

// Save creates the UserDevice in the database.
func (udc *UserDeviceCreate) Save(ctx context.Context) (*UserDevice, error) {
	udc.defaults()
	return withHooks(ctx, udc.sqlSave, udc.mutation, udc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (udc *UserDeviceCreate) SaveX(ctx context.Context) *UserDevice {
	v, err := udc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udc *UserDeviceCreate) Exec(ctx context.Context) error {
	_, err := udc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udc *UserDeviceCreate) ExecX(ctx context.Context) {
	if err := udc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (udc *UserDeviceCreate) defaults() {
	if _, ok := udc.mutation.UpdatedAt(); !ok {
		v := userdevice.DefaultUpdatedAt()
		udc.mutation.SetUpdatedAt(v)
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		v := userdevice.DefaultCreatedAt()
		udc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udc *UserDeviceCreate) check() error {
	if _, ok := udc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserDevice.user_id"`)}
	}
	if _, ok := udc.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`ent: missing required field "UserDevice.device_id"`)}
	}
	if _, ok := udc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserDevice.updated_at"`)}
	}
	if _, ok := udc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserDevice.created_at"`)}
	}
	if _, ok := udc.mutation.DeviceInfoID(); !ok {
		return &ValidationError{Name: "device_info", err: errors.New(`ent: missing required edge "UserDevice.device_info"`)}
	}
	if _, ok := udc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserDevice.user"`)}
	}
	return nil
}

func (udc *UserDeviceCreate) sqlSave(ctx context.Context) (*UserDevice, error) {
	if err := udc.check(); err != nil {
		return nil, err
	}
	_node, _spec := udc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	udc.mutation.id = &_node.ID
	udc.mutation.done = true
	return _node, nil
}

func (udc *UserDeviceCreate) createSpec() (*UserDevice, *sqlgraph.CreateSpec) {
	var (
		_node = &UserDevice{config: udc.config}
		_spec = sqlgraph.NewCreateSpec(userdevice.Table, sqlgraph.NewFieldSpec(userdevice.FieldID, field.TypeInt))
	)
	_spec.OnConflict = udc.conflict
	if value, ok := udc.mutation.UpdatedAt(); ok {
		_spec.SetField(userdevice.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := udc.mutation.CreatedAt(); ok {
		_spec.SetField(userdevice.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := udc.mutation.DeviceInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userdevice.DeviceInfoTable,
			Columns: []string{userdevice.DeviceInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deviceinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DeviceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := udc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userdevice.UserTable,
			Columns: []string{userdevice.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserDevice.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDeviceUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (udc *UserDeviceCreate) OnConflict(opts ...sql.ConflictOption) *UserDeviceUpsertOne {
	udc.conflict = opts
	return &UserDeviceUpsertOne{
		create: udc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (udc *UserDeviceCreate) OnConflictColumns(columns ...string) *UserDeviceUpsertOne {
	udc.conflict = append(udc.conflict, sql.ConflictColumns(columns...))
	return &UserDeviceUpsertOne{
		create: udc,
	}
}

type (
	// UserDeviceUpsertOne is the builder for "upsert"-ing
	//  one UserDevice node.
	UserDeviceUpsertOne struct {
		create *UserDeviceCreate
	}

	// UserDeviceUpsert is the "OnConflict" setter.
	UserDeviceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsert) SetUserID(v model.InternalID) *UserDeviceUpsert {
	u.Set(userdevice.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateUserID() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldUserID)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsert) SetDeviceID(v model.InternalID) *UserDeviceUpsert {
	u.Set(userdevice.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateDeviceID() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldDeviceID)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsert) SetUpdatedAt(v time.Time) *UserDeviceUpsert {
	u.Set(userdevice.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateUpdatedAt() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserDeviceUpsert) SetCreatedAt(v time.Time) *UserDeviceUpsert {
	u.Set(userdevice.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserDeviceUpsert) UpdateCreatedAt() *UserDeviceUpsert {
	u.SetExcluded(userdevice.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserDeviceUpsertOne) UpdateNewValues() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserDeviceUpsertOne) Ignore() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDeviceUpsertOne) DoNothing() *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDeviceCreate.OnConflict
// documentation for more info.
func (u *UserDeviceUpsertOne) Update(set func(*UserDeviceUpsert)) *UserDeviceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsertOne) SetUserID(v model.InternalID) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateUserID() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUserID()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsertOne) SetDeviceID(v model.InternalID) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateDeviceID() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsertOne) SetUpdatedAt(v time.Time) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateUpdatedAt() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UserDeviceUpsertOne) SetCreatedAt(v time.Time) *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserDeviceUpsertOne) UpdateCreatedAt() *UserDeviceUpsertOne {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *UserDeviceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserDeviceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDeviceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserDeviceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserDeviceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserDeviceCreateBulk is the builder for creating many UserDevice entities in bulk.
type UserDeviceCreateBulk struct {
	config
	err      error
	builders []*UserDeviceCreate
	conflict []sql.ConflictOption
}

// Save creates the UserDevice entities in the database.
func (udcb *UserDeviceCreateBulk) Save(ctx context.Context) ([]*UserDevice, error) {
	if udcb.err != nil {
		return nil, udcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(udcb.builders))
	nodes := make([]*UserDevice, len(udcb.builders))
	mutators := make([]Mutator, len(udcb.builders))
	for i := range udcb.builders {
		func(i int, root context.Context) {
			builder := udcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDeviceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, udcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = udcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, udcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) SaveX(ctx context.Context) []*UserDevice {
	v, err := udcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udcb *UserDeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := udcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udcb *UserDeviceCreateBulk) ExecX(ctx context.Context) {
	if err := udcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserDevice.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserDeviceUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (udcb *UserDeviceCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserDeviceUpsertBulk {
	udcb.conflict = opts
	return &UserDeviceUpsertBulk{
		create: udcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (udcb *UserDeviceCreateBulk) OnConflictColumns(columns ...string) *UserDeviceUpsertBulk {
	udcb.conflict = append(udcb.conflict, sql.ConflictColumns(columns...))
	return &UserDeviceUpsertBulk{
		create: udcb,
	}
}

// UserDeviceUpsertBulk is the builder for "upsert"-ing
// a bulk of UserDevice nodes.
type UserDeviceUpsertBulk struct {
	create *UserDeviceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserDeviceUpsertBulk) UpdateNewValues() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserDevice.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserDeviceUpsertBulk) Ignore() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserDeviceUpsertBulk) DoNothing() *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserDeviceCreateBulk.OnConflict
// documentation for more info.
func (u *UserDeviceUpsertBulk) Update(set func(*UserDeviceUpsert)) *UserDeviceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserDeviceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserDeviceUpsertBulk) SetUserID(v model.InternalID) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateUserID() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUserID()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *UserDeviceUpsertBulk) SetDeviceID(v model.InternalID) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateDeviceID() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateDeviceID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserDeviceUpsertBulk) SetUpdatedAt(v time.Time) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateUpdatedAt() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UserDeviceUpsertBulk) SetCreatedAt(v time.Time) *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserDeviceUpsertBulk) UpdateCreatedAt() *UserDeviceUpsertBulk {
	return u.Update(func(s *UserDeviceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *UserDeviceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserDeviceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserDeviceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserDeviceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}