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
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/deviceinfo"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/app/sephirah/internal/data/internal/ent/usersession"
	"github.com/tuihub/librarian/internal/model"
)

// DeviceInfoQuery is the builder for querying DeviceInfo entities.
type DeviceInfoQuery struct {
	config
	ctx             *QueryContext
	order           []deviceinfo.OrderOption
	inters          []Interceptor
	predicates      []predicate.DeviceInfo
	withUserSession *UserSessionQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceInfoQuery builder.
func (diq *DeviceInfoQuery) Where(ps ...predicate.DeviceInfo) *DeviceInfoQuery {
	diq.predicates = append(diq.predicates, ps...)
	return diq
}

// Limit the number of records to be returned by this query.
func (diq *DeviceInfoQuery) Limit(limit int) *DeviceInfoQuery {
	diq.ctx.Limit = &limit
	return diq
}

// Offset to start from.
func (diq *DeviceInfoQuery) Offset(offset int) *DeviceInfoQuery {
	diq.ctx.Offset = &offset
	return diq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (diq *DeviceInfoQuery) Unique(unique bool) *DeviceInfoQuery {
	diq.ctx.Unique = &unique
	return diq
}

// Order specifies how the records should be ordered.
func (diq *DeviceInfoQuery) Order(o ...deviceinfo.OrderOption) *DeviceInfoQuery {
	diq.order = append(diq.order, o...)
	return diq
}

// QueryUserSession chains the current query on the "user_session" edge.
func (diq *DeviceInfoQuery) QueryUserSession() *UserSessionQuery {
	query := (&UserSessionClient{config: diq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := diq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := diq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deviceinfo.Table, deviceinfo.FieldID, selector),
			sqlgraph.To(usersession.Table, usersession.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, deviceinfo.UserSessionTable, deviceinfo.UserSessionColumn),
		)
		fromU = sqlgraph.SetNeighbors(diq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeviceInfo entity from the query.
// Returns a *NotFoundError when no DeviceInfo was found.
func (diq *DeviceInfoQuery) First(ctx context.Context) (*DeviceInfo, error) {
	nodes, err := diq.Limit(1).All(setContextOp(ctx, diq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deviceinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (diq *DeviceInfoQuery) FirstX(ctx context.Context) *DeviceInfo {
	node, err := diq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceInfo ID from the query.
// Returns a *NotFoundError when no DeviceInfo ID was found.
func (diq *DeviceInfoQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = diq.Limit(1).IDs(setContextOp(ctx, diq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deviceinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (diq *DeviceInfoQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := diq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceInfo entity is found.
// Returns a *NotFoundError when no DeviceInfo entities are found.
func (diq *DeviceInfoQuery) Only(ctx context.Context) (*DeviceInfo, error) {
	nodes, err := diq.Limit(2).All(setContextOp(ctx, diq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deviceinfo.Label}
	default:
		return nil, &NotSingularError{deviceinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (diq *DeviceInfoQuery) OnlyX(ctx context.Context) *DeviceInfo {
	node, err := diq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceInfo ID in the query.
// Returns a *NotSingularError when more than one DeviceInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (diq *DeviceInfoQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = diq.Limit(2).IDs(setContextOp(ctx, diq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deviceinfo.Label}
	default:
		err = &NotSingularError{deviceinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (diq *DeviceInfoQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := diq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceInfos.
func (diq *DeviceInfoQuery) All(ctx context.Context) ([]*DeviceInfo, error) {
	ctx = setContextOp(ctx, diq.ctx, "All")
	if err := diq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DeviceInfo, *DeviceInfoQuery]()
	return withInterceptors[[]*DeviceInfo](ctx, diq, qr, diq.inters)
}

// AllX is like All, but panics if an error occurs.
func (diq *DeviceInfoQuery) AllX(ctx context.Context) []*DeviceInfo {
	nodes, err := diq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceInfo IDs.
func (diq *DeviceInfoQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if diq.ctx.Unique == nil && diq.path != nil {
		diq.Unique(true)
	}
	ctx = setContextOp(ctx, diq.ctx, "IDs")
	if err = diq.Select(deviceinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (diq *DeviceInfoQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := diq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (diq *DeviceInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, diq.ctx, "Count")
	if err := diq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, diq, querierCount[*DeviceInfoQuery](), diq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (diq *DeviceInfoQuery) CountX(ctx context.Context) int {
	count, err := diq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (diq *DeviceInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, diq.ctx, "Exist")
	switch _, err := diq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (diq *DeviceInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := diq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (diq *DeviceInfoQuery) Clone() *DeviceInfoQuery {
	if diq == nil {
		return nil
	}
	return &DeviceInfoQuery{
		config:          diq.config,
		ctx:             diq.ctx.Clone(),
		order:           append([]deviceinfo.OrderOption{}, diq.order...),
		inters:          append([]Interceptor{}, diq.inters...),
		predicates:      append([]predicate.DeviceInfo{}, diq.predicates...),
		withUserSession: diq.withUserSession.Clone(),
		// clone intermediate query.
		sql:  diq.sql.Clone(),
		path: diq.path,
	}
}

// WithUserSession tells the query-builder to eager-load the nodes that are connected to
// the "user_session" edge. The optional arguments are used to configure the query builder of the edge.
func (diq *DeviceInfoQuery) WithUserSession(opts ...func(*UserSessionQuery)) *DeviceInfoQuery {
	query := (&UserSessionClient{config: diq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	diq.withUserSession = query
	return diq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceName string `json:"device_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceInfo.Query().
//		GroupBy(deviceinfo.FieldDeviceName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (diq *DeviceInfoQuery) GroupBy(field string, fields ...string) *DeviceInfoGroupBy {
	diq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeviceInfoGroupBy{build: diq}
	grbuild.flds = &diq.ctx.Fields
	grbuild.label = deviceinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceName string `json:"device_name,omitempty"`
//	}
//
//	client.DeviceInfo.Query().
//		Select(deviceinfo.FieldDeviceName).
//		Scan(ctx, &v)
func (diq *DeviceInfoQuery) Select(fields ...string) *DeviceInfoSelect {
	diq.ctx.Fields = append(diq.ctx.Fields, fields...)
	sbuild := &DeviceInfoSelect{DeviceInfoQuery: diq}
	sbuild.label = deviceinfo.Label
	sbuild.flds, sbuild.scan = &diq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeviceInfoSelect configured with the given aggregations.
func (diq *DeviceInfoQuery) Aggregate(fns ...AggregateFunc) *DeviceInfoSelect {
	return diq.Select().Aggregate(fns...)
}

func (diq *DeviceInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range diq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, diq); err != nil {
				return err
			}
		}
	}
	for _, f := range diq.ctx.Fields {
		if !deviceinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if diq.path != nil {
		prev, err := diq.path(ctx)
		if err != nil {
			return err
		}
		diq.sql = prev
	}
	return nil
}

func (diq *DeviceInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceInfo, error) {
	var (
		nodes       = []*DeviceInfo{}
		withFKs     = diq.withFKs
		_spec       = diq.querySpec()
		loadedTypes = [1]bool{
			diq.withUserSession != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DeviceInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DeviceInfo{config: diq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, diq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := diq.withUserSession; query != nil {
		if err := diq.loadUserSession(ctx, query, nodes,
			func(n *DeviceInfo) { n.Edges.UserSession = []*UserSession{} },
			func(n *DeviceInfo, e *UserSession) { n.Edges.UserSession = append(n.Edges.UserSession, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (diq *DeviceInfoQuery) loadUserSession(ctx context.Context, query *UserSessionQuery, nodes []*DeviceInfo, init func(*DeviceInfo), assign func(*DeviceInfo, *UserSession)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[model.InternalID]*DeviceInfo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserSession(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(deviceinfo.UserSessionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.device_info_user_session
		if fk == nil {
			return fmt.Errorf(`foreign-key "device_info_user_session" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "device_info_user_session" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (diq *DeviceInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := diq.querySpec()
	_spec.Node.Columns = diq.ctx.Fields
	if len(diq.ctx.Fields) > 0 {
		_spec.Unique = diq.ctx.Unique != nil && *diq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, diq.driver, _spec)
}

func (diq *DeviceInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(deviceinfo.Table, deviceinfo.Columns, sqlgraph.NewFieldSpec(deviceinfo.FieldID, field.TypeInt64))
	_spec.From = diq.sql
	if unique := diq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if diq.path != nil {
		_spec.Unique = true
	}
	if fields := diq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinfo.FieldID)
		for i := range fields {
			if fields[i] != deviceinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := diq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := diq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := diq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := diq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (diq *DeviceInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(diq.driver.Dialect())
	t1 := builder.Table(deviceinfo.Table)
	columns := diq.ctx.Fields
	if len(columns) == 0 {
		columns = deviceinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if diq.sql != nil {
		selector = diq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if diq.ctx.Unique != nil && *diq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range diq.predicates {
		p(selector)
	}
	for _, p := range diq.order {
		p(selector)
	}
	if offset := diq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := diq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceInfoGroupBy is the group-by builder for DeviceInfo entities.
type DeviceInfoGroupBy struct {
	selector
	build *DeviceInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (digb *DeviceInfoGroupBy) Aggregate(fns ...AggregateFunc) *DeviceInfoGroupBy {
	digb.fns = append(digb.fns, fns...)
	return digb
}

// Scan applies the selector query and scans the result into the given value.
func (digb *DeviceInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, digb.build.ctx, "GroupBy")
	if err := digb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceInfoQuery, *DeviceInfoGroupBy](ctx, digb.build, digb, digb.build.inters, v)
}

func (digb *DeviceInfoGroupBy) sqlScan(ctx context.Context, root *DeviceInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(digb.fns))
	for _, fn := range digb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*digb.flds)+len(digb.fns))
		for _, f := range *digb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*digb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := digb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeviceInfoSelect is the builder for selecting fields of DeviceInfo entities.
type DeviceInfoSelect struct {
	*DeviceInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dis *DeviceInfoSelect) Aggregate(fns ...AggregateFunc) *DeviceInfoSelect {
	dis.fns = append(dis.fns, fns...)
	return dis
}

// Scan applies the selector query and scans the result into the given value.
func (dis *DeviceInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dis.ctx, "Select")
	if err := dis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceInfoQuery, *DeviceInfoSelect](ctx, dis.DeviceInfoQuery, dis, dis.inters, v)
}

func (dis *DeviceInfoSelect) sqlScan(ctx context.Context, root *DeviceInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dis.fns))
	for _, fn := range dis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
