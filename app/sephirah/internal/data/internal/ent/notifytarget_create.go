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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// NotifyTargetCreate is the builder for creating a NotifyTarget entity.
type NotifyTargetCreate struct {
	config
	mutation *NotifyTargetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetToken sets the "token" field.
func (ntc *NotifyTargetCreate) SetToken(s string) *NotifyTargetCreate {
	ntc.mutation.SetToken(s)
	return ntc
}

// SetName sets the "name" field.
func (ntc *NotifyTargetCreate) SetName(s string) *NotifyTargetCreate {
	ntc.mutation.SetName(s)
	return ntc
}

// SetDescription sets the "description" field.
func (ntc *NotifyTargetCreate) SetDescription(s string) *NotifyTargetCreate {
	ntc.mutation.SetDescription(s)
	return ntc
}

// SetType sets the "type" field.
func (ntc *NotifyTargetCreate) SetType(n notifytarget.Type) *NotifyTargetCreate {
	ntc.mutation.SetType(n)
	return ntc
}

// SetStatus sets the "status" field.
func (ntc *NotifyTargetCreate) SetStatus(n notifytarget.Status) *NotifyTargetCreate {
	ntc.mutation.SetStatus(n)
	return ntc
}

// SetUpdatedAt sets the "updated_at" field.
func (ntc *NotifyTargetCreate) SetUpdatedAt(t time.Time) *NotifyTargetCreate {
	ntc.mutation.SetUpdatedAt(t)
	return ntc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ntc *NotifyTargetCreate) SetNillableUpdatedAt(t *time.Time) *NotifyTargetCreate {
	if t != nil {
		ntc.SetUpdatedAt(*t)
	}
	return ntc
}

// SetCreatedAt sets the "created_at" field.
func (ntc *NotifyTargetCreate) SetCreatedAt(t time.Time) *NotifyTargetCreate {
	ntc.mutation.SetCreatedAt(t)
	return ntc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ntc *NotifyTargetCreate) SetNillableCreatedAt(t *time.Time) *NotifyTargetCreate {
	if t != nil {
		ntc.SetCreatedAt(*t)
	}
	return ntc
}

// SetID sets the "id" field.
func (ntc *NotifyTargetCreate) SetID(mi model.InternalID) *NotifyTargetCreate {
	ntc.mutation.SetID(mi)
	return ntc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ntc *NotifyTargetCreate) SetOwnerID(id model.InternalID) *NotifyTargetCreate {
	ntc.mutation.SetOwnerID(id)
	return ntc
}

// SetOwner sets the "owner" edge to the User entity.
func (ntc *NotifyTargetCreate) SetOwner(u *User) *NotifyTargetCreate {
	return ntc.SetOwnerID(u.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (ntc *NotifyTargetCreate) AddNotifyFlowIDs(ids ...model.InternalID) *NotifyTargetCreate {
	ntc.mutation.AddNotifyFlowIDs(ids...)
	return ntc
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (ntc *NotifyTargetCreate) AddNotifyFlow(n ...*NotifyFlow) *NotifyTargetCreate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntc.AddNotifyFlowIDs(ids...)
}

// AddNotifyFlowTargetIDs adds the "notify_flow_target" edge to the NotifyFlowTarget entity by IDs.
func (ntc *NotifyTargetCreate) AddNotifyFlowTargetIDs(ids ...int) *NotifyTargetCreate {
	ntc.mutation.AddNotifyFlowTargetIDs(ids...)
	return ntc
}

// AddNotifyFlowTarget adds the "notify_flow_target" edges to the NotifyFlowTarget entity.
func (ntc *NotifyTargetCreate) AddNotifyFlowTarget(n ...*NotifyFlowTarget) *NotifyTargetCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ntc.AddNotifyFlowTargetIDs(ids...)
}

// Mutation returns the NotifyTargetMutation object of the builder.
func (ntc *NotifyTargetCreate) Mutation() *NotifyTargetMutation {
	return ntc.mutation
}

// Save creates the NotifyTarget in the database.
func (ntc *NotifyTargetCreate) Save(ctx context.Context) (*NotifyTarget, error) {
	ntc.defaults()
	return withHooks(ctx, ntc.sqlSave, ntc.mutation, ntc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ntc *NotifyTargetCreate) SaveX(ctx context.Context) *NotifyTarget {
	v, err := ntc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntc *NotifyTargetCreate) Exec(ctx context.Context) error {
	_, err := ntc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntc *NotifyTargetCreate) ExecX(ctx context.Context) {
	if err := ntc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ntc *NotifyTargetCreate) defaults() {
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		v := notifytarget.DefaultUpdatedAt()
		ntc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		v := notifytarget.DefaultCreatedAt()
		ntc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ntc *NotifyTargetCreate) check() error {
	if _, ok := ntc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "NotifyTarget.token"`)}
	}
	if _, ok := ntc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "NotifyTarget.name"`)}
	}
	if _, ok := ntc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "NotifyTarget.description"`)}
	}
	if _, ok := ntc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "NotifyTarget.type"`)}
	}
	if v, ok := ntc.mutation.GetType(); ok {
		if err := notifytarget.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.type": %w`, err)}
		}
	}
	if _, ok := ntc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "NotifyTarget.status"`)}
	}
	if v, ok := ntc.mutation.Status(); ok {
		if err := notifytarget.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "NotifyTarget.status": %w`, err)}
		}
	}
	if _, ok := ntc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "NotifyTarget.updated_at"`)}
	}
	if _, ok := ntc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "NotifyTarget.created_at"`)}
	}
	if _, ok := ntc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "NotifyTarget.owner"`)}
	}
	return nil
}

func (ntc *NotifyTargetCreate) sqlSave(ctx context.Context) (*NotifyTarget, error) {
	if err := ntc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ntc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ntc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	ntc.mutation.id = &_node.ID
	ntc.mutation.done = true
	return _node, nil
}

func (ntc *NotifyTargetCreate) createSpec() (*NotifyTarget, *sqlgraph.CreateSpec) {
	var (
		_node = &NotifyTarget{config: ntc.config}
		_spec = sqlgraph.NewCreateSpec(notifytarget.Table, sqlgraph.NewFieldSpec(notifytarget.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ntc.conflict
	if id, ok := ntc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ntc.mutation.Token(); ok {
		_spec.SetField(notifytarget.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := ntc.mutation.Name(); ok {
		_spec.SetField(notifytarget.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ntc.mutation.Description(); ok {
		_spec.SetField(notifytarget.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ntc.mutation.GetType(); ok {
		_spec.SetField(notifytarget.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ntc.mutation.Status(); ok {
		_spec.SetField(notifytarget.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ntc.mutation.UpdatedAt(); ok {
		_spec.SetField(notifytarget.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ntc.mutation.CreatedAt(); ok {
		_spec.SetField(notifytarget.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ntc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifytarget.OwnerTable,
			Columns: []string{notifytarget.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_notify_target = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ntc.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTable,
			Columns: notifytarget.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowTargetCreate{config: ntc.config, mutation: newNotifyFlowTargetMutation(ntc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ntc.mutation.NotifyFlowTargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifytarget.NotifyFlowTargetTable,
			Columns: []string{notifytarget.NotifyFlowTargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowtarget.FieldID, field.TypeInt),
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
//	client.NotifyTarget.Create().
//		SetToken(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyTargetUpsert) {
//			SetToken(v+v).
//		}).
//		Exec(ctx)
func (ntc *NotifyTargetCreate) OnConflict(opts ...sql.ConflictOption) *NotifyTargetUpsertOne {
	ntc.conflict = opts
	return &NotifyTargetUpsertOne{
		create: ntc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ntc *NotifyTargetCreate) OnConflictColumns(columns ...string) *NotifyTargetUpsertOne {
	ntc.conflict = append(ntc.conflict, sql.ConflictColumns(columns...))
	return &NotifyTargetUpsertOne{
		create: ntc,
	}
}

type (
	// NotifyTargetUpsertOne is the builder for "upsert"-ing
	//  one NotifyTarget node.
	NotifyTargetUpsertOne struct {
		create *NotifyTargetCreate
	}

	// NotifyTargetUpsert is the "OnConflict" setter.
	NotifyTargetUpsert struct {
		*sql.UpdateSet
	}
)

// SetToken sets the "token" field.
func (u *NotifyTargetUpsert) SetToken(v string) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateToken() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldToken)
	return u
}

// SetName sets the "name" field.
func (u *NotifyTargetUpsert) SetName(v string) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateName() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *NotifyTargetUpsert) SetDescription(v string) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateDescription() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldDescription)
	return u
}

// SetType sets the "type" field.
func (u *NotifyTargetUpsert) SetType(v notifytarget.Type) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateType() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldType)
	return u
}

// SetStatus sets the "status" field.
func (u *NotifyTargetUpsert) SetStatus(v notifytarget.Status) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateStatus() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldStatus)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyTargetUpsert) SetUpdatedAt(v time.Time) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateUpdatedAt() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyTargetUpsert) SetCreatedAt(v time.Time) *NotifyTargetUpsert {
	u.Set(notifytarget.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyTargetUpsert) UpdateCreatedAt() *NotifyTargetUpsert {
	u.SetExcluded(notifytarget.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifytarget.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifyTargetUpsertOne) UpdateNewValues() *NotifyTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(notifytarget.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *NotifyTargetUpsertOne) Ignore() *NotifyTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyTargetUpsertOne) DoNothing() *NotifyTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyTargetCreate.OnConflict
// documentation for more info.
func (u *NotifyTargetUpsertOne) Update(set func(*NotifyTargetUpsert)) *NotifyTargetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyTargetUpsert{UpdateSet: update})
	}))
	return u
}

// SetToken sets the "token" field.
func (u *NotifyTargetUpsertOne) SetToken(v string) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateToken() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateToken()
	})
}

// SetName sets the "name" field.
func (u *NotifyTargetUpsertOne) SetName(v string) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateName() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *NotifyTargetUpsertOne) SetDescription(v string) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateDescription() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateDescription()
	})
}

// SetType sets the "type" field.
func (u *NotifyTargetUpsertOne) SetType(v notifytarget.Type) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateType() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateType()
	})
}

// SetStatus sets the "status" field.
func (u *NotifyTargetUpsertOne) SetStatus(v notifytarget.Status) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateStatus() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyTargetUpsertOne) SetUpdatedAt(v time.Time) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateUpdatedAt() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyTargetUpsertOne) SetCreatedAt(v time.Time) *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyTargetUpsertOne) UpdateCreatedAt() *NotifyTargetUpsertOne {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyTargetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyTargetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyTargetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *NotifyTargetUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *NotifyTargetUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// NotifyTargetCreateBulk is the builder for creating many NotifyTarget entities in bulk.
type NotifyTargetCreateBulk struct {
	config
	err      error
	builders []*NotifyTargetCreate
	conflict []sql.ConflictOption
}

// Save creates the NotifyTarget entities in the database.
func (ntcb *NotifyTargetCreateBulk) Save(ctx context.Context) ([]*NotifyTarget, error) {
	if ntcb.err != nil {
		return nil, ntcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ntcb.builders))
	nodes := make([]*NotifyTarget, len(ntcb.builders))
	mutators := make([]Mutator, len(ntcb.builders))
	for i := range ntcb.builders {
		func(i int, root context.Context) {
			builder := ntcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NotifyTargetMutation)
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
					_, err = mutators[i+1].Mutate(root, ntcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ntcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ntcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ntcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ntcb *NotifyTargetCreateBulk) SaveX(ctx context.Context) []*NotifyTarget {
	v, err := ntcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ntcb *NotifyTargetCreateBulk) Exec(ctx context.Context) error {
	_, err := ntcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ntcb *NotifyTargetCreateBulk) ExecX(ctx context.Context) {
	if err := ntcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.NotifyTarget.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.NotifyTargetUpsert) {
//			SetToken(v+v).
//		}).
//		Exec(ctx)
func (ntcb *NotifyTargetCreateBulk) OnConflict(opts ...sql.ConflictOption) *NotifyTargetUpsertBulk {
	ntcb.conflict = opts
	return &NotifyTargetUpsertBulk{
		create: ntcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ntcb *NotifyTargetCreateBulk) OnConflictColumns(columns ...string) *NotifyTargetUpsertBulk {
	ntcb.conflict = append(ntcb.conflict, sql.ConflictColumns(columns...))
	return &NotifyTargetUpsertBulk{
		create: ntcb,
	}
}

// NotifyTargetUpsertBulk is the builder for "upsert"-ing
// a bulk of NotifyTarget nodes.
type NotifyTargetUpsertBulk struct {
	create *NotifyTargetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(notifytarget.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *NotifyTargetUpsertBulk) UpdateNewValues() *NotifyTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(notifytarget.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.NotifyTarget.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *NotifyTargetUpsertBulk) Ignore() *NotifyTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *NotifyTargetUpsertBulk) DoNothing() *NotifyTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the NotifyTargetCreateBulk.OnConflict
// documentation for more info.
func (u *NotifyTargetUpsertBulk) Update(set func(*NotifyTargetUpsert)) *NotifyTargetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&NotifyTargetUpsert{UpdateSet: update})
	}))
	return u
}

// SetToken sets the "token" field.
func (u *NotifyTargetUpsertBulk) SetToken(v string) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateToken() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateToken()
	})
}

// SetName sets the "name" field.
func (u *NotifyTargetUpsertBulk) SetName(v string) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateName() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *NotifyTargetUpsertBulk) SetDescription(v string) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateDescription() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateDescription()
	})
}

// SetType sets the "type" field.
func (u *NotifyTargetUpsertBulk) SetType(v notifytarget.Type) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateType() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateType()
	})
}

// SetStatus sets the "status" field.
func (u *NotifyTargetUpsertBulk) SetStatus(v notifytarget.Status) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateStatus() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *NotifyTargetUpsertBulk) SetUpdatedAt(v time.Time) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateUpdatedAt() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *NotifyTargetUpsertBulk) SetCreatedAt(v time.Time) *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *NotifyTargetUpsertBulk) UpdateCreatedAt() *NotifyTargetUpsertBulk {
	return u.Update(func(s *NotifyTargetUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *NotifyTargetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the NotifyTargetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for NotifyTargetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *NotifyTargetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
