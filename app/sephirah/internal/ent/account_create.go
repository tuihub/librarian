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
	"github.com/tuihub/librarian/app/sephirah/internal/ent/account"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetInternalID sets the "internal_id" field.
func (ac *AccountCreate) SetInternalID(i int64) *AccountCreate {
	ac.mutation.SetInternalID(i)
	return ac
}

// SetPlatform sets the "platform" field.
func (ac *AccountCreate) SetPlatform(a account.Platform) *AccountCreate {
	ac.mutation.SetPlatform(a)
	return ac
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (ac *AccountCreate) SetPlatformAccountID(s string) *AccountCreate {
	ac.mutation.SetPlatformAccountID(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetProfileURL sets the "profile_url" field.
func (ac *AccountCreate) SetProfileURL(s string) *AccountCreate {
	ac.mutation.SetProfileURL(s)
	return ac
}

// SetAvatarURL sets the "avatar_url" field.
func (ac *AccountCreate) SetAvatarURL(s string) *AccountCreate {
	ac.mutation.SetAvatarURL(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.InternalID(); !ok {
		return &ValidationError{Name: "internal_id", err: errors.New(`ent: missing required field "Account.internal_id"`)}
	}
	if _, ok := ac.mutation.Platform(); !ok {
		return &ValidationError{Name: "platform", err: errors.New(`ent: missing required field "Account.platform"`)}
	}
	if v, ok := ac.mutation.Platform(); ok {
		if err := account.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Account.platform": %w`, err)}
		}
	}
	if _, ok := ac.mutation.PlatformAccountID(); !ok {
		return &ValidationError{Name: "platform_account_id", err: errors.New(`ent: missing required field "Account.platform_account_id"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Account.name"`)}
	}
	if _, ok := ac.mutation.ProfileURL(); !ok {
		return &ValidationError{Name: "profile_url", err: errors.New(`ent: missing required field "Account.profile_url"`)}
	}
	if _, ok := ac.mutation.AvatarURL(); !ok {
		return &ValidationError{Name: "avatar_url", err: errors.New(`ent: missing required field "Account.avatar_url"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Account.created_at"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.InternalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: account.FieldInternalID,
		})
		_node.InternalID = value
	}
	if value, ok := ac.mutation.Platform(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: account.FieldPlatform,
		})
		_node.Platform = value
	}
	if value, ok := ac.mutation.PlatformAccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldPlatformAccountID,
		})
		_node.PlatformAccountID = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.ProfileURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldProfileURL,
		})
		_node.ProfileURL = value
	}
	if value, ok := ac.mutation.AvatarURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: account.FieldAvatarURL,
		})
		_node.AvatarURL = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: account.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetInternalID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetInternalID(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetInternalID sets the "internal_id" field.
func (u *AccountUpsert) SetInternalID(v int64) *AccountUpsert {
	u.Set(account.FieldInternalID, v)
	return u
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AccountUpsert) UpdateInternalID() *AccountUpsert {
	u.SetExcluded(account.FieldInternalID)
	return u
}

// AddInternalID adds v to the "internal_id" field.
func (u *AccountUpsert) AddInternalID(v int64) *AccountUpsert {
	u.Add(account.FieldInternalID, v)
	return u
}

// SetPlatform sets the "platform" field.
func (u *AccountUpsert) SetPlatform(v account.Platform) *AccountUpsert {
	u.Set(account.FieldPlatform, v)
	return u
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AccountUpsert) UpdatePlatform() *AccountUpsert {
	u.SetExcluded(account.FieldPlatform)
	return u
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (u *AccountUpsert) SetPlatformAccountID(v string) *AccountUpsert {
	u.Set(account.FieldPlatformAccountID, v)
	return u
}

// UpdatePlatformAccountID sets the "platform_account_id" field to the value that was provided on create.
func (u *AccountUpsert) UpdatePlatformAccountID() *AccountUpsert {
	u.SetExcluded(account.FieldPlatformAccountID)
	return u
}

// SetName sets the "name" field.
func (u *AccountUpsert) SetName(v string) *AccountUpsert {
	u.Set(account.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsert) UpdateName() *AccountUpsert {
	u.SetExcluded(account.FieldName)
	return u
}

// SetProfileURL sets the "profile_url" field.
func (u *AccountUpsert) SetProfileURL(v string) *AccountUpsert {
	u.Set(account.FieldProfileURL, v)
	return u
}

// UpdateProfileURL sets the "profile_url" field to the value that was provided on create.
func (u *AccountUpsert) UpdateProfileURL() *AccountUpsert {
	u.SetExcluded(account.FieldProfileURL)
	return u
}

// SetAvatarURL sets the "avatar_url" field.
func (u *AccountUpsert) SetAvatarURL(v string) *AccountUpsert {
	u.Set(account.FieldAvatarURL, v)
	return u
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *AccountUpsert) UpdateAvatarURL() *AccountUpsert {
	u.SetExcluded(account.FieldAvatarURL)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AccountUpsert) SetCreatedAt(v time.Time) *AccountUpsert {
	u.Set(account.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AccountUpsert) UpdateCreatedAt() *AccountUpsert {
	u.SetExcluded(account.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Account.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetInternalID sets the "internal_id" field.
func (u *AccountUpsertOne) SetInternalID(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetInternalID(v)
	})
}

// AddInternalID adds v to the "internal_id" field.
func (u *AccountUpsertOne) AddInternalID(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddInternalID(v)
	})
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateInternalID() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateInternalID()
	})
}

// SetPlatform sets the "platform" field.
func (u *AccountUpsertOne) SetPlatform(v account.Platform) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdatePlatform() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePlatform()
	})
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (u *AccountUpsertOne) SetPlatformAccountID(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetPlatformAccountID(v)
	})
}

// UpdatePlatformAccountID sets the "platform_account_id" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdatePlatformAccountID() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePlatformAccountID()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertOne) SetName(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateName() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetProfileURL sets the "profile_url" field.
func (u *AccountUpsertOne) SetProfileURL(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetProfileURL(v)
	})
}

// UpdateProfileURL sets the "profile_url" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateProfileURL() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateProfileURL()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *AccountUpsertOne) SetAvatarURL(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateAvatarURL() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAvatarURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AccountUpsertOne) SetCreatedAt(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateCreatedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetInternalID(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetInternalID sets the "internal_id" field.
func (u *AccountUpsertBulk) SetInternalID(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetInternalID(v)
	})
}

// AddInternalID adds v to the "internal_id" field.
func (u *AccountUpsertBulk) AddInternalID(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddInternalID(v)
	})
}

// UpdateInternalID sets the "internal_id" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateInternalID() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateInternalID()
	})
}

// SetPlatform sets the "platform" field.
func (u *AccountUpsertBulk) SetPlatform(v account.Platform) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdatePlatform() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePlatform()
	})
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (u *AccountUpsertBulk) SetPlatformAccountID(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetPlatformAccountID(v)
	})
}

// UpdatePlatformAccountID sets the "platform_account_id" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdatePlatformAccountID() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePlatformAccountID()
	})
}

// SetName sets the "name" field.
func (u *AccountUpsertBulk) SetName(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateName() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetProfileURL sets the "profile_url" field.
func (u *AccountUpsertBulk) SetProfileURL(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetProfileURL(v)
	})
}

// UpdateProfileURL sets the "profile_url" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateProfileURL() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateProfileURL()
	})
}

// SetAvatarURL sets the "avatar_url" field.
func (u *AccountUpsertBulk) SetAvatarURL(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetAvatarURL(v)
	})
}

// UpdateAvatarURL sets the "avatar_url" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateAvatarURL() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAvatarURL()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AccountUpsertBulk) SetCreatedAt(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateCreatedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}