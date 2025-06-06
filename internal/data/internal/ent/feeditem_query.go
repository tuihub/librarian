// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/feed"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditem"
	"github.com/tuihub/librarian/internal/data/internal/ent/feeditemcollection"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// FeedItemQuery is the builder for querying FeedItem entities.
type FeedItemQuery struct {
	config
	ctx                    *QueryContext
	order                  []feeditem.OrderOption
	inters                 []Interceptor
	predicates             []predicate.FeedItem
	withFeed               *FeedQuery
	withFeedItemCollection *FeedItemCollectionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeedItemQuery builder.
func (fiq *FeedItemQuery) Where(ps ...predicate.FeedItem) *FeedItemQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit the number of records to be returned by this query.
func (fiq *FeedItemQuery) Limit(limit int) *FeedItemQuery {
	fiq.ctx.Limit = &limit
	return fiq
}

// Offset to start from.
func (fiq *FeedItemQuery) Offset(offset int) *FeedItemQuery {
	fiq.ctx.Offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FeedItemQuery) Unique(unique bool) *FeedItemQuery {
	fiq.ctx.Unique = &unique
	return fiq
}

// Order specifies how the records should be ordered.
func (fiq *FeedItemQuery) Order(o ...feeditem.OrderOption) *FeedItemQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QueryFeed chains the current query on the "feed" edge.
func (fiq *FeedItemQuery) QueryFeed() *FeedQuery {
	query := (&FeedClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feeditem.Table, feeditem.FieldID, selector),
			sqlgraph.To(feed.Table, feed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, feeditem.FeedTable, feeditem.FeedColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeedItemCollection chains the current query on the "feed_item_collection" edge.
func (fiq *FeedItemQuery) QueryFeedItemCollection() *FeedItemCollectionQuery {
	query := (&FeedItemCollectionClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feeditem.Table, feeditem.FieldID, selector),
			sqlgraph.To(feeditemcollection.Table, feeditemcollection.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, feeditem.FeedItemCollectionTable, feeditem.FeedItemCollectionPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FeedItem entity from the query.
// Returns a *NotFoundError when no FeedItem was found.
func (fiq *FeedItemQuery) First(ctx context.Context) (*FeedItem, error) {
	nodes, err := fiq.Limit(1).All(setContextOp(ctx, fiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feeditem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FeedItemQuery) FirstX(ctx context.Context) *FeedItem {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeedItem ID from the query.
// Returns a *NotFoundError when no FeedItem ID was found.
func (fiq *FeedItemQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fiq.Limit(1).IDs(setContextOp(ctx, fiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feeditem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FeedItemQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeedItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeedItem entity is found.
// Returns a *NotFoundError when no FeedItem entities are found.
func (fiq *FeedItemQuery) Only(ctx context.Context) (*FeedItem, error) {
	nodes, err := fiq.Limit(2).All(setContextOp(ctx, fiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feeditem.Label}
	default:
		return nil, &NotSingularError{feeditem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FeedItemQuery) OnlyX(ctx context.Context) *FeedItem {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeedItem ID in the query.
// Returns a *NotSingularError when more than one FeedItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FeedItemQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = fiq.Limit(2).IDs(setContextOp(ctx, fiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feeditem.Label}
	default:
		err = &NotSingularError{feeditem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FeedItemQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeedItems.
func (fiq *FeedItemQuery) All(ctx context.Context) ([]*FeedItem, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryAll)
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FeedItem, *FeedItemQuery]()
	return withInterceptors[[]*FeedItem](ctx, fiq, qr, fiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FeedItemQuery) AllX(ctx context.Context) []*FeedItem {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeedItem IDs.
func (fiq *FeedItemQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if fiq.ctx.Unique == nil && fiq.path != nil {
		fiq.Unique(true)
	}
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryIDs)
	if err = fiq.Select(feeditem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FeedItemQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FeedItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryCount)
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fiq, querierCount[*FeedItemQuery](), fiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FeedItemQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FeedItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fiq.ctx, ent.OpQueryExist)
	switch _, err := fiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FeedItemQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeedItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FeedItemQuery) Clone() *FeedItemQuery {
	if fiq == nil {
		return nil
	}
	return &FeedItemQuery{
		config:                 fiq.config,
		ctx:                    fiq.ctx.Clone(),
		order:                  append([]feeditem.OrderOption{}, fiq.order...),
		inters:                 append([]Interceptor{}, fiq.inters...),
		predicates:             append([]predicate.FeedItem{}, fiq.predicates...),
		withFeed:               fiq.withFeed.Clone(),
		withFeedItemCollection: fiq.withFeedItemCollection.Clone(),
		// clone intermediate query.
		sql:  fiq.sql.Clone(),
		path: fiq.path,
	}
}

// WithFeed tells the query-builder to eager-load the nodes that are connected to
// the "feed" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FeedItemQuery) WithFeed(opts ...func(*FeedQuery)) *FeedItemQuery {
	query := (&FeedClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFeed = query
	return fiq
}

// WithFeedItemCollection tells the query-builder to eager-load the nodes that are connected to
// the "feed_item_collection" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FeedItemQuery) WithFeedItemCollection(opts ...func(*FeedItemCollectionQuery)) *FeedItemQuery {
	query := (&FeedItemCollectionClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withFeedItemCollection = query
	return fiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FeedID model.InternalID `json:"feed_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FeedItem.Query().
//		GroupBy(feeditem.FieldFeedID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fiq *FeedItemQuery) GroupBy(field string, fields ...string) *FeedItemGroupBy {
	fiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeedItemGroupBy{build: fiq}
	grbuild.flds = &fiq.ctx.Fields
	grbuild.label = feeditem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FeedID model.InternalID `json:"feed_id,omitempty"`
//	}
//
//	client.FeedItem.Query().
//		Select(feeditem.FieldFeedID).
//		Scan(ctx, &v)
func (fiq *FeedItemQuery) Select(fields ...string) *FeedItemSelect {
	fiq.ctx.Fields = append(fiq.ctx.Fields, fields...)
	sbuild := &FeedItemSelect{FeedItemQuery: fiq}
	sbuild.label = feeditem.Label
	sbuild.flds, sbuild.scan = &fiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeedItemSelect configured with the given aggregations.
func (fiq *FeedItemQuery) Aggregate(fns ...AggregateFunc) *FeedItemSelect {
	return fiq.Select().Aggregate(fns...)
}

func (fiq *FeedItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fiq); err != nil {
				return err
			}
		}
	}
	for _, f := range fiq.ctx.Fields {
		if !feeditem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FeedItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeedItem, error) {
	var (
		nodes       = []*FeedItem{}
		_spec       = fiq.querySpec()
		loadedTypes = [2]bool{
			fiq.withFeed != nil,
			fiq.withFeedItemCollection != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FeedItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FeedItem{config: fiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fiq.withFeed; query != nil {
		if err := fiq.loadFeed(ctx, query, nodes, nil,
			func(n *FeedItem, e *Feed) { n.Edges.Feed = e }); err != nil {
			return nil, err
		}
	}
	if query := fiq.withFeedItemCollection; query != nil {
		if err := fiq.loadFeedItemCollection(ctx, query, nodes,
			func(n *FeedItem) { n.Edges.FeedItemCollection = []*FeedItemCollection{} },
			func(n *FeedItem, e *FeedItemCollection) {
				n.Edges.FeedItemCollection = append(n.Edges.FeedItemCollection, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fiq *FeedItemQuery) loadFeed(ctx context.Context, query *FeedQuery, nodes []*FeedItem, init func(*FeedItem), assign func(*FeedItem, *Feed)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*FeedItem)
	for i := range nodes {
		fk := nodes[i].FeedID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(feed.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "feed_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fiq *FeedItemQuery) loadFeedItemCollection(ctx context.Context, query *FeedItemCollectionQuery, nodes []*FeedItem, init func(*FeedItem), assign func(*FeedItem, *FeedItemCollection)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*FeedItem)
	nids := make(map[model.InternalID]map[*FeedItem]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(feeditem.FeedItemCollectionTable)
		s.Join(joinT).On(s.C(feeditemcollection.FieldID), joinT.C(feeditem.FeedItemCollectionPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(feeditem.FeedItemCollectionPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(feeditem.FeedItemCollectionPrimaryKey[0]))
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
					nids[inValue] = map[*FeedItem]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*FeedItemCollection](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "feed_item_collection" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (fiq *FeedItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	_spec.Node.Columns = fiq.ctx.Fields
	if len(fiq.ctx.Fields) > 0 {
		_spec.Unique = fiq.ctx.Unique != nil && *fiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FeedItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(feeditem.Table, feeditem.Columns, sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt64))
	_spec.From = fiq.sql
	if unique := fiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fiq.path != nil {
		_spec.Unique = true
	}
	if fields := fiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feeditem.FieldID)
		for i := range fields {
			if fields[i] != feeditem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fiq.withFeed != nil {
			_spec.Node.AddColumnOnce(feeditem.FieldFeedID)
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FeedItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(feeditem.Table)
	columns := fiq.ctx.Fields
	if len(columns) == 0 {
		columns = feeditem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fiq.ctx.Unique != nil && *fiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeedItemGroupBy is the group-by builder for FeedItem entities.
type FeedItemGroupBy struct {
	selector
	build *FeedItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FeedItemGroupBy) Aggregate(fns ...AggregateFunc) *FeedItemGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the selector query and scans the result into the given value.
func (figb *FeedItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, figb.build.ctx, ent.OpQueryGroupBy)
	if err := figb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedItemQuery, *FeedItemGroupBy](ctx, figb.build, figb, figb.build.inters, v)
}

func (figb *FeedItemGroupBy) sqlScan(ctx context.Context, root *FeedItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*figb.flds)+len(figb.fns))
		for _, f := range *figb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*figb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeedItemSelect is the builder for selecting fields of FeedItem entities.
type FeedItemSelect struct {
	*FeedItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fis *FeedItemSelect) Aggregate(fns ...AggregateFunc) *FeedItemSelect {
	fis.fns = append(fis.fns, fns...)
	return fis
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FeedItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fis.ctx, ent.OpQuerySelect)
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeedItemQuery, *FeedItemSelect](ctx, fis.FeedItemQuery, fis, fis.inters, v)
}

func (fis *FeedItemSelect) sqlScan(ctx context.Context, root *FeedItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fis.fns))
	for _, fn := range fis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
