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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinstruntime"
	"github.com/tuihub/librarian/internal/model"
)

// AppInstRunTimeCreate is the builder for creating a AppInstRunTime entity.
type AppInstRunTimeCreate struct {
	config
	mutation *AppInstRunTimeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (airtc *AppInstRunTimeCreate) SetUserID(mi model.InternalID) *AppInstRunTimeCreate {
	airtc.mutation.SetUserID(mi)
	return airtc
}

// SetAppID sets the "app_id" field.
func (airtc *AppInstRunTimeCreate) SetAppID(mi model.InternalID) *AppInstRunTimeCreate {
	airtc.mutation.SetAppID(mi)
	return airtc
}

// SetStartTime sets the "start_time" field.
func (airtc *AppInstRunTimeCreate) SetStartTime(t time.Time) *AppInstRunTimeCreate {
	airtc.mutation.SetStartTime(t)
	return airtc
}

// SetRunDuration sets the "run_duration" field.
func (airtc *AppInstRunTimeCreate) SetRunDuration(t time.Duration) *AppInstRunTimeCreate {
	airtc.mutation.SetRunDuration(t)
	return airtc
}

// SetUpdatedAt sets the "updated_at" field.
func (airtc *AppInstRunTimeCreate) SetUpdatedAt(t time.Time) *AppInstRunTimeCreate {
	airtc.mutation.SetUpdatedAt(t)
	return airtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (airtc *AppInstRunTimeCreate) SetNillableUpdatedAt(t *time.Time) *AppInstRunTimeCreate {
	if t != nil {
		airtc.SetUpdatedAt(*t)
	}
	return airtc
}

// SetCreatedAt sets the "created_at" field.
func (airtc *AppInstRunTimeCreate) SetCreatedAt(t time.Time) *AppInstRunTimeCreate {
	airtc.mutation.SetCreatedAt(t)
	return airtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (airtc *AppInstRunTimeCreate) SetNillableCreatedAt(t *time.Time) *AppInstRunTimeCreate {
	if t != nil {
		airtc.SetCreatedAt(*t)
	}
	return airtc
}

// Mutation returns the AppInstRunTimeMutation object of the builder.
func (airtc *AppInstRunTimeCreate) Mutation() *AppInstRunTimeMutation {
	return airtc.mutation
}

// Save creates the AppInstRunTime in the database.
func (airtc *AppInstRunTimeCreate) Save(ctx context.Context) (*AppInstRunTime, error) {
	airtc.defaults()
	return withHooks(ctx, airtc.sqlSave, airtc.mutation, airtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (airtc *AppInstRunTimeCreate) SaveX(ctx context.Context) *AppInstRunTime {
	v, err := airtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (airtc *AppInstRunTimeCreate) Exec(ctx context.Context) error {
	_, err := airtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (airtc *AppInstRunTimeCreate) ExecX(ctx context.Context) {
	if err := airtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (airtc *AppInstRunTimeCreate) defaults() {
	if _, ok := airtc.mutation.UpdatedAt(); !ok {
		v := appinstruntime.DefaultUpdatedAt()
		airtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := airtc.mutation.CreatedAt(); !ok {
		v := appinstruntime.DefaultCreatedAt()
		airtc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (airtc *AppInstRunTimeCreate) check() error {
	if _, ok := airtc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppInstRunTime.user_id"`)}
	}
	if _, ok := airtc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppInstRunTime.app_id"`)}
	}
	if _, ok := airtc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "AppInstRunTime.start_time"`)}
	}
	if _, ok := airtc.mutation.RunDuration(); !ok {
		return &ValidationError{Name: "run_duration", err: errors.New(`ent: missing required field "AppInstRunTime.run_duration"`)}
	}
	if _, ok := airtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppInstRunTime.updated_at"`)}
	}
	if _, ok := airtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppInstRunTime.created_at"`)}
	}
	return nil
}

func (airtc *AppInstRunTimeCreate) sqlSave(ctx context.Context) (*AppInstRunTime, error) {
	if err := airtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := airtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, airtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	airtc.mutation.id = &_node.ID
	airtc.mutation.done = true
	return _node, nil
}

func (airtc *AppInstRunTimeCreate) createSpec() (*AppInstRunTime, *sqlgraph.CreateSpec) {
	var (
		_node = &AppInstRunTime{config: airtc.config}
		_spec = sqlgraph.NewCreateSpec(appinstruntime.Table, sqlgraph.NewFieldSpec(appinstruntime.FieldID, field.TypeInt))
	)
	_spec.OnConflict = airtc.conflict
	if value, ok := airtc.mutation.UserID(); ok {
		_spec.SetField(appinstruntime.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := airtc.mutation.AppID(); ok {
		_spec.SetField(appinstruntime.FieldAppID, field.TypeInt64, value)
		_node.AppID = value
	}
	if value, ok := airtc.mutation.StartTime(); ok {
		_spec.SetField(appinstruntime.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := airtc.mutation.RunDuration(); ok {
		_spec.SetField(appinstruntime.FieldRunDuration, field.TypeInt64, value)
		_node.RunDuration = value
	}
	if value, ok := airtc.mutation.UpdatedAt(); ok {
		_spec.SetField(appinstruntime.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := airtc.mutation.CreatedAt(); ok {
		_spec.SetField(appinstruntime.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppInstRunTime.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppInstRunTimeUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (airtc *AppInstRunTimeCreate) OnConflict(opts ...sql.ConflictOption) *AppInstRunTimeUpsertOne {
	airtc.conflict = opts
	return &AppInstRunTimeUpsertOne{
		create: airtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (airtc *AppInstRunTimeCreate) OnConflictColumns(columns ...string) *AppInstRunTimeUpsertOne {
	airtc.conflict = append(airtc.conflict, sql.ConflictColumns(columns...))
	return &AppInstRunTimeUpsertOne{
		create: airtc,
	}
}

type (
	// AppInstRunTimeUpsertOne is the builder for "upsert"-ing
	//  one AppInstRunTime node.
	AppInstRunTimeUpsertOne struct {
		create *AppInstRunTimeCreate
	}

	// AppInstRunTimeUpsert is the "OnConflict" setter.
	AppInstRunTimeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *AppInstRunTimeUpsert) SetUserID(v model.InternalID) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateUserID() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldUserID)
	return u
}

// AddUserID adds v to the "user_id" field.
func (u *AppInstRunTimeUpsert) AddUserID(v model.InternalID) *AppInstRunTimeUpsert {
	u.Add(appinstruntime.FieldUserID, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppInstRunTimeUpsert) SetAppID(v model.InternalID) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateAppID() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldAppID)
	return u
}

// AddAppID adds v to the "app_id" field.
func (u *AppInstRunTimeUpsert) AddAppID(v model.InternalID) *AppInstRunTimeUpsert {
	u.Add(appinstruntime.FieldAppID, v)
	return u
}

// SetStartTime sets the "start_time" field.
func (u *AppInstRunTimeUpsert) SetStartTime(v time.Time) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldStartTime, v)
	return u
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateStartTime() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldStartTime)
	return u
}

// SetRunDuration sets the "run_duration" field.
func (u *AppInstRunTimeUpsert) SetRunDuration(v time.Duration) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldRunDuration, v)
	return u
}

// UpdateRunDuration sets the "run_duration" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateRunDuration() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldRunDuration)
	return u
}

// AddRunDuration adds v to the "run_duration" field.
func (u *AppInstRunTimeUpsert) AddRunDuration(v time.Duration) *AppInstRunTimeUpsert {
	u.Add(appinstruntime.FieldRunDuration, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppInstRunTimeUpsert) SetUpdatedAt(v time.Time) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateUpdatedAt() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldUpdatedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppInstRunTimeUpsert) SetCreatedAt(v time.Time) *AppInstRunTimeUpsert {
	u.Set(appinstruntime.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsert) UpdateCreatedAt() *AppInstRunTimeUpsert {
	u.SetExcluded(appinstruntime.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AppInstRunTimeUpsertOne) UpdateNewValues() *AppInstRunTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppInstRunTimeUpsertOne) Ignore() *AppInstRunTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppInstRunTimeUpsertOne) DoNothing() *AppInstRunTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppInstRunTimeCreate.OnConflict
// documentation for more info.
func (u *AppInstRunTimeUpsertOne) Update(set func(*AppInstRunTimeUpsert)) *AppInstRunTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppInstRunTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppInstRunTimeUpsertOne) SetUserID(v model.InternalID) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *AppInstRunTimeUpsertOne) AddUserID(v model.InternalID) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateUserID() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateUserID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppInstRunTimeUpsertOne) SetAppID(v model.InternalID) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *AppInstRunTimeUpsertOne) AddAppID(v model.InternalID) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateAppID() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateAppID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *AppInstRunTimeUpsertOne) SetStartTime(v time.Time) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateStartTime() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateStartTime()
	})
}

// SetRunDuration sets the "run_duration" field.
func (u *AppInstRunTimeUpsertOne) SetRunDuration(v time.Duration) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetRunDuration(v)
	})
}

// AddRunDuration adds v to the "run_duration" field.
func (u *AppInstRunTimeUpsertOne) AddRunDuration(v time.Duration) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddRunDuration(v)
	})
}

// UpdateRunDuration sets the "run_duration" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateRunDuration() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateRunDuration()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppInstRunTimeUpsertOne) SetUpdatedAt(v time.Time) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateUpdatedAt() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppInstRunTimeUpsertOne) SetCreatedAt(v time.Time) *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertOne) UpdateCreatedAt() *AppInstRunTimeUpsertOne {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppInstRunTimeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppInstRunTimeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppInstRunTimeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppInstRunTimeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppInstRunTimeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppInstRunTimeCreateBulk is the builder for creating many AppInstRunTime entities in bulk.
type AppInstRunTimeCreateBulk struct {
	config
	err      error
	builders []*AppInstRunTimeCreate
	conflict []sql.ConflictOption
}

// Save creates the AppInstRunTime entities in the database.
func (airtcb *AppInstRunTimeCreateBulk) Save(ctx context.Context) ([]*AppInstRunTime, error) {
	if airtcb.err != nil {
		return nil, airtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(airtcb.builders))
	nodes := make([]*AppInstRunTime, len(airtcb.builders))
	mutators := make([]Mutator, len(airtcb.builders))
	for i := range airtcb.builders {
		func(i int, root context.Context) {
			builder := airtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppInstRunTimeMutation)
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
					_, err = mutators[i+1].Mutate(root, airtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = airtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, airtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, airtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (airtcb *AppInstRunTimeCreateBulk) SaveX(ctx context.Context) []*AppInstRunTime {
	v, err := airtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (airtcb *AppInstRunTimeCreateBulk) Exec(ctx context.Context) error {
	_, err := airtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (airtcb *AppInstRunTimeCreateBulk) ExecX(ctx context.Context) {
	if err := airtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppInstRunTime.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppInstRunTimeUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (airtcb *AppInstRunTimeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppInstRunTimeUpsertBulk {
	airtcb.conflict = opts
	return &AppInstRunTimeUpsertBulk{
		create: airtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (airtcb *AppInstRunTimeCreateBulk) OnConflictColumns(columns ...string) *AppInstRunTimeUpsertBulk {
	airtcb.conflict = append(airtcb.conflict, sql.ConflictColumns(columns...))
	return &AppInstRunTimeUpsertBulk{
		create: airtcb,
	}
}

// AppInstRunTimeUpsertBulk is the builder for "upsert"-ing
// a bulk of AppInstRunTime nodes.
type AppInstRunTimeUpsertBulk struct {
	create *AppInstRunTimeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AppInstRunTimeUpsertBulk) UpdateNewValues() *AppInstRunTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppInstRunTime.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppInstRunTimeUpsertBulk) Ignore() *AppInstRunTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppInstRunTimeUpsertBulk) DoNothing() *AppInstRunTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppInstRunTimeCreateBulk.OnConflict
// documentation for more info.
func (u *AppInstRunTimeUpsertBulk) Update(set func(*AppInstRunTimeUpsert)) *AppInstRunTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppInstRunTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppInstRunTimeUpsertBulk) SetUserID(v model.InternalID) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *AppInstRunTimeUpsertBulk) AddUserID(v model.InternalID) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateUserID() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateUserID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppInstRunTimeUpsertBulk) SetAppID(v model.InternalID) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *AppInstRunTimeUpsertBulk) AddAppID(v model.InternalID) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateAppID() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateAppID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *AppInstRunTimeUpsertBulk) SetStartTime(v time.Time) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateStartTime() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateStartTime()
	})
}

// SetRunDuration sets the "run_duration" field.
func (u *AppInstRunTimeUpsertBulk) SetRunDuration(v time.Duration) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetRunDuration(v)
	})
}

// AddRunDuration adds v to the "run_duration" field.
func (u *AppInstRunTimeUpsertBulk) AddRunDuration(v time.Duration) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.AddRunDuration(v)
	})
}

// UpdateRunDuration sets the "run_duration" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateRunDuration() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateRunDuration()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppInstRunTimeUpsertBulk) SetUpdatedAt(v time.Time) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateUpdatedAt() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AppInstRunTimeUpsertBulk) SetCreatedAt(v time.Time) *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppInstRunTimeUpsertBulk) UpdateCreatedAt() *AppInstRunTimeUpsertBulk {
	return u.Update(func(s *AppInstRunTimeUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AppInstRunTimeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppInstRunTimeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppInstRunTimeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppInstRunTimeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}