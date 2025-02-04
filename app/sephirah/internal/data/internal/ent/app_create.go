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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AppCreate is the builder for creating a App entity.
type AppCreate struct {
	config
	mutation *AppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AppCreate) SetName(s string) *AppCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AppCreate) SetDescription(s string) *AppCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetDeviceID sets the "device_id" field.
func (ac *AppCreate) SetDeviceID(mi model.InternalID) *AppCreate {
	ac.mutation.SetDeviceID(mi)
	return ac
}

// SetPublic sets the "public" field.
func (ac *AppCreate) SetPublic(b bool) *AppCreate {
	ac.mutation.SetPublic(b)
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AppCreate) SetUpdatedAt(t time.Time) *AppCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableUpdatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppCreate) SetCreatedAt(t time.Time) *AppCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedAt(t *time.Time) *AppCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AppCreate) SetID(mi model.InternalID) *AppCreate {
	ac.mutation.SetID(mi)
	return ac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ac *AppCreate) SetOwnerID(id model.InternalID) *AppCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetOwner sets the "owner" edge to the User entity.
func (ac *AppCreate) SetOwner(u *User) *AppCreate {
	return ac.SetOwnerID(u.ID)
}

// SetAppInfoID sets the "app_info" edge to the AppInfo entity by ID.
func (ac *AppCreate) SetAppInfoID(id model.InternalID) *AppCreate {
	ac.mutation.SetAppInfoID(id)
	return ac
}

// SetNillableAppInfoID sets the "app_info" edge to the AppInfo entity by ID if the given value is not nil.
func (ac *AppCreate) SetNillableAppInfoID(id *model.InternalID) *AppCreate {
	if id != nil {
		ac = ac.SetAppInfoID(*id)
	}
	return ac
}

// SetAppInfo sets the "app_info" edge to the AppInfo entity.
func (ac *AppCreate) SetAppInfo(a *AppInfo) *AppCreate {
	return ac.SetAppInfoID(a.ID)
}

// Mutation returns the AppMutation object of the builder.
func (ac *AppCreate) Mutation() *AppMutation {
	return ac.mutation
}

// Save creates the App in the database.
func (ac *AppCreate) Save(ctx context.Context) (*App, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppCreate) SaveX(ctx context.Context) *App {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppCreate) defaults() {
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := app.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := app.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "App.name"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "App.description"`)}
	}
	if _, ok := ac.mutation.DeviceID(); !ok {
		return &ValidationError{Name: "device_id", err: errors.New(`ent: missing required field "App.device_id"`)}
	}
	if _, ok := ac.mutation.Public(); !ok {
		return &ValidationError{Name: "public", err: errors.New(`ent: missing required field "App.public"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "App.updated_at"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "App.created_at"`)}
	}
	if len(ac.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "App.owner"`)}
	}
	return nil
}

func (ac *AppCreate) sqlSave(ctx context.Context) (*App, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AppCreate) createSpec() (*App, *sqlgraph.CreateSpec) {
	var (
		_node = &App{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(app.Table, sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(app.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(app.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.DeviceID(); ok {
		_spec.SetField(app.FieldDeviceID, field.TypeInt64, value)
		_node.DeviceID = value
	}
	if value, ok := ac.mutation.Public(); ok {
		_spec.SetField(app.FieldPublic, field.TypeBool, value)
		_node.Public = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(app.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(app.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   app.OwnerTable,
			Columns: []string{app.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_app = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AppInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   app.AppInfoTable,
			Columns: []string{app.AppInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.app_info_app = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AppCreate) OnConflict(opts ...sql.ConflictOption) *AppUpsertOne {
	ac.conflict = opts
	return &AppUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AppCreate) OnConflictColumns(columns ...string) *AppUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertOne{
		create: ac,
	}
}

type (
	// AppUpsertOne is the builder for "upsert"-ing
	//  one App node.
	AppUpsertOne struct {
		create *AppCreate
	}

	// AppUpsert is the "OnConflict" setter.
	AppUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AppUpsert) SetName(v string) *AppUpsert {
	u.Set(app.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsert) UpdateName() *AppUpsert {
	u.SetExcluded(app.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *AppUpsert) SetDescription(v string) *AppUpsert {
	u.Set(app.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsert) UpdateDescription() *AppUpsert {
	u.SetExcluded(app.FieldDescription)
	return u
}

// SetDeviceID sets the "device_id" field.
func (u *AppUpsert) SetDeviceID(v model.InternalID) *AppUpsert {
	u.Set(app.FieldDeviceID, v)
	return u
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *AppUpsert) UpdateDeviceID() *AppUpsert {
	u.SetExcluded(app.FieldDeviceID)
	return u
}

// AddDeviceID adds v to the "device_id" field.
func (u *AppUpsert) AddDeviceID(v model.InternalID) *AppUpsert {
	u.Add(app.FieldDeviceID, v)
	return u
}

// SetPublic sets the "public" field.
func (u *AppUpsert) SetPublic(v bool) *AppUpsert {
	u.Set(app.FieldPublic, v)
	return u
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *AppUpsert) UpdatePublic() *AppUpsert {
	u.SetExcluded(app.FieldPublic)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsert) SetUpdatedAt(v time.Time) *AppUpsert {
	u.Set(app.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateUpdatedAt() *AppUpsert {
	u.SetExcluded(app.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsert) SetCreatedAt(v time.Time) *AppUpsert {
	u.Set(app.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateCreatedAt() *AppUpsert {
	u.SetExcluded(app.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppUpsertOne) UpdateNewValues() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(app.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppUpsertOne) Ignore() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertOne) DoNothing() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreate.OnConflict
// documentation for more info.
func (u *AppUpsertOne) Update(set func(*AppUpsert)) *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AppUpsertOne) SetName(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateName() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertOne) SetDescription(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *AppUpsertOne) SetDeviceID(v model.InternalID) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDeviceID(v)
	})
}

// AddDeviceID adds v to the "device_id" field.
func (u *AppUpsertOne) AddDeviceID(v model.InternalID) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.AddDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDeviceID() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeviceID()
	})
}

// SetPublic sets the "public" field.
func (u *AppUpsertOne) SetPublic(v bool) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *AppUpsertOne) UpdatePublic() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdatePublic()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertOne) SetUpdatedAt(v time.Time) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateUpdatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertOne) SetCreatedAt(v time.Time) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateCreatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCreateBulk is the builder for creating many App entities in bulk.
type AppCreateBulk struct {
	config
	err      error
	builders []*AppCreate
	conflict []sql.ConflictOption
}

// Save creates the App entities in the database.
func (acb *AppCreateBulk) Save(ctx context.Context) ([]*App, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*App, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = model.InternalID(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppCreateBulk) SaveX(ctx context.Context) []*App {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AppCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUpsertBulk {
	acb.conflict = opts
	return &AppUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AppCreateBulk) OnConflictColumns(columns ...string) *AppUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertBulk{
		create: acb,
	}
}

// AppUpsertBulk is the builder for "upsert"-ing
// a bulk of App nodes.
type AppUpsertBulk struct {
	create *AppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppUpsertBulk) UpdateNewValues() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(app.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppUpsertBulk) Ignore() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertBulk) DoNothing() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreateBulk.OnConflict
// documentation for more info.
func (u *AppUpsertBulk) Update(set func(*AppUpsert)) *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AppUpsertBulk) SetName(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateName() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertBulk) SetDescription(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// SetDeviceID sets the "device_id" field.
func (u *AppUpsertBulk) SetDeviceID(v model.InternalID) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDeviceID(v)
	})
}

// AddDeviceID adds v to the "device_id" field.
func (u *AppUpsertBulk) AddDeviceID(v model.InternalID) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.AddDeviceID(v)
	})
}

// UpdateDeviceID sets the "device_id" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDeviceID() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeviceID()
	})
}

// SetPublic sets the "public" field.
func (u *AppUpsertBulk) SetPublic(v bool) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetPublic(v)
	})
}

// UpdatePublic sets the "public" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdatePublic() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdatePublic()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertBulk) SetUpdatedAt(v time.Time) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateUpdatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertBulk) SetCreatedAt(v time.Time) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateCreatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
