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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyFlowTargetCreate is the builder for creating a NotifyFlowTarget entity.
type NotifyFlowTargetCreate struct {
	config
	mutation *NotifyFlowTargetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (nftc *NotifyFlowTargetCreate) SetNotifyFlowID(mi model.InternalID) *NotifyFlowTargetCreate {
	nftc.mutation.SetNotifyFlowID(mi)
	return nftc
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (nftc *NotifyFlowTargetCreate) SetNotifyTargetID(mi model.InternalID) *NotifyFlowTargetCreate {
	nftc.mutation.SetNotifyTargetID(mi)
	return nftc
}

// SetChannelID sets the "channel_id" field.
func (nftc *NotifyFlowTargetCreate) SetChannelID(s string) *NotifyFlowTargetCreate {
	nftc.mutation.SetChannelID(s)
	return nftc
}

// SetUpdatedAt sets the "updated_at" field.
func (nftc *NotifyFlowTargetCreate) SetUpdatedAt(t time.Time) *NotifyFlowTargetCreate {
	nftc.mutation.SetUpdatedAt(t)
	return nftc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nftc *NotifyFlowTargetCreate) SetNillableUpdatedAt(t *time.Time) *NotifyFlowTargetCreate {
	if t != nil {
		nftc.SetUpdatedAt(*t)
	}
	return nftc
}

// SetCreatedAt sets the "created_at" field.
func (nftc *NotifyFlowTargetCreate) SetCreatedAt(t time.Time) *NotifyFlowTargetCreate {
	nftc.mutation.SetCreatedAt(t)
	return nftc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nftc *NotifyFlowTargetCreate) SetNillableCreatedAt(t *time.Time) *NotifyFlowTargetCreate {
	if t != nil {
		nftc.SetCreatedAt(*t)
	}
	return nftc
}

// SetNotifyFlow sets the "notify_flow" edge to the NotifyFlow entity.
func (nftc *NotifyFlowTargetCreate) SetNotifyFlow(n *NotifyFlow) *NotifyFlowTargetCreate {
	return nftc.SetNotifyFlowID(n.ID)
}

// SetNotifyTarget sets the "notify_target" edge to the NotifyTarget entity.
func (nftc *NotifyFlowTargetCreate) SetNotifyTarget(n *NotifyTarget) *NotifyFlowTargetCreate {
	return nftc.SetNotifyTargetID(n.ID)
}

// Mutation returns the NotifyFlowTargetMutation object of the builder.
func (nftc *NotifyFlowTargetCreate) Mutation() *NotifyFlowTargetMutation {
	return nftc.mutation
}

// Save creates the NotifyFlowTarget in the database.
func (nftc *NotifyFlowTargetCreate) Save(ctx context.Context) (*NotifyFlowTarget, error) {
	nftc.defaults()
	return withHooks[*NotifyFlowTarget, NotifyFlowTargetMutation](ctx, nftc.sqlSave, nftc.mutation, nftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nftc *NotifyFlowTargetCreate) SaveX(ctx context.Context) *NotifyFlowTarget {
	v, err := nftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nftc *NotifyFlowTargetCreate) Exec(ctx context.Context) error {
	_, err := nftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nftc *NotifyFlowTargetCreate) ExecX(ctx context.Context) {
	if err := nftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nftc *NotifyFlowTargetCreate) defaults() {
	if _, ok := nftc.mutation.UpdatedAt(); !ok {
		v := notifyflowtarget.DefaultUpdatedAt()
		nftc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nftc.mutation.CreatedAt(); !ok {
		v := notifyflowtarget.DefaultCreatedAt()
		nftc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nftc *NotifyFlowTargetCreate) check() error {
	if _, ok := nftc.mutation.NotifyFlowID(); !ok {
		return &ValidationError{Name: "notify_flow_id", err: errors.New(`ent: missing required field "NotifyFlowTarget.notify_flow_id"`)}
	}
	if _, ok := nftc.mutation.NotifyTargetID(); !ok {
		return &ValidationError{Name: "notify_target_id", err: errors.New(`ent: missing required field "NotifyFlowTarget.notify_target_id"`)}
	}
	if _, ok := nftc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "NotifyFlowTarget.channel_id"`)}
	}
	if _, ok := nftc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotifyFlowTarget.updated_at"`)}
	}
	if _, ok := nftc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifyFlowTarget.created_at"`)}
	}
	if _, ok := nftc.mutation.NotifyFlowID(); !ok {
		return &ValidationError{Name: "notify_flow", err: errors.New(`ent: missing required edge "NotifyFlowTarget.notify_flow"`)}
	}
	if _, ok := nftc.mutation.NotifyTargetID(); !ok {
		return &ValidationError{Name: "notify_target", err: errors.New(`ent: missing required edge "NotifyFlowTarget.notify_target"`)}
	}
	return nil
}

func (nftc *NotifyFlowTargetCreate) sqlSave(ctx context.Context) (*NotifyFlowTarget, error) {
	if err := nftc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nftc.mutation.id = &_node.ID
	nftc.mutation.done = true
	return _node, nil
}

func (nftc *NotifyFlowTargetCreate) createSpec() (*NotifyFlowTarget, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifyFlowTarget{config: nftc.config}
		_spec = sqlgraph.NewCreateSpec(notifyflowtarget.Table, sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt))
	)
	_spec.OnConflict = nftc.conflict
	if value, ok := nftc.mutation.ChannelID(); ok {
		_spec.SetField(notifyflowtarget.FieldChannelID, field.TypeString, value)
		_node.ChannelID = value
	}
	if value, ok := nftc.mutation.UpdatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nftc.mutation.CreatedAt(); ok {
		_spec.SetField(notifyflowtarget.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := nftc.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyFlowTable,
			Columns: []string{notifyflowtarget.NotifyFlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NotifyFlowID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nftc.mutation.NotifyTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   notifyflowtarget.NotifyTargetTable,
			Columns: []string{notifyflowtarget.NotifyTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NotifyTargetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlowTarget.Create().
//		SetNotifyFlowID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowTargetUpsert) {
//			SetNotifyFlowID(v+v).
//		}).
//		Exec(ctx)
func (nftc *NotifyFlowTargetCreate) OnConflict(opts ...sql.ConflictOption) *NotifyFlowTargetUpsertOne {
	nftc.conflict = opts
	return &NotifyFlowTargetUpsertOne{
		create: nftc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nftc *NotifyFlowTargetCreate) OnConflictColumns(columns ...string) *NotifyFlowTargetUpsertOne {
	nftc.conflict = append(nftc.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowTargetUpsertOne{
		create: nftc,
	}
}

type (
	// NotifyFlowTargetUpsertOne is the builder for "upsert"-ing
	//  one NotifyFlowTarget node.
	NotifyFlowTargetUpsertOne struct {
		create *NotifyFlowTargetCreate
	}

	// NotifyFlowTargetUpsert is the "OnConflict" setter.
	NotifyFlowTargetUpsert struct {
		*sql.UpdateSet
	}
)

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowTargetUpsert) SetNotifyFlowID(v model.InternalID) *NotifyFlowTargetUpsert {
	u.Set(notifyflowtarget.FieldNotifyFlowID, v)
	return u
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsert) UpdateNotifyFlowID() *NotifyFlowTargetUpsert {
	u.SetExcluded(notifyflowtarget.FieldNotifyFlowID)
	return u
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (u *NotifyFlowTargetUpsert) SetNotifyTargetID(v model.InternalID) *NotifyFlowTargetUpsert {
	u.Set(notifyflowtarget.FieldNotifyTargetID, v)
	return u
}

// UpdateNotifyTargetID sets the "notify_target_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsert) UpdateNotifyTargetID() *NotifyFlowTargetUpsert {
	u.SetExcluded(notifyflowtarget.FieldNotifyTargetID)
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *NotifyFlowTargetUpsert) SetChannelID(v string) *NotifyFlowTargetUpsert {
	u.Set(notifyflowtarget.FieldChannelID, v)
	return u
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsert) UpdateChannelID() *NotifyFlowTargetUpsert {
	u.SetExcluded(notifyflowtarget.FieldChannelID)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowTargetUpsert) SetUpdatedAt(v time.Time) *NotifyFlowTargetUpsert {
	u.Set(notifyflowtarget.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsert) UpdateUpdatedAt() *NotifyFlowTargetUpsert {
	u.SetExcluded(notifyflowtarget.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowTargetUpsert) SetCreatedAt(v time.Time) *NotifyFlowTargetUpsert {
	u.Set(notifyflowtarget.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsert) UpdateCreatedAt() *NotifyFlowTargetUpsert {
	u.SetExcluded(notifyflowtarget.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NotifyFlowTargetUpsertOne) UpdateNewValues() *NotifyFlowTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotifyFlowTargetUpsertOne) Ignore() *NotifyFlowTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowTargetUpsertOne) DoNothing() *NotifyFlowTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowTargetCreate.OnConflict
// documentation for more info.
func (u *NotifyFlowTargetUpsertOne) Update(set func(*NotifyFlowTargetUpsert)) *NotifyFlowTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowTargetUpsert{UpdateSet: update})
	}))
	return u
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowTargetUpsertOne) SetNotifyFlowID(v model.InternalID) *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetNotifyFlowID(v)
	})
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertOne) UpdateNotifyFlowID() *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateNotifyFlowID()
	})
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (u *NotifyFlowTargetUpsertOne) SetNotifyTargetID(v model.InternalID) *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetNotifyTargetID(v)
	})
}

// UpdateNotifyTargetID sets the "notify_target_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertOne) UpdateNotifyTargetID() *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateNotifyTargetID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *NotifyFlowTargetUpsertOne) SetChannelID(v string) *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertOne) UpdateChannelID() *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateChannelID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowTargetUpsertOne) SetUpdatedAt(v time.Time) *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertOne) UpdateUpdatedAt() *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowTargetUpsertOne) SetCreatedAt(v time.Time) *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertOne) UpdateCreatedAt() *NotifyFlowTargetUpsertOne {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowTargetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowTargetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowTargetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifyFlowTargetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifyFlowTargetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifyFlowTargetCreateBulk is the builder for creating many NotifyFlowTarget entities in bulk.
type NotifyFlowTargetCreateBulk struct {
	config
	builders []*NotifyFlowTargetCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifyFlowTarget entities in the database.
func (nftcb *NotifyFlowTargetCreateBulk) Save(ctx context.Context) ([]*NotifyFlowTarget, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nftcb.builders))
	nodes := make([]*NotifyFlowTarget, len(nftcb.builders))
	mutators := make([]Mutator, len(nftcb.builders))
	for i := range nftcb.builders {
		func(i int, root context.Context) {
			builder := nftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifyFlowTargetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = nftcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nftcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nftcb *NotifyFlowTargetCreateBulk) SaveX(ctx context.Context) []*NotifyFlowTarget {
	v, err := nftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nftcb *NotifyFlowTargetCreateBulk) Exec(ctx context.Context) error {
	_, err := nftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nftcb *NotifyFlowTargetCreateBulk) ExecX(ctx context.Context) {
	if err := nftcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyFlowTarget.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyFlowTargetUpsert) {
//			SetNotifyFlowID(v+v).
//		}).
//		Exec(ctx)
func (nftcb *NotifyFlowTargetCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifyFlowTargetUpsertBulk {
	nftcb.conflict = opts
	return &NotifyFlowTargetUpsertBulk{
		create: nftcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (nftcb *NotifyFlowTargetCreateBulk) OnConflictColumns(columns ...string) *NotifyFlowTargetUpsertBulk {
	nftcb.conflict = append(nftcb.conflict, sql.ConflictColumns(columns...))
	return &NotifyFlowTargetUpsertBulk{
		create: nftcb,
	}
}

// NotifyFlowTargetUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifyFlowTarget nodes.
type NotifyFlowTargetUpsertBulk struct {
	create *NotifyFlowTargetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *NotifyFlowTargetUpsertBulk) UpdateNewValues() *NotifyFlowTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyFlowTarget.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotifyFlowTargetUpsertBulk) Ignore() *NotifyFlowTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyFlowTargetUpsertBulk) DoNothing() *NotifyFlowTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyFlowTargetCreateBulk.OnConflict
// documentation for more info.
func (u *NotifyFlowTargetUpsertBulk) Update(set func(*NotifyFlowTargetUpsert)) *NotifyFlowTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyFlowTargetUpsert{UpdateSet: update})
	}))
	return u
}

// SetNotifyFlowID sets the "notify_flow_id" field.
func (u *NotifyFlowTargetUpsertBulk) SetNotifyFlowID(v model.InternalID) *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetNotifyFlowID(v)
	})
}

// UpdateNotifyFlowID sets the "notify_flow_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertBulk) UpdateNotifyFlowID() *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateNotifyFlowID()
	})
}

// SetNotifyTargetID sets the "notify_target_id" field.
func (u *NotifyFlowTargetUpsertBulk) SetNotifyTargetID(v model.InternalID) *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetNotifyTargetID(v)
	})
}

// UpdateNotifyTargetID sets the "notify_target_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertBulk) UpdateNotifyTargetID() *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateNotifyTargetID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *NotifyFlowTargetUpsertBulk) SetChannelID(v string) *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertBulk) UpdateChannelID() *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateChannelID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyFlowTargetUpsertBulk) SetUpdatedAt(v time.Time) *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertBulk) UpdateUpdatedAt() *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyFlowTargetUpsertBulk) SetCreatedAt(v time.Time) *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyFlowTargetUpsertBulk) UpdateCreatedAt() *NotifyFlowTargetUpsertBulk {
	return u.Update(func(s *NotifyFlowTargetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyFlowTargetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifyFlowTargetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyFlowTargetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyFlowTargetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}