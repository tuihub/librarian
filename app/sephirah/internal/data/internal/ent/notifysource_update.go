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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifysource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// NotifySourceUpdate is the builder for updating NotifySource entities.
type NotifySourceUpdate struct {
	config
	hooks    []Hook
	mutation *NotifySourceMutation
}

// Where appends a list predicates to the NotifySourceUpdate builder.
func (nsu *NotifySourceUpdate) Where(ps ...predicate.NotifySource) *NotifySourceUpdate {
	nsu.mutation.Where(ps...)
	return nsu
}

// SetFeedConfigID sets the "feed_config_id" field.
func (nsu *NotifySourceUpdate) SetFeedConfigID(mi model.InternalID) *NotifySourceUpdate {
	nsu.mutation.SetFeedConfigID(mi)
	return nsu
}

// SetNillableFeedConfigID sets the "feed_config_id" field if the given value is not nil.
func (nsu *NotifySourceUpdate) SetNillableFeedConfigID(mi *model.InternalID) *NotifySourceUpdate {
	if mi != nil {
		nsu.SetFeedConfigID(*mi)
	}
	return nsu
}

// ClearFeedConfigID clears the value of the "feed_config_id" field.
func (nsu *NotifySourceUpdate) ClearFeedConfigID() *NotifySourceUpdate {
	nsu.mutation.ClearFeedConfigID()
	return nsu
}

// SetFeedItemCollectionID sets the "feed_item_collection_id" field.
func (nsu *NotifySourceUpdate) SetFeedItemCollectionID(mi model.InternalID) *NotifySourceUpdate {
	nsu.mutation.SetFeedItemCollectionID(mi)
	return nsu
}

// SetNillableFeedItemCollectionID sets the "feed_item_collection_id" field if the given value is not nil.
func (nsu *NotifySourceUpdate) SetNillableFeedItemCollectionID(mi *model.InternalID) *NotifySourceUpdate {
	if mi != nil {
		nsu.SetFeedItemCollectionID(*mi)
	}
	return nsu
}

// ClearFeedItemCollectionID clears the value of the "feed_item_collection_id" field.
func (nsu *NotifySourceUpdate) ClearFeedItemCollectionID() *NotifySourceUpdate {
	nsu.mutation.ClearFeedItemCollectionID()
	return nsu
}

// SetUpdatedAt sets the "updated_at" field.
func (nsu *NotifySourceUpdate) SetUpdatedAt(t time.Time) *NotifySourceUpdate {
	nsu.mutation.SetUpdatedAt(t)
	return nsu
}

// SetCreatedAt sets the "created_at" field.
func (nsu *NotifySourceUpdate) SetCreatedAt(t time.Time) *NotifySourceUpdate {
	nsu.mutation.SetCreatedAt(t)
	return nsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nsu *NotifySourceUpdate) SetNillableCreatedAt(t *time.Time) *NotifySourceUpdate {
	if t != nil {
		nsu.SetCreatedAt(*t)
	}
	return nsu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (nsu *NotifySourceUpdate) SetOwnerID(id model.InternalID) *NotifySourceUpdate {
	nsu.mutation.SetOwnerID(id)
	return nsu
}

// SetOwner sets the "owner" edge to the User entity.
func (nsu *NotifySourceUpdate) SetOwner(u *User) *NotifySourceUpdate {
	return nsu.SetOwnerID(u.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (nsu *NotifySourceUpdate) AddNotifyFlowIDs(ids ...model.InternalID) *NotifySourceUpdate {
	nsu.mutation.AddNotifyFlowIDs(ids...)
	return nsu
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (nsu *NotifySourceUpdate) AddNotifyFlow(n ...*NotifyFlow) *NotifySourceUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsu.AddNotifyFlowIDs(ids...)
}

// SetFeedConfig sets the "feed_config" edge to the FeedConfig entity.
func (nsu *NotifySourceUpdate) SetFeedConfig(f *FeedConfig) *NotifySourceUpdate {
	return nsu.SetFeedConfigID(f.ID)
}

// SetFeedItemCollection sets the "feed_item_collection" edge to the FeedItemCollection entity.
func (nsu *NotifySourceUpdate) SetFeedItemCollection(f *FeedItemCollection) *NotifySourceUpdate {
	return nsu.SetFeedItemCollectionID(f.ID)
}

// AddNotifyFlowSourceIDs adds the "notify_flow_source" edge to the NotifyFlowSource entity by IDs.
func (nsu *NotifySourceUpdate) AddNotifyFlowSourceIDs(ids ...int) *NotifySourceUpdate {
	nsu.mutation.AddNotifyFlowSourceIDs(ids...)
	return nsu
}

// AddNotifyFlowSource adds the "notify_flow_source" edges to the NotifyFlowSource entity.
func (nsu *NotifySourceUpdate) AddNotifyFlowSource(n ...*NotifyFlowSource) *NotifySourceUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsu.AddNotifyFlowSourceIDs(ids...)
}

// Mutation returns the NotifySourceMutation object of the builder.
func (nsu *NotifySourceUpdate) Mutation() *NotifySourceMutation {
	return nsu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (nsu *NotifySourceUpdate) ClearOwner() *NotifySourceUpdate {
	nsu.mutation.ClearOwner()
	return nsu
}

// ClearNotifyFlow clears all "notify_flow" edges to the NotifyFlow entity.
func (nsu *NotifySourceUpdate) ClearNotifyFlow() *NotifySourceUpdate {
	nsu.mutation.ClearNotifyFlow()
	return nsu
}

// RemoveNotifyFlowIDs removes the "notify_flow" edge to NotifyFlow entities by IDs.
func (nsu *NotifySourceUpdate) RemoveNotifyFlowIDs(ids ...model.InternalID) *NotifySourceUpdate {
	nsu.mutation.RemoveNotifyFlowIDs(ids...)
	return nsu
}

// RemoveNotifyFlow removes "notify_flow" edges to NotifyFlow entities.
func (nsu *NotifySourceUpdate) RemoveNotifyFlow(n ...*NotifyFlow) *NotifySourceUpdate {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsu.RemoveNotifyFlowIDs(ids...)
}

// ClearFeedConfig clears the "feed_config" edge to the FeedConfig entity.
func (nsu *NotifySourceUpdate) ClearFeedConfig() *NotifySourceUpdate {
	nsu.mutation.ClearFeedConfig()
	return nsu
}

// ClearFeedItemCollection clears the "feed_item_collection" edge to the FeedItemCollection entity.
func (nsu *NotifySourceUpdate) ClearFeedItemCollection() *NotifySourceUpdate {
	nsu.mutation.ClearFeedItemCollection()
	return nsu
}

// ClearNotifyFlowSource clears all "notify_flow_source" edges to the NotifyFlowSource entity.
func (nsu *NotifySourceUpdate) ClearNotifyFlowSource() *NotifySourceUpdate {
	nsu.mutation.ClearNotifyFlowSource()
	return nsu
}

// RemoveNotifyFlowSourceIDs removes the "notify_flow_source" edge to NotifyFlowSource entities by IDs.
func (nsu *NotifySourceUpdate) RemoveNotifyFlowSourceIDs(ids ...int) *NotifySourceUpdate {
	nsu.mutation.RemoveNotifyFlowSourceIDs(ids...)
	return nsu
}

// RemoveNotifyFlowSource removes "notify_flow_source" edges to NotifyFlowSource entities.
func (nsu *NotifySourceUpdate) RemoveNotifyFlowSource(n ...*NotifyFlowSource) *NotifySourceUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsu.RemoveNotifyFlowSourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nsu *NotifySourceUpdate) Save(ctx context.Context) (int, error) {
	nsu.defaults()
	return withHooks(ctx, nsu.sqlSave, nsu.mutation, nsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsu *NotifySourceUpdate) SaveX(ctx context.Context) int {
	affected, err := nsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nsu *NotifySourceUpdate) Exec(ctx context.Context) error {
	_, err := nsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsu *NotifySourceUpdate) ExecX(ctx context.Context) {
	if err := nsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nsu *NotifySourceUpdate) defaults() {
	if _, ok := nsu.mutation.UpdatedAt(); !ok {
		v := notifysource.UpdateDefaultUpdatedAt()
		nsu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nsu *NotifySourceUpdate) check() error {
	if _, ok := nsu.mutation.OwnerID(); nsu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotifySource.owner"`)
	}
	return nil
}

func (nsu *NotifySourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifysource.Table, notifysource.Columns, sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64))
	if ps := nsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsu.mutation.UpdatedAt(); ok {
		_spec.SetField(notifysource.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nsu.mutation.CreatedAt(); ok {
		_spec.SetField(notifysource.FieldCreatedAt, field.TypeTime, value)
	}
	if nsu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.OwnerTable,
			Columns: []string{notifysource.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.OwnerTable,
			Columns: []string{notifysource.OwnerColumn},
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
	if nsu.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		createE := &NotifyFlowSourceCreate{config: nsu.config, mutation: newNotifyFlowSourceMutation(nsu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.RemovedNotifyFlowIDs(); len(nodes) > 0 && !nsu.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowSourceCreate{config: nsu.config, mutation: newNotifyFlowSourceMutation(nsu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowSourceCreate{config: nsu.config, mutation: newNotifyFlowSourceMutation(nsu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsu.mutation.FeedConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedConfigTable,
			Columns: []string{notifysource.FeedConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.FeedConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedConfigTable,
			Columns: []string{notifysource.FeedConfigColumn},
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
	if nsu.mutation.FeedItemCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedItemCollectionTable,
			Columns: []string{notifysource.FeedItemCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.FeedItemCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedItemCollectionTable,
			Columns: []string{notifysource.FeedItemCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsu.mutation.NotifyFlowSourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.RemovedNotifyFlowSourceIDs(); len(nodes) > 0 && !nsu.mutation.NotifyFlowSourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.NotifyFlowSourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifysource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nsu.mutation.done = true
	return n, nil
}

// NotifySourceUpdateOne is the builder for updating a single NotifySource entity.
type NotifySourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NotifySourceMutation
}

// SetFeedConfigID sets the "feed_config_id" field.
func (nsuo *NotifySourceUpdateOne) SetFeedConfigID(mi model.InternalID) *NotifySourceUpdateOne {
	nsuo.mutation.SetFeedConfigID(mi)
	return nsuo
}

// SetNillableFeedConfigID sets the "feed_config_id" field if the given value is not nil.
func (nsuo *NotifySourceUpdateOne) SetNillableFeedConfigID(mi *model.InternalID) *NotifySourceUpdateOne {
	if mi != nil {
		nsuo.SetFeedConfigID(*mi)
	}
	return nsuo
}

// ClearFeedConfigID clears the value of the "feed_config_id" field.
func (nsuo *NotifySourceUpdateOne) ClearFeedConfigID() *NotifySourceUpdateOne {
	nsuo.mutation.ClearFeedConfigID()
	return nsuo
}

// SetFeedItemCollectionID sets the "feed_item_collection_id" field.
func (nsuo *NotifySourceUpdateOne) SetFeedItemCollectionID(mi model.InternalID) *NotifySourceUpdateOne {
	nsuo.mutation.SetFeedItemCollectionID(mi)
	return nsuo
}

// SetNillableFeedItemCollectionID sets the "feed_item_collection_id" field if the given value is not nil.
func (nsuo *NotifySourceUpdateOne) SetNillableFeedItemCollectionID(mi *model.InternalID) *NotifySourceUpdateOne {
	if mi != nil {
		nsuo.SetFeedItemCollectionID(*mi)
	}
	return nsuo
}

// ClearFeedItemCollectionID clears the value of the "feed_item_collection_id" field.
func (nsuo *NotifySourceUpdateOne) ClearFeedItemCollectionID() *NotifySourceUpdateOne {
	nsuo.mutation.ClearFeedItemCollectionID()
	return nsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nsuo *NotifySourceUpdateOne) SetUpdatedAt(t time.Time) *NotifySourceUpdateOne {
	nsuo.mutation.SetUpdatedAt(t)
	return nsuo
}

// SetCreatedAt sets the "created_at" field.
func (nsuo *NotifySourceUpdateOne) SetCreatedAt(t time.Time) *NotifySourceUpdateOne {
	nsuo.mutation.SetCreatedAt(t)
	return nsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nsuo *NotifySourceUpdateOne) SetNillableCreatedAt(t *time.Time) *NotifySourceUpdateOne {
	if t != nil {
		nsuo.SetCreatedAt(*t)
	}
	return nsuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (nsuo *NotifySourceUpdateOne) SetOwnerID(id model.InternalID) *NotifySourceUpdateOne {
	nsuo.mutation.SetOwnerID(id)
	return nsuo
}

// SetOwner sets the "owner" edge to the User entity.
func (nsuo *NotifySourceUpdateOne) SetOwner(u *User) *NotifySourceUpdateOne {
	return nsuo.SetOwnerID(u.ID)
}

// AddNotifyFlowIDs adds the "notify_flow" edge to the NotifyFlow entity by IDs.
func (nsuo *NotifySourceUpdateOne) AddNotifyFlowIDs(ids ...model.InternalID) *NotifySourceUpdateOne {
	nsuo.mutation.AddNotifyFlowIDs(ids...)
	return nsuo
}

// AddNotifyFlow adds the "notify_flow" edges to the NotifyFlow entity.
func (nsuo *NotifySourceUpdateOne) AddNotifyFlow(n ...*NotifyFlow) *NotifySourceUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsuo.AddNotifyFlowIDs(ids...)
}

// SetFeedConfig sets the "feed_config" edge to the FeedConfig entity.
func (nsuo *NotifySourceUpdateOne) SetFeedConfig(f *FeedConfig) *NotifySourceUpdateOne {
	return nsuo.SetFeedConfigID(f.ID)
}

// SetFeedItemCollection sets the "feed_item_collection" edge to the FeedItemCollection entity.
func (nsuo *NotifySourceUpdateOne) SetFeedItemCollection(f *FeedItemCollection) *NotifySourceUpdateOne {
	return nsuo.SetFeedItemCollectionID(f.ID)
}

// AddNotifyFlowSourceIDs adds the "notify_flow_source" edge to the NotifyFlowSource entity by IDs.
func (nsuo *NotifySourceUpdateOne) AddNotifyFlowSourceIDs(ids ...int) *NotifySourceUpdateOne {
	nsuo.mutation.AddNotifyFlowSourceIDs(ids...)
	return nsuo
}

// AddNotifyFlowSource adds the "notify_flow_source" edges to the NotifyFlowSource entity.
func (nsuo *NotifySourceUpdateOne) AddNotifyFlowSource(n ...*NotifyFlowSource) *NotifySourceUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsuo.AddNotifyFlowSourceIDs(ids...)
}

// Mutation returns the NotifySourceMutation object of the builder.
func (nsuo *NotifySourceUpdateOne) Mutation() *NotifySourceMutation {
	return nsuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (nsuo *NotifySourceUpdateOne) ClearOwner() *NotifySourceUpdateOne {
	nsuo.mutation.ClearOwner()
	return nsuo
}

// ClearNotifyFlow clears all "notify_flow" edges to the NotifyFlow entity.
func (nsuo *NotifySourceUpdateOne) ClearNotifyFlow() *NotifySourceUpdateOne {
	nsuo.mutation.ClearNotifyFlow()
	return nsuo
}

// RemoveNotifyFlowIDs removes the "notify_flow" edge to NotifyFlow entities by IDs.
func (nsuo *NotifySourceUpdateOne) RemoveNotifyFlowIDs(ids ...model.InternalID) *NotifySourceUpdateOne {
	nsuo.mutation.RemoveNotifyFlowIDs(ids...)
	return nsuo
}

// RemoveNotifyFlow removes "notify_flow" edges to NotifyFlow entities.
func (nsuo *NotifySourceUpdateOne) RemoveNotifyFlow(n ...*NotifyFlow) *NotifySourceUpdateOne {
	ids := make([]model.InternalID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsuo.RemoveNotifyFlowIDs(ids...)
}

// ClearFeedConfig clears the "feed_config" edge to the FeedConfig entity.
func (nsuo *NotifySourceUpdateOne) ClearFeedConfig() *NotifySourceUpdateOne {
	nsuo.mutation.ClearFeedConfig()
	return nsuo
}

// ClearFeedItemCollection clears the "feed_item_collection" edge to the FeedItemCollection entity.
func (nsuo *NotifySourceUpdateOne) ClearFeedItemCollection() *NotifySourceUpdateOne {
	nsuo.mutation.ClearFeedItemCollection()
	return nsuo
}

// ClearNotifyFlowSource clears all "notify_flow_source" edges to the NotifyFlowSource entity.
func (nsuo *NotifySourceUpdateOne) ClearNotifyFlowSource() *NotifySourceUpdateOne {
	nsuo.mutation.ClearNotifyFlowSource()
	return nsuo
}

// RemoveNotifyFlowSourceIDs removes the "notify_flow_source" edge to NotifyFlowSource entities by IDs.
func (nsuo *NotifySourceUpdateOne) RemoveNotifyFlowSourceIDs(ids ...int) *NotifySourceUpdateOne {
	nsuo.mutation.RemoveNotifyFlowSourceIDs(ids...)
	return nsuo
}

// RemoveNotifyFlowSource removes "notify_flow_source" edges to NotifyFlowSource entities.
func (nsuo *NotifySourceUpdateOne) RemoveNotifyFlowSource(n ...*NotifyFlowSource) *NotifySourceUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nsuo.RemoveNotifyFlowSourceIDs(ids...)
}

// Where appends a list predicates to the NotifySourceUpdate builder.
func (nsuo *NotifySourceUpdateOne) Where(ps ...predicate.NotifySource) *NotifySourceUpdateOne {
	nsuo.mutation.Where(ps...)
	return nsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nsuo *NotifySourceUpdateOne) Select(field string, fields ...string) *NotifySourceUpdateOne {
	nsuo.fields = append([]string{field}, fields...)
	return nsuo
}

// Save executes the query and returns the updated NotifySource entity.
func (nsuo *NotifySourceUpdateOne) Save(ctx context.Context) (*NotifySource, error) {
	nsuo.defaults()
	return withHooks(ctx, nsuo.sqlSave, nsuo.mutation, nsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsuo *NotifySourceUpdateOne) SaveX(ctx context.Context) *NotifySource {
	node, err := nsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nsuo *NotifySourceUpdateOne) Exec(ctx context.Context) error {
	_, err := nsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsuo *NotifySourceUpdateOne) ExecX(ctx context.Context) {
	if err := nsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nsuo *NotifySourceUpdateOne) defaults() {
	if _, ok := nsuo.mutation.UpdatedAt(); !ok {
		v := notifysource.UpdateDefaultUpdatedAt()
		nsuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nsuo *NotifySourceUpdateOne) check() error {
	if _, ok := nsuo.mutation.OwnerID(); nsuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NotifySource.owner"`)
	}
	return nil
}

func (nsuo *NotifySourceUpdateOne) sqlSave(ctx context.Context) (_node *NotifySource, err error) {
	if err := nsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(notifysource.Table, notifysource.Columns, sqlgraph.NewFieldSpec(notifysource.FieldID, field.TypeInt64))
	id, ok := nsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NotifySource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifysource.FieldID)
		for _, f := range fields {
			if !notifysource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != notifysource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(notifysource.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nsuo.mutation.CreatedAt(); ok {
		_spec.SetField(notifysource.FieldCreatedAt, field.TypeTime, value)
	}
	if nsuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.OwnerTable,
			Columns: []string{notifysource.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.OwnerTable,
			Columns: []string{notifysource.OwnerColumn},
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
	if nsuo.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		createE := &NotifyFlowSourceCreate{config: nsuo.config, mutation: newNotifyFlowSourceMutation(nsuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.RemovedNotifyFlowIDs(); len(nodes) > 0 && !nsuo.mutation.NotifyFlowCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowSourceCreate{config: nsuo.config, mutation: newNotifyFlowSourceMutation(nsuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.NotifyFlowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowTable,
			Columns: notifysource.NotifyFlowPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &NotifyFlowSourceCreate{config: nsuo.config, mutation: newNotifyFlowSourceMutation(nsuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsuo.mutation.FeedConfigCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedConfigTable,
			Columns: []string{notifysource.FeedConfigColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.FeedConfigIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedConfigTable,
			Columns: []string{notifysource.FeedConfigColumn},
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
	if nsuo.mutation.FeedItemCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedItemCollectionTable,
			Columns: []string{notifysource.FeedItemCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.FeedItemCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   notifysource.FeedItemCollectionTable,
			Columns: []string{notifysource.FeedItemCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feeditemcollection.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nsuo.mutation.NotifyFlowSourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.RemovedNotifyFlowSourceIDs(); len(nodes) > 0 && !nsuo.mutation.NotifyFlowSourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.NotifyFlowSourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   notifysource.NotifyFlowSourceTable,
			Columns: []string{notifysource.NotifyFlowSourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notifyflowsource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NotifySource{config: nsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{notifysource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nsuo.mutation.done = true
	return _node, nil
}
