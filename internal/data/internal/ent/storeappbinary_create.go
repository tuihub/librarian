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
	"github.com/tuihub/librarian/internal/data/internal/ent/storeappbinary"
	"github.com/tuihub/librarian/internal/model"
)

// StoreAppBinaryCreate is the builder for creating a StoreAppBinary entity.
type StoreAppBinaryCreate struct {
	config
	mutation *StoreAppBinaryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sabc *StoreAppBinaryCreate) SetName(s string) *StoreAppBinaryCreate {
	sabc.mutation.SetName(s)
	return sabc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sabc *StoreAppBinaryCreate) SetNillableName(s *string) *StoreAppBinaryCreate {
	if s != nil {
		sabc.SetName(*s)
	}
	return sabc
}

// SetSizeBytes sets the "size_bytes" field.
func (sabc *StoreAppBinaryCreate) SetSizeBytes(i int64) *StoreAppBinaryCreate {
	sabc.mutation.SetSizeBytes(i)
	return sabc
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (sabc *StoreAppBinaryCreate) SetNillableSizeBytes(i *int64) *StoreAppBinaryCreate {
	if i != nil {
		sabc.SetSizeBytes(*i)
	}
	return sabc
}

// SetPublicURL sets the "public_url" field.
func (sabc *StoreAppBinaryCreate) SetPublicURL(s string) *StoreAppBinaryCreate {
	sabc.mutation.SetPublicURL(s)
	return sabc
}

// SetNillablePublicURL sets the "public_url" field if the given value is not nil.
func (sabc *StoreAppBinaryCreate) SetNillablePublicURL(s *string) *StoreAppBinaryCreate {
	if s != nil {
		sabc.SetPublicURL(*s)
	}
	return sabc
}

// SetSha256 sets the "sha256" field.
func (sabc *StoreAppBinaryCreate) SetSha256(b []byte) *StoreAppBinaryCreate {
	sabc.mutation.SetSha256(b)
	return sabc
}

// SetUpdatedAt sets the "updated_at" field.
func (sabc *StoreAppBinaryCreate) SetUpdatedAt(t time.Time) *StoreAppBinaryCreate {
	sabc.mutation.SetUpdatedAt(t)
	return sabc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sabc *StoreAppBinaryCreate) SetNillableUpdatedAt(t *time.Time) *StoreAppBinaryCreate {
	if t != nil {
		sabc.SetUpdatedAt(*t)
	}
	return sabc
}

// SetCreatedAt sets the "created_at" field.
func (sabc *StoreAppBinaryCreate) SetCreatedAt(t time.Time) *StoreAppBinaryCreate {
	sabc.mutation.SetCreatedAt(t)
	return sabc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sabc *StoreAppBinaryCreate) SetNillableCreatedAt(t *time.Time) *StoreAppBinaryCreate {
	if t != nil {
		sabc.SetCreatedAt(*t)
	}
	return sabc
}

// SetID sets the "id" field.
func (sabc *StoreAppBinaryCreate) SetID(mi model.InternalID) *StoreAppBinaryCreate {
	sabc.mutation.SetID(mi)
	return sabc
}

// Mutation returns the StoreAppBinaryMutation object of the builder.
func (sabc *StoreAppBinaryCreate) Mutation() *StoreAppBinaryMutation {
	return sabc.mutation
}

// Save creates the StoreAppBinary in the database.
func (sabc *StoreAppBinaryCreate) Save(ctx context.Context) (*StoreAppBinary, error) {
	sabc.defaults()
	return withHooks(ctx, sabc.sqlSave, sabc.mutation, sabc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sabc *StoreAppBinaryCreate) SaveX(ctx context.Context) *StoreAppBinary {
	v, err := sabc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sabc *StoreAppBinaryCreate) Exec(ctx context.Context) error {
	_, err := sabc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabc *StoreAppBinaryCreate) ExecX(ctx context.Context) {
	if err := sabc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sabc *StoreAppBinaryCreate) defaults() {
	if _, ok := sabc.mutation.UpdatedAt(); !ok {
		v := storeappbinary.DefaultUpdatedAt()
		sabc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sabc.mutation.CreatedAt(); !ok {
		v := storeappbinary.DefaultCreatedAt()
		sabc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sabc *StoreAppBinaryCreate) check() error {
	if _, ok := sabc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StoreAppBinary.updated_at"`)}
	}
	if _, ok := sabc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StoreAppBinary.created_at"`)}
	}
	return nil
}

func (sabc *StoreAppBinaryCreate) sqlSave(ctx context.Context) (*StoreAppBinary, error) {
	if err := sabc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sabc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sabc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	sabc.mutation.id = &_node.ID
	sabc.mutation.done = true
	return _node, nil
}

func (sabc *StoreAppBinaryCreate) createSpec() (*StoreAppBinary, *sqlgraph.CreateSpec) {
	var (
		_node = &StoreAppBinary{config: sabc.config}
		_spec = sqlgraph.NewCreateSpec(storeappbinary.Table, sqlgraph.NewFieldSpec(storeappbinary.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = sabc.conflict
	if id, ok := sabc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sabc.mutation.Name(); ok {
		_spec.SetField(storeappbinary.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sabc.mutation.SizeBytes(); ok {
		_spec.SetField(storeappbinary.FieldSizeBytes, field.TypeInt64, value)
		_node.SizeBytes = value
	}
	if value, ok := sabc.mutation.PublicURL(); ok {
		_spec.SetField(storeappbinary.FieldPublicURL, field.TypeString, value)
		_node.PublicURL = value
	}
	if value, ok := sabc.mutation.Sha256(); ok {
		_spec.SetField(storeappbinary.FieldSha256, field.TypeBytes, value)
		_node.Sha256 = value
	}
	if value, ok := sabc.mutation.UpdatedAt(); ok {
		_spec.SetField(storeappbinary.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sabc.mutation.CreatedAt(); ok {
		_spec.SetField(storeappbinary.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StoreAppBinary.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StoreAppBinaryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sabc *StoreAppBinaryCreate) OnConflict(opts ...sql.ConflictOption) *StoreAppBinaryUpsertOne {
	sabc.conflict = opts
	return &StoreAppBinaryUpsertOne{
		create: sabc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sabc *StoreAppBinaryCreate) OnConflictColumns(columns ...string) *StoreAppBinaryUpsertOne {
	sabc.conflict = append(sabc.conflict, sql.ConflictColumns(columns...))
	return &StoreAppBinaryUpsertOne{
		create: sabc,
	}
}

type (
	// StoreAppBinaryUpsertOne is the builder for "upsert"-ing
	//  one StoreAppBinary node.
	StoreAppBinaryUpsertOne struct {
		create *StoreAppBinaryCreate
	}

	// StoreAppBinaryUpsert is the "OnConflict" setter.
	StoreAppBinaryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StoreAppBinaryUpsert) SetName(v string) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdateName() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *StoreAppBinaryUpsert) ClearName() *StoreAppBinaryUpsert {
	u.SetNull(storeappbinary.FieldName)
	return u
}

// SetSizeBytes sets the "size_bytes" field.
func (u *StoreAppBinaryUpsert) SetSizeBytes(v int64) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldSizeBytes, v)
	return u
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdateSizeBytes() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldSizeBytes)
	return u
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *StoreAppBinaryUpsert) AddSizeBytes(v int64) *StoreAppBinaryUpsert {
	u.Add(storeappbinary.FieldSizeBytes, v)
	return u
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *StoreAppBinaryUpsert) ClearSizeBytes() *StoreAppBinaryUpsert {
	u.SetNull(storeappbinary.FieldSizeBytes)
	return u
}

// SetPublicURL sets the "public_url" field.
func (u *StoreAppBinaryUpsert) SetPublicURL(v string) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldPublicURL, v)
	return u
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdatePublicURL() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldPublicURL)
	return u
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *StoreAppBinaryUpsert) ClearPublicURL() *StoreAppBinaryUpsert {
	u.SetNull(storeappbinary.FieldPublicURL)
	return u
}

// SetSha256 sets the "sha256" field.
func (u *StoreAppBinaryUpsert) SetSha256(v []byte) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldSha256, v)
	return u
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdateSha256() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldSha256)
	return u
}

// ClearSha256 clears the value of the "sha256" field.
func (u *StoreAppBinaryUpsert) ClearSha256() *StoreAppBinaryUpsert {
	u.SetNull(storeappbinary.FieldSha256)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppBinaryUpsert) SetUpdatedAt(v time.Time) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdateUpdatedAt() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppBinaryUpsert) SetCreatedAt(v time.Time) *StoreAppBinaryUpsert {
	u.Set(storeappbinary.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsert) UpdateCreatedAt() *StoreAppBinaryUpsert {
	u.SetExcluded(storeappbinary.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(storeappbinary.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StoreAppBinaryUpsertOne) UpdateNewValues() *StoreAppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(storeappbinary.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StoreAppBinaryUpsertOne) Ignore() *StoreAppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StoreAppBinaryUpsertOne) DoNothing() *StoreAppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StoreAppBinaryCreate.OnConflict
// documentation for more info.
func (u *StoreAppBinaryUpsertOne) Update(set func(*StoreAppBinaryUpsert)) *StoreAppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StoreAppBinaryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StoreAppBinaryUpsertOne) SetName(v string) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdateName() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *StoreAppBinaryUpsertOne) ClearName() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearName()
	})
}

// SetSizeBytes sets the "size_bytes" field.
func (u *StoreAppBinaryUpsertOne) SetSizeBytes(v int64) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetSizeBytes(v)
	})
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *StoreAppBinaryUpsertOne) AddSizeBytes(v int64) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.AddSizeBytes(v)
	})
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdateSizeBytes() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateSizeBytes()
	})
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *StoreAppBinaryUpsertOne) ClearSizeBytes() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearSizeBytes()
	})
}

// SetPublicURL sets the "public_url" field.
func (u *StoreAppBinaryUpsertOne) SetPublicURL(v string) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetPublicURL(v)
	})
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdatePublicURL() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdatePublicURL()
	})
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *StoreAppBinaryUpsertOne) ClearPublicURL() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearPublicURL()
	})
}

// SetSha256 sets the "sha256" field.
func (u *StoreAppBinaryUpsertOne) SetSha256(v []byte) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdateSha256() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateSha256()
	})
}

// ClearSha256 clears the value of the "sha256" field.
func (u *StoreAppBinaryUpsertOne) ClearSha256() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppBinaryUpsertOne) SetUpdatedAt(v time.Time) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdateUpdatedAt() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppBinaryUpsertOne) SetCreatedAt(v time.Time) *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertOne) UpdateCreatedAt() *StoreAppBinaryUpsertOne {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StoreAppBinaryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StoreAppBinaryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StoreAppBinaryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StoreAppBinaryUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StoreAppBinaryUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StoreAppBinaryCreateBulk is the builder for creating many StoreAppBinary entities in bulk.
type StoreAppBinaryCreateBulk struct {
	config
	err      error
	builders []*StoreAppBinaryCreate
	conflict []sql.ConflictOption
}

// Save creates the StoreAppBinary entities in the database.
func (sabcb *StoreAppBinaryCreateBulk) Save(ctx context.Context) ([]*StoreAppBinary, error) {
	if sabcb.err != nil {
		return nil, sabcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sabcb.builders))
	nodes := make([]*StoreAppBinary, len(sabcb.builders))
	mutators := make([]Mutator, len(sabcb.builders))
	for i := range sabcb.builders {
		func(i int, root context.Context) {
			builder := sabcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoreAppBinaryMutation)
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
					_, err = mutators[i+1].Mutate(root, sabcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = sabcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sabcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sabcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sabcb *StoreAppBinaryCreateBulk) SaveX(ctx context.Context) []*StoreAppBinary {
	v, err := sabcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sabcb *StoreAppBinaryCreateBulk) Exec(ctx context.Context) error {
	_, err := sabcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sabcb *StoreAppBinaryCreateBulk) ExecX(ctx context.Context) {
	if err := sabcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StoreAppBinary.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StoreAppBinaryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sabcb *StoreAppBinaryCreateBulk) OnConflict(opts ...sql.ConflictOption) *StoreAppBinaryUpsertBulk {
	sabcb.conflict = opts
	return &StoreAppBinaryUpsertBulk{
		create: sabcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sabcb *StoreAppBinaryCreateBulk) OnConflictColumns(columns ...string) *StoreAppBinaryUpsertBulk {
	sabcb.conflict = append(sabcb.conflict, sql.ConflictColumns(columns...))
	return &StoreAppBinaryUpsertBulk{
		create: sabcb,
	}
}

// StoreAppBinaryUpsertBulk is the builder for "upsert"-ing
// a bulk of StoreAppBinary nodes.
type StoreAppBinaryUpsertBulk struct {
	create *StoreAppBinaryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(storeappbinary.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StoreAppBinaryUpsertBulk) UpdateNewValues() *StoreAppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(storeappbinary.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StoreAppBinary.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StoreAppBinaryUpsertBulk) Ignore() *StoreAppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StoreAppBinaryUpsertBulk) DoNothing() *StoreAppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StoreAppBinaryCreateBulk.OnConflict
// documentation for more info.
func (u *StoreAppBinaryUpsertBulk) Update(set func(*StoreAppBinaryUpsert)) *StoreAppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StoreAppBinaryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StoreAppBinaryUpsertBulk) SetName(v string) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdateName() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *StoreAppBinaryUpsertBulk) ClearName() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearName()
	})
}

// SetSizeBytes sets the "size_bytes" field.
func (u *StoreAppBinaryUpsertBulk) SetSizeBytes(v int64) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetSizeBytes(v)
	})
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *StoreAppBinaryUpsertBulk) AddSizeBytes(v int64) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.AddSizeBytes(v)
	})
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdateSizeBytes() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateSizeBytes()
	})
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *StoreAppBinaryUpsertBulk) ClearSizeBytes() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearSizeBytes()
	})
}

// SetPublicURL sets the "public_url" field.
func (u *StoreAppBinaryUpsertBulk) SetPublicURL(v string) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetPublicURL(v)
	})
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdatePublicURL() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdatePublicURL()
	})
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *StoreAppBinaryUpsertBulk) ClearPublicURL() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearPublicURL()
	})
}

// SetSha256 sets the "sha256" field.
func (u *StoreAppBinaryUpsertBulk) SetSha256(v []byte) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdateSha256() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateSha256()
	})
}

// ClearSha256 clears the value of the "sha256" field.
func (u *StoreAppBinaryUpsertBulk) ClearSha256() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.ClearSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StoreAppBinaryUpsertBulk) SetUpdatedAt(v time.Time) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdateUpdatedAt() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *StoreAppBinaryUpsertBulk) SetCreatedAt(v time.Time) *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StoreAppBinaryUpsertBulk) UpdateCreatedAt() *StoreAppBinaryUpsertBulk {
	return u.Update(func(s *StoreAppBinaryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *StoreAppBinaryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StoreAppBinaryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StoreAppBinaryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StoreAppBinaryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
