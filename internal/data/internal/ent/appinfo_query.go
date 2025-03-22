// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/appinfo"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppInfoQuery is the builder for querying AppInfo entities.
type AppInfoQuery struct {
	config
	ctx        *QueryContext
	order      []appinfo.OrderOption
	inters     []Interceptor
	predicates []predicate.AppInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppInfoQuery builder.
func (aiq *AppInfoQuery) Where(ps ...predicate.AppInfo) *AppInfoQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *AppInfoQuery) Limit(limit int) *AppInfoQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *AppInfoQuery) Offset(offset int) *AppInfoQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *AppInfoQuery) Unique(unique bool) *AppInfoQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *AppInfoQuery) Order(o ...appinfo.OrderOption) *AppInfoQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// First returns the first AppInfo entity from the query.
// Returns a *NotFoundError when no AppInfo was found.
func (aiq *AppInfoQuery) First(ctx context.Context) (*AppInfo, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *AppInfoQuery) FirstX(ctx context.Context) *AppInfo {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppInfo ID from the query.
// Returns a *NotFoundError when no AppInfo ID was found.
func (aiq *AppInfoQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *AppInfoQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppInfo entity is found.
// Returns a *NotFoundError when no AppInfo entities are found.
func (aiq *AppInfoQuery) Only(ctx context.Context) (*AppInfo, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appinfo.Label}
	default:
		return nil, &NotSingularError{appinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *AppInfoQuery) OnlyX(ctx context.Context) *AppInfo {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppInfo ID in the query.
// Returns a *NotSingularError when more than one AppInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *AppInfoQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appinfo.Label}
	default:
		err = &NotSingularError{appinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *AppInfoQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppInfos.
func (aiq *AppInfoQuery) All(ctx context.Context) ([]*AppInfo, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryAll)
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppInfo, *AppInfoQuery]()
	return withInterceptors[[]*AppInfo](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *AppInfoQuery) AllX(ctx context.Context) []*AppInfo {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppInfo IDs.
func (aiq *AppInfoQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryIDs)
	if err = aiq.Select(appinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *AppInfoQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *AppInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryCount)
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*AppInfoQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *AppInfoQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *AppInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, ent.OpQueryExist)
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
func (aiq *AppInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *AppInfoQuery) Clone() *AppInfoQuery {
	if aiq == nil {
		return nil
	}
	return &AppInfoQuery{
		config:     aiq.config,
		ctx:        aiq.ctx.Clone(),
		order:      append([]appinfo.OrderOption{}, aiq.order...),
		inters:     append([]Interceptor{}, aiq.inters...),
		predicates: append([]predicate.AppInfo{}, aiq.predicates...),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Source string `json:"source,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppInfo.Query().
//		GroupBy(appinfo.FieldSource).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aiq *AppInfoQuery) GroupBy(field string, fields ...string) *AppInfoGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppInfoGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = appinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Source string `json:"source,omitempty"`
//	}
//
//	client.AppInfo.Query().
//		Select(appinfo.FieldSource).
//		Scan(ctx, &v)
func (aiq *AppInfoQuery) Select(fields ...string) *AppInfoSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &AppInfoSelect{AppInfoQuery: aiq}
	sbuild.label = appinfo.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppInfoSelect configured with the given aggregations.
func (aiq *AppInfoQuery) Aggregate(fns ...AggregateFunc) *AppInfoSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *AppInfoQuery) prepareQuery(ctx context.Context) error {
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
		if !appinfo.ValidColumn(f) {
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

func (aiq *AppInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppInfo, error) {
	var (
		nodes = []*AppInfo{}
		_spec = aiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppInfo{config: aiq.config}
		nodes = append(nodes, node)
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
	return nodes, nil
}

func (aiq *AppInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *AppInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appinfo.Table, appinfo.Columns, sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeInt64))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appinfo.FieldID)
		for i := range fields {
			if fields[i] != appinfo.FieldID {
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

func (aiq *AppInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(appinfo.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = appinfo.Columns
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

// AppInfoGroupBy is the group-by builder for AppInfo entities.
type AppInfoGroupBy struct {
	selector
	build *AppInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *AppInfoGroupBy) Aggregate(fns ...AggregateFunc) *AppInfoGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *AppInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, ent.OpQueryGroupBy)
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInfoQuery, *AppInfoGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *AppInfoGroupBy) sqlScan(ctx context.Context, root *AppInfoQuery, v any) error {
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

// AppInfoSelect is the builder for selecting fields of AppInfo entities.
type AppInfoSelect struct {
	*AppInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *AppInfoSelect) Aggregate(fns ...AggregateFunc) *AppInfoSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *AppInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, ent.OpQuerySelect)
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppInfoQuery, *AppInfoSelect](ctx, ais.AppInfoQuery, ais, ais.inters, v)
}

func (ais *AppInfoSelect) sqlScan(ctx context.Context, root *AppInfoQuery, v any) error {
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
