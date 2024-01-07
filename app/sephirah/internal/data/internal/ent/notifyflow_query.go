// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowtarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifytarget"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/model"
)

// NotifyFlowQuery is the builder for querying NotifyFlow entities.
type NotifyFlowQuery struct {
	config
	ctx                  *QueryContext
	order                []notifyflow.OrderOption
	inters               []Interceptor
	predicates           []predicate.NotifyFlow
	withOwner            *UserQuery
	withNotifyTarget     *NotifyTargetQuery
	withFeedConfig       *FeedConfigQuery
	withNotifyFlowTarget *NotifyFlowTargetQuery
	withNotifyFlowSource *NotifyFlowSourceQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotifyFlowQuery builder.
func (nfq *NotifyFlowQuery) Where(ps ...predicate.NotifyFlow) *NotifyFlowQuery {
	nfq.predicates = append(nfq.predicates, ps...)
	return nfq
}

// Limit the number of records to be returned by this query.
func (nfq *NotifyFlowQuery) Limit(limit int) *NotifyFlowQuery {
	nfq.ctx.Limit = &limit
	return nfq
}

// Offset to start from.
func (nfq *NotifyFlowQuery) Offset(offset int) *NotifyFlowQuery {
	nfq.ctx.Offset = &offset
	return nfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nfq *NotifyFlowQuery) Unique(unique bool) *NotifyFlowQuery {
	nfq.ctx.Unique = &unique
	return nfq
}

// Order specifies how the records should be ordered.
func (nfq *NotifyFlowQuery) Order(o ...notifyflow.OrderOption) *NotifyFlowQuery {
	nfq.order = append(nfq.order, o...)
	return nfq
}

// QueryOwner chains the current query on the "owner" edge.
func (nfq *NotifyFlowQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: nfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifyflow.Table, notifyflow.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notifyflow.OwnerTable, notifyflow.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(nfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyTarget chains the current query on the "notify_target" edge.
func (nfq *NotifyFlowQuery) QueryNotifyTarget() *NotifyTargetQuery {
	query := (&NotifyTargetClient{config: nfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifyflow.Table, notifyflow.FieldID, selector),
			sqlgraph.To(notifytarget.Table, notifytarget.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, notifyflow.NotifyTargetTable, notifyflow.NotifyTargetPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeedConfig chains the current query on the "feed_config" edge.
func (nfq *NotifyFlowQuery) QueryFeedConfig() *FeedConfigQuery {
	query := (&FeedConfigClient{config: nfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifyflow.Table, notifyflow.FieldID, selector),
			sqlgraph.To(feedconfig.Table, feedconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, notifyflow.FeedConfigTable, notifyflow.FeedConfigPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyFlowTarget chains the current query on the "notify_flow_target" edge.
func (nfq *NotifyFlowQuery) QueryNotifyFlowTarget() *NotifyFlowTargetQuery {
	query := (&NotifyFlowTargetClient{config: nfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifyflow.Table, notifyflow.FieldID, selector),
			sqlgraph.To(notifyflowtarget.Table, notifyflowtarget.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, notifyflow.NotifyFlowTargetTable, notifyflow.NotifyFlowTargetColumn),
		)
		fromU = sqlgraph.SetNeighbors(nfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyFlowSource chains the current query on the "notify_flow_source" edge.
func (nfq *NotifyFlowQuery) QueryNotifyFlowSource() *NotifyFlowSourceQuery {
	query := (&NotifyFlowSourceClient{config: nfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifyflow.Table, notifyflow.FieldID, selector),
			sqlgraph.To(notifyflowsource.Table, notifyflowsource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, notifyflow.NotifyFlowSourceTable, notifyflow.NotifyFlowSourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(nfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NotifyFlow entity from the query.
// Returns a *NotFoundError when no NotifyFlow was found.
func (nfq *NotifyFlowQuery) First(ctx context.Context) (*NotifyFlow, error) {
	nodes, err := nfq.Limit(1).All(setContextOp(ctx, nfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notifyflow.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nfq *NotifyFlowQuery) FirstX(ctx context.Context) *NotifyFlow {
	node, err := nfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotifyFlow ID from the query.
// Returns a *NotFoundError when no NotifyFlow ID was found.
func (nfq *NotifyFlowQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = nfq.Limit(1).IDs(setContextOp(ctx, nfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notifyflow.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nfq *NotifyFlowQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := nfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotifyFlow entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotifyFlow entity is found.
// Returns a *NotFoundError when no NotifyFlow entities are found.
func (nfq *NotifyFlowQuery) Only(ctx context.Context) (*NotifyFlow, error) {
	nodes, err := nfq.Limit(2).All(setContextOp(ctx, nfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notifyflow.Label}
	default:
		return nil, &NotSingularError{notifyflow.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nfq *NotifyFlowQuery) OnlyX(ctx context.Context) *NotifyFlow {
	node, err := nfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotifyFlow ID in the query.
// Returns a *NotSingularError when more than one NotifyFlow ID is found.
// Returns a *NotFoundError when no entities are found.
func (nfq *NotifyFlowQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = nfq.Limit(2).IDs(setContextOp(ctx, nfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notifyflow.Label}
	default:
		err = &NotSingularError{notifyflow.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nfq *NotifyFlowQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := nfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotifyFlows.
func (nfq *NotifyFlowQuery) All(ctx context.Context) ([]*NotifyFlow, error) {
	ctx = setContextOp(ctx, nfq.ctx, "All")
	if err := nfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NotifyFlow, *NotifyFlowQuery]()
	return withInterceptors[[]*NotifyFlow](ctx, nfq, qr, nfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nfq *NotifyFlowQuery) AllX(ctx context.Context) []*NotifyFlow {
	nodes, err := nfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotifyFlow IDs.
func (nfq *NotifyFlowQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if nfq.ctx.Unique == nil && nfq.path != nil {
		nfq.Unique(true)
	}
	ctx = setContextOp(ctx, nfq.ctx, "IDs")
	if err = nfq.Select(notifyflow.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nfq *NotifyFlowQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := nfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nfq *NotifyFlowQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nfq.ctx, "Count")
	if err := nfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nfq, querierCount[*NotifyFlowQuery](), nfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nfq *NotifyFlowQuery) CountX(ctx context.Context) int {
	count, err := nfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nfq *NotifyFlowQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nfq.ctx, "Exist")
	switch _, err := nfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nfq *NotifyFlowQuery) ExistX(ctx context.Context) bool {
	exist, err := nfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotifyFlowQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nfq *NotifyFlowQuery) Clone() *NotifyFlowQuery {
	if nfq == nil {
		return nil
	}
	return &NotifyFlowQuery{
		config:               nfq.config,
		ctx:                  nfq.ctx.Clone(),
		order:                append([]notifyflow.OrderOption{}, nfq.order...),
		inters:               append([]Interceptor{}, nfq.inters...),
		predicates:           append([]predicate.NotifyFlow{}, nfq.predicates...),
		withOwner:            nfq.withOwner.Clone(),
		withNotifyTarget:     nfq.withNotifyTarget.Clone(),
		withFeedConfig:       nfq.withFeedConfig.Clone(),
		withNotifyFlowTarget: nfq.withNotifyFlowTarget.Clone(),
		withNotifyFlowSource: nfq.withNotifyFlowSource.Clone(),
		// clone intermediate query.
		sql:  nfq.sql.Clone(),
		path: nfq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (nfq *NotifyFlowQuery) WithOwner(opts ...func(*UserQuery)) *NotifyFlowQuery {
	query := (&UserClient{config: nfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nfq.withOwner = query
	return nfq
}

// WithNotifyTarget tells the query-builder to eager-load the nodes that are connected to
// the "notify_target" edge. The optional arguments are used to configure the query builder of the edge.
func (nfq *NotifyFlowQuery) WithNotifyTarget(opts ...func(*NotifyTargetQuery)) *NotifyFlowQuery {
	query := (&NotifyTargetClient{config: nfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nfq.withNotifyTarget = query
	return nfq
}

// WithFeedConfig tells the query-builder to eager-load the nodes that are connected to
// the "feed_config" edge. The optional arguments are used to configure the query builder of the edge.
func (nfq *NotifyFlowQuery) WithFeedConfig(opts ...func(*FeedConfigQuery)) *NotifyFlowQuery {
	query := (&FeedConfigClient{config: nfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nfq.withFeedConfig = query
	return nfq
}

// WithNotifyFlowTarget tells the query-builder to eager-load the nodes that are connected to
// the "notify_flow_target" edge. The optional arguments are used to configure the query builder of the edge.
func (nfq *NotifyFlowQuery) WithNotifyFlowTarget(opts ...func(*NotifyFlowTargetQuery)) *NotifyFlowQuery {
	query := (&NotifyFlowTargetClient{config: nfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nfq.withNotifyFlowTarget = query
	return nfq
}

// WithNotifyFlowSource tells the query-builder to eager-load the nodes that are connected to
// the "notify_flow_source" edge. The optional arguments are used to configure the query builder of the edge.
func (nfq *NotifyFlowQuery) WithNotifyFlowSource(opts ...func(*NotifyFlowSourceQuery)) *NotifyFlowQuery {
	query := (&NotifyFlowSourceClient{config: nfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nfq.withNotifyFlowSource = query
	return nfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NotifyFlow.Query().
//		GroupBy(notifyflow.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nfq *NotifyFlowQuery) GroupBy(field string, fields ...string) *NotifyFlowGroupBy {
	nfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotifyFlowGroupBy{build: nfq}
	grbuild.flds = &nfq.ctx.Fields
	grbuild.label = notifyflow.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.NotifyFlow.Query().
//		Select(notifyflow.FieldName).
//		Scan(ctx, &v)
func (nfq *NotifyFlowQuery) Select(fields ...string) *NotifyFlowSelect {
	nfq.ctx.Fields = append(nfq.ctx.Fields, fields...)
	sbuild := &NotifyFlowSelect{NotifyFlowQuery: nfq}
	sbuild.label = notifyflow.Label
	sbuild.flds, sbuild.scan = &nfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotifyFlowSelect configured with the given aggregations.
func (nfq *NotifyFlowQuery) Aggregate(fns ...AggregateFunc) *NotifyFlowSelect {
	return nfq.Select().Aggregate(fns...)
}

func (nfq *NotifyFlowQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nfq); err != nil {
				return err
			}
		}
	}
	for _, f := range nfq.ctx.Fields {
		if !notifyflow.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nfq.path != nil {
		prev, err := nfq.path(ctx)
		if err != nil {
			return err
		}
		nfq.sql = prev
	}
	return nil
}

func (nfq *NotifyFlowQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotifyFlow, error) {
	var (
		nodes       = []*NotifyFlow{}
		withFKs     = nfq.withFKs
		_spec       = nfq.querySpec()
		loadedTypes = [5]bool{
			nfq.withOwner != nil,
			nfq.withNotifyTarget != nil,
			nfq.withFeedConfig != nil,
			nfq.withNotifyFlowTarget != nil,
			nfq.withNotifyFlowSource != nil,
		}
	)
	if nfq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, notifyflow.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NotifyFlow).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NotifyFlow{config: nfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nfq.withOwner; query != nil {
		if err := nfq.loadOwner(ctx, query, nodes, nil,
			func(n *NotifyFlow, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := nfq.withNotifyTarget; query != nil {
		if err := nfq.loadNotifyTarget(ctx, query, nodes,
			func(n *NotifyFlow) { n.Edges.NotifyTarget = []*NotifyTarget{} },
			func(n *NotifyFlow, e *NotifyTarget) { n.Edges.NotifyTarget = append(n.Edges.NotifyTarget, e) }); err != nil {
			return nil, err
		}
	}
	if query := nfq.withFeedConfig; query != nil {
		if err := nfq.loadFeedConfig(ctx, query, nodes,
			func(n *NotifyFlow) { n.Edges.FeedConfig = []*FeedConfig{} },
			func(n *NotifyFlow, e *FeedConfig) { n.Edges.FeedConfig = append(n.Edges.FeedConfig, e) }); err != nil {
			return nil, err
		}
	}
	if query := nfq.withNotifyFlowTarget; query != nil {
		if err := nfq.loadNotifyFlowTarget(ctx, query, nodes,
			func(n *NotifyFlow) { n.Edges.NotifyFlowTarget = []*NotifyFlowTarget{} },
			func(n *NotifyFlow, e *NotifyFlowTarget) {
				n.Edges.NotifyFlowTarget = append(n.Edges.NotifyFlowTarget, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := nfq.withNotifyFlowSource; query != nil {
		if err := nfq.loadNotifyFlowSource(ctx, query, nodes,
			func(n *NotifyFlow) { n.Edges.NotifyFlowSource = []*NotifyFlowSource{} },
			func(n *NotifyFlow, e *NotifyFlowSource) {
				n.Edges.NotifyFlowSource = append(n.Edges.NotifyFlowSource, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nfq *NotifyFlowQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*NotifyFlow, init func(*NotifyFlow), assign func(*NotifyFlow, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*NotifyFlow)
	for i := range nodes {
		if nodes[i].user_notify_flow == nil {
			continue
		}
		fk := *nodes[i].user_notify_flow
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_notify_flow" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nfq *NotifyFlowQuery) loadNotifyTarget(ctx context.Context, query *NotifyTargetQuery, nodes []*NotifyFlow, init func(*NotifyFlow), assign func(*NotifyFlow, *NotifyTarget)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*NotifyFlow)
	nids := make(map[model.InternalID]map[*NotifyFlow]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notifyflow.NotifyTargetTable)
		s.Join(joinT).On(s.C(notifytarget.FieldID), joinT.C(notifyflow.NotifyTargetPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(notifyflow.NotifyTargetPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notifyflow.NotifyTargetPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := model.InternalID(values[0].(*sql.NullInt64).Int64)
				inValue := model.InternalID(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*NotifyFlow]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*NotifyTarget](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "notify_target" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nfq *NotifyFlowQuery) loadFeedConfig(ctx context.Context, query *FeedConfigQuery, nodes []*NotifyFlow, init func(*NotifyFlow), assign func(*NotifyFlow, *FeedConfig)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*NotifyFlow)
	nids := make(map[model.InternalID]map[*NotifyFlow]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notifyflow.FeedConfigTable)
		s.Join(joinT).On(s.C(feedconfig.FieldID), joinT.C(notifyflow.FeedConfigPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(notifyflow.FeedConfigPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notifyflow.FeedConfigPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := model.InternalID(values[0].(*sql.NullInt64).Int64)
				inValue := model.InternalID(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*NotifyFlow]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*FeedConfig](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "feed_config" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nfq *NotifyFlowQuery) loadNotifyFlowTarget(ctx context.Context, query *NotifyFlowTargetQuery, nodes []*NotifyFlow, init func(*NotifyFlow), assign func(*NotifyFlow, *NotifyFlowTarget)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*NotifyFlow)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(notifyflowtarget.FieldNotifyFlowID)
	}
	query.Where(predicate.NotifyFlowTarget(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notifyflow.NotifyFlowTargetColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifyFlowID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notify_flow_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (nfq *NotifyFlowQuery) loadNotifyFlowSource(ctx context.Context, query *NotifyFlowSourceQuery, nodes []*NotifyFlow, init func(*NotifyFlow), assign func(*NotifyFlow, *NotifyFlowSource)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*NotifyFlow)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(notifyflowsource.FieldNotifyFlowID)
	}
	query.Where(predicate.NotifyFlowSource(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notifyflow.NotifyFlowSourceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifyFlowID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notify_flow_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nfq *NotifyFlowQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nfq.querySpec()
	_spec.Node.Columns = nfq.ctx.Fields
	if len(nfq.ctx.Fields) > 0 {
		_spec.Unique = nfq.ctx.Unique != nil && *nfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nfq.driver, _spec)
}

func (nfq *NotifyFlowQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notifyflow.Table, notifyflow.Columns, sqlgraph.NewFieldSpec(notifyflow.FieldID, field.TypeInt64))
	_spec.From = nfq.sql
	if unique := nfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nfq.path != nil {
		_spec.Unique = true
	}
	if fields := nfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifyflow.FieldID)
		for i := range fields {
			if fields[i] != notifyflow.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nfq *NotifyFlowQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nfq.driver.Dialect())
	t1 := builder.Table(notifyflow.Table)
	columns := nfq.ctx.Fields
	if len(columns) == 0 {
		columns = notifyflow.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nfq.sql != nil {
		selector = nfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nfq.ctx.Unique != nil && *nfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nfq.predicates {
		p(selector)
	}
	for _, p := range nfq.order {
		p(selector)
	}
	if offset := nfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotifyFlowGroupBy is the group-by builder for NotifyFlow entities.
type NotifyFlowGroupBy struct {
	selector
	build *NotifyFlowQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nfgb *NotifyFlowGroupBy) Aggregate(fns ...AggregateFunc) *NotifyFlowGroupBy {
	nfgb.fns = append(nfgb.fns, fns...)
	return nfgb
}

// Scan applies the selector query and scans the result into the given value.
func (nfgb *NotifyFlowGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nfgb.build.ctx, "GroupBy")
	if err := nfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifyFlowQuery, *NotifyFlowGroupBy](ctx, nfgb.build, nfgb, nfgb.build.inters, v)
}

func (nfgb *NotifyFlowGroupBy) sqlScan(ctx context.Context, root *NotifyFlowQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nfgb.fns))
	for _, fn := range nfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nfgb.flds)+len(nfgb.fns))
		for _, f := range *nfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotifyFlowSelect is the builder for selecting fields of NotifyFlow entities.
type NotifyFlowSelect struct {
	*NotifyFlowQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nfs *NotifyFlowSelect) Aggregate(fns ...AggregateFunc) *NotifyFlowSelect {
	nfs.fns = append(nfs.fns, fns...)
	return nfs
}

// Scan applies the selector query and scans the result into the given value.
func (nfs *NotifyFlowSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nfs.ctx, "Select")
	if err := nfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifyFlowQuery, *NotifyFlowSelect](ctx, nfs.NotifyFlowQuery, nfs, nfs.inters, v)
}

func (nfs *NotifyFlowSelect) sqlScan(ctx context.Context, root *NotifyFlowQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nfs.fns))
	for _, fn := range nfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
