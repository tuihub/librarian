// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/appinst"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// AppInstQuery is the builder for querying AppInst entities.
type AppInstQuery struct {
	config
	ctx        *QueryContext
	order      []appinst.OrderOption
	inters     []Interceptor
	predicates []predicate.AppInst
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppInstQuery builder.
func (aiq *AppInstQuery) Where(ps ...predicate.AppInst) *AppInstQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *AppInstQuery) Limit(limit int) *AppInstQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *AppInstQuery) Offset(offset int) *AppInstQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *AppInstQuery) Unique(unique bool) *AppInstQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *AppInstQuery) Order(o ...appinst.OrderOption) *AppInstQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// QueryOwner chains the current query on the "owner" edge.
func (aiq *AppInstQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: aiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appinst.Table, appinst.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appinst.OwnerTable, appinst.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(aiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppInst entity from the query.
// Returns a *NotFoundError when no AppInst was found.
func (aiq *AppInstQuery) First(ctx context.Context) (*AppInst, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appinst.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *AppInstQuery) FirstX(ctx context.Context) *AppInst {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppInst ID from the query.
// Returns a *NotFoundError when no AppInst ID was found.
func (aiq *AppInstQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appinst.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *AppInstQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppInst entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppInst entity is found.
// Returns a *NotFoundError when no AppInst entities are found.
func (aiq *AppInstQuery) Only(ctx context.Context) (*AppInst, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appinst.Label}
	default:
		return nil, &NotSingularError{appinst.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *AppInstQuery) OnlyX(ctx context.Context) *AppInst {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppInst ID in the query.
// Returns a *NotSingularError when more than one AppInst ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *AppInstQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appinst.Label}
	default:
		err = &NotSingularError{appinst.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *AppInstQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppInsts.
func (aiq *AppInstQuery) All(ctx context.Context) ([]*AppInst, error) {
	ctx = setContextOp(ctx, aiq.ctx, "All")
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppInst, *AppInstQuery]()
	return withInterceptors[[]*AppInst](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *AppInstQuery) AllX(ctx context.Context) []*AppInst {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppInst IDs.
func (aiq *AppInstQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, "IDs")
	if err = aiq.Select(appinst.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *AppInstQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *AppInstQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Count")
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*AppInstQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *AppInstQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *AppInstQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Exist")
	switch _, err := aiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *AppInstQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppInstQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *AppInstQuery) Clone() *AppInstQuery {
	if aiq == nil {
		return nil
	}
	return &AppInstQuery{
		config:     aiq.config,
		ctx:        aiq.ctx.Clone(),
		order:      append([]appinst.OrderOption{}, aiq.order...),
		inters:     append([]Interceptor{}, aiq.inters...),
		predicates: append([]predicate.AppInst{}, aiq.predicates...),
		withOwner:  aiq.withOwner.Clone(),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (aiq *AppInstQuery) WithOwner(opts ...func(*UserQuery)) *AppInstQuery {
	query := (&UserClient{config: aiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aiq.withOwner = query
	return aiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceID model.InternalID `json:"device_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppInst.Query().
//		GroupBy(appinst.FieldDeviceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aiq *AppInstQuery) GroupBy(field string, fields ...string) *AppInstGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppInstGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = appinst.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceID model.InternalID `json:"device_id,omitempty"`
//	}
//
//	client.AppInst.Query().
//		Select(appinst.FieldDeviceID).
//		Scan(ctx, &v)
func (aiq *AppInstQuery) Select(fields ...string) *AppInstSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &AppInstSelect{AppInstQuery: aiq}
	sbuild.label = appinst.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppInstSelect configured with the given aggregations.
func (aiq *AppInstQuery) Aggregate(fns ...AggregateFunc) *AppInstSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *AppInstQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aiq); err != nil {
				return err
			}
		}
	}
	for _, f := range aiq.ctx.Fields {
		if !appinst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *AppInstQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppInst, error) {
	var (
		nodes       = []*AppInst{}
		withFKs     = aiq.withFKs
		_spec       = aiq.querySpec()
		loadedTypes = [1]bool{
			aiq.withOwner != nil,
		}
	)
	if aiq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, appinst.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppInst).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppInst{config: aiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aiq.withOwner; query != nil {
		if err := aiq.loadOwner(ctx, query, nodes, nil,
			func(n *AppInst, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aiq *AppInstQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*AppInst, init func(*AppInst), assign func(*AppInst, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*AppInst)
	for i := range nodes {
		if nodes[i].user_app_inst == nil {
			continue
		}
		fk := *nodes[i].user_app_inst
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
			return fmt.Errorf(`unexpected foreign-key "user_app_inst" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aiq *AppInstQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *AppInstQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appinst.Table, appinst.Columns, sqlgraph.NewFieldSpec(appinst.FieldID, field.TypeInt64))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinst.FieldID)
		for i := range fields {
			if fields[i] != appinst.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aiq *AppInstQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(appinst.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = appinst.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aiq.ctx.Unique != nil && *aiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector)
	}
	if offset := aiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppInstGroupBy is the group-by builder for AppInst entities.
type AppInstGroupBy struct {
	selector
	build *AppInstQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *AppInstGroupBy) Aggregate(fns ...AggregateFunc) *AppInstGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *AppInstGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, "GroupBy")
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInstQuery, *AppInstGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *AppInstGroupBy) sqlScan(ctx context.Context, root *AppInstQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aigb.fns))
	for _, fn := range aigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aigb.flds)+len(aigb.fns))
		for _, f := range *aigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppInstSelect is the builder for selecting fields of AppInst entities.
type AppInstSelect struct {
	*AppInstQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *AppInstSelect) Aggregate(fns ...AggregateFunc) *AppInstSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *AppInstSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, "Select")
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInstQuery, *AppInstSelect](ctx, ais.AppInstQuery, ais, ais.inters, v)
}

func (ais *AppInstSelect) sqlScan(ctx context.Context, root *AppInstQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ais.fns))
	for _, fn := range ais.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ais.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
