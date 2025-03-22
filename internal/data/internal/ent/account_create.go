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
	"github.com/tuihub/librarian/internal/data/internal/ent/account"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPlatform sets the "platform" field.
func (ac *AccountCreate) SetPlatform(s string) *AccountCreate {
	ac.mutation.SetPlatform(s)
	return ac
}

// SetPlatformAccountID sets the "platform_account_id" field.
func (ac *AccountCreate) SetPlatformAccountID(s string) *AccountCreate {
	ac.mutation.SetPlatformAccountID(s)
	return ac
}

// SetBoundUserID sets the "bound_user_id" field.
func (ac *AccountCreate) SetBoundUserID(mi model.InternalID) *AccountCreate {
	ac.mutation.SetBoundUserID(mi)
	return ac
}

// SetNillableBoundUserID sets the "bound_user_id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableBoundUserID(mi *model.InternalID) *AccountCreate {
	if mi != nil {
		ac.SetBoundUserID(*mi)
	}
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

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
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

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(mi model.InternalID) *AccountCreate {
	ac.mutation.SetID(mi)
	return ac
}

// SetBoundUser sets the "bound_user" edge to the User entity.
func (ac *AccountCreate) SetBoundUser(u *User) *AccountCreate {
	return ac.SetBoundUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
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
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.Platform(); !ok {
		return &ValidationError{Name: "platform", err: errors.New(`ent: missing required field "Account.platform"`)}
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
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Account.updated_at"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Account.created_at"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = model.InternalID(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Platform(); ok {
		_spec.SetField(account.FieldPlatform, field.TypeString, value)
		_node.Platform = value
	}
	if value, ok := ac.mutation.PlatformAccountID(); ok {
		_spec.SetField(account.FieldPlatformAccountID, field.TypeString, value)
		_node.PlatformAccountID = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.ProfileURL(); ok {
		_spec.SetField(account.FieldProfileURL, field.TypeString, value)
		_node.ProfileURL = value
	}
	if value, ok := ac.mutation.AvatarURL(); ok {
		_spec.SetField(account.FieldAvatarURL, field.TypeString, value)
		_node.AvatarURL = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.BoundUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.BoundUserTable,
			Columns: []string{account.BoundUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BoundUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetPlatform(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetPlatform(v+v).
//		}).
//		Exec(ctx)
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

// SetPlatform sets the "platform" field.
func (u *AccountUpsert) SetPlatform(v string) *AccountUpsert {
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

// SetBoundUserID sets the "bound_user_id" field.
func (u *AccountUpsert) SetBoundUserID(v model.InternalID) *AccountUpsert {
	u.Set(account.FieldBoundUserID, v)
	return u
}

// UpdateBoundUserID sets the "bound_user_id" field to the value that was provided on create.
func (u *AccountUpsert) UpdateBoundUserID() *AccountUpsert {
	u.SetExcluded(account.FieldBoundUserID)
	return u
}

// ClearBoundUserID clears the value of the "bound_user_id" field.
func (u *AccountUpsert) ClearBoundUserID() *AccountUpsert {
	u.SetNull(account.FieldBoundUserID)
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

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsert) SetUpdatedAt(v time.Time) *AccountUpsert {
	u.Set(account.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdatedAt() *AccountUpsert {
	u.SetExcluded(account.FieldUpdatedAt)
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

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(account.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
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

// SetPlatform sets the "platform" field.
func (u *AccountUpsertOne) SetPlatform(v string) *AccountUpsertOne {
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

// SetBoundUserID sets the "bound_user_id" field.
func (u *AccountUpsertOne) SetBoundUserID(v model.InternalID) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetBoundUserID(v)
	})
}

// UpdateBoundUserID sets the "bound_user_id" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateBoundUserID() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateBoundUserID()
	})
}

// ClearBoundUserID clears the value of the "bound_user_id" field.
func (u *AccountUpsertOne) ClearBoundUserID() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.ClearBoundUserID()
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

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsertOne) SetUpdatedAt(v time.Time) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdatedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
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
func (u *AccountUpsertOne) ID(ctx context.Context) (id model.InternalID, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) model.InternalID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
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
				var err error
				nodes[i], specs[i] = builder.createSpec()
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
//			SetPlatform(v+v).
//		}).
//		Exec(ctx)
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
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(account.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
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

// SetPlatform sets the "platform" field.
func (u *AccountUpsertBulk) SetPlatform(v string) *AccountUpsertBulk {
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

// SetBoundUserID sets the "bound_user_id" field.
func (u *AccountUpsertBulk) SetBoundUserID(v model.InternalID) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetBoundUserID(v)
	})
}

// UpdateBoundUserID sets the "bound_user_id" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateBoundUserID() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateBoundUserID()
	})
}

// ClearBoundUserID clears the value of the "bound_user_id" field.
func (u *AccountUpsertBulk) ClearBoundUserID() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.ClearBoundUserID()
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

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsertBulk) SetUpdatedAt(v time.Time) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdatedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
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
	if u.create.err != nil {
		return u.create.err
	}
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
