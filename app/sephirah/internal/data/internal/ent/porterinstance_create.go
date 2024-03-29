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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/porterinstance"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
	"github.com/tuihub/librarian/internal/model"
)

// PorterInstanceCreate is the builder for creating a PorterInstance entity.
type PorterInstanceCreate struct {
	config
	mutation *PorterInstanceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pic *PorterInstanceCreate) SetName(s string) *PorterInstanceCreate {
	pic.mutation.SetName(s)
	return pic
}

// SetVersion sets the "version" field.
func (pic *PorterInstanceCreate) SetVersion(s string) *PorterInstanceCreate {
	pic.mutation.SetVersion(s)
	return pic
}

// SetGlobalName sets the "global_name" field.
func (pic *PorterInstanceCreate) SetGlobalName(s string) *PorterInstanceCreate {
	pic.mutation.SetGlobalName(s)
	return pic
}

// SetAddress sets the "address" field.
func (pic *PorterInstanceCreate) SetAddress(s string) *PorterInstanceCreate {
	pic.mutation.SetAddress(s)
	return pic
}

// SetFeatureSummary sets the "feature_summary" field.
func (pic *PorterInstanceCreate) SetFeatureSummary(mfs *modeltiphereth.PorterFeatureSummary) *PorterInstanceCreate {
	pic.mutation.SetFeatureSummary(mfs)
	return pic
}

// SetStatus sets the "status" field.
func (pic *PorterInstanceCreate) SetStatus(po porterinstance.Status) *PorterInstanceCreate {
	pic.mutation.SetStatus(po)
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *PorterInstanceCreate) SetUpdatedAt(t time.Time) *PorterInstanceCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *PorterInstanceCreate) SetNillableUpdatedAt(t *time.Time) *PorterInstanceCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetCreatedAt sets the "created_at" field.
func (pic *PorterInstanceCreate) SetCreatedAt(t time.Time) *PorterInstanceCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *PorterInstanceCreate) SetNillableCreatedAt(t *time.Time) *PorterInstanceCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PorterInstanceCreate) SetID(mi model.InternalID) *PorterInstanceCreate {
	pic.mutation.SetID(mi)
	return pic
}

// Mutation returns the PorterInstanceMutation object of the builder.
func (pic *PorterInstanceCreate) Mutation() *PorterInstanceMutation {
	return pic.mutation
}

// Save creates the PorterInstance in the database.
func (pic *PorterInstanceCreate) Save(ctx context.Context) (*PorterInstance, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PorterInstanceCreate) SaveX(ctx context.Context) *PorterInstance {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PorterInstanceCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PorterInstanceCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PorterInstanceCreate) defaults() {
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		v := porterinstance.DefaultUpdatedAt()
		pic.mutation.SetUpdatedAt(v)
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		v := porterinstance.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PorterInstanceCreate) check() error {
	if _, ok := pic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PorterInstance.name"`)}
	}
	if _, ok := pic.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "PorterInstance.version"`)}
	}
	if _, ok := pic.mutation.GlobalName(); !ok {
		return &ValidationError{Name: "global_name", err: errors.New(`ent: missing required field "PorterInstance.global_name"`)}
	}
	if _, ok := pic.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "PorterInstance.address"`)}
	}
	if _, ok := pic.mutation.FeatureSummary(); !ok {
		return &ValidationError{Name: "feature_summary", err: errors.New(`ent: missing required field "PorterInstance.feature_summary"`)}
	}
	if _, ok := pic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "PorterInstance.status"`)}
	}
	if v, ok := pic.mutation.Status(); ok {
		if err := porterinstance.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "PorterInstance.status": %w`, err)}
		}
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PorterInstance.updated_at"`)}
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PorterInstance.created_at"`)}
	}
	return nil
}

func (pic *PorterInstanceCreate) sqlSave(ctx context.Context) (*PorterInstance, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *PorterInstanceCreate) createSpec() (*PorterInstance, *sqlgraph.CreateSpec) {
	var (
		_node = &PorterInstance{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(porterinstance.Table, sqlgraph.NewFieldSpec(porterinstance.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = pic.conflict
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.Name(); ok {
		_spec.SetField(porterinstance.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pic.mutation.Version(); ok {
		_spec.SetField(porterinstance.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := pic.mutation.GlobalName(); ok {
		_spec.SetField(porterinstance.FieldGlobalName, field.TypeString, value)
		_node.GlobalName = value
	}
	if value, ok := pic.mutation.Address(); ok {
		_spec.SetField(porterinstance.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := pic.mutation.FeatureSummary(); ok {
		_spec.SetField(porterinstance.FieldFeatureSummary, field.TypeJSON, value)
		_node.FeatureSummary = value
	}
	if value, ok := pic.mutation.Status(); ok {
		_spec.SetField(porterinstance.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.SetField(porterinstance.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.SetField(porterinstance.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PorterInstance.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PorterInstanceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pic *PorterInstanceCreate) OnConflict(opts ...sql.ConflictOption) *PorterInstanceUpsertOne {
	pic.conflict = opts
	return &PorterInstanceUpsertOne{
		create: pic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pic *PorterInstanceCreate) OnConflictColumns(columns ...string) *PorterInstanceUpsertOne {
	pic.conflict = append(pic.conflict, sql.ConflictColumns(columns...))
	return &PorterInstanceUpsertOne{
		create: pic,
	}
}

type (
	// PorterInstanceUpsertOne is the builder for "upsert"-ing
	//  one PorterInstance node.
	PorterInstanceUpsertOne struct {
		create *PorterInstanceCreate
	}

	// PorterInstanceUpsert is the "OnConflict" setter.
	PorterInstanceUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PorterInstanceUpsert) SetName(v string) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateName() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldName)
	return u
}

// SetVersion sets the "version" field.
func (u *PorterInstanceUpsert) SetVersion(v string) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateVersion() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldVersion)
	return u
}

// SetGlobalName sets the "global_name" field.
func (u *PorterInstanceUpsert) SetGlobalName(v string) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldGlobalName, v)
	return u
}

// UpdateGlobalName sets the "global_name" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateGlobalName() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldGlobalName)
	return u
}

// SetAddress sets the "address" field.
func (u *PorterInstanceUpsert) SetAddress(v string) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateAddress() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldAddress)
	return u
}

// SetFeatureSummary sets the "feature_summary" field.
func (u *PorterInstanceUpsert) SetFeatureSummary(v *modeltiphereth.PorterFeatureSummary) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldFeatureSummary, v)
	return u
}

// UpdateFeatureSummary sets the "feature_summary" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateFeatureSummary() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldFeatureSummary)
	return u
}

// SetStatus sets the "status" field.
func (u *PorterInstanceUpsert) SetStatus(v porterinstance.Status) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateStatus() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldStatus)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PorterInstanceUpsert) SetUpdatedAt(v time.Time) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateUpdatedAt() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PorterInstanceUpsert) SetCreatedAt(v time.Time) *PorterInstanceUpsert {
	u.Set(porterinstance.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PorterInstanceUpsert) UpdateCreatedAt() *PorterInstanceUpsert {
	u.SetExcluded(porterinstance.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(porterinstance.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PorterInstanceUpsertOne) UpdateNewValues() *PorterInstanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(porterinstance.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PorterInstanceUpsertOne) Ignore() *PorterInstanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PorterInstanceUpsertOne) DoNothing() *PorterInstanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PorterInstanceCreate.OnConflict
// documentation for more info.
func (u *PorterInstanceUpsertOne) Update(set func(*PorterInstanceUpsert)) *PorterInstanceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PorterInstanceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PorterInstanceUpsertOne) SetName(v string) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateName() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *PorterInstanceUpsertOne) SetVersion(v string) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateVersion() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateVersion()
	})
}

// SetGlobalName sets the "global_name" field.
func (u *PorterInstanceUpsertOne) SetGlobalName(v string) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetGlobalName(v)
	})
}

// UpdateGlobalName sets the "global_name" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateGlobalName() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateGlobalName()
	})
}

// SetAddress sets the "address" field.
func (u *PorterInstanceUpsertOne) SetAddress(v string) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateAddress() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateAddress()
	})
}

// SetFeatureSummary sets the "feature_summary" field.
func (u *PorterInstanceUpsertOne) SetFeatureSummary(v *modeltiphereth.PorterFeatureSummary) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetFeatureSummary(v)
	})
}

// UpdateFeatureSummary sets the "feature_summary" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateFeatureSummary() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateFeatureSummary()
	})
}

// SetStatus sets the "status" field.
func (u *PorterInstanceUpsertOne) SetStatus(v porterinstance.Status) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateStatus() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PorterInstanceUpsertOne) SetUpdatedAt(v time.Time) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateUpdatedAt() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PorterInstanceUpsertOne) SetCreatedAt(v time.Time) *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PorterInstanceUpsertOne) UpdateCreatedAt() *PorterInstanceUpsertOne {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *PorterInstanceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PorterInstanceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PorterInstanceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PorterInstanceUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PorterInstanceUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PorterInstanceCreateBulk is the builder for creating many PorterInstance entities in bulk.
type PorterInstanceCreateBulk struct {
	config
	err      error
	builders []*PorterInstanceCreate
	conflict []sql.ConflictOption
}

// Save creates the PorterInstance entities in the database.
func (picb *PorterInstanceCreateBulk) Save(ctx context.Context) ([]*PorterInstance, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PorterInstance, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PorterInstanceMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = picb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PorterInstanceCreateBulk) SaveX(ctx context.Context) []*PorterInstance {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PorterInstanceCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PorterInstanceCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PorterInstance.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PorterInstanceUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (picb *PorterInstanceCreateBulk) OnConflict(opts ...sql.ConflictOption) *PorterInstanceUpsertBulk {
	picb.conflict = opts
	return &PorterInstanceUpsertBulk{
		create: picb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (picb *PorterInstanceCreateBulk) OnConflictColumns(columns ...string) *PorterInstanceUpsertBulk {
	picb.conflict = append(picb.conflict, sql.ConflictColumns(columns...))
	return &PorterInstanceUpsertBulk{
		create: picb,
	}
}

// PorterInstanceUpsertBulk is the builder for "upsert"-ing
// a bulk of PorterInstance nodes.
type PorterInstanceUpsertBulk struct {
	create *PorterInstanceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(porterinstance.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PorterInstanceUpsertBulk) UpdateNewValues() *PorterInstanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(porterinstance.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PorterInstance.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PorterInstanceUpsertBulk) Ignore() *PorterInstanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PorterInstanceUpsertBulk) DoNothing() *PorterInstanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PorterInstanceCreateBulk.OnConflict
// documentation for more info.
func (u *PorterInstanceUpsertBulk) Update(set func(*PorterInstanceUpsert)) *PorterInstanceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PorterInstanceUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PorterInstanceUpsertBulk) SetName(v string) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateName() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateName()
	})
}

// SetVersion sets the "version" field.
func (u *PorterInstanceUpsertBulk) SetVersion(v string) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateVersion() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateVersion()
	})
}

// SetGlobalName sets the "global_name" field.
func (u *PorterInstanceUpsertBulk) SetGlobalName(v string) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetGlobalName(v)
	})
}

// UpdateGlobalName sets the "global_name" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateGlobalName() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateGlobalName()
	})
}

// SetAddress sets the "address" field.
func (u *PorterInstanceUpsertBulk) SetAddress(v string) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateAddress() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateAddress()
	})
}

// SetFeatureSummary sets the "feature_summary" field.
func (u *PorterInstanceUpsertBulk) SetFeatureSummary(v *modeltiphereth.PorterFeatureSummary) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetFeatureSummary(v)
	})
}

// UpdateFeatureSummary sets the "feature_summary" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateFeatureSummary() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateFeatureSummary()
	})
}

// SetStatus sets the "status" field.
func (u *PorterInstanceUpsertBulk) SetStatus(v porterinstance.Status) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateStatus() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PorterInstanceUpsertBulk) SetUpdatedAt(v time.Time) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateUpdatedAt() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *PorterInstanceUpsertBulk) SetCreatedAt(v time.Time) *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PorterInstanceUpsertBulk) UpdateCreatedAt() *PorterInstanceUpsertBulk {
	return u.Update(func(s *PorterInstanceUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *PorterInstanceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PorterInstanceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PorterInstanceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PorterInstanceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
