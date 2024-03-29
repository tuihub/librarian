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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appbinary"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/model"
)

// AppBinaryCreate is the builder for creating a AppBinary entity.
type AppBinaryCreate struct {
	config
	mutation *AppBinaryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (abc *AppBinaryCreate) SetName(s string) *AppBinaryCreate {
	abc.mutation.SetName(s)
	return abc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (abc *AppBinaryCreate) SetNillableName(s *string) *AppBinaryCreate {
	if s != nil {
		abc.SetName(*s)
	}
	return abc
}

// SetSizeBytes sets the "size_bytes" field.
func (abc *AppBinaryCreate) SetSizeBytes(i int64) *AppBinaryCreate {
	abc.mutation.SetSizeBytes(i)
	return abc
}

// SetNillableSizeBytes sets the "size_bytes" field if the given value is not nil.
func (abc *AppBinaryCreate) SetNillableSizeBytes(i *int64) *AppBinaryCreate {
	if i != nil {
		abc.SetSizeBytes(*i)
	}
	return abc
}

// SetPublicURL sets the "public_url" field.
func (abc *AppBinaryCreate) SetPublicURL(s string) *AppBinaryCreate {
	abc.mutation.SetPublicURL(s)
	return abc
}

// SetNillablePublicURL sets the "public_url" field if the given value is not nil.
func (abc *AppBinaryCreate) SetNillablePublicURL(s *string) *AppBinaryCreate {
	if s != nil {
		abc.SetPublicURL(*s)
	}
	return abc
}

// SetSha256 sets the "sha256" field.
func (abc *AppBinaryCreate) SetSha256(b []byte) *AppBinaryCreate {
	abc.mutation.SetSha256(b)
	return abc
}

// SetUpdatedAt sets the "updated_at" field.
func (abc *AppBinaryCreate) SetUpdatedAt(t time.Time) *AppBinaryCreate {
	abc.mutation.SetUpdatedAt(t)
	return abc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (abc *AppBinaryCreate) SetNillableUpdatedAt(t *time.Time) *AppBinaryCreate {
	if t != nil {
		abc.SetUpdatedAt(*t)
	}
	return abc
}

// SetCreatedAt sets the "created_at" field.
func (abc *AppBinaryCreate) SetCreatedAt(t time.Time) *AppBinaryCreate {
	abc.mutation.SetCreatedAt(t)
	return abc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (abc *AppBinaryCreate) SetNillableCreatedAt(t *time.Time) *AppBinaryCreate {
	if t != nil {
		abc.SetCreatedAt(*t)
	}
	return abc
}

// SetID sets the "id" field.
func (abc *AppBinaryCreate) SetID(mi model.InternalID) *AppBinaryCreate {
	abc.mutation.SetID(mi)
	return abc
}

// SetAppInfoID sets the "app_info" edge to the AppInfo entity by ID.
func (abc *AppBinaryCreate) SetAppInfoID(id model.InternalID) *AppBinaryCreate {
	abc.mutation.SetAppInfoID(id)
	return abc
}

// SetNillableAppInfoID sets the "app_info" edge to the AppInfo entity by ID if the given value is not nil.
func (abc *AppBinaryCreate) SetNillableAppInfoID(id *model.InternalID) *AppBinaryCreate {
	if id != nil {
		abc = abc.SetAppInfoID(*id)
	}
	return abc
}

// SetAppInfo sets the "app_info" edge to the AppInfo entity.
func (abc *AppBinaryCreate) SetAppInfo(a *AppInfo) *AppBinaryCreate {
	return abc.SetAppInfoID(a.ID)
}

// Mutation returns the AppBinaryMutation object of the builder.
func (abc *AppBinaryCreate) Mutation() *AppBinaryMutation {
	return abc.mutation
}

// Save creates the AppBinary in the database.
func (abc *AppBinaryCreate) Save(ctx context.Context) (*AppBinary, error) {
	abc.defaults()
	return withHooks(ctx, abc.sqlSave, abc.mutation, abc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (abc *AppBinaryCreate) SaveX(ctx context.Context) *AppBinary {
	v, err := abc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abc *AppBinaryCreate) Exec(ctx context.Context) error {
	_, err := abc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abc *AppBinaryCreate) ExecX(ctx context.Context) {
	if err := abc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (abc *AppBinaryCreate) defaults() {
	if _, ok := abc.mutation.UpdatedAt(); !ok {
		v := appbinary.DefaultUpdatedAt()
		abc.mutation.SetUpdatedAt(v)
	}
	if _, ok := abc.mutation.CreatedAt(); !ok {
		v := appbinary.DefaultCreatedAt()
		abc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (abc *AppBinaryCreate) check() error {
	if _, ok := abc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppBinary.updated_at"`)}
	}
	if _, ok := abc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppBinary.created_at"`)}
	}
	return nil
}

func (abc *AppBinaryCreate) sqlSave(ctx context.Context) (*AppBinary, error) {
	if err := abc.check(); err != nil {
		return nil, err
	}
	_node, _spec := abc.createSpec()
	if err := sqlgraph.CreateNode(ctx, abc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	abc.mutation.id = &_node.ID
	abc.mutation.done = true
	return _node, nil
}

func (abc *AppBinaryCreate) createSpec() (*AppBinary, *sqlgraph.CreateSpec) {
	var (
		_node = &AppBinary{config: abc.config}
		_spec = sqlgraph.NewCreateSpec(appbinary.Table, sqlgraph.NewFieldSpec(appbinary.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = abc.conflict
	if id, ok := abc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := abc.mutation.Name(); ok {
		_spec.SetField(appbinary.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := abc.mutation.SizeBytes(); ok {
		_spec.SetField(appbinary.FieldSizeBytes, field.TypeInt64, value)
		_node.SizeBytes = value
	}
	if value, ok := abc.mutation.PublicURL(); ok {
		_spec.SetField(appbinary.FieldPublicURL, field.TypeString, value)
		_node.PublicURL = value
	}
	if value, ok := abc.mutation.Sha256(); ok {
		_spec.SetField(appbinary.FieldSha256, field.TypeBytes, value)
		_node.Sha256 = value
	}
	if value, ok := abc.mutation.UpdatedAt(); ok {
		_spec.SetField(appbinary.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := abc.mutation.CreatedAt(); ok {
		_spec.SetField(appbinary.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := abc.mutation.AppInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appbinary.AppInfoTable,
			Columns: []string{appbinary.AppInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.app_info_app_binary = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppBinary.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppBinaryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (abc *AppBinaryCreate) OnConflict(opts ...sql.ConflictOption) *AppBinaryUpsertOne {
	abc.conflict = opts
	return &AppBinaryUpsertOne{
		create: abc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (abc *AppBinaryCreate) OnConflictColumns(columns ...string) *AppBinaryUpsertOne {
	abc.conflict = append(abc.conflict, sql.ConflictColumns(columns...))
	return &AppBinaryUpsertOne{
		create: abc,
	}
}

type (
	// AppBinaryUpsertOne is the builder for "upsert"-ing
	//  one AppBinary node.
	AppBinaryUpsertOne struct {
		create *AppBinaryCreate
	}

	// AppBinaryUpsert is the "OnConflict" setter.
	AppBinaryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AppBinaryUpsert) SetName(v string) *AppBinaryUpsert {
	u.Set(appbinary.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdateName() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *AppBinaryUpsert) ClearName() *AppBinaryUpsert {
	u.SetNull(appbinary.FieldName)
	return u
}

// SetSizeBytes sets the "size_bytes" field.
func (u *AppBinaryUpsert) SetSizeBytes(v int64) *AppBinaryUpsert {
	u.Set(appbinary.FieldSizeBytes, v)
	return u
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdateSizeBytes() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldSizeBytes)
	return u
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *AppBinaryUpsert) AddSizeBytes(v int64) *AppBinaryUpsert {
	u.Add(appbinary.FieldSizeBytes, v)
	return u
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *AppBinaryUpsert) ClearSizeBytes() *AppBinaryUpsert {
	u.SetNull(appbinary.FieldSizeBytes)
	return u
}

// SetPublicURL sets the "public_url" field.
func (u *AppBinaryUpsert) SetPublicURL(v string) *AppBinaryUpsert {
	u.Set(appbinary.FieldPublicURL, v)
	return u
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdatePublicURL() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldPublicURL)
	return u
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *AppBinaryUpsert) ClearPublicURL() *AppBinaryUpsert {
	u.SetNull(appbinary.FieldPublicURL)
	return u
}

// SetSha256 sets the "sha256" field.
func (u *AppBinaryUpsert) SetSha256(v []byte) *AppBinaryUpsert {
	u.Set(appbinary.FieldSha256, v)
	return u
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdateSha256() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldSha256)
	return u
}

// ClearSha256 clears the value of the "sha256" field.
func (u *AppBinaryUpsert) ClearSha256() *AppBinaryUpsert {
	u.SetNull(appbinary.FieldSha256)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppBinaryUpsert) SetUpdatedAt(v time.Time) *AppBinaryUpsert {
	u.Set(appbinary.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdateUpdatedAt() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppBinaryUpsert) SetCreatedAt(v time.Time) *AppBinaryUpsert {
	u.Set(appbinary.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppBinaryUpsert) UpdateCreatedAt() *AppBinaryUpsert {
	u.SetExcluded(appbinary.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appbinary.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppBinaryUpsertOne) UpdateNewValues() *AppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appbinary.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppBinaryUpsertOne) Ignore() *AppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppBinaryUpsertOne) DoNothing() *AppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppBinaryCreate.OnConflict
// documentation for more info.
func (u *AppBinaryUpsertOne) Update(set func(*AppBinaryUpsert)) *AppBinaryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppBinaryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AppBinaryUpsertOne) SetName(v string) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdateName() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppBinaryUpsertOne) ClearName() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearName()
	})
}

// SetSizeBytes sets the "size_bytes" field.
func (u *AppBinaryUpsertOne) SetSizeBytes(v int64) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetSizeBytes(v)
	})
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *AppBinaryUpsertOne) AddSizeBytes(v int64) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.AddSizeBytes(v)
	})
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdateSizeBytes() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateSizeBytes()
	})
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *AppBinaryUpsertOne) ClearSizeBytes() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearSizeBytes()
	})
}

// SetPublicURL sets the "public_url" field.
func (u *AppBinaryUpsertOne) SetPublicURL(v string) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetPublicURL(v)
	})
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdatePublicURL() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdatePublicURL()
	})
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *AppBinaryUpsertOne) ClearPublicURL() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearPublicURL()
	})
}

// SetSha256 sets the "sha256" field.
func (u *AppBinaryUpsertOne) SetSha256(v []byte) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdateSha256() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateSha256()
	})
}

// ClearSha256 clears the value of the "sha256" field.
func (u *AppBinaryUpsertOne) ClearSha256() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppBinaryUpsertOne) SetUpdatedAt(v time.Time) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdateUpdatedAt() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppBinaryUpsertOne) SetCreatedAt(v time.Time) *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppBinaryUpsertOne) UpdateCreatedAt() *AppBinaryUpsertOne {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppBinaryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppBinaryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppBinaryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppBinaryUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppBinaryUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppBinaryCreateBulk is the builder for creating many AppBinary entities in bulk.
type AppBinaryCreateBulk struct {
	config
	err      error
	builders []*AppBinaryCreate
	conflict []sql.ConflictOption
}

// Save creates the AppBinary entities in the database.
func (abcb *AppBinaryCreateBulk) Save(ctx context.Context) ([]*AppBinary, error) {
	if abcb.err != nil {
		return nil, abcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(abcb.builders))
	nodes := make([]*AppBinary, len(abcb.builders))
	mutators := make([]Mutator, len(abcb.builders))
	for i := range abcb.builders {
		func(i int, root context.Context) {
			builder := abcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppBinaryMutation)
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
					_, err = mutators[i+1].Mutate(root, abcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = abcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, abcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, abcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (abcb *AppBinaryCreateBulk) SaveX(ctx context.Context) []*AppBinary {
	v, err := abcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abcb *AppBinaryCreateBulk) Exec(ctx context.Context) error {
	_, err := abcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abcb *AppBinaryCreateBulk) ExecX(ctx context.Context) {
	if err := abcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppBinary.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppBinaryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (abcb *AppBinaryCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppBinaryUpsertBulk {
	abcb.conflict = opts
	return &AppBinaryUpsertBulk{
		create: abcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (abcb *AppBinaryCreateBulk) OnConflictColumns(columns ...string) *AppBinaryUpsertBulk {
	abcb.conflict = append(abcb.conflict, sql.ConflictColumns(columns...))
	return &AppBinaryUpsertBulk{
		create: abcb,
	}
}

// AppBinaryUpsertBulk is the builder for "upsert"-ing
// a bulk of AppBinary nodes.
type AppBinaryUpsertBulk struct {
	create *AppBinaryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appbinary.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppBinaryUpsertBulk) UpdateNewValues() *AppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appbinary.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppBinary.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppBinaryUpsertBulk) Ignore() *AppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppBinaryUpsertBulk) DoNothing() *AppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppBinaryCreateBulk.OnConflict
// documentation for more info.
func (u *AppBinaryUpsertBulk) Update(set func(*AppBinaryUpsert)) *AppBinaryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppBinaryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AppBinaryUpsertBulk) SetName(v string) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdateName() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppBinaryUpsertBulk) ClearName() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearName()
	})
}

// SetSizeBytes sets the "size_bytes" field.
func (u *AppBinaryUpsertBulk) SetSizeBytes(v int64) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetSizeBytes(v)
	})
}

// AddSizeBytes adds v to the "size_bytes" field.
func (u *AppBinaryUpsertBulk) AddSizeBytes(v int64) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.AddSizeBytes(v)
	})
}

// UpdateSizeBytes sets the "size_bytes" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdateSizeBytes() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateSizeBytes()
	})
}

// ClearSizeBytes clears the value of the "size_bytes" field.
func (u *AppBinaryUpsertBulk) ClearSizeBytes() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearSizeBytes()
	})
}

// SetPublicURL sets the "public_url" field.
func (u *AppBinaryUpsertBulk) SetPublicURL(v string) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetPublicURL(v)
	})
}

// UpdatePublicURL sets the "public_url" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdatePublicURL() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdatePublicURL()
	})
}

// ClearPublicURL clears the value of the "public_url" field.
func (u *AppBinaryUpsertBulk) ClearPublicURL() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearPublicURL()
	})
}

// SetSha256 sets the "sha256" field.
func (u *AppBinaryUpsertBulk) SetSha256(v []byte) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetSha256(v)
	})
}

// UpdateSha256 sets the "sha256" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdateSha256() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateSha256()
	})
}

// ClearSha256 clears the value of the "sha256" field.
func (u *AppBinaryUpsertBulk) ClearSha256() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.ClearSha256()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppBinaryUpsertBulk) SetUpdatedAt(v time.Time) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdateUpdatedAt() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppBinaryUpsertBulk) SetCreatedAt(v time.Time) *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppBinaryUpsertBulk) UpdateCreatedAt() *AppBinaryUpsertBulk {
	return u.Update(func(s *AppBinaryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppBinaryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppBinaryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppBinaryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppBinaryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
