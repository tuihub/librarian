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
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/sentinelappbinary"
)

// SentinelAppBinaryQuery is the builder for querying SentinelAppBinary entities.
type SentinelAppBinaryQuery struct {
	config
	ctx        *QueryContext
	order      []sentinelappbinary.OrderOption
	inters     []Interceptor
	predicates []predicate.SentinelAppBinary
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SentinelAppBinaryQuery builder.
func (sabq *SentinelAppBinaryQuery) Where(ps ...predicate.SentinelAppBinary) *SentinelAppBinaryQuery {
	sabq.predicates = append(sabq.predicates, ps...)
	return sabq
}

// Limit the number of records to be returned by this query.
func (sabq *SentinelAppBinaryQuery) Limit(limit int) *SentinelAppBinaryQuery {
	sabq.ctx.Limit = &limit
	return sabq
}

// Offset to start from.
func (sabq *SentinelAppBinaryQuery) Offset(offset int) *SentinelAppBinaryQuery {
	sabq.ctx.Offset = &offset
	return sabq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sabq *SentinelAppBinaryQuery) Unique(unique bool) *SentinelAppBinaryQuery {
	sabq.ctx.Unique = &unique
	return sabq
}

// Order specifies how the records should be ordered.
func (sabq *SentinelAppBinaryQuery) Order(o ...sentinelappbinary.OrderOption) *SentinelAppBinaryQuery {
	sabq.order = append(sabq.order, o...)
	return sabq
}

// First returns the first SentinelAppBinary entity from the query.
// Returns a *NotFoundError when no SentinelAppBinary was found.
func (sabq *SentinelAppBinaryQuery) First(ctx context.Context) (*SentinelAppBinary, error) {
	nodes, err := sabq.Limit(1).All(setContextOp(ctx, sabq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sentinelappbinary.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) FirstX(ctx context.Context) *SentinelAppBinary {
	node, err := sabq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SentinelAppBinary ID from the query.
// Returns a *NotFoundError when no SentinelAppBinary ID was found.
func (sabq *SentinelAppBinaryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sabq.Limit(1).IDs(setContextOp(ctx, sabq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sentinelappbinary.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) FirstIDX(ctx context.Context) int {
	id, err := sabq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SentinelAppBinary entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SentinelAppBinary entity is found.
// Returns a *NotFoundError when no SentinelAppBinary entities are found.
func (sabq *SentinelAppBinaryQuery) Only(ctx context.Context) (*SentinelAppBinary, error) {
	nodes, err := sabq.Limit(2).All(setContextOp(ctx, sabq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sentinelappbinary.Label}
	default:
		return nil, &NotSingularError{sentinelappbinary.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) OnlyX(ctx context.Context) *SentinelAppBinary {
	node, err := sabq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SentinelAppBinary ID in the query.
// Returns a *NotSingularError when more than one SentinelAppBinary ID is found.
// Returns a *NotFoundError when no entities are found.
func (sabq *SentinelAppBinaryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sabq.Limit(2).IDs(setContextOp(ctx, sabq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sentinelappbinary.Label}
	default:
		err = &NotSingularError{sentinelappbinary.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) OnlyIDX(ctx context.Context) int {
	id, err := sabq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SentinelAppBinaries.
func (sabq *SentinelAppBinaryQuery) All(ctx context.Context) ([]*SentinelAppBinary, error) {
	ctx = setContextOp(ctx, sabq.ctx, ent.OpQueryAll)
	if err := sabq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SentinelAppBinary, *SentinelAppBinaryQuery]()
	return withInterceptors[[]*SentinelAppBinary](ctx, sabq, qr, sabq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) AllX(ctx context.Context) []*SentinelAppBinary {
	nodes, err := sabq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SentinelAppBinary IDs.
func (sabq *SentinelAppBinaryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sabq.ctx.Unique == nil && sabq.path != nil {
		sabq.Unique(true)
	}
	ctx = setContextOp(ctx, sabq.ctx, ent.OpQueryIDs)
	if err = sabq.Select(sentinelappbinary.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) IDsX(ctx context.Context) []int {
	ids, err := sabq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sabq *SentinelAppBinaryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sabq.ctx, ent.OpQueryCount)
	if err := sabq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sabq, querierCount[*SentinelAppBinaryQuery](), sabq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) CountX(ctx context.Context) int {
	count, err := sabq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sabq *SentinelAppBinaryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sabq.ctx, ent.OpQueryExist)
	switch _, err := sabq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sabq *SentinelAppBinaryQuery) ExistX(ctx context.Context) bool {
	exist, err := sabq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SentinelAppBinaryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sabq *SentinelAppBinaryQuery) Clone() *SentinelAppBinaryQuery {
	if sabq == nil {
		return nil
	}
	return &SentinelAppBinaryQuery{
		config:     sabq.config,
		ctx:        sabq.ctx.Clone(),
		order:      append([]sentinelappbinary.OrderOption{}, sabq.order...),
		inters:     append([]Interceptor{}, sabq.inters...),
		predicates: append([]predicate.SentinelAppBinary{}, sabq.predicates...),
		// clone intermediate query.
		sql:  sabq.sql.Clone(),
		path: sabq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SentinelInfoID model.InternalID `json:"sentinel_info_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SentinelAppBinary.Query().
//		GroupBy(sentinelappbinary.FieldSentinelInfoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sabq *SentinelAppBinaryQuery) GroupBy(field string, fields ...string) *SentinelAppBinaryGroupBy {
	sabq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SentinelAppBinaryGroupBy{build: sabq}
	grbuild.flds = &sabq.ctx.Fields
	grbuild.label = sentinelappbinary.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SentinelInfoID model.InternalID `json:"sentinel_info_id,omitempty"`
//	}
//
//	client.SentinelAppBinary.Query().
//		Select(sentinelappbinary.FieldSentinelInfoID).
//		Scan(ctx, &v)
func (sabq *SentinelAppBinaryQuery) Select(fields ...string) *SentinelAppBinarySelect {
	sabq.ctx.Fields = append(sabq.ctx.Fields, fields...)
	sbuild := &SentinelAppBinarySelect{SentinelAppBinaryQuery: sabq}
	sbuild.label = sentinelappbinary.Label
	sbuild.flds, sbuild.scan = &sabq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SentinelAppBinarySelect configured with the given aggregations.
func (sabq *SentinelAppBinaryQuery) Aggregate(fns ...AggregateFunc) *SentinelAppBinarySelect {
	return sabq.Select().Aggregate(fns...)
}

func (sabq *SentinelAppBinaryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sabq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sabq); err != nil {
				return err
			}
		}
	}
	for _, f := range sabq.ctx.Fields {
		if !sentinelappbinary.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sabq.path != nil {
		prev, err := sabq.path(ctx)
		if err != nil {
			return err
		}
		sabq.sql = prev
	}
	return nil
}

func (sabq *SentinelAppBinaryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SentinelAppBinary, error) {
	var (
		nodes = []*SentinelAppBinary{}
		_spec = sabq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SentinelAppBinary).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SentinelAppBinary{config: sabq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sabq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sabq *SentinelAppBinaryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sabq.querySpec()
	_spec.Node.Columns = sabq.ctx.Fields
	if len(sabq.ctx.Fields) > 0 {
		_spec.Unique = sabq.ctx.Unique != nil && *sabq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sabq.driver, _spec)
}

func (sabq *SentinelAppBinaryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sentinelappbinary.Table, sentinelappbinary.Columns, sqlgraph.NewFieldSpec(sentinelappbinary.FieldID, field.TypeInt))
	_spec.From = sabq.sql
	if unique := sabq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sabq.path != nil {
		_spec.Unique = true
	}
	if fields := sabq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sentinelappbinary.FieldID)
		for i := range fields {
			if fields[i] != sentinelappbinary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sabq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sabq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sabq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sabq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sabq *SentinelAppBinaryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sabq.driver.Dialect())
	t1 := builder.Table(sentinelappbinary.Table)
	columns := sabq.ctx.Fields
	if len(columns) == 0 {
		columns = sentinelappbinary.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sabq.sql != nil {
		selector = sabq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sabq.ctx.Unique != nil && *sabq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sabq.predicates {
		p(selector)
	}
	for _, p := range sabq.order {
		p(selector)
	}
	if offset := sabq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sabq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SentinelAppBinaryGroupBy is the group-by builder for SentinelAppBinary entities.
type SentinelAppBinaryGroupBy struct {
	selector
	build *SentinelAppBinaryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sabgb *SentinelAppBinaryGroupBy) Aggregate(fns ...AggregateFunc) *SentinelAppBinaryGroupBy {
	sabgb.fns = append(sabgb.fns, fns...)
	return sabgb
}

// Scan applies the selector query and scans the result into the given value.
func (sabgb *SentinelAppBinaryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sabgb.build.ctx, ent.OpQueryGroupBy)
	if err := sabgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SentinelAppBinaryQuery, *SentinelAppBinaryGroupBy](ctx, sabgb.build, sabgb, sabgb.build.inters, v)
}

func (sabgb *SentinelAppBinaryGroupBy) sqlScan(ctx context.Context, root *SentinelAppBinaryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sabgb.fns))
	for _, fn := range sabgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sabgb.flds)+len(sabgb.fns))
		for _, f := range *sabgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sabgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sabgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SentinelAppBinarySelect is the builder for selecting fields of SentinelAppBinary entities.
type SentinelAppBinarySelect struct {
	*SentinelAppBinaryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sabs *SentinelAppBinarySelect) Aggregate(fns ...AggregateFunc) *SentinelAppBinarySelect {
	sabs.fns = append(sabs.fns, fns...)
	return sabs
}

// Scan applies the selector query and scans the result into the given value.
func (sabs *SentinelAppBinarySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sabs.ctx, ent.OpQuerySelect)
	if err := sabs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SentinelAppBinaryQuery, *SentinelAppBinarySelect](ctx, sabs.SentinelAppBinaryQuery, sabs, sabs.inters, v)
}

func (sabs *SentinelAppBinarySelect) sqlScan(ctx context.Context, root *SentinelAppBinaryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sabs.fns))
	for _, fn := range sabs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sabs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sabs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
