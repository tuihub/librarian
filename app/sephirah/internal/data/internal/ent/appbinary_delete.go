// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appbinary"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
)

// AppBinaryDelete is the builder for deleting a AppBinary entity.
type AppBinaryDelete struct {
	config
	hooks    []Hook
	mutation *AppBinaryMutation
}

// Where appends a list predicates to the AppBinaryDelete builder.
func (abd *AppBinaryDelete) Where(ps ...predicate.AppBinary) *AppBinaryDelete {
	abd.mutation.Where(ps...)
	return abd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (abd *AppBinaryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, abd.sqlExec, abd.mutation, abd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (abd *AppBinaryDelete) ExecX(ctx context.Context) int {
	n, err := abd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (abd *AppBinaryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appbinary.Table, sqlgraph.NewFieldSpec(appbinary.FieldID, field.TypeInt64))
	if ps := abd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, abd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	abd.mutation.done = true
	return affected, err
}

// AppBinaryDeleteOne is the builder for deleting a single AppBinary entity.
type AppBinaryDeleteOne struct {
	abd *AppBinaryDelete
}

// Where appends a list predicates to the AppBinaryDelete builder.
func (abdo *AppBinaryDeleteOne) Where(ps ...predicate.AppBinary) *AppBinaryDeleteOne {
	abdo.abd.mutation.Where(ps...)
	return abdo
}

// Exec executes the deletion query.
func (abdo *AppBinaryDeleteOne) Exec(ctx context.Context) error {
	n, err := abdo.abd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appbinary.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (abdo *AppBinaryDeleteOne) ExecX(ctx context.Context) {
	if err := abdo.Exec(ctx); err != nil {
		panic(err)
	}
}
