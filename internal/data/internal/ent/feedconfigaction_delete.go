// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/feedconfigaction"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
)

// FeedConfigActionDelete is the builder for deleting a FeedConfigAction entity.
type FeedConfigActionDelete struct {
	config
	hooks    []Hook
	mutation *FeedConfigActionMutation
}

// Where appends a list predicates to the FeedConfigActionDelete builder.
func (fcad *FeedConfigActionDelete) Where(ps ...predicate.FeedConfigAction) *FeedConfigActionDelete {
	fcad.mutation.Where(ps...)
	return fcad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fcad *FeedConfigActionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fcad.sqlExec, fcad.mutation, fcad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fcad *FeedConfigActionDelete) ExecX(ctx context.Context) int {
	n, err := fcad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fcad *FeedConfigActionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feedconfigaction.Table, sqlgraph.NewFieldSpec(feedconfigaction.FieldID, field.TypeInt))
	if ps := fcad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fcad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fcad.mutation.done = true
	return affected, err
}

// FeedConfigActionDeleteOne is the builder for deleting a single FeedConfigAction entity.
type FeedConfigActionDeleteOne struct {
	fcad *FeedConfigActionDelete
}

// Where appends a list predicates to the FeedConfigActionDelete builder.
func (fcado *FeedConfigActionDeleteOne) Where(ps ...predicate.FeedConfigAction) *FeedConfigActionDeleteOne {
	fcado.fcad.mutation.Where(ps...)
	return fcado
}

// Exec executes the deletion query.
func (fcado *FeedConfigActionDeleteOne) Exec(ctx context.Context) error {
	n, err := fcado.fcad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feedconfigaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fcado *FeedConfigActionDeleteOne) ExecX(ctx context.Context) {
	if err := fcado.Exec(ctx); err != nil {
		panic(err)
	}
}
