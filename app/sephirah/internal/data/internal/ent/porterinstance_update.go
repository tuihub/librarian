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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/model/modeltiphereth"
)

// PorterInstanceUpdate is the builder for updating PorterInstance entities.
type PorterInstanceUpdate struct {
	config
	hooks    []Hook
	mutation *PorterInstanceMutation
}

// Where appends a list predicates to the PorterInstanceUpdate builder.
func (piu *PorterInstanceUpdate) Where(ps ...predicate.PorterInstance) *PorterInstanceUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetName sets the "name" field.
func (piu *PorterInstanceUpdate) SetName(s string) *PorterInstanceUpdate {
	piu.mutation.SetName(s)
	return piu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableName(s *string) *PorterInstanceUpdate {
	if s != nil {
		piu.SetName(*s)
	}
	return piu
}

// SetVersion sets the "version" field.
func (piu *PorterInstanceUpdate) SetVersion(s string) *PorterInstanceUpdate {
	piu.mutation.SetVersion(s)
	return piu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableVersion(s *string) *PorterInstanceUpdate {
	if s != nil {
		piu.SetVersion(*s)
	}
	return piu
}

// SetGlobalName sets the "global_name" field.
func (piu *PorterInstanceUpdate) SetGlobalName(s string) *PorterInstanceUpdate {
	piu.mutation.SetGlobalName(s)
	return piu
}

// SetNillableGlobalName sets the "global_name" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableGlobalName(s *string) *PorterInstanceUpdate {
	if s != nil {
		piu.SetGlobalName(*s)
	}
	return piu
}

// SetAddress sets the "address" field.
func (piu *PorterInstanceUpdate) SetAddress(s string) *PorterInstanceUpdate {
	piu.mutation.SetAddress(s)
	return piu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableAddress(s *string) *PorterInstanceUpdate {
	if s != nil {
		piu.SetAddress(*s)
	}
	return piu
}

// SetFeatureSummary sets the "feature_summary" field.
func (piu *PorterInstanceUpdate) SetFeatureSummary(mfs *modeltiphereth.PorterFeatureSummary) *PorterInstanceUpdate {
	piu.mutation.SetFeatureSummary(mfs)
	return piu
}

// SetStatus sets the "status" field.
func (piu *PorterInstanceUpdate) SetStatus(po porterinstance.Status) *PorterInstanceUpdate {
	piu.mutation.SetStatus(po)
	return piu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableStatus(po *porterinstance.Status) *PorterInstanceUpdate {
	if po != nil {
		piu.SetStatus(*po)
	}
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *PorterInstanceUpdate) SetUpdatedAt(t time.Time) *PorterInstanceUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetCreatedAt sets the "created_at" field.
func (piu *PorterInstanceUpdate) SetCreatedAt(t time.Time) *PorterInstanceUpdate {
	piu.mutation.SetCreatedAt(t)
	return piu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piu *PorterInstanceUpdate) SetNillableCreatedAt(t *time.Time) *PorterInstanceUpdate {
	if t != nil {
		piu.SetCreatedAt(*t)
	}
	return piu
}

// Mutation returns the PorterInstanceMutation object of the builder.
func (piu *PorterInstanceUpdate) Mutation() *PorterInstanceMutation {
	return piu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PorterInstanceUpdate) Save(ctx context.Context) (int, error) {
	piu.defaults()
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PorterInstanceUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PorterInstanceUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PorterInstanceUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piu *PorterInstanceUpdate) defaults() {
	if _, ok := piu.mutation.UpdatedAt(); !ok {
		v := porterinstance.UpdateDefaultUpdatedAt()
		piu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *PorterInstanceUpdate) check() error {
	if v, ok := piu.mutation.Status(); ok {
		if err := porterinstance.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "PorterInstance.status": %w`, err)}
		}
	}
	return nil
}

func (piu *PorterInstanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(porterinstance.Table, porterinstance.Columns, sqlgraph.NewFieldSpec(porterinstance.FieldID, field.TypeInt64))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.Name(); ok {
		_spec.SetField(porterinstance.FieldName, field.TypeString, value)
	}
	if value, ok := piu.mutation.Version(); ok {
		_spec.SetField(porterinstance.FieldVersion, field.TypeString, value)
	}
	if value, ok := piu.mutation.GlobalName(); ok {
		_spec.SetField(porterinstance.FieldGlobalName, field.TypeString, value)
	}
	if value, ok := piu.mutation.Address(); ok {
		_spec.SetField(porterinstance.FieldAddress, field.TypeString, value)
	}
	if value, ok := piu.mutation.FeatureSummary(); ok {
		_spec.SetField(porterinstance.FieldFeatureSummary, field.TypeJSON, value)
	}
	if value, ok := piu.mutation.Status(); ok {
		_spec.SetField(porterinstance.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(porterinstance.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := piu.mutation.CreatedAt(); ok {
		_spec.SetField(porterinstance.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{porterinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// PorterInstanceUpdateOne is the builder for updating a single PorterInstance entity.
type PorterInstanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PorterInstanceMutation
}

// SetName sets the "name" field.
func (piuo *PorterInstanceUpdateOne) SetName(s string) *PorterInstanceUpdateOne {
	piuo.mutation.SetName(s)
	return piuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableName(s *string) *PorterInstanceUpdateOne {
	if s != nil {
		piuo.SetName(*s)
	}
	return piuo
}

// SetVersion sets the "version" field.
func (piuo *PorterInstanceUpdateOne) SetVersion(s string) *PorterInstanceUpdateOne {
	piuo.mutation.SetVersion(s)
	return piuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableVersion(s *string) *PorterInstanceUpdateOne {
	if s != nil {
		piuo.SetVersion(*s)
	}
	return piuo
}

// SetGlobalName sets the "global_name" field.
func (piuo *PorterInstanceUpdateOne) SetGlobalName(s string) *PorterInstanceUpdateOne {
	piuo.mutation.SetGlobalName(s)
	return piuo
}

// SetNillableGlobalName sets the "global_name" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableGlobalName(s *string) *PorterInstanceUpdateOne {
	if s != nil {
		piuo.SetGlobalName(*s)
	}
	return piuo
}

// SetAddress sets the "address" field.
func (piuo *PorterInstanceUpdateOne) SetAddress(s string) *PorterInstanceUpdateOne {
	piuo.mutation.SetAddress(s)
	return piuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableAddress(s *string) *PorterInstanceUpdateOne {
	if s != nil {
		piuo.SetAddress(*s)
	}
	return piuo
}

// SetFeatureSummary sets the "feature_summary" field.
func (piuo *PorterInstanceUpdateOne) SetFeatureSummary(mfs *modeltiphereth.PorterFeatureSummary) *PorterInstanceUpdateOne {
	piuo.mutation.SetFeatureSummary(mfs)
	return piuo
}

// SetStatus sets the "status" field.
func (piuo *PorterInstanceUpdateOne) SetStatus(po porterinstance.Status) *PorterInstanceUpdateOne {
	piuo.mutation.SetStatus(po)
	return piuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableStatus(po *porterinstance.Status) *PorterInstanceUpdateOne {
	if po != nil {
		piuo.SetStatus(*po)
	}
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *PorterInstanceUpdateOne) SetUpdatedAt(t time.Time) *PorterInstanceUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetCreatedAt sets the "created_at" field.
func (piuo *PorterInstanceUpdateOne) SetCreatedAt(t time.Time) *PorterInstanceUpdateOne {
	piuo.mutation.SetCreatedAt(t)
	return piuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piuo *PorterInstanceUpdateOne) SetNillableCreatedAt(t *time.Time) *PorterInstanceUpdateOne {
	if t != nil {
		piuo.SetCreatedAt(*t)
	}
	return piuo
}

// Mutation returns the PorterInstanceMutation object of the builder.
func (piuo *PorterInstanceUpdateOne) Mutation() *PorterInstanceMutation {
	return piuo.mutation
}

// Where appends a list predicates to the PorterInstanceUpdate builder.
func (piuo *PorterInstanceUpdateOne) Where(ps ...predicate.PorterInstance) *PorterInstanceUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PorterInstanceUpdateOne) Select(field string, fields ...string) *PorterInstanceUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PorterInstance entity.
func (piuo *PorterInstanceUpdateOne) Save(ctx context.Context) (*PorterInstance, error) {
	piuo.defaults()
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PorterInstanceUpdateOne) SaveX(ctx context.Context) *PorterInstance {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PorterInstanceUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PorterInstanceUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piuo *PorterInstanceUpdateOne) defaults() {
	if _, ok := piuo.mutation.UpdatedAt(); !ok {
		v := porterinstance.UpdateDefaultUpdatedAt()
		piuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *PorterInstanceUpdateOne) check() error {
	if v, ok := piuo.mutation.Status(); ok {
		if err := porterinstance.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "PorterInstance.status": %w`, err)}
		}
	}
	return nil
}

func (piuo *PorterInstanceUpdateOne) sqlSave(ctx context.Context) (_node *PorterInstance, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(porterinstance.Table, porterinstance.Columns, sqlgraph.NewFieldSpec(porterinstance.FieldID, field.TypeInt64))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PorterInstance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, porterinstance.FieldID)
		for _, f := range fields {
			if !porterinstance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != porterinstance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.Name(); ok {
		_spec.SetField(porterinstance.FieldName, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Version(); ok {
		_spec.SetField(porterinstance.FieldVersion, field.TypeString, value)
	}
	if value, ok := piuo.mutation.GlobalName(); ok {
		_spec.SetField(porterinstance.FieldGlobalName, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Address(); ok {
		_spec.SetField(porterinstance.FieldAddress, field.TypeString, value)
	}
	if value, ok := piuo.mutation.FeatureSummary(); ok {
		_spec.SetField(porterinstance.FieldFeatureSummary, field.TypeJSON, value)
	}
	if value, ok := piuo.mutation.Status(); ok {
		_spec.SetField(porterinstance.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(porterinstance.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.CreatedAt(); ok {
		_spec.SetField(porterinstance.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &PorterInstance{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{porterinstance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
