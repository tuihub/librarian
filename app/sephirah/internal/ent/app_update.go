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
	"github.com/tuihub/librarian/app/sephirah/internal/ent/app"
	"github.com/tuihub/librarian/app/sephirah/internal/ent/predicate"
)

// AppUpdate is the builder for updating App entities.
type AppUpdate struct {
	config
	hooks    []Hook
	mutation *AppMutation
}

// Where appends a list predicates to the AppUpdate builder.
func (au *AppUpdate) Where(ps ...predicate.App) *AppUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetInternalID sets the "internal_id" field.
func (au *AppUpdate) SetInternalID(i int64) *AppUpdate {
	au.mutation.ResetInternalID()
	au.mutation.SetInternalID(i)
	return au
}

// AddInternalID adds i to the "internal_id" field.
func (au *AppUpdate) AddInternalID(i int64) *AppUpdate {
	au.mutation.AddInternalID(i)
	return au
}

// SetSource sets the "source" field.
func (au *AppUpdate) SetSource(a app.Source) *AppUpdate {
	au.mutation.SetSource(a)
	return au
}

// SetSourceAppID sets the "source_app_id" field.
func (au *AppUpdate) SetSourceAppID(s string) *AppUpdate {
	au.mutation.SetSourceAppID(s)
	return au
}

// SetSourceURL sets the "source_url" field.
func (au *AppUpdate) SetSourceURL(s string) *AppUpdate {
	au.mutation.SetSourceURL(s)
	return au
}

// SetName sets the "name" field.
func (au *AppUpdate) SetName(s string) *AppUpdate {
	au.mutation.SetName(s)
	return au
}

// SetType sets the "type" field.
func (au *AppUpdate) SetType(a app.Type) *AppUpdate {
	au.mutation.SetType(a)
	return au
}

// SetShortDescription sets the "short_description" field.
func (au *AppUpdate) SetShortDescription(s string) *AppUpdate {
	au.mutation.SetShortDescription(s)
	return au
}

// SetDescription sets the "description" field.
func (au *AppUpdate) SetDescription(s string) *AppUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetImageURL sets the "image_url" field.
func (au *AppUpdate) SetImageURL(s string) *AppUpdate {
	au.mutation.SetImageURL(s)
	return au
}

// SetReleaseDate sets the "release_date" field.
func (au *AppUpdate) SetReleaseDate(s string) *AppUpdate {
	au.mutation.SetReleaseDate(s)
	return au
}

// SetDeveloper sets the "developer" field.
func (au *AppUpdate) SetDeveloper(s string) *AppUpdate {
	au.mutation.SetDeveloper(s)
	return au
}

// SetPublisher sets the "publisher" field.
func (au *AppUpdate) SetPublisher(s string) *AppUpdate {
	au.mutation.SetPublisher(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AppUpdate) SetCreatedAt(t time.Time) *AppUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AppUpdate) SetNillableCreatedAt(t *time.Time) *AppUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// Mutation returns the AppMutation object of the builder.
func (au *AppUpdate) Mutation() *AppMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AppUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AppUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AppUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AppUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AppUpdate) check() error {
	if v, ok := au.mutation.Source(); ok {
		if err := app.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "App.source": %w`, err)}
		}
	}
	if v, ok := au.mutation.GetType(); ok {
		if err := app.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "App.type": %w`, err)}
		}
	}
	return nil
}

func (au *AppUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.InternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: app.FieldInternalID,
		})
	}
	if value, ok := au.mutation.AddedInternalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: app.FieldInternalID,
		})
	}
	if value, ok := au.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: app.FieldSource,
		})
	}
	if value, ok := au.mutation.SourceAppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldSourceAppID,
		})
	}
	if value, ok := au.mutation.SourceURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldSourceURL,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldName,
		})
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: app.FieldType,
		})
	}
	if value, ok := au.mutation.ShortDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldShortDescription,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDescription,
		})
	}
	if value, ok := au.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldImageURL,
		})
	}
	if value, ok := au.mutation.ReleaseDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldReleaseDate,
		})
	}
	if value, ok := au.mutation.Developer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDeveloper,
		})
	}
	if value, ok := au.mutation.Publisher(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldPublisher,
		})
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: app.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppUpdateOne is the builder for updating a single App entity.
type AppUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppMutation
}

// SetInternalID sets the "internal_id" field.
func (auo *AppUpdateOne) SetInternalID(i int64) *AppUpdateOne {
	auo.mutation.ResetInternalID()
	auo.mutation.SetInternalID(i)
	return auo
}

// AddInternalID adds i to the "internal_id" field.
func (auo *AppUpdateOne) AddInternalID(i int64) *AppUpdateOne {
	auo.mutation.AddInternalID(i)
	return auo
}

// SetSource sets the "source" field.
func (auo *AppUpdateOne) SetSource(a app.Source) *AppUpdateOne {
	auo.mutation.SetSource(a)
	return auo
}

// SetSourceAppID sets the "source_app_id" field.
func (auo *AppUpdateOne) SetSourceAppID(s string) *AppUpdateOne {
	auo.mutation.SetSourceAppID(s)
	return auo
}

// SetSourceURL sets the "source_url" field.
func (auo *AppUpdateOne) SetSourceURL(s string) *AppUpdateOne {
	auo.mutation.SetSourceURL(s)
	return auo
}

// SetName sets the "name" field.
func (auo *AppUpdateOne) SetName(s string) *AppUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetType sets the "type" field.
func (auo *AppUpdateOne) SetType(a app.Type) *AppUpdateOne {
	auo.mutation.SetType(a)
	return auo
}

// SetShortDescription sets the "short_description" field.
func (auo *AppUpdateOne) SetShortDescription(s string) *AppUpdateOne {
	auo.mutation.SetShortDescription(s)
	return auo
}

// SetDescription sets the "description" field.
func (auo *AppUpdateOne) SetDescription(s string) *AppUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetImageURL sets the "image_url" field.
func (auo *AppUpdateOne) SetImageURL(s string) *AppUpdateOne {
	auo.mutation.SetImageURL(s)
	return auo
}

// SetReleaseDate sets the "release_date" field.
func (auo *AppUpdateOne) SetReleaseDate(s string) *AppUpdateOne {
	auo.mutation.SetReleaseDate(s)
	return auo
}

// SetDeveloper sets the "developer" field.
func (auo *AppUpdateOne) SetDeveloper(s string) *AppUpdateOne {
	auo.mutation.SetDeveloper(s)
	return auo
}

// SetPublisher sets the "publisher" field.
func (auo *AppUpdateOne) SetPublisher(s string) *AppUpdateOne {
	auo.mutation.SetPublisher(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AppUpdateOne) SetCreatedAt(t time.Time) *AppUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AppUpdateOne) SetNillableCreatedAt(t *time.Time) *AppUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// Mutation returns the AppMutation object of the builder.
func (auo *AppUpdateOne) Mutation() *AppMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AppUpdateOne) Select(field string, fields ...string) *AppUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated App entity.
func (auo *AppUpdateOne) Save(ctx context.Context) (*App, error) {
	var (
		err  error
		node *App
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*App)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AppUpdateOne) SaveX(ctx context.Context) *App {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AppUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AppUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AppUpdateOne) check() error {
	if v, ok := auo.mutation.Source(); ok {
		if err := app.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`ent: validator failed for field "App.source": %w`, err)}
		}
	}
	if v, ok := auo.mutation.GetType(); ok {
		if err := app.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "App.type": %w`, err)}
		}
	}
	return nil
}

func (auo *AppUpdateOne) sqlSave(ctx context.Context) (_node *App, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   app.Table,
			Columns: app.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: app.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "App.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, app.FieldID)
		for _, f := range fields {
			if !app.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != app.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.InternalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: app.FieldInternalID,
		})
	}
	if value, ok := auo.mutation.AddedInternalID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: app.FieldInternalID,
		})
	}
	if value, ok := auo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: app.FieldSource,
		})
	}
	if value, ok := auo.mutation.SourceAppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldSourceAppID,
		})
	}
	if value, ok := auo.mutation.SourceURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldSourceURL,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldName,
		})
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: app.FieldType,
		})
	}
	if value, ok := auo.mutation.ShortDescription(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldShortDescription,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDescription,
		})
	}
	if value, ok := auo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldImageURL,
		})
	}
	if value, ok := auo.mutation.ReleaseDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldReleaseDate,
		})
	}
	if value, ok := auo.mutation.Developer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDeveloper,
		})
	}
	if value, ok := auo.mutation.Publisher(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldPublisher,
		})
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: app.FieldCreatedAt,
		})
	}
	_node = &App{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{app.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}