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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/feedconfig"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflow"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/notifyflowsource"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// FeedConfigQuery is the builder for querying FeedConfig entities.
type FeedConfigQuery struct {
	config
	ctx                  *QueryContext
	order                []feedconfig.OrderOption
	inters               []Interceptor
	predicates           []predicate.FeedConfig
	withOwner            *UserQuery
	withFeed             *FeedQuery
	withNotifyFlow       *NotifyFlowQuery
	withNotifyFlowSource *NotifyFlowSourceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedConfigQuery builder.
func (fcq *FeedConfigQuery) Where(ps ...predicate.FeedConfig) *FeedConfigQuery {
	fcq.predicates = append(fcq.predicates, ps...)
	return fcq
}

// Limit the number of records to be returned by this query.
func (fcq *FeedConfigQuery) Limit(limit int) *FeedConfigQuery {
	fcq.ctx.Limit = &limit
	return fcq
}

// Offset to start from.
func (fcq *FeedConfigQuery) Offset(offset int) *FeedConfigQuery {
	fcq.ctx.Offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FeedConfigQuery) Unique(unique bool) *FeedConfigQuery {
	fcq.ctx.Unique = &unique
	return fcq
}

// Order specifies how the records should be ordered.
func (fcq *FeedConfigQuery) Order(o ...feedconfig.OrderOption) *FeedConfigQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// QueryOwner chains the current query on the "owner" edge.
func (fcq *FeedConfigQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedconfig.Table, feedconfig.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, feedconfig.OwnerTable, feedconfig.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeed chains the current query on the "feed" edge.
func (fcq *FeedConfigQuery) QueryFeed() *FeedQuery {
	query := (&FeedClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedconfig.Table, feedconfig.FieldID, selector),
			sqlgraph.To(feed.Table, feed.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, feedconfig.FeedTable, feedconfig.FeedColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyFlow chains the current query on the "notify_flow" edge.
func (fcq *FeedConfigQuery) QueryNotifyFlow() *NotifyFlowQuery {
	query := (&NotifyFlowClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedconfig.Table, feedconfig.FieldID, selector),
			sqlgraph.To(notifyflow.Table, notifyflow.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, feedconfig.NotifyFlowTable, feedconfig.NotifyFlowPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotifyFlowSource chains the current query on the "notify_flow_source" edge.
func (fcq *FeedConfigQuery) QueryNotifyFlowSource() *NotifyFlowSourceQuery {
	query := (&NotifyFlowSourceClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feedconfig.Table, feedconfig.FieldID, selector),
			sqlgraph.To(notifyflowsource.Table, notifyflowsource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, feedconfig.NotifyFlowSourceTable, feedconfig.NotifyFlowSourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FeedConfig entity from the query.
// Returns a *NotFoundError when no FeedConfig was found.
func (fcq *FeedConfigQuery) First(ctx context.Context) (*FeedConfig, error) {
	nodes, err := fcq.Limit(1).All(setContextOp(ctx, fcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feedconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcq *FeedConfigQuery) FirstX(ctx context.Context) *FeedConfig {
	node, err := fcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeedConfig ID from the query.
// Returns a *NotFoundError when no FeedConfig ID was found.
func (fcq *FeedConfigQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fcq.Limit(1).IDs(setContextOp(ctx, fcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feedconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcq *FeedConfigQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := fcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeedConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeedConfig entity is found.
// Returns a *NotFoundError when no FeedConfig entities are found.
func (fcq *FeedConfigQuery) Only(ctx context.Context) (*FeedConfig, error) {
	nodes, err := fcq.Limit(2).All(setContextOp(ctx, fcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feedconfig.Label}
	default:
		return nil, &NotSingularError{feedconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcq *FeedConfigQuery) OnlyX(ctx context.Context) *FeedConfig {
	node, err := fcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeedConfig ID in the query.
// Returns a *NotSingularError when more than one FeedConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcq *FeedConfigQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fcq.Limit(2).IDs(setContextOp(ctx, fcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feedconfig.Label}
	default:
		err = &NotSingularError{feedconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcq *FeedConfigQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := fcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeedConfigs.
func (fcq *FeedConfigQuery) All(ctx context.Context) ([]*FeedConfig, error) {
	ctx = setContextOp(ctx, fcq.ctx, "All")
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FeedConfig, *FeedConfigQuery]()
	return withInterceptors[[]*FeedConfig](ctx, fcq, qr, fcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fcq *FeedConfigQuery) AllX(ctx context.Context) []*FeedConfig {
	nodes, err := fcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeedConfig IDs.
func (fcq *FeedConfigQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if fcq.ctx.Unique == nil && fcq.path != nil {
		fcq.Unique(true)
	}
	ctx = setContextOp(ctx, fcq.ctx, "IDs")
	if err = fcq.Select(feedconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcq *FeedConfigQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := fcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcq *FeedConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Count")
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fcq, querierCount[*FeedConfigQuery](), fcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fcq *FeedConfigQuery) CountX(ctx context.Context) int {
	count, err := fcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcq *FeedConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Exist")
	switch _, err := fcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fcq *FeedConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := fcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcq *FeedConfigQuery) Clone() *FeedConfigQuery {
	if fcq == nil {
		return nil
	}
	return &FeedConfigQuery{
		config:               fcq.config,
		ctx:                  fcq.ctx.Clone(),
		order:                append([]feedconfig.OrderOption{}, fcq.order...),
		inters:               append([]Interceptor{}, fcq.inters...),
		predicates:           append([]predicate.FeedConfig{}, fcq.predicates...),
		withOwner:            fcq.withOwner.Clone(),
		withFeed:             fcq.withFeed.Clone(),
		withNotifyFlow:       fcq.withNotifyFlow.Clone(),
		withNotifyFlowSource: fcq.withNotifyFlowSource.Clone(),
		// clone intermediate query.
		sql:  fcq.sql.Clone(),
		path: fcq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FeedConfigQuery) WithOwner(opts ...func(*UserQuery)) *FeedConfigQuery {
	query := (&UserClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withOwner = query
	return fcq
}

// WithFeed tells the query-builder to eager-load the nodes that are connected to
// the "feed" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FeedConfigQuery) WithFeed(opts ...func(*FeedQuery)) *FeedConfigQuery {
	query := (&FeedClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withFeed = query
	return fcq
}

// WithNotifyFlow tells the query-builder to eager-load the nodes that are connected to
// the "notify_flow" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FeedConfigQuery) WithNotifyFlow(opts ...func(*NotifyFlowQuery)) *FeedConfigQuery {
	query := (&NotifyFlowClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withNotifyFlow = query
	return fcq
}

// WithNotifyFlowSource tells the query-builder to eager-load the nodes that are connected to
// the "notify_flow_source" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FeedConfigQuery) WithNotifyFlowSource(opts ...func(*NotifyFlowSourceQuery)) *FeedConfigQuery {
	query := (&NotifyFlowSourceClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withNotifyFlowSource = query
	return fcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserFeedConfig model.InternalID `json:"user_feed_config,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FeedConfig.Query().
//		GroupBy(feedconfig.FieldUserFeedConfig).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fcq *FeedConfigQuery) GroupBy(field string, fields ...string) *FeedConfigGroupBy {
	fcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeedConfigGroupBy{build: fcq}
	grbuild.flds = &fcq.ctx.Fields
	grbuild.label = feedconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserFeedConfig model.InternalID `json:"user_feed_config,omitempty"`
//	}
//
//	client.FeedConfig.Query().
//		Select(feedconfig.FieldUserFeedConfig).
//		Scan(ctx, &v)
func (fcq *FeedConfigQuery) Select(fields ...string) *FeedConfigSelect {
	fcq.ctx.Fields = append(fcq.ctx.Fields, fields...)
	sbuild := &FeedConfigSelect{FeedConfigQuery: fcq}
	sbuild.label = feedconfig.Label
	sbuild.flds, sbuild.scan = &fcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeedConfigSelect configured with the given aggregations.
func (fcq *FeedConfigQuery) Aggregate(fns ...AggregateFunc) *FeedConfigSelect {
	return fcq.Select().Aggregate(fns...)
}

func (fcq *FeedConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fcq); err != nil {
				return err
			}
		}
	}
	for _, f := range fcq.ctx.Fields {
		if !feedconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcq.path != nil {
		prev, err := fcq.path(ctx)
		if err != nil {
			return err
		}
		fcq.sql = prev
	}
	return nil
}

func (fcq *FeedConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeedConfig, error) {
	var (
		nodes       = []*FeedConfig{}
		_spec       = fcq.querySpec()
		loadedTypes = [4]bool{
			fcq.withOwner != nil,
			fcq.withFeed != nil,
			fcq.withNotifyFlow != nil,
			fcq.withNotifyFlowSource != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FeedConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FeedConfig{config: fcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fcq.withOwner; query != nil {
		if err := fcq.loadOwner(ctx, query, nodes, nil,
			func(n *FeedConfig, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withFeed; query != nil {
		if err := fcq.loadFeed(ctx, query, nodes, nil,
			func(n *FeedConfig, e *Feed) { n.Edges.Feed = e }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withNotifyFlow; query != nil {
		if err := fcq.loadNotifyFlow(ctx, query, nodes,
			func(n *FeedConfig) { n.Edges.NotifyFlow = []*NotifyFlow{} },
			func(n *FeedConfig, e *NotifyFlow) { n.Edges.NotifyFlow = append(n.Edges.NotifyFlow, e) }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withNotifyFlowSource; query != nil {
		if err := fcq.loadNotifyFlowSource(ctx, query, nodes,
			func(n *FeedConfig) { n.Edges.NotifyFlowSource = []*NotifyFlowSource{} },
			func(n *FeedConfig, e *NotifyFlowSource) {
				n.Edges.NotifyFlowSource = append(n.Edges.NotifyFlowSource, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fcq *FeedConfigQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*FeedConfig, init func(*FeedConfig), assign func(*FeedConfig, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*FeedConfig)
	for i := range nodes {
		fk := nodes[i].UserFeedConfig
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
			return fmt.Errorf(`unexpected foreign-key "user_feed_config" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fcq *FeedConfigQuery) loadFeed(ctx context.Context, query *FeedQuery, nodes []*FeedConfig, init func(*FeedConfig), assign func(*FeedConfig, *Feed)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*FeedConfig)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Feed(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(feedconfig.FeedColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.feed_config_feed
		if fk == nil {
			return fmt.Errorf(`foreign-key "feed_config_feed" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "feed_config_feed" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (fcq *FeedConfigQuery) loadNotifyFlow(ctx context.Context, query *NotifyFlowQuery, nodes []*FeedConfig, init func(*FeedConfig), assign func(*FeedConfig, *NotifyFlow)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*FeedConfig)
	nids := make(map[model.InternalID]map[*FeedConfig]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(feedconfig.NotifyFlowTable)
		s.Join(joinT).On(s.C(notifyflow.FieldID), joinT.C(feedconfig.NotifyFlowPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(feedconfig.NotifyFlowPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(feedconfig.NotifyFlowPrimaryKey[0]))
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
					nids[inValue] = map[*FeedConfig]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*NotifyFlow](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "notify_flow" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (fcq *FeedConfigQuery) loadNotifyFlowSource(ctx context.Context, query *NotifyFlowSourceQuery, nodes []*FeedConfig, init func(*FeedConfig), assign func(*FeedConfig, *NotifyFlowSource)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*FeedConfig)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(notifyflowsource.FieldNotifySourceID)
	}
	query.Where(predicate.NotifyFlowSource(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(feedconfig.NotifyFlowSourceColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.NotifySourceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notify_source_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fcq *FeedConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	_spec.Node.Columns = fcq.ctx.Fields
	if len(fcq.ctx.Fields) > 0 {
		_spec.Unique = fcq.ctx.Unique != nil && *fcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FeedConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(feedconfig.Table, feedconfig.Columns, sqlgraph.NewFieldSpec(feedconfig.FieldID, field.TypeInt64))
	_spec.From = fcq.sql
	if unique := fcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fcq.path != nil {
		_spec.Unique = true
	}
	if fields := fcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feedconfig.FieldID)
		for i := range fields {
			if fields[i] != feedconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fcq.withOwner != nil {
			_spec.Node.AddColumnOnce(feedconfig.FieldUserFeedConfig)
		}
	}
	if ps := fcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcq *FeedConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcq.driver.Dialect())
	t1 := builder.Table(feedconfig.Table)
	columns := fcq.ctx.Fields
	if len(columns) == 0 {
		columns = feedconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.ctx.Unique != nil && *fcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedConfigGroupBy is the group-by builder for FeedConfig entities.
type FeedConfigGroupBy struct {
	selector
	build *FeedConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FeedConfigGroupBy) Aggregate(fns ...AggregateFunc) *FeedConfigGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the selector query and scans the result into the given value.
func (fcgb *FeedConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcgb.build.ctx, "GroupBy")
	if err := fcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedConfigQuery, *FeedConfigGroupBy](ctx, fcgb.build, fcgb, fcgb.build.inters, v)
}

func (fcgb *FeedConfigGroupBy) sqlScan(ctx context.Context, root *FeedConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fcgb.flds)+len(fcgb.fns))
		for _, f := range *fcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeedConfigSelect is the builder for selecting fields of FeedConfig entities.
type FeedConfigSelect struct {
	*FeedConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fcs *FeedConfigSelect) Aggregate(fns ...AggregateFunc) *FeedConfigSelect {
	fcs.fns = append(fcs.fns, fns...)
	return fcs
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FeedConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcs.ctx, "Select")
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedConfigQuery, *FeedConfigSelect](ctx, fcs.FeedConfigQuery, fcs, fcs.inters, v)
}

func (fcs *FeedConfigSelect) sqlScan(ctx context.Context, root *FeedConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fcs.fns))
	for _, fn := range fcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
