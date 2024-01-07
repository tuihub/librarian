// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/model"
	"github.com/tuihub/librarian/model/modelfeed"
)

// FeedUpdate is the builder for updating Feed entities.
type FeedUpdate struct {
	config
	hooks    []Hook
	mutation *FeedMutation
}

// Where appends a list predicates to the FeedUpdate builder.
func (fu *FeedUpdate) Where(ps ...predicate.Feed) *FeedUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetTitle sets the "title" field.
func (fu *FeedUpdate) SetTitle(s string) *FeedUpdate {
	fu.mutation.SetTitle(s)
	return fu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fu *FeedUpdate) SetNillableTitle(s *string) *FeedUpdate {
	if s != nil {
		fu.SetTitle(*s)
	}
	return fu
}

// ClearTitle clears the value of the "title" field.
func (fu *FeedUpdate) ClearTitle() *FeedUpdate {
	fu.mutation.ClearTitle()
	return fu
}

// SetLink sets the "link" field.
func (fu *FeedUpdate) SetLink(s string) *FeedUpdate {
	fu.mutation.SetLink(s)
	return fu
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (fu *FeedUpdate) SetNillableLink(s *string) *FeedUpdate {
	if s != nil {
		fu.SetLink(*s)
	}
	return fu
}

// ClearLink clears the value of the "link" field.
func (fu *FeedUpdate) ClearLink() *FeedUpdate {
	fu.mutation.ClearLink()
	return fu
}

// SetDescription sets the "description" field.
func (fu *FeedUpdate) SetDescription(s string) *FeedUpdate {
	fu.mutation.SetDescription(s)
	return fu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fu *FeedUpdate) SetNillableDescription(s *string) *FeedUpdate {
	if s != nil {
		fu.SetDescription(*s)
	}
	return fu
}

// ClearDescription clears the value of the "description" field.
func (fu *FeedUpdate) ClearDescription() *FeedUpdate {
	fu.mutation.ClearDescription()
	return fu
}

// SetLanguage sets the "language" field.
func (fu *FeedUpdate) SetLanguage(s string) *FeedUpdate {
	fu.mutation.SetLanguage(s)
	return fu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (fu *FeedUpdate) SetNillableLanguage(s *string) *FeedUpdate {
	if s != nil {
		fu.SetLanguage(*s)
	}
	return fu
}

// ClearLanguage clears the value of the "language" field.
func (fu *FeedUpdate) ClearLanguage() *FeedUpdate {
	fu.mutation.ClearLanguage()
	return fu
}

// SetAuthors sets the "authors" field.
func (fu *FeedUpdate) SetAuthors(m []*modelfeed.Person) *FeedUpdate {
	fu.mutation.SetAuthors(m)
	return fu
}

// AppendAuthors appends m to the "authors" field.
func (fu *FeedUpdate) AppendAuthors(m []*modelfeed.Person) *FeedUpdate {
	fu.mutation.AppendAuthors(m)
	return fu
}

// ClearAuthors clears the value of the "authors" field.
func (fu *FeedUpdate) ClearAuthors() *FeedUpdate {
	fu.mutation.ClearAuthors()
	return fu
}

// SetImage sets the "image" field.
func (fu *FeedUpdate) SetImage(m *modelfeed.Image) *FeedUpdate {
	fu.mutation.SetImage(m)
	return fu
}

// ClearImage clears the value of the "image" field.
func (fu *FeedUpdate) ClearImage() *FeedUpdate {
	fu.mutation.ClearImage()
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FeedUpdate) SetUpdatedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FeedUpdate) SetCreatedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FeedUpdate) SetNillableCreatedAt(t *time.Time) *FeedUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// AddItemIDs adds the "item" edge to the FeedItem entity by IDs.
func (fu *FeedUpdate) AddItemIDs(ids ...model.InternalID) *FeedUpdate {
	fu.mutation.AddItemIDs(ids...)
	return fu
}

// AddItem adds the "item" edges to the FeedItem entity.
func (fu *FeedUpdate) AddItem(f ...*FeedItem) *FeedUpdate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddItemIDs(ids...)
}

// SetConfigID sets the "config" edge to the FeedConfig entity by ID.
func (fu *FeedUpdate) SetConfigID(id model.InternalID) *FeedUpdate {
	fu.mutation.SetConfigID(id)
	return fu
}

// SetConfig sets the "config" edge to the FeedConfig entity.
func (fu *FeedUpdate) SetConfig(f *FeedConfig) *FeedUpdate {
	return fu.SetConfigID(f.ID)
}

// Mutation returns the FeedMutation object of the builder.
func (fu *FeedUpdate) Mutation() *FeedMutation {
	return fu.mutation
}

// ClearItem clears all "item" edges to the FeedItem entity.
func (fu *FeedUpdate) ClearItem() *FeedUpdate {
	fu.mutation.ClearItem()
	return fu
}

// RemoveItemIDs removes the "item" edge to FeedItem entities by IDs.
func (fu *FeedUpdate) RemoveItemIDs(ids ...model.InternalID) *FeedUpdate {
	fu.mutation.RemoveItemIDs(ids...)
	return fu
}

// RemoveItem removes "item" edges to FeedItem entities.
func (fu *FeedUpdate) RemoveItem(f ...*FeedItem) *FeedUpdate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveItemIDs(ids...)
}

// ClearConfig clears the "config" edge to the FeedConfig entity.
func (fu *FeedUpdate) ClearConfig() *FeedUpdate {
	fu.mutation.ClearConfig()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FeedUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeedUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeedUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeedUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FeedUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := feed.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FeedUpdate) check() error {
	if _, ok := fu.mutation.ConfigID(); fu.mutation.ConfigCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Feed.config"`)
	}
	return nil
}

func (fu *FeedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feed.Table, feed.Columns, sqlgraph.NewFieldSpec(feed.FieldID, field.TypeInt64))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Title(); ok {
		_spec.SetField(feed.FieldTitle, field.TypeString, value)
	}
	if fu.mutation.TitleCleared() {
		_spec.ClearField(feed.FieldTitle, field.TypeString)
	}
	if value, ok := fu.mutation.Link(); ok {
		_spec.SetField(feed.FieldLink, field.TypeString, value)
	}
	if fu.mutation.LinkCleared() {
		_spec.ClearField(feed.FieldLink, field.TypeString)
	}
	if value, ok := fu.mutation.Description(); ok {
		_spec.SetField(feed.FieldDescription, field.TypeString, value)
	}
	if fu.mutation.DescriptionCleared() {
		_spec.ClearField(feed.FieldDescription, field.TypeString)
	}
	if value, ok := fu.mutation.Language(); ok {
		_spec.SetField(feed.FieldLanguage, field.TypeString, value)
	}
	if fu.mutation.LanguageCleared() {
		_spec.ClearField(feed.FieldLanguage, field.TypeString)
	}
	if value, ok := fu.mutation.Authors(); ok {
		_spec.SetField(feed.FieldAuthors, field.TypeJSON, value)
	}
	if value, ok := fu.mutation.AppendedAuthors(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, feed.FieldAuthors, value)
		})
	}
	if fu.mutation.AuthorsCleared() {
		_spec.ClearField(feed.FieldAuthors, field.TypeJSON)
	}
	if value, ok := fu.mutation.Image(); ok {
		_spec.SetField(feed.FieldImage, field.TypeJSON, value)
	}
	if fu.mutation.ImageCleared() {
		_spec.ClearField(feed.FieldImage, field.TypeJSON)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(feed.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.SetField(feed.FieldCreatedAt, field.TypeTime, value)
	}
	if fu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedItemIDs(); len(nodes) > 0 && !fu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.ConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   feed.ConfigTable,
			Columns: []string{feed.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   feed.ConfigTable,
			Columns: []string{feed.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FeedUpdateOne is the builder for updating a single Feed entity.
type FeedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedMutation
}

// SetTitle sets the "title" field.
func (fuo *FeedUpdateOne) SetTitle(s string) *FeedUpdateOne {
	fuo.mutation.SetTitle(s)
	return fuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableTitle(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetTitle(*s)
	}
	return fuo
}

// ClearTitle clears the value of the "title" field.
func (fuo *FeedUpdateOne) ClearTitle() *FeedUpdateOne {
	fuo.mutation.ClearTitle()
	return fuo
}

// SetLink sets the "link" field.
func (fuo *FeedUpdateOne) SetLink(s string) *FeedUpdateOne {
	fuo.mutation.SetLink(s)
	return fuo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableLink(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetLink(*s)
	}
	return fuo
}

// ClearLink clears the value of the "link" field.
func (fuo *FeedUpdateOne) ClearLink() *FeedUpdateOne {
	fuo.mutation.ClearLink()
	return fuo
}

// SetDescription sets the "description" field.
func (fuo *FeedUpdateOne) SetDescription(s string) *FeedUpdateOne {
	fuo.mutation.SetDescription(s)
	return fuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableDescription(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetDescription(*s)
	}
	return fuo
}

// ClearDescription clears the value of the "description" field.
func (fuo *FeedUpdateOne) ClearDescription() *FeedUpdateOne {
	fuo.mutation.ClearDescription()
	return fuo
}

// SetLanguage sets the "language" field.
func (fuo *FeedUpdateOne) SetLanguage(s string) *FeedUpdateOne {
	fuo.mutation.SetLanguage(s)
	return fuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableLanguage(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetLanguage(*s)
	}
	return fuo
}

// ClearLanguage clears the value of the "language" field.
func (fuo *FeedUpdateOne) ClearLanguage() *FeedUpdateOne {
	fuo.mutation.ClearLanguage()
	return fuo
}

// SetAuthors sets the "authors" field.
func (fuo *FeedUpdateOne) SetAuthors(m []*modelfeed.Person) *FeedUpdateOne {
	fuo.mutation.SetAuthors(m)
	return fuo
}

// AppendAuthors appends m to the "authors" field.
func (fuo *FeedUpdateOne) AppendAuthors(m []*modelfeed.Person) *FeedUpdateOne {
	fuo.mutation.AppendAuthors(m)
	return fuo
}

// ClearAuthors clears the value of the "authors" field.
func (fuo *FeedUpdateOne) ClearAuthors() *FeedUpdateOne {
	fuo.mutation.ClearAuthors()
	return fuo
}

// SetImage sets the "image" field.
func (fuo *FeedUpdateOne) SetImage(m *modelfeed.Image) *FeedUpdateOne {
	fuo.mutation.SetImage(m)
	return fuo
}

// ClearImage clears the value of the "image" field.
func (fuo *FeedUpdateOne) ClearImage() *FeedUpdateOne {
	fuo.mutation.ClearImage()
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FeedUpdateOne) SetUpdatedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FeedUpdateOne) SetCreatedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableCreatedAt(t *time.Time) *FeedUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// AddItemIDs adds the "item" edge to the FeedItem entity by IDs.
func (fuo *FeedUpdateOne) AddItemIDs(ids ...model.InternalID) *FeedUpdateOne {
	fuo.mutation.AddItemIDs(ids...)
	return fuo
}

// AddItem adds the "item" edges to the FeedItem entity.
func (fuo *FeedUpdateOne) AddItem(f ...*FeedItem) *FeedUpdateOne {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddItemIDs(ids...)
}

// SetConfigID sets the "config" edge to the FeedConfig entity by ID.
func (fuo *FeedUpdateOne) SetConfigID(id model.InternalID) *FeedUpdateOne {
	fuo.mutation.SetConfigID(id)
	return fuo
}

// SetConfig sets the "config" edge to the FeedConfig entity.
func (fuo *FeedUpdateOne) SetConfig(f *FeedConfig) *FeedUpdateOne {
	return fuo.SetConfigID(f.ID)
}

// Mutation returns the FeedMutation object of the builder.
func (fuo *FeedUpdateOne) Mutation() *FeedMutation {
	return fuo.mutation
}

// ClearItem clears all "item" edges to the FeedItem entity.
func (fuo *FeedUpdateOne) ClearItem() *FeedUpdateOne {
	fuo.mutation.ClearItem()
	return fuo
}

// RemoveItemIDs removes the "item" edge to FeedItem entities by IDs.
func (fuo *FeedUpdateOne) RemoveItemIDs(ids ...model.InternalID) *FeedUpdateOne {
	fuo.mutation.RemoveItemIDs(ids...)
	return fuo
}

// RemoveItem removes "item" edges to FeedItem entities.
func (fuo *FeedUpdateOne) RemoveItem(f ...*FeedItem) *FeedUpdateOne {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveItemIDs(ids...)
}

// ClearConfig clears the "config" edge to the FeedConfig entity.
func (fuo *FeedUpdateOne) ClearConfig() *FeedUpdateOne {
	fuo.mutation.ClearConfig()
	return fuo
}

// Where appends a list predicates to the FeedUpdate builder.
func (fuo *FeedUpdateOne) Where(ps ...predicate.Feed) *FeedUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FeedUpdateOne) Select(field string, fields ...string) *FeedUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Feed entity.
func (fuo *FeedUpdateOne) Save(ctx context.Context) (*Feed, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeedUpdateOne) SaveX(ctx context.Context) *Feed {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FeedUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeedUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FeedUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := feed.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FeedUpdateOne) check() error {
	if _, ok := fuo.mutation.ConfigID(); fuo.mutation.ConfigCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Feed.config"`)
	}
	return nil
}

func (fuo *FeedUpdateOne) sqlSave(ctx context.Context) (_node *Feed, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feed.Table, feed.Columns, sqlgraph.NewFieldSpec(feed.FieldID, field.TypeInt64))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Feed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feed.FieldID)
		for _, f := range fields {
			if !feed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Title(); ok {
		_spec.SetField(feed.FieldTitle, field.TypeString, value)
	}
	if fuo.mutation.TitleCleared() {
		_spec.ClearField(feed.FieldTitle, field.TypeString)
	}
	if value, ok := fuo.mutation.Link(); ok {
		_spec.SetField(feed.FieldLink, field.TypeString, value)
	}
	if fuo.mutation.LinkCleared() {
		_spec.ClearField(feed.FieldLink, field.TypeString)
	}
	if value, ok := fuo.mutation.Description(); ok {
		_spec.SetField(feed.FieldDescription, field.TypeString, value)
	}
	if fuo.mutation.DescriptionCleared() {
		_spec.ClearField(feed.FieldDescription, field.TypeString)
	}
	if value, ok := fuo.mutation.Language(); ok {
		_spec.SetField(feed.FieldLanguage, field.TypeString, value)
	}
	if fuo.mutation.LanguageCleared() {
		_spec.ClearField(feed.FieldLanguage, field.TypeString)
	}
	if value, ok := fuo.mutation.Authors(); ok {
		_spec.SetField(feed.FieldAuthors, field.TypeJSON, value)
	}
	if value, ok := fuo.mutation.AppendedAuthors(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, feed.FieldAuthors, value)
		})
	}
	if fuo.mutation.AuthorsCleared() {
		_spec.ClearField(feed.FieldAuthors, field.TypeJSON)
	}
	if value, ok := fuo.mutation.Image(); ok {
		_spec.SetField(feed.FieldImage, field.TypeJSON, value)
	}
	if fuo.mutation.ImageCleared() {
		_spec.ClearField(feed.FieldImage, field.TypeJSON)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feed.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.SetField(feed.FieldCreatedAt, field.TypeTime, value)
	}
	if fuo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedItemIDs(); len(nodes) > 0 && !fuo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.ItemTable,
			Columns: []string{feed.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.ConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   feed.ConfigTable,
			Columns: []string{feed.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   feed.ConfigTable,
			Columns: []string{feed.ConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Feed{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
