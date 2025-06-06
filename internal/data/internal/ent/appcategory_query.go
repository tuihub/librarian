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
	"github.com/tuihub/librarian/internal/data/internal/ent/app"
	"github.com/tuihub/librarian/internal/data/internal/ent/appappcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/appcategory"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/model"
)

// AppCategoryQuery is the builder for querying AppCategory entities.
type AppCategoryQuery struct {
	config
	ctx                *QueryContext
	order              []appcategory.OrderOption
	inters             []Interceptor
	predicates         []predicate.AppCategory
	withApp            *AppQuery
	withAppAppCategory *AppAppCategoryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppCategoryQuery builder.
func (acq *AppCategoryQuery) Where(ps ...predicate.AppCategory) *AppCategoryQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AppCategoryQuery) Limit(limit int) *AppCategoryQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AppCategoryQuery) Offset(offset int) *AppCategoryQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AppCategoryQuery) Unique(unique bool) *AppCategoryQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AppCategoryQuery) Order(o ...appcategory.OrderOption) *AppCategoryQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryApp chains the current query on the "app" edge.
func (acq *AppCategoryQuery) QueryApp() *AppQuery {
	query := (&AppClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appcategory.Table, appcategory.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, appcategory.AppTable, appcategory.AppPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAppAppCategory chains the current query on the "app_app_category" edge.
func (acq *AppCategoryQuery) QueryAppAppCategory() *AppAppCategoryQuery {
	query := (&AppAppCategoryClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appcategory.Table, appcategory.FieldID, selector),
			sqlgraph.To(appappcategory.Table, appappcategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, appcategory.AppAppCategoryTable, appcategory.AppAppCategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppCategory entity from the query.
// Returns a *NotFoundError when no AppCategory was found.
func (acq *AppCategoryQuery) First(ctx context.Context) (*AppCategory, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AppCategoryQuery) FirstX(ctx context.Context) *AppCategory {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppCategory ID from the query.
// Returns a *NotFoundError when no AppCategory ID was found.
func (acq *AppCategoryQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AppCategoryQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppCategory entity is found.
// Returns a *NotFoundError when no AppCategory entities are found.
func (acq *AppCategoryQuery) Only(ctx context.Context) (*AppCategory, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appcategory.Label}
	default:
		return nil, &NotSingularError{appcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AppCategoryQuery) OnlyX(ctx context.Context) *AppCategory {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppCategory ID in the query.
// Returns a *NotSingularError when more than one AppCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AppCategoryQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appcategory.Label}
	default:
		err = &NotSingularError{appcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AppCategoryQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppCategories.
func (acq *AppCategoryQuery) All(ctx context.Context) ([]*AppCategory, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryAll)
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppCategory, *AppCategoryQuery]()
	return withInterceptors[[]*AppCategory](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AppCategoryQuery) AllX(ctx context.Context) []*AppCategory {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppCategory IDs.
func (acq *AppCategoryQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryIDs)
	if err = acq.Select(appcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AppCategoryQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AppCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryCount)
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AppCategoryQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AppCategoryQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AppCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, ent.OpQueryExist)
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AppCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AppCategoryQuery) Clone() *AppCategoryQuery {
	if acq == nil {
		return nil
	}
	return &AppCategoryQuery{
		config:             acq.config,
		ctx:                acq.ctx.Clone(),
		order:              append([]appcategory.OrderOption{}, acq.order...),
		inters:             append([]Interceptor{}, acq.inters...),
		predicates:         append([]predicate.AppCategory{}, acq.predicates...),
		withApp:            acq.withApp.Clone(),
		withAppAppCategory: acq.withAppAppCategory.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AppCategoryQuery) WithApp(opts ...func(*AppQuery)) *AppCategoryQuery {
	query := (&AppClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withApp = query
	return acq
}

// WithAppAppCategory tells the query-builder to eager-load the nodes that are connected to
// the "app_app_category" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AppCategoryQuery) WithAppAppCategory(opts ...func(*AppAppCategoryQuery)) *AppCategoryQuery {
	query := (&AppAppCategoryClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withAppAppCategory = query
	return acq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID model.InternalID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppCategory.Query().
//		GroupBy(appcategory.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AppCategoryQuery) GroupBy(field string, fields ...string) *AppCategoryGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppCategoryGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = appcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID model.InternalID `json:"user_id,omitempty"`
//	}
//
//	client.AppCategory.Query().
//		Select(appcategory.FieldUserID).
//		Scan(ctx, &v)
func (acq *AppCategoryQuery) Select(fields ...string) *AppCategorySelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AppCategorySelect{AppCategoryQuery: acq}
	sbuild.label = appcategory.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppCategorySelect configured with the given aggregations.
func (acq *AppCategoryQuery) Aggregate(fns ...AggregateFunc) *AppCategorySelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AppCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !appcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AppCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppCategory, error) {
	var (
		nodes       = []*AppCategory{}
		_spec       = acq.querySpec()
		loadedTypes = [2]bool{
			acq.withApp != nil,
			acq.withAppAppCategory != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppCategory{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withApp; query != nil {
		if err := acq.loadApp(ctx, query, nodes,
			func(n *AppCategory) { n.Edges.App = []*App{} },
			func(n *AppCategory, e *App) { n.Edges.App = append(n.Edges.App, e) }); err != nil {
			return nil, err
		}
	}
	if query := acq.withAppAppCategory; query != nil {
		if err := acq.loadAppAppCategory(ctx, query, nodes,
			func(n *AppCategory) { n.Edges.AppAppCategory = []*AppAppCategory{} },
			func(n *AppCategory, e *AppAppCategory) { n.Edges.AppAppCategory = append(n.Edges.AppAppCategory, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AppCategoryQuery) loadApp(ctx context.Context, query *AppQuery, nodes []*AppCategory, init func(*AppCategory), assign func(*AppCategory, *App)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[model.InternalID]*AppCategory)
	nids := make(map[model.InternalID]map[*AppCategory]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(appcategory.AppTable)
		s.Join(joinT).On(s.C(app.FieldID), joinT.C(appcategory.AppPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(appcategory.AppPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(appcategory.AppPrimaryKey[0]))
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
					nids[inValue] = map[*AppCategory]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*App](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "app" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (acq *AppCategoryQuery) loadAppAppCategory(ctx context.Context, query *AppAppCategoryQuery, nodes []*AppCategory, init func(*AppCategory), assign func(*AppCategory, *AppAppCategory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*AppCategory)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(appappcategory.FieldAppCategoryID)
	}
	query.Where(predicate.AppAppCategory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(appcategory.AppAppCategoryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AppCategoryID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "app_category_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (acq *AppCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AppCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appcategory.Table, appcategory.Columns, sqlgraph.NewFieldSpec(appcategory.FieldID, field.TypeInt64))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcategory.FieldID)
		for i := range fields {
			if fields[i] != appcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AppCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(appcategory.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = appcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppCategoryGroupBy is the group-by builder for AppCategory entities.
type AppCategoryGroupBy struct {
	selector
	build *AppCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AppCategoryGroupBy) Aggregate(fns ...AggregateFunc) *AppCategoryGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AppCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, ent.OpQueryGroupBy)
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppCategoryQuery, *AppCategoryGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AppCategoryGroupBy) sqlScan(ctx context.Context, root *AppCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppCategorySelect is the builder for selecting fields of AppCategory entities.
type AppCategorySelect struct {
	*AppCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AppCategorySelect) Aggregate(fns ...AggregateFunc) *AppCategorySelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AppCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, ent.OpQuerySelect)
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppCategoryQuery, *AppCategorySelect](ctx, acs.AppCategoryQuery, acs, acs.inters, v)
}

func (acs *AppCategorySelect) sqlScan(ctx context.Context, root *AppCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
