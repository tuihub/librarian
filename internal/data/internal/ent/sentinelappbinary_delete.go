// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
)

// SentinelAppBinaryDelete is the builder for deleting a SentinelAppBinary entity.
type SentinelAppBinaryDelete struct {
	config
	hooks    []Hook
	mutation *SentinelAppBinaryMutation
}

// Where appends a list predicates to the SentinelAppBinaryDelete builder.
func (sabd *SentinelAppBinaryDelete) Where(ps ...predicate.SentinelAppBinary) *SentinelAppBinaryDelete {
	sabd.mutation.Where(ps...)
	return sabd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sabd *SentinelAppBinaryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sabd.sqlExec, sabd.mutation, sabd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sabd *SentinelAppBinaryDelete) ExecX(ctx context.Context) int {
	n, err := sabd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sabd *SentinelAppBinaryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sentinelappbinary.Table, sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt))
	if ps := sabd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sabd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sabd.mutation.done = true
	return affected, err
}

// SentinelAppBinaryDeleteOne is the builder for deleting a single SentinelAppBinary entity.
type SentinelAppBinaryDeleteOne struct {
	sabd *SentinelAppBinaryDelete
}

// Where appends a list predicates to the SentinelAppBinaryDelete builder.
func (sabdo *SentinelAppBinaryDeleteOne) Where(ps ...predicate.SentinelAppBinary) *SentinelAppBinaryDeleteOne {
	sabdo.sabd.mutation.Where(ps...)
	return sabdo
}

// Exec executes the deletion query.
func (sabdo *SentinelAppBinaryDeleteOne) Exec(ctx context.Context) error {
	n, err := sabdo.sabd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sentinelappbinary.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sabdo *SentinelAppBinaryDeleteOne) ExecX(ctx context.Context) {
	if err := sabdo.Exec(ctx); err != nil {
		panic(err)
	}
}
