// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelinfo"
)

// SentinelInfoDelete is the builder for deleting a SentinelInfo entity.
type SentinelInfoDelete struct {
	config
	hooks    []Hook
	mutation *SentinelInfoMutation
}

// Where appends a list predicates to the SentinelInfoDelete builder.
func (sid *SentinelInfoDelete) Where(ps ...predicate.SentinelInfo) *SentinelInfoDelete {
	sid.mutation.Where(ps...)
	return sid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sid *SentinelInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sid.sqlExec, sid.mutation, sid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sid *SentinelInfoDelete) ExecX(ctx context.Context) int {
	n, err := sid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sid *SentinelInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sentinelinfo.Table, sqlgraph.NewFieldSpec(sentinelinfo.FieldID, field.TypeInt))
	if ps := sid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sid.mutation.done = true
	return affected, err
}

// SentinelInfoDeleteOne is the builder for deleting a single SentinelInfo entity.
type SentinelInfoDeleteOne struct {
	sid *SentinelInfoDelete
}

// Where appends a list predicates to the SentinelInfoDelete builder.
func (sido *SentinelInfoDeleteOne) Where(ps ...predicate.SentinelInfo) *SentinelInfoDeleteOne {
	sido.sid.mutation.Where(ps...)
	return sido
}

// Exec executes the deletion query.
func (sido *SentinelInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := sido.sid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sentinelinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sido *SentinelInfoDeleteOne) ExecX(ctx context.Context) {
	if err := sido.Exec(ctx); err != nil {
		panic(err)
	}
}
