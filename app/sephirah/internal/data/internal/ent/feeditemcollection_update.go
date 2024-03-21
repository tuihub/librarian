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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedItemCollectionUpdate is the builder for updating FeedItemCollection entities.
type FeedItemCollectionUpdate struct {
	config
	hooks    []Hook
	mutation *FeedItemCollectionMutation
}

// Where appends a list predicates to the FeedItemCollectionUpdate builder.
func (ficu *FeedItemCollectionUpdate) Where(ps ...predicate.FeedItemCollection) *FeedItemCollectionUpdate {
	ficu.mutation.Where(ps...)
	return ficu
}

// SetName sets the "name" field.
func (ficu *FeedItemCollectionUpdate) SetName(s string) *FeedItemCollectionUpdate {
	ficu.mutation.SetName(s)
	return ficu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ficu *FeedItemCollectionUpdate) SetNillableName(s *string) *FeedItemCollectionUpdate {
	if s != nil {
		ficu.SetName(*s)
	}
	return ficu
}

// SetDescription sets the "description" field.
func (ficu *FeedItemCollectionUpdate) SetDescription(s string) *FeedItemCollectionUpdate {
	ficu.mutation.SetDescription(s)
	return ficu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ficu *FeedItemCollectionUpdate) SetNillableDescription(s *string) *FeedItemCollectionUpdate {
	if s != nil {
		ficu.SetDescription(*s)
	}
	return ficu
}

// SetCategory sets the "category" field.
func (ficu *FeedItemCollectionUpdate) SetCategory(s string) *FeedItemCollectionUpdate {
	ficu.mutation.SetCategory(s)
	return ficu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (ficu *FeedItemCollectionUpdate) SetNillableCategory(s *string) *FeedItemCollectionUpdate {
	if s != nil {
		ficu.SetCategory(*s)
	}
	return ficu
}

// SetUpdatedAt sets the "updated_at" field.
func (ficu *FeedItemCollectionUpdate) SetUpdatedAt(t time.Time) *FeedItemCollectionUpdate {
	ficu.mutation.SetUpdatedAt(t)
	return ficu
}

// SetCreatedAt sets the "created_at" field.
func (ficu *FeedItemCollectionUpdate) SetCreatedAt(t time.Time) *FeedItemCollectionUpdate {
	ficu.mutation.SetCreatedAt(t)
	return ficu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ficu *FeedItemCollectionUpdate) SetNillableCreatedAt(t *time.Time) *FeedItemCollectionUpdate {
	if t != nil {
		ficu.SetCreatedAt(*t)
	}
	return ficu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ficu *FeedItemCollectionUpdate) SetOwnerID(id model.InternalID) *FeedItemCollectionUpdate {
	ficu.mutation.SetOwnerID(id)
	return ficu
}

// SetOwner sets the "owner" edge to the User entity.
func (ficu *FeedItemCollectionUpdate) SetOwner(u *User) *FeedItemCollectionUpdate {
	return ficu.SetOwnerID(u.ID)
}

// AddFeedItemIDs adds the "feed_item" edge to the FeedItem entity by IDs.
func (ficu *FeedItemCollectionUpdate) AddFeedItemIDs(ids ...model.InternalID) *FeedItemCollectionUpdate {
	ficu.mutation.AddFeedItemIDs(ids...)
	return ficu
}

// AddFeedItem adds the "feed_item" edges to the FeedItem entity.
func (ficu *FeedItemCollectionUpdate) AddFeedItem(f ...*FeedItem) *FeedItemCollectionUpdate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ficu.AddFeedItemIDs(ids...)
}

// AddNotifySourceIDs adds the "notify_source" edge to the NotifySource entity by IDs.
func (ficu *FeedItemCollectionUpdate) AddNotifySourceIDs(ids ...model.InternalID) *FeedItemCollectionUpdate {
	ficu.mutation.AddNotifySourceIDs(ids...)
	return ficu
}

// AddNotifySource adds the "notify_source" edges to the NotifySource entity.
func (ficu *FeedItemCollectionUpdate) AddNotifySource(n ...*NotifySource) *FeedItemCollectionUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ficu.AddNotifySourceIDs(ids...)
}

// Mutation returns the FeedItemCollectionMutation object of the builder.
func (ficu *FeedItemCollectionUpdate) Mutation() *FeedItemCollectionMutation {
	return ficu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ficu *FeedItemCollectionUpdate) ClearOwner() *FeedItemCollectionUpdate {
	ficu.mutation.ClearOwner()
	return ficu
}

// ClearFeedItem clears all "feed_item" edges to the FeedItem entity.
func (ficu *FeedItemCollectionUpdate) ClearFeedItem() *FeedItemCollectionUpdate {
	ficu.mutation.ClearFeedItem()
	return ficu
}

// RemoveFeedItemIDs removes the "feed_item" edge to FeedItem entities by IDs.
func (ficu *FeedItemCollectionUpdate) RemoveFeedItemIDs(ids ...model.InternalID) *FeedItemCollectionUpdate {
	ficu.mutation.RemoveFeedItemIDs(ids...)
	return ficu
}

// RemoveFeedItem removes "feed_item" edges to FeedItem entities.
func (ficu *FeedItemCollectionUpdate) RemoveFeedItem(f ...*FeedItem) *FeedItemCollectionUpdate {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ficu.RemoveFeedItemIDs(ids...)
}

// ClearNotifySource clears all "notify_source" edges to the NotifySource entity.
func (ficu *FeedItemCollectionUpdate) ClearNotifySource() *FeedItemCollectionUpdate {
	ficu.mutation.ClearNotifySource()
	return ficu
}

// RemoveNotifySourceIDs removes the "notify_source" edge to NotifySource entities by IDs.
func (ficu *FeedItemCollectionUpdate) RemoveNotifySourceIDs(ids ...model.InternalID) *FeedItemCollectionUpdate {
	ficu.mutation.RemoveNotifySourceIDs(ids...)
	return ficu
}

// RemoveNotifySource removes "notify_source" edges to NotifySource entities.
func (ficu *FeedItemCollectionUpdate) RemoveNotifySource(n ...*NotifySource) *FeedItemCollectionUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ficu.RemoveNotifySourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ficu *FeedItemCollectionUpdate) Save(ctx context.Context) (int, error) {
	ficu.defaults()
	return withHooks(ctx, ficu.sqlSave, ficu.mutation, ficu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ficu *FeedItemCollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := ficu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ficu *FeedItemCollectionUpdate) Exec(ctx context.Context) error {
	_, err := ficu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficu *FeedItemCollectionUpdate) ExecX(ctx context.Context) {
	if err := ficu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ficu *FeedItemCollectionUpdate) defaults() {
	if _, ok := ficu.mutation.UpdatedAt(); !ok {
		v := feeditemcollection.UpdateDefaultUpdatedAt()
		ficu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ficu *FeedItemCollectionUpdate) check() error {
	if _, ok := ficu.mutation.OwnerID(); ficu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedItemCollection.owner"`)
	}
	return nil
}

func (ficu *FeedItemCollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ficu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feeditemcollection.Table, feeditemcollection.Columns, sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64))
	if ps := ficu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ficu.mutation.Name(); ok {
		_spec.SetField(feeditemcollection.FieldName, field.TypeString, value)
	}
	if value, ok := ficu.mutation.Description(); ok {
		_spec.SetField(feeditemcollection.FieldDescription, field.TypeString, value)
	}
	if value, ok := ficu.mutation.Category(); ok {
		_spec.SetField(feeditemcollection.FieldCategory, field.TypeString, value)
	}
	if value, ok := ficu.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ficu.mutation.CreatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldCreatedAt, field.TypeTime, value)
	}
	if ficu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditemcollection.OwnerTable,
			Columns: []string{feeditemcollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditemcollection.OwnerTable,
			Columns: []string{feeditemcollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ficu.mutation.FeedItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficu.mutation.RemovedFeedItemIDs(); len(nodes) > 0 && !ficu.mutation.FeedItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
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
	if nodes := ficu.mutation.FeedItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
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
	if ficu.mutation.NotifySourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficu.mutation.RemovedNotifySourceIDs(); len(nodes) > 0 && !ficu.mutation.NotifySourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficu.mutation.NotifySourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ficu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeditemcollection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ficu.mutation.done = true
	return n, nil
}

// FeedItemCollectionUpdateOne is the builder for updating a single FeedItemCollection entity.
type FeedItemCollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedItemCollectionMutation
}

// SetName sets the "name" field.
func (ficuo *FeedItemCollectionUpdateOne) SetName(s string) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetName(s)
	return ficuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ficuo *FeedItemCollectionUpdateOne) SetNillableName(s *string) *FeedItemCollectionUpdateOne {
	if s != nil {
		ficuo.SetName(*s)
	}
	return ficuo
}

// SetDescription sets the "description" field.
func (ficuo *FeedItemCollectionUpdateOne) SetDescription(s string) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetDescription(s)
	return ficuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ficuo *FeedItemCollectionUpdateOne) SetNillableDescription(s *string) *FeedItemCollectionUpdateOne {
	if s != nil {
		ficuo.SetDescription(*s)
	}
	return ficuo
}

// SetCategory sets the "category" field.
func (ficuo *FeedItemCollectionUpdateOne) SetCategory(s string) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetCategory(s)
	return ficuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (ficuo *FeedItemCollectionUpdateOne) SetNillableCategory(s *string) *FeedItemCollectionUpdateOne {
	if s != nil {
		ficuo.SetCategory(*s)
	}
	return ficuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ficuo *FeedItemCollectionUpdateOne) SetUpdatedAt(t time.Time) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetUpdatedAt(t)
	return ficuo
}

// SetCreatedAt sets the "created_at" field.
func (ficuo *FeedItemCollectionUpdateOne) SetCreatedAt(t time.Time) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetCreatedAt(t)
	return ficuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ficuo *FeedItemCollectionUpdateOne) SetNillableCreatedAt(t *time.Time) *FeedItemCollectionUpdateOne {
	if t != nil {
		ficuo.SetCreatedAt(*t)
	}
	return ficuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ficuo *FeedItemCollectionUpdateOne) SetOwnerID(id model.InternalID) *FeedItemCollectionUpdateOne {
	ficuo.mutation.SetOwnerID(id)
	return ficuo
}

// SetOwner sets the "owner" edge to the User entity.
func (ficuo *FeedItemCollectionUpdateOne) SetOwner(u *User) *FeedItemCollectionUpdateOne {
	return ficuo.SetOwnerID(u.ID)
}

// AddFeedItemIDs adds the "feed_item" edge to the FeedItem entity by IDs.
func (ficuo *FeedItemCollectionUpdateOne) AddFeedItemIDs(ids ...model.InternalID) *FeedItemCollectionUpdateOne {
	ficuo.mutation.AddFeedItemIDs(ids...)
	return ficuo
}

// AddFeedItem adds the "feed_item" edges to the FeedItem entity.
func (ficuo *FeedItemCollectionUpdateOne) AddFeedItem(f ...*FeedItem) *FeedItemCollectionUpdateOne {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ficuo.AddFeedItemIDs(ids...)
}

// AddNotifySourceIDs adds the "notify_source" edge to the NotifySource entity by IDs.
func (ficuo *FeedItemCollectionUpdateOne) AddNotifySourceIDs(ids ...model.InternalID) *FeedItemCollectionUpdateOne {
	ficuo.mutation.AddNotifySourceIDs(ids...)
	return ficuo
}

// AddNotifySource adds the "notify_source" edges to the NotifySource entity.
func (ficuo *FeedItemCollectionUpdateOne) AddNotifySource(n ...*NotifySource) *FeedItemCollectionUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ficuo.AddNotifySourceIDs(ids...)
}

// Mutation returns the FeedItemCollectionMutation object of the builder.
func (ficuo *FeedItemCollectionUpdateOne) Mutation() *FeedItemCollectionMutation {
	return ficuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ficuo *FeedItemCollectionUpdateOne) ClearOwner() *FeedItemCollectionUpdateOne {
	ficuo.mutation.ClearOwner()
	return ficuo
}

// ClearFeedItem clears all "feed_item" edges to the FeedItem entity.
func (ficuo *FeedItemCollectionUpdateOne) ClearFeedItem() *FeedItemCollectionUpdateOne {
	ficuo.mutation.ClearFeedItem()
	return ficuo
}

// RemoveFeedItemIDs removes the "feed_item" edge to FeedItem entities by IDs.
func (ficuo *FeedItemCollectionUpdateOne) RemoveFeedItemIDs(ids ...model.InternalID) *FeedItemCollectionUpdateOne {
	ficuo.mutation.RemoveFeedItemIDs(ids...)
	return ficuo
}

// RemoveFeedItem removes "feed_item" edges to FeedItem entities.
func (ficuo *FeedItemCollectionUpdateOne) RemoveFeedItem(f ...*FeedItem) *FeedItemCollectionUpdateOne {
	ids := make([]model.InternalID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ficuo.RemoveFeedItemIDs(ids...)
}

// ClearNotifySource clears all "notify_source" edges to the NotifySource entity.
func (ficuo *FeedItemCollectionUpdateOne) ClearNotifySource() *FeedItemCollectionUpdateOne {
	ficuo.mutation.ClearNotifySource()
	return ficuo
}

// RemoveNotifySourceIDs removes the "notify_source" edge to NotifySource entities by IDs.
func (ficuo *FeedItemCollectionUpdateOne) RemoveNotifySourceIDs(ids ...model.InternalID) *FeedItemCollectionUpdateOne {
	ficuo.mutation.RemoveNotifySourceIDs(ids...)
	return ficuo
}

// RemoveNotifySource removes "notify_source" edges to NotifySource entities.
func (ficuo *FeedItemCollectionUpdateOne) RemoveNotifySource(n ...*NotifySource) *FeedItemCollectionUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ficuo.RemoveNotifySourceIDs(ids...)
}

// Where appends a list predicates to the FeedItemCollectionUpdate builder.
func (ficuo *FeedItemCollectionUpdateOne) Where(ps ...predicate.FeedItemCollection) *FeedItemCollectionUpdateOne {
	ficuo.mutation.Where(ps...)
	return ficuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ficuo *FeedItemCollectionUpdateOne) Select(field string, fields ...string) *FeedItemCollectionUpdateOne {
	ficuo.fields = append([]string{field}, fields...)
	return ficuo
}

// Save executes the query and returns the updated FeedItemCollection entity.
func (ficuo *FeedItemCollectionUpdateOne) Save(ctx context.Context) (*FeedItemCollection, error) {
	ficuo.defaults()
	return withHooks(ctx, ficuo.sqlSave, ficuo.mutation, ficuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ficuo *FeedItemCollectionUpdateOne) SaveX(ctx context.Context) *FeedItemCollection {
	node, err := ficuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ficuo *FeedItemCollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := ficuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ficuo *FeedItemCollectionUpdateOne) ExecX(ctx context.Context) {
	if err := ficuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ficuo *FeedItemCollectionUpdateOne) defaults() {
	if _, ok := ficuo.mutation.UpdatedAt(); !ok {
		v := feeditemcollection.UpdateDefaultUpdatedAt()
		ficuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ficuo *FeedItemCollectionUpdateOne) check() error {
	if _, ok := ficuo.mutation.OwnerID(); ficuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedItemCollection.owner"`)
	}
	return nil
}

func (ficuo *FeedItemCollectionUpdateOne) sqlSave(ctx context.Context) (_node *FeedItemCollection, err error) {
	if err := ficuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feeditemcollection.Table, feeditemcollection.Columns, sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64))
	id, ok := ficuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FeedItemCollection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ficuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feeditemcollection.FieldID)
		for _, f := range fields {
			if !feeditemcollection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feeditemcollection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ficuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ficuo.mutation.Name(); ok {
		_spec.SetField(feeditemcollection.FieldName, field.TypeString, value)
	}
	if value, ok := ficuo.mutation.Description(); ok {
		_spec.SetField(feeditemcollection.FieldDescription, field.TypeString, value)
	}
	if value, ok := ficuo.mutation.Category(); ok {
		_spec.SetField(feeditemcollection.FieldCategory, field.TypeString, value)
	}
	if value, ok := ficuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ficuo.mutation.CreatedAt(); ok {
		_spec.SetField(feeditemcollection.FieldCreatedAt, field.TypeTime, value)
	}
	if ficuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditemcollection.OwnerTable,
			Columns: []string{feeditemcollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   feeditemcollection.OwnerTable,
			Columns: []string{feeditemcollection.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ficuo.mutation.FeedItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficuo.mutation.RemovedFeedItemIDs(); len(nodes) > 0 && !ficuo.mutation.FeedItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
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
	if nodes := ficuo.mutation.FeedItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   feeditemcollection.FeedItemTable,
			Columns: feeditemcollection.FeedItemPrimaryKey,
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
	if ficuo.mutation.NotifySourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficuo.mutation.RemovedNotifySourceIDs(); len(nodes) > 0 && !ficuo.mutation.NotifySourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ficuo.mutation.NotifySourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feeditemcollection.NotifySourceTable,
			Columns: []string{feeditemcollection.NotifySourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FeedItemCollection{config: ficuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ficuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeditemcollection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ficuo.mutation.done = true
	return _node, nil
}
