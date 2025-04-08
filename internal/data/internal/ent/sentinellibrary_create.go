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
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinellibrary"
)

// SentinelLibraryCreate is the builder for creating a SentinelLibrary entity.
type SentinelLibraryCreate struct {
	config
	mutation *SentinelLibraryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetSentinelInfoID sets the "sentinel_info_id" field.
func (slc *SentinelLibraryCreate) SetSentinelInfoID(i int) *SentinelLibraryCreate {
	slc.mutation.SetSentinelInfoID(i)
	return slc
}

// SetReportedID sets the "reported_id" field.
func (slc *SentinelLibraryCreate) SetReportedID(i int64) *SentinelLibraryCreate {
	slc.mutation.SetReportedID(i)
	return slc
}

// SetDownloadBasePath sets the "download_base_path" field.
func (slc *SentinelLibraryCreate) SetDownloadBasePath(s string) *SentinelLibraryCreate {
	slc.mutation.SetDownloadBasePath(s)
	return slc
}

// SetUpdatedAt sets the "updated_at" field.
func (slc *SentinelLibraryCreate) SetUpdatedAt(t time.Time) *SentinelLibraryCreate {
	slc.mutation.SetUpdatedAt(t)
	return slc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (slc *SentinelLibraryCreate) SetNillableUpdatedAt(t *time.Time) *SentinelLibraryCreate {
	if t != nil {
		slc.SetUpdatedAt(*t)
	}
	return slc
}

// SetCreatedAt sets the "created_at" field.
func (slc *SentinelLibraryCreate) SetCreatedAt(t time.Time) *SentinelLibraryCreate {
	slc.mutation.SetCreatedAt(t)
	return slc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (slc *SentinelLibraryCreate) SetNillableCreatedAt(t *time.Time) *SentinelLibraryCreate {
	if t != nil {
		slc.SetCreatedAt(*t)
	}
	return slc
}

// SetSentinelInfo sets the "sentinel_info" edge to the SentinelInfo entity.
func (slc *SentinelLibraryCreate) SetSentinelInfo(s *SentinelInfo) *SentinelLibraryCreate {
	return slc.SetSentinelInfoID(s.ID)
}

// AddSentinelAppBinaryIDs adds the "sentinel_app_binary" edge to the SentinelAppBinary entity by IDs.
func (slc *SentinelLibraryCreate) AddSentinelAppBinaryIDs(ids ...int) *SentinelLibraryCreate {
	slc.mutation.AddSentinelAppBinaryIDs(ids...)
	return slc
}

// AddSentinelAppBinary adds the "sentinel_app_binary" edges to the SentinelAppBinary entity.
func (slc *SentinelLibraryCreate) AddSentinelAppBinary(s ...*SentinelAppBinary) *SentinelLibraryCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return slc.AddSentinelAppBinaryIDs(ids...)
}

// Mutation returns the SentinelLibraryMutation object of the builder.
func (slc *SentinelLibraryCreate) Mutation() *SentinelLibraryMutation {
	return slc.mutation
}

// Save creates the SentinelLibrary in the database.
func (slc *SentinelLibraryCreate) Save(ctx context.Context) (*SentinelLibrary, error) {
	slc.defaults()
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *SentinelLibraryCreate) SaveX(ctx context.Context) *SentinelLibrary {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *SentinelLibraryCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *SentinelLibraryCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *SentinelLibraryCreate) defaults() {
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		v := sentinellibrary.DefaultUpdatedAt()
		slc.mutation.SetUpdatedAt(v)
	}
	if _, ok := slc.mutation.CreatedAt(); !ok {
		v := sentinellibrary.DefaultCreatedAt()
		slc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *SentinelLibraryCreate) check() error {
	if _, ok := slc.mutation.SentinelInfoID(); !ok {
		return &ValidationError{Name: "sentinel_info_id", err: errors.New(`ent: missing required field "SentinelLibrary.sentinel_info_id"`)}
	}
	if _, ok := slc.mutation.ReportedID(); !ok {
		return &ValidationError{Name: "reported_id", err: errors.New(`ent: missing required field "SentinelLibrary.reported_id"`)}
	}
	if _, ok := slc.mutation.DownloadBasePath(); !ok {
		return &ValidationError{Name: "download_base_path", err: errors.New(`ent: missing required field "SentinelLibrary.download_base_path"`)}
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SentinelLibrary.updated_at"`)}
	}
	if _, ok := slc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SentinelLibrary.created_at"`)}
	}
	if len(slc.mutation.SentinelInfoIDs()) == 0 {
		return &ValidationError{Name: "sentinel_info", err: errors.New(`ent: missing required edge "SentinelLibrary.sentinel_info"`)}
	}
	return nil
}

func (slc *SentinelLibraryCreate) sqlSave(ctx context.Context) (*SentinelLibrary, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *SentinelLibraryCreate) createSpec() (*SentinelLibrary, *sqlgraph.CreateSpec) {
	var (
		_node = &SentinelLibrary{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(sentinellibrary.Table, sqlgraph.NewFieldSpec(sentinellibrary.FieldID, field.TypeInt))
	)
	_spec.OnConflict = slc.conflict
	if value, ok := slc.mutation.ReportedID(); ok {
		_spec.SetField(sentinellibrary.FieldReportedID, field.TypeInt64, value)
		_node.ReportedID = value
	}
	if value, ok := slc.mutation.DownloadBasePath(); ok {
		_spec.SetField(sentinellibrary.FieldDownloadBasePath, field.TypeString, value)
		_node.DownloadBasePath = value
	}
	if value, ok := slc.mutation.UpdatedAt(); ok {
		_spec.SetField(sentinellibrary.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := slc.mutation.CreatedAt(); ok {
		_spec.SetField(sentinellibrary.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := slc.mutation.SentinelInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sentinellibrary.SentinelInfoTable,
			Columns: []string{sentinellibrary.SentinelInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelinfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SentinelInfoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := slc.mutation.SentinelAppBinaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sentinellibrary.SentinelAppBinaryTable,
			Columns: []string{sentinellibrary.SentinelAppBinaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt),
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
//	client.SentinelLibrary.Create().
//		SetSentinelInfoID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SentinelLibraryUpsert) {
//			SetSentinelInfoID(v+v).
//		}).
//		Exec(ctx)
func (slc *SentinelLibraryCreate) OnConflict(opts ...sql.ConflictOption) *SentinelLibraryUpsertOne {
	slc.conflict = opts
	return &SentinelLibraryUpsertOne{
		create: slc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (slc *SentinelLibraryCreate) OnConflictColumns(columns ...string) *SentinelLibraryUpsertOne {
	slc.conflict = append(slc.conflict, sql.ConflictColumns(columns...))
	return &SentinelLibraryUpsertOne{
		create: slc,
	}
}

type (
	// SentinelLibraryUpsertOne is the builder for "upsert"-ing
	//  one SentinelLibrary node.
	SentinelLibraryUpsertOne struct {
		create *SentinelLibraryCreate
	}

	// SentinelLibraryUpsert is the "OnConflict" setter.
	SentinelLibraryUpsert struct {
		*sql.UpdateSet
	}
)

// SetSentinelInfoID sets the "sentinel_info_id" field.
func (u *SentinelLibraryUpsert) SetSentinelInfoID(v int) *SentinelLibraryUpsert {
	u.Set(sentinellibrary.FieldSentinelInfoID, v)
	return u
}

// UpdateSentinelInfoID sets the "sentinel_info_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsert) UpdateSentinelInfoID() *SentinelLibraryUpsert {
	u.SetExcluded(sentinellibrary.FieldSentinelInfoID)
	return u
}

// SetReportedID sets the "reported_id" field.
func (u *SentinelLibraryUpsert) SetReportedID(v int64) *SentinelLibraryUpsert {
	u.Set(sentinellibrary.FieldReportedID, v)
	return u
}

// UpdateReportedID sets the "reported_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsert) UpdateReportedID() *SentinelLibraryUpsert {
	u.SetExcluded(sentinellibrary.FieldReportedID)
	return u
}

// AddReportedID adds v to the "reported_id" field.
func (u *SentinelLibraryUpsert) AddReportedID(v int64) *SentinelLibraryUpsert {
	u.Add(sentinellibrary.FieldReportedID, v)
	return u
}

// SetDownloadBasePath sets the "download_base_path" field.
func (u *SentinelLibraryUpsert) SetDownloadBasePath(v string) *SentinelLibraryUpsert {
	u.Set(sentinellibrary.FieldDownloadBasePath, v)
	return u
}

// UpdateDownloadBasePath sets the "download_base_path" field to the value that was provided on create.
func (u *SentinelLibraryUpsert) UpdateDownloadBasePath() *SentinelLibraryUpsert {
	u.SetExcluded(sentinellibrary.FieldDownloadBasePath)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SentinelLibraryUpsert) SetUpdatedAt(v time.Time) *SentinelLibraryUpsert {
	u.Set(sentinellibrary.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsert) UpdateUpdatedAt() *SentinelLibraryUpsert {
	u.SetExcluded(sentinellibrary.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SentinelLibraryUpsert) SetCreatedAt(v time.Time) *SentinelLibraryUpsert {
	u.Set(sentinellibrary.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsert) UpdateCreatedAt() *SentinelLibraryUpsert {
	u.SetExcluded(sentinellibrary.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SentinelLibraryUpsertOne) UpdateNewValues() *SentinelLibraryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SentinelLibraryUpsertOne) Ignore() *SentinelLibraryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SentinelLibraryUpsertOne) DoNothing() *SentinelLibraryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SentinelLibraryCreate.OnConflict
// documentation for more info.
func (u *SentinelLibraryUpsertOne) Update(set func(*SentinelLibraryUpsert)) *SentinelLibraryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SentinelLibraryUpsert{UpdateSet: update})
	}))
	return u
}

// SetSentinelInfoID sets the "sentinel_info_id" field.
func (u *SentinelLibraryUpsertOne) SetSentinelInfoID(v int) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetSentinelInfoID(v)
	})
}

// UpdateSentinelInfoID sets the "sentinel_info_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsertOne) UpdateSentinelInfoID() *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateSentinelInfoID()
	})
}

// SetReportedID sets the "reported_id" field.
func (u *SentinelLibraryUpsertOne) SetReportedID(v int64) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetReportedID(v)
	})
}

// AddReportedID adds v to the "reported_id" field.
func (u *SentinelLibraryUpsertOne) AddReportedID(v int64) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.AddReportedID(v)
	})
}

// UpdateReportedID sets the "reported_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsertOne) UpdateReportedID() *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateReportedID()
	})
}

// SetDownloadBasePath sets the "download_base_path" field.
func (u *SentinelLibraryUpsertOne) SetDownloadBasePath(v string) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetDownloadBasePath(v)
	})
}

// UpdateDownloadBasePath sets the "download_base_path" field to the value that was provided on create.
func (u *SentinelLibraryUpsertOne) UpdateDownloadBasePath() *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateDownloadBasePath()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SentinelLibraryUpsertOne) SetUpdatedAt(v time.Time) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsertOne) UpdateUpdatedAt() *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SentinelLibraryUpsertOne) SetCreatedAt(v time.Time) *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsertOne) UpdateCreatedAt() *SentinelLibraryUpsertOne {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *SentinelLibraryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SentinelLibraryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SentinelLibraryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SentinelLibraryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SentinelLibraryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SentinelLibraryCreateBulk is the builder for creating many SentinelLibrary entities in bulk.
type SentinelLibraryCreateBulk struct {
	config
	err      error
	builders []*SentinelLibraryCreate
	conflict []sql.ConflictOption
}

// Save creates the SentinelLibrary entities in the database.
func (slcb *SentinelLibraryCreateBulk) Save(ctx context.Context) ([]*SentinelLibrary, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*SentinelLibrary, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SentinelLibraryMutation)
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
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = slcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *SentinelLibraryCreateBulk) SaveX(ctx context.Context) []*SentinelLibrary {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *SentinelLibraryCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *SentinelLibraryCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SentinelLibrary.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SentinelLibraryUpsert) {
//			SetSentinelInfoID(v+v).
//		}).
//		Exec(ctx)
func (slcb *SentinelLibraryCreateBulk) OnConflict(opts ...sql.ConflictOption) *SentinelLibraryUpsertBulk {
	slcb.conflict = opts
	return &SentinelLibraryUpsertBulk{
		create: slcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (slcb *SentinelLibraryCreateBulk) OnConflictColumns(columns ...string) *SentinelLibraryUpsertBulk {
	slcb.conflict = append(slcb.conflict, sql.ConflictColumns(columns...))
	return &SentinelLibraryUpsertBulk{
		create: slcb,
	}
}

// SentinelLibraryUpsertBulk is the builder for "upsert"-ing
// a bulk of SentinelLibrary nodes.
type SentinelLibraryUpsertBulk struct {
	create *SentinelLibraryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SentinelLibraryUpsertBulk) UpdateNewValues() *SentinelLibraryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SentinelLibrary.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SentinelLibraryUpsertBulk) Ignore() *SentinelLibraryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SentinelLibraryUpsertBulk) DoNothing() *SentinelLibraryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SentinelLibraryCreateBulk.OnConflict
// documentation for more info.
func (u *SentinelLibraryUpsertBulk) Update(set func(*SentinelLibraryUpsert)) *SentinelLibraryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SentinelLibraryUpsert{UpdateSet: update})
	}))
	return u
}

// SetSentinelInfoID sets the "sentinel_info_id" field.
func (u *SentinelLibraryUpsertBulk) SetSentinelInfoID(v int) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetSentinelInfoID(v)
	})
}

// UpdateSentinelInfoID sets the "sentinel_info_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsertBulk) UpdateSentinelInfoID() *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateSentinelInfoID()
	})
}

// SetReportedID sets the "reported_id" field.
func (u *SentinelLibraryUpsertBulk) SetReportedID(v int64) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetReportedID(v)
	})
}

// AddReportedID adds v to the "reported_id" field.
func (u *SentinelLibraryUpsertBulk) AddReportedID(v int64) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.AddReportedID(v)
	})
}

// UpdateReportedID sets the "reported_id" field to the value that was provided on create.
func (u *SentinelLibraryUpsertBulk) UpdateReportedID() *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateReportedID()
	})
}

// SetDownloadBasePath sets the "download_base_path" field.
func (u *SentinelLibraryUpsertBulk) SetDownloadBasePath(v string) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetDownloadBasePath(v)
	})
}

// UpdateDownloadBasePath sets the "download_base_path" field to the value that was provided on create.
func (u *SentinelLibraryUpsertBulk) UpdateDownloadBasePath() *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateDownloadBasePath()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SentinelLibraryUpsertBulk) SetUpdatedAt(v time.Time) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsertBulk) UpdateUpdatedAt() *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SentinelLibraryUpsertBulk) SetCreatedAt(v time.Time) *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SentinelLibraryUpsertBulk) UpdateCreatedAt() *SentinelLibraryUpsertBulk {
	return u.Update(func(s *SentinelLibraryUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *SentinelLibraryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SentinelLibraryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SentinelLibraryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SentinelLibraryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
