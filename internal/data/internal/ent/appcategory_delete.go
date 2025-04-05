// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/appcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
)

// AppCategoryDelete is the builder for deleting a AppCategory entity.
type AppCategoryDelete struct {
	config
	hooks    []Hook
	mutation *AppCategoryMutation
}

// Where appends a list predicates to the AppCategoryDelete builder.
func (acd *AppCategoryDelete) Where(ps ...predicate.AppCategory) *AppCategoryDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *AppCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *AppCategoryDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *AppCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appcategory.Table, sqlgraph.NewFieldSpec(appcategory.FieldID, field.TypeInt64))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// AppCategoryDeleteOne is the builder for deleting a single AppCategory entity.
type AppCategoryDeleteOne struct {
	acd *AppCategoryDelete
}

// Where appends a list predicates to the AppCategoryDelete builder.
func (acdo *AppCategoryDeleteOne) Where(ps ...predicate.AppCategory) *AppCategoryDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *AppCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *AppCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}
