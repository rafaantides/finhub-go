// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"finhub-go/internal/ent/category"
	"finhub-go/internal/ent/debt"
	"finhub-go/internal/ent/invoice"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DebtQuery is the builder for querying Debt entities.
type DebtQuery struct {
	config
	ctx          *QueryContext
	order        []debt.OrderOption
	inters       []Interceptor
	predicates   []predicate.Debt
	withInvoice  *InvoiceQuery
	withCategory *CategoryQuery
	withStatus   *PaymentStatusQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DebtQuery builder.
func (dq *DebtQuery) Where(ps ...predicate.Debt) *DebtQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DebtQuery) Limit(limit int) *DebtQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DebtQuery) Offset(offset int) *DebtQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DebtQuery) Unique(unique bool) *DebtQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DebtQuery) Order(o ...debt.OrderOption) *DebtQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (dq *DebtQuery) QueryInvoice() *InvoiceQuery {
	query := (&InvoiceClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, selector),
			sqlgraph.To(invoice.Table, invoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, debt.InvoiceTable, debt.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (dq *DebtQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, debt.CategoryTable, debt.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatus chains the current query on the "status" edge.
func (dq *DebtQuery) QueryStatus() *PaymentStatusQuery {
	query := (&PaymentStatusClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(debt.Table, debt.FieldID, selector),
			sqlgraph.To(paymentstatus.Table, paymentstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, debt.StatusTable, debt.StatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Debt entity from the query.
// Returns a *NotFoundError when no Debt was found.
func (dq *DebtQuery) First(ctx context.Context) (*Debt, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{debt.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DebtQuery) FirstX(ctx context.Context) *Debt {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Debt ID from the query.
// Returns a *NotFoundError when no Debt ID was found.
func (dq *DebtQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{debt.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DebtQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Debt entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Debt entity is found.
// Returns a *NotFoundError when no Debt entities are found.
func (dq *DebtQuery) Only(ctx context.Context) (*Debt, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{debt.Label}
	default:
		return nil, &NotSingularError{debt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DebtQuery) OnlyX(ctx context.Context) *Debt {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Debt ID in the query.
// Returns a *NotSingularError when more than one Debt ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DebtQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{debt.Label}
	default:
		err = &NotSingularError{debt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DebtQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Debts.
func (dq *DebtQuery) All(ctx context.Context) ([]*Debt, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Debt, *DebtQuery]()
	return withInterceptors[[]*Debt](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DebtQuery) AllX(ctx context.Context) []*Debt {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Debt IDs.
func (dq *DebtQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(debt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DebtQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DebtQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DebtQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DebtQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DebtQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DebtQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DebtQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DebtQuery) Clone() *DebtQuery {
	if dq == nil {
		return nil
	}
	return &DebtQuery{
		config:       dq.config,
		ctx:          dq.ctx.Clone(),
		order:        append([]debt.OrderOption{}, dq.order...),
		inters:       append([]Interceptor{}, dq.inters...),
		predicates:   append([]predicate.Debt{}, dq.predicates...),
		withInvoice:  dq.withInvoice.Clone(),
		withCategory: dq.withCategory.Clone(),
		withStatus:   dq.withStatus.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DebtQuery) WithInvoice(opts ...func(*InvoiceQuery)) *DebtQuery {
	query := (&InvoiceClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withInvoice = query
	return dq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DebtQuery) WithCategory(opts ...func(*CategoryQuery)) *DebtQuery {
	query := (&CategoryClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withCategory = query
	return dq
}

// WithStatus tells the query-builder to eager-load the nodes that are connected to
// the "status" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DebtQuery) WithStatus(opts ...func(*PaymentStatusQuery)) *DebtQuery {
	query := (&PaymentStatusClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withStatus = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Debt.Query().
//		GroupBy(debt.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DebtQuery) GroupBy(field string, fields ...string) *DebtGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DebtGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = debt.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Debt.Query().
//		Select(debt.FieldCreatedAt).
//		Scan(ctx, &v)
func (dq *DebtQuery) Select(fields ...string) *DebtSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DebtSelect{DebtQuery: dq}
	sbuild.label = debt.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DebtSelect configured with the given aggregations.
func (dq *DebtQuery) Aggregate(fns ...AggregateFunc) *DebtSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DebtQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !debt.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DebtQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Debt, error) {
	var (
		nodes       = []*Debt{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withInvoice != nil,
			dq.withCategory != nil,
			dq.withStatus != nil,
		}
	)
	if dq.withInvoice != nil || dq.withCategory != nil || dq.withStatus != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, debt.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Debt).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Debt{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withInvoice; query != nil {
		if err := dq.loadInvoice(ctx, query, nodes, nil,
			func(n *Debt, e *Invoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withCategory; query != nil {
		if err := dq.loadCategory(ctx, query, nodes, nil,
			func(n *Debt, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withStatus; query != nil {
		if err := dq.loadStatus(ctx, query, nodes, nil,
			func(n *Debt, e *PaymentStatus) { n.Edges.Status = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DebtQuery) loadInvoice(ctx context.Context, query *InvoiceQuery, nodes []*Debt, init func(*Debt), assign func(*Debt, *Invoice)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Debt)
	for i := range nodes {
		if nodes[i].invoice_id == nil {
			continue
		}
		fk := *nodes[i].invoice_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(invoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DebtQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*Debt, init func(*Debt), assign func(*Debt, *Category)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Debt)
	for i := range nodes {
		if nodes[i].category_id == nil {
			continue
		}
		fk := *nodes[i].category_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DebtQuery) loadStatus(ctx context.Context, query *PaymentStatusQuery, nodes []*Debt, init func(*Debt), assign func(*Debt, *PaymentStatus)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Debt)
	for i := range nodes {
		if nodes[i].status_id == nil {
			continue
		}
		fk := *nodes[i].status_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(paymentstatus.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "status_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DebtQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DebtQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(debt.Table, debt.Columns, sqlgraph.NewFieldSpec(debt.FieldID, field.TypeUUID))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, debt.FieldID)
		for i := range fields {
			if fields[i] != debt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DebtQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(debt.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = debt.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DebtGroupBy is the group-by builder for Debt entities.
type DebtGroupBy struct {
	selector
	build *DebtQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DebtGroupBy) Aggregate(fns ...AggregateFunc) *DebtGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DebtGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebtQuery, *DebtGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DebtGroupBy) sqlScan(ctx context.Context, root *DebtQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DebtSelect is the builder for selecting fields of Debt entities.
type DebtSelect struct {
	*DebtQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DebtSelect) Aggregate(fns ...AggregateFunc) *DebtSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DebtSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DebtQuery, *DebtSelect](ctx, ds.DebtQuery, ds, ds.inters, v)
}

func (ds *DebtSelect) sqlScan(ctx context.Context, root *DebtQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
