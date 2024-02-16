// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinstruntime"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
)

// AppInstRunTimeDelete is the builder for deleting a AppInstRunTime entity.
type AppInstRunTimeDelete struct {
	config
	hooks    []Hook
	mutation *AppInstRunTimeMutation
}

// Where appends a list predicates to the AppInstRunTimeDelete builder.
func (airtd *AppInstRunTimeDelete) Where(ps ...predicate.AppInstRunTime) *AppInstRunTimeDelete {
	airtd.mutation.Where(ps...)
	return airtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (airtd *AppInstRunTimeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, airtd.sqlExec, airtd.mutation, airtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (airtd *AppInstRunTimeDelete) ExecX(ctx context.Context) int {
	n, err := airtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (airtd *AppInstRunTimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appinstruntime.Table, sqlgraph.NewFieldSpec(appinstruntime.FieldID, field.TypeInt))
	if ps := airtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, airtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	airtd.mutation.done = true
	return affected, err
}

// AppInstRunTimeDeleteOne is the builder for deleting a single AppInstRunTime entity.
type AppInstRunTimeDeleteOne struct {
	airtd *AppInstRunTimeDelete
}

// Where appends a list predicates to the AppInstRunTimeDelete builder.
func (airtdo *AppInstRunTimeDeleteOne) Where(ps ...predicate.AppInstRunTime) *AppInstRunTimeDeleteOne {
	airtdo.airtd.mutation.Where(ps...)
	return airtdo
}

// Exec executes the deletion query.
func (airtdo *AppInstRunTimeDeleteOne) Exec(ctx context.Context) error {
	n, err := airtdo.airtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appinstruntime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (airtdo *AppInstRunTimeDeleteOne) ExecX(ctx context.Context) {
	if err := airtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
