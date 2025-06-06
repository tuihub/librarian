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
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
	"github.com/tuihub/librarian/internal/data/internal/ent/storeapp"
	"github.com/tuihub/librarian/internal/data/internal/ent/storeappbinary"
	"github.com/tuihub/librarian/internal/model"
)

// StoreAppCreate is the builder for creating a StoreApp entity.
type StoreAppCreate struct {
	config
	mutation *StoreAppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sac *StoreAppCreate) SetName(s string) *StoreAppCreate {
	sac.mutation.SetName(s)
	return sac
}

// SetDescription sets the "description" field.
func (sac *StoreAppCreate) SetDescription(s string) *StoreAppCreate {
	sac.mutation.SetDescription(s)
	return sac
}

// SetUpdatedAt sets the "updated_at" field.
func (sac *StoreAppCreate) SetUpdatedAt(t time.Time) *StoreAppCreate {
	sac.mutation.SetUpdatedAt(t)
	return sac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sac *StoreAppCreate) SetNillableUpdatedAt(t *time.Time) *StoreAppCreate {
	if t != nil {
		sac.SetUpdatedAt(*t)
	}
	return sac
}

// SetCreatedAt sets the "created_at" field.
func (sac *StoreAppCreate) SetCreatedAt(t time.Time) *StoreAppCreate {
	sac.mutation.SetCreatedAt(t)
	return sac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sac *StoreAppCreate) SetNillableCreatedAt(t *time.Time) *StoreAppCreate {
	if t != nil {
		sac.SetCreatedAt(*t)
	}
	return sac
}

// SetID sets the "id" field.
func (sac *StoreAppCreate) SetID(mi model.InternalID) *StoreAppCreate {
	sac.mutation.SetID(mi)
	return sac
}

// AddAppBinaryIDs adds the "app_binary" edge to the SentinelAppBinary entity by IDs.
func (sac *StoreAppCreate) AddAppBinaryIDs(ids ...model.InternalID) *StoreAppCreate {
	sac.mutation.AddAppBinaryIDs(ids...)
	return sac
}

// AddAppBinary adds the "app_binary" edges to the SentinelAppBinary entity.
func (sac *StoreAppCreate) AddAppBinary(s ...*SentinelAppBinary) *StoreAppCreate {
	ids := make([]model.InternalID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sac.AddAppBinaryIDs(ids...)
}

// AddStoreAppBinaryIDs adds the "store_app_binary" edge to the StoreAppBinary entity by IDs.
func (sac *StoreAppCreate) AddStoreAppBinaryIDs(ids ...int) *StoreAppCreate {
	sac.mutation.AddStoreAppBinaryIDs(ids...)
	return sac
}

// AddStoreAppBinary adds the "store_app_binary" edges to the StoreAppBinary entity.
func (sac *StoreAppCreate) AddStoreAppBinary(s ...*StoreAppBinary) *StoreAppCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sac.AddStoreAppBinaryIDs(ids...)
}

// Mutation returns the StoreAppMutation object of the builder.
func (sac *StoreAppCreate) Mutation() *StoreAppMutation {
	return sac.mutation
}

// Save creates the StoreApp in the database.
func (sac *StoreAppCreate) Save(ctx context.Context) (*StoreApp, error) {
	sac.defaults()
	return withHooks(ctx, sac.sqlSave, sac.mutation, sac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sac *StoreAppCreate) SaveX(ctx context.Context) *StoreApp {
	v, err := sac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sac *StoreAppCreate) Exec(ctx context.Context) error {
	_, err := sac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sac *StoreAppCreate) ExecX(ctx context.Context) {
	if err := sac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sac *StoreAppCreate) defaults() {
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		v := storeapp.DefaultUpdatedAt()
		sac.mutation.SetUpdatedAt(v)
	}
	if _, ok := sac.mutation.CreatedAt(); !ok {
		v := storeapp.DefaultCreatedAt()
		sac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sac *StoreAppCreate) check() error {
	if _, ok := sac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "StoreApp.name"`)}
	}
	if _, ok := sac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "StoreApp.description"`)}
	}
	if _, ok := sac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StoreApp.updated_at"`)}
	}
	if _, ok := sac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StoreApp.created_at"`)}
	}
	return nil
}

func (sac *StoreAppCreate) sqlSave(ctx context.Context) (*StoreApp, error) {
	if err := sac.check(); err != nil {
		return nil, err
	}
	_node, _spec := sac.createSpec()
	if err := sqlgraph.CreateNode(ctx, sac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	sac.mutation.id = &_node.ID
	sac.mutation.done = true
	return _node, nil
}

func (sac *StoreAppCreate) createSpec() (*StoreApp, *sqlgraph.CreateSpec) {
	var (
		_node = &StoreApp{config: sac.config}
		_spec = sqlgraph.NewCreateSpec(storeapp.Table, sqlgraph.NewFieldSpec(storeapp.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = sac.conflict
	if id, ok := sac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sac.mutation.Name(); ok {
		_spec.SetField(storeapp.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sac.mutation.Description(); ok {
		_spec.SetField(storeapp.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sac.mutation.UpdatedAt(); ok {
		_spec.SetField(storeapp.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sac.mutation.CreatedAt(); ok {
		_spec.SetField(storeapp.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sac.mutation.AppBinaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   storeapp.AppBinaryTable,
			Columns: storeapp.AppBinaryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &StoreAppBinaryCreate{config: sac.config, mutation: newStoreAppBinaryMutation(sac.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sac.mutation.StoreAppBinaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   storeapp.StoreAppBinaryTable,
			Columns: []string{storeapp.StoreAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(storeappbinary.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StoreApp.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StoreAppUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sac *StoreAppCreate) OnConflict(opts ...sql.ConflictOption) *StoreAppUpsertOne {
	sac.conflict = opts
	return &StoreAppUpsertOne{
		create: sac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sac *StoreAppCreate) OnConflictColumns(columns ...string) *StoreAppUpsertOne {
	sac.conflict = append(sac.conflict, sql.ConflictColumns(columns...))
	return &StoreAppUpsertOne{
		create: sac,
	}
}

type (
	// StoreAppUpsertOne is the builder for "upsert"-ing
	//  one StoreApp node.
	StoreAppUpsertOne struct {
		create *StoreAppCreate
	}

	// StoreAppUpsert is the "OnConflict" setter.
	StoreAppUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StoreAppUpsert) SetName(v string) *StoreAppUpsert {
	u.Set(storeapp.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppUpsert) UpdateName() *StoreAppUpsert {
	u.SetExcluded(storeapp.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *StoreAppUpsert) SetDescription(v string) *StoreAppUpsert {
	u.Set(storeapp.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *StoreAppUpsert) UpdateDescription() *StoreAppUpsert {
	u.SetExcluded(storeapp.FieldDescription)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppUpsert) SetUpdatedAt(v time.Time) *StoreAppUpsert {
	u.Set(storeapp.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppUpsert) UpdateUpdatedAt() *StoreAppUpsert {
	u.SetExcluded(storeapp.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppUpsert) SetCreatedAt(v time.Time) *StoreAppUpsert {
	u.Set(storeapp.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppUpsert) UpdateCreatedAt() *StoreAppUpsert {
	u.SetExcluded(storeapp.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(storeapp.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StoreAppUpsertOne) UpdateNewValues() *StoreAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(storeapp.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StoreAppUpsertOne) Ignore() *StoreAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StoreAppUpsertOne) DoNothing() *StoreAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StoreAppCreate.OnConflict
// documentation for more info.
func (u *StoreAppUpsertOne) Update(set func(*StoreAppUpsert)) *StoreAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StoreAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StoreAppUpsertOne) SetName(v string) *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppUpsertOne) UpdateName() *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *StoreAppUpsertOne) SetDescription(v string) *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *StoreAppUpsertOne) UpdateDescription() *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateDescription()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppUpsertOne) SetUpdatedAt(v time.Time) *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppUpsertOne) UpdateUpdatedAt() *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppUpsertOne) SetCreatedAt(v time.Time) *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppUpsertOne) UpdateCreatedAt() *StoreAppUpsertOne {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StoreAppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StoreAppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StoreAppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StoreAppUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StoreAppUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StoreAppCreateBulk is the builder for creating many StoreApp entities in bulk.
type StoreAppCreateBulk struct {
	config
	err      error
	builders []*StoreAppCreate
	conflict []sql.ConflictOption
}

// Save creates the StoreApp entities in the database.
func (sacb *StoreAppCreateBulk) Save(ctx context.Context) ([]*StoreApp, error) {
	if sacb.err != nil {
		return nil, sacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sacb.builders))
	nodes := make([]*StoreApp, len(sacb.builders))
	mutators := make([]Mutator, len(sacb.builders))
	for i := range sacb.builders {
		func(i int, root context.Context) {
			builder := sacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoreAppMutation)
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
					_, err = mutators[i+1].Mutate(root, sacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sacb *StoreAppCreateBulk) SaveX(ctx context.Context) []*StoreApp {
	v, err := sacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sacb *StoreAppCreateBulk) Exec(ctx context.Context) error {
	_, err := sacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sacb *StoreAppCreateBulk) ExecX(ctx context.Context) {
	if err := sacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StoreApp.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StoreAppUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sacb *StoreAppCreateBulk) OnConflict(opts ...sql.ConflictOption) *StoreAppUpsertBulk {
	sacb.conflict = opts
	return &StoreAppUpsertBulk{
		create: sacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sacb *StoreAppCreateBulk) OnConflictColumns(columns ...string) *StoreAppUpsertBulk {
	sacb.conflict = append(sacb.conflict, sql.ConflictColumns(columns...))
	return &StoreAppUpsertBulk{
		create: sacb,
	}
}

// StoreAppUpsertBulk is the builder for "upsert"-ing
// a bulk of StoreApp nodes.
type StoreAppUpsertBulk struct {
	create *StoreAppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(storeapp.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StoreAppUpsertBulk) UpdateNewValues() *StoreAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(storeapp.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StoreApp.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StoreAppUpsertBulk) Ignore() *StoreAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StoreAppUpsertBulk) DoNothing() *StoreAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StoreAppCreateBulk.OnConflict
// documentation for more info.
func (u *StoreAppUpsertBulk) Update(set func(*StoreAppUpsert)) *StoreAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StoreAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StoreAppUpsertBulk) SetName(v string) *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppUpsertBulk) UpdateName() *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *StoreAppUpsertBulk) SetDescription(v string) *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *StoreAppUpsertBulk) UpdateDescription() *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateDescription()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppUpsertBulk) SetUpdatedAt(v time.Time) *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppUpsertBulk) UpdateUpdatedAt() *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppUpsertBulk) SetCreatedAt(v time.Time) *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppUpsertBulk) UpdateCreatedAt() *StoreAppUpsertBulk {
	return u.Update(func(s *StoreAppUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StoreAppUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StoreAppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StoreAppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StoreAppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
