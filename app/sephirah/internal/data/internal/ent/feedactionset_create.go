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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedactionset"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modelsupervisor"
	"github.com/tuihub/librarian/internal/model"
)

// FeedActionSetCreate is the builder for creating a FeedActionSet entity.
type FeedActionSetCreate struct {
	config
	mutation *FeedActionSetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (fasc *FeedActionSetCreate) SetName(s string) *FeedActionSetCreate {
	fasc.mutation.SetName(s)
	return fasc
}

// SetDescription sets the "description" field.
func (fasc *FeedActionSetCreate) SetDescription(s string) *FeedActionSetCreate {
	fasc.mutation.SetDescription(s)
	return fasc
}

// SetActions sets the "actions" field.
func (fasc *FeedActionSetCreate) SetActions(mr []*modelsupervisor.FeatureRequest) *FeedActionSetCreate {
	fasc.mutation.SetActions(mr)
	return fasc
}

// SetUpdatedAt sets the "updated_at" field.
func (fasc *FeedActionSetCreate) SetUpdatedAt(t time.Time) *FeedActionSetCreate {
	fasc.mutation.SetUpdatedAt(t)
	return fasc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fasc *FeedActionSetCreate) SetNillableUpdatedAt(t *time.Time) *FeedActionSetCreate {
	if t != nil {
		fasc.SetUpdatedAt(*t)
	}
	return fasc
}

// SetCreatedAt sets the "created_at" field.
func (fasc *FeedActionSetCreate) SetCreatedAt(t time.Time) *FeedActionSetCreate {
	fasc.mutation.SetCreatedAt(t)
	return fasc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fasc *FeedActionSetCreate) SetNillableCreatedAt(t *time.Time) *FeedActionSetCreate {
	if t != nil {
		fasc.SetCreatedAt(*t)
	}
	return fasc
}

// SetID sets the "id" field.
func (fasc *FeedActionSetCreate) SetID(mi model.InternalID) *FeedActionSetCreate {
	fasc.mutation.SetID(mi)
	return fasc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fasc *FeedActionSetCreate) SetOwnerID(id model.InternalID) *FeedActionSetCreate {
	fasc.mutation.SetOwnerID(id)
	return fasc
}

// SetOwner sets the "owner" edge to the User entity.
func (fasc *FeedActionSetCreate) SetOwner(u *User) *FeedActionSetCreate {
	return fasc.SetOwnerID(u.ID)
}

// AddFeedConfigIDs adds the "feed_config" edge to the FeedConfig entity by IDs.
func (fasc *FeedActionSetCreate) AddFeedConfigIDs(ids ...model.InternalID) *FeedActionSetCreate {
	fasc.mutation.AddFeedConfigIDs(ids...)
	return fasc
}

// AddFeedConfig adds the "feed_config" edges to the FeedConfig entity.
func (fasc *FeedActionSetCreate) AddFeedConfig(f ...*FeedConfig) *FeedActionSetCreate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fasc.AddFeedConfigIDs(ids...)
}

// Mutation returns the FeedActionSetMutation object of the builder.
func (fasc *FeedActionSetCreate) Mutation() *FeedActionSetMutation {
	return fasc.mutation
}

// Save creates the FeedActionSet in the database.
func (fasc *FeedActionSetCreate) Save(ctx context.Context) (*FeedActionSet, error) {
	fasc.defaults()
	return withHooks(ctx, fasc.sqlSave, fasc.mutation, fasc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fasc *FeedActionSetCreate) SaveX(ctx context.Context) *FeedActionSet {
	v, err := fasc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fasc *FeedActionSetCreate) Exec(ctx context.Context) error {
	_, err := fasc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fasc *FeedActionSetCreate) ExecX(ctx context.Context) {
	if err := fasc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fasc *FeedActionSetCreate) defaults() {
	if _, ok := fasc.mutation.UpdatedAt(); !ok {
		v := feedactionset.DefaultUpdatedAt()
		fasc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fasc.mutation.CreatedAt(); !ok {
		v := feedactionset.DefaultCreatedAt()
		fasc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fasc *FeedActionSetCreate) check() error {
	if _, ok := fasc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FeedActionSet.name"`)}
	}
	if _, ok := fasc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "FeedActionSet.description"`)}
	}
	if _, ok := fasc.mutation.Actions(); !ok {
		return &ValidationError{Name: "actions", err: errors.New(`ent: missing required field "FeedActionSet.actions"`)}
	}
	if _, ok := fasc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeedActionSet.updated_at"`)}
	}
	if _, ok := fasc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeedActionSet.created_at"`)}
	}
	if _, ok := fasc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "FeedActionSet.owner"`)}
	}
	return nil
}

func (fasc *FeedActionSetCreate) sqlSave(ctx context.Context) (*FeedActionSet, error) {
	if err := fasc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fasc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fasc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	fasc.mutation.id = &_node.ID
	fasc.mutation.done = true
	return _node, nil
}

func (fasc *FeedActionSetCreate) createSpec() (*FeedActionSet, *sqlgraph.CreateSpec) {
	var (
		_node = &FeedActionSet{config: fasc.config}
		_spec = sqlgraph.NewCreateSpec(feedactionset.Table, sqlgraph.NewFieldSpec(feedactionset.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = fasc.conflict
	if id, ok := fasc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fasc.mutation.Name(); ok {
		_spec.SetField(feedactionset.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fasc.mutation.Description(); ok {
		_spec.SetField(feedactionset.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := fasc.mutation.Actions(); ok {
		_spec.SetField(feedactionset.FieldActions, field.TypeJSON, value)
		_node.Actions = value
	}
	if value, ok := fasc.mutation.UpdatedAt(); ok {
		_spec.SetField(feedactionset.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fasc.mutation.CreatedAt(); ok {
		_spec.SetField(feedactionset.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := fasc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feedactionset.OwnerTable,
			Columns: []string{feedactionset.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_feed_action_set = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fasc.mutation.FeedConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feedactionset.FeedConfigTable,
			Columns: feedactionset.FeedConfigPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
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
//	client.FeedActionSet.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedActionSetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fasc *FeedActionSetCreate) OnConflict(opts ...sql.ConflictOption) *FeedActionSetUpsertOne {
	fasc.conflict = opts
	return &FeedActionSetUpsertOne{
		create: fasc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fasc *FeedActionSetCreate) OnConflictColumns(columns ...string) *FeedActionSetUpsertOne {
	fasc.conflict = append(fasc.conflict, sql.ConflictColumns(columns...))
	return &FeedActionSetUpsertOne{
		create: fasc,
	}
}

type (
	// FeedActionSetUpsertOne is the builder for "upsert"-ing
	//  one FeedActionSet node.
	FeedActionSetUpsertOne struct {
		create *FeedActionSetCreate
	}

	// FeedActionSetUpsert is the "OnConflict" setter.
	FeedActionSetUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *FeedActionSetUpsert) SetName(v string) *FeedActionSetUpsert {
	u.Set(feedactionset.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedActionSetUpsert) UpdateName() *FeedActionSetUpsert {
	u.SetExcluded(feedactionset.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *FeedActionSetUpsert) SetDescription(v string) *FeedActionSetUpsert {
	u.Set(feedactionset.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedActionSetUpsert) UpdateDescription() *FeedActionSetUpsert {
	u.SetExcluded(feedactionset.FieldDescription)
	return u
}

// SetActions sets the "actions" field.
func (u *FeedActionSetUpsert) SetActions(v []*modelsupervisor.FeatureRequest) *FeedActionSetUpsert {
	u.Set(feedactionset.FieldActions, v)
	return u
}

// UpdateActions sets the "actions" field to the value that was provided on create.
func (u *FeedActionSetUpsert) UpdateActions() *FeedActionSetUpsert {
	u.SetExcluded(feedactionset.FieldActions)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedActionSetUpsert) SetUpdatedAt(v time.Time) *FeedActionSetUpsert {
	u.Set(feedactionset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedActionSetUpsert) UpdateUpdatedAt() *FeedActionSetUpsert {
	u.SetExcluded(feedactionset.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedActionSetUpsert) SetCreatedAt(v time.Time) *FeedActionSetUpsert {
	u.Set(feedactionset.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedActionSetUpsert) UpdateCreatedAt() *FeedActionSetUpsert {
	u.SetExcluded(feedactionset.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedactionset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedActionSetUpsertOne) UpdateNewValues() *FeedActionSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(feedactionset.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FeedActionSetUpsertOne) Ignore() *FeedActionSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedActionSetUpsertOne) DoNothing() *FeedActionSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedActionSetCreate.OnConflict
// documentation for more info.
func (u *FeedActionSetUpsertOne) Update(set func(*FeedActionSetUpsert)) *FeedActionSetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedActionSetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FeedActionSetUpsertOne) SetName(v string) *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedActionSetUpsertOne) UpdateName() *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *FeedActionSetUpsertOne) SetDescription(v string) *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedActionSetUpsertOne) UpdateDescription() *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateDescription()
	})
}

// SetActions sets the "actions" field.
func (u *FeedActionSetUpsertOne) SetActions(v []*modelsupervisor.FeatureRequest) *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetActions(v)
	})
}

// UpdateActions sets the "actions" field to the value that was provided on create.
func (u *FeedActionSetUpsertOne) UpdateActions() *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateActions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedActionSetUpsertOne) SetUpdatedAt(v time.Time) *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedActionSetUpsertOne) UpdateUpdatedAt() *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedActionSetUpsertOne) SetCreatedAt(v time.Time) *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedActionSetUpsertOne) UpdateCreatedAt() *FeedActionSetUpsertOne {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedActionSetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedActionSetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedActionSetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FeedActionSetUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FeedActionSetUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FeedActionSetCreateBulk is the builder for creating many FeedActionSet entities in bulk.
type FeedActionSetCreateBulk struct {
	config
	err      error
	builders []*FeedActionSetCreate
	conflict []sql.ConflictOption
}

// Save creates the FeedActionSet entities in the database.
func (fascb *FeedActionSetCreateBulk) Save(ctx context.Context) ([]*FeedActionSet, error) {
	if fascb.err != nil {
		return nil, fascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fascb.builders))
	nodes := make([]*FeedActionSet, len(fascb.builders))
	mutators := make([]Mutator, len(fascb.builders))
	for i := range fascb.builders {
		func(i int, root context.Context) {
			builder := fascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedActionSetMutation)
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
					_, err = mutators[i+1].Mutate(root, fascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fascb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fascb *FeedActionSetCreateBulk) SaveX(ctx context.Context) []*FeedActionSet {
	v, err := fascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fascb *FeedActionSetCreateBulk) Exec(ctx context.Context) error {
	_, err := fascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fascb *FeedActionSetCreateBulk) ExecX(ctx context.Context) {
	if err := fascb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FeedActionSet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FeedActionSetUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fascb *FeedActionSetCreateBulk) OnConflict(opts ...sql.ConflictOption) *FeedActionSetUpsertBulk {
	fascb.conflict = opts
	return &FeedActionSetUpsertBulk{
		create: fascb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fascb *FeedActionSetCreateBulk) OnConflictColumns(columns ...string) *FeedActionSetUpsertBulk {
	fascb.conflict = append(fascb.conflict, sql.ConflictColumns(columns...))
	return &FeedActionSetUpsertBulk{
		create: fascb,
	}
}

// FeedActionSetUpsertBulk is the builder for "upsert"-ing
// a bulk of FeedActionSet nodes.
type FeedActionSetUpsertBulk struct {
	create *FeedActionSetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(feedactionset.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FeedActionSetUpsertBulk) UpdateNewValues() *FeedActionSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(feedactionset.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FeedActionSet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FeedActionSetUpsertBulk) Ignore() *FeedActionSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FeedActionSetUpsertBulk) DoNothing() *FeedActionSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FeedActionSetCreateBulk.OnConflict
// documentation for more info.
func (u *FeedActionSetUpsertBulk) Update(set func(*FeedActionSetUpsert)) *FeedActionSetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FeedActionSetUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *FeedActionSetUpsertBulk) SetName(v string) *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FeedActionSetUpsertBulk) UpdateName() *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *FeedActionSetUpsertBulk) SetDescription(v string) *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *FeedActionSetUpsertBulk) UpdateDescription() *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateDescription()
	})
}

// SetActions sets the "actions" field.
func (u *FeedActionSetUpsertBulk) SetActions(v []*modelsupervisor.FeatureRequest) *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetActions(v)
	})
}

// UpdateActions sets the "actions" field to the value that was provided on create.
func (u *FeedActionSetUpsertBulk) UpdateActions() *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateActions()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FeedActionSetUpsertBulk) SetUpdatedAt(v time.Time) *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FeedActionSetUpsertBulk) UpdateUpdatedAt() *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *FeedActionSetUpsertBulk) SetCreatedAt(v time.Time) *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FeedActionSetUpsertBulk) UpdateCreatedAt() *FeedActionSetUpsertBulk {
	return u.Update(func(s *FeedActionSetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *FeedActionSetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FeedActionSetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FeedActionSetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FeedActionSetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
