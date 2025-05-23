// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
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

// InvoiceQuery is the builder for querying Invoice entities.
type InvoiceQuery struct {
	config
	ctx        *QueryContext
	order      []invoice.OrderOption
	inters     []Interceptor
	predicates []predicate.Invoice
	withStatus *PaymentStatusQuery
	withDebts  *DebtQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InvoiceQuery builder.
func (iq *InvoiceQuery) Where(ps ...predicate.Invoice) *InvoiceQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InvoiceQuery) Limit(limit int) *InvoiceQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InvoiceQuery) Offset(offset int) *InvoiceQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InvoiceQuery) Unique(unique bool) *InvoiceQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InvoiceQuery) Order(o ...invoice.OrderOption) *InvoiceQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryStatus chains the current query on the "status" edge.
func (iq *InvoiceQuery) QueryStatus() *PaymentStatusQuery {
	query := (&PaymentStatusClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(paymentstatus.Table, paymentstatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, invoice.StatusTable, invoice.StatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDebts chains the current query on the "debts" edge.
func (iq *InvoiceQuery) QueryDebts() *DebtQuery {
	query := (&DebtClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(invoice.Table, invoice.FieldID, selector),
			sqlgraph.To(debt.Table, debt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, invoice.DebtsTable, invoice.DebtsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Invoice entity from the query.
// Returns a *NotFoundError when no Invoice was found.
func (iq *InvoiceQuery) First(ctx context.Context) (*Invoice, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{invoice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InvoiceQuery) FirstX(ctx context.Context) *Invoice {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Invoice ID from the query.
// Returns a *NotFoundError when no Invoice ID was found.
func (iq *InvoiceQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{invoice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InvoiceQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Invoice entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Invoice entity is found.
// Returns a *NotFoundError when no Invoice entities are found.
func (iq *InvoiceQuery) Only(ctx context.Context) (*Invoice, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{invoice.Label}
	default:
		return nil, &NotSingularError{invoice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InvoiceQuery) OnlyX(ctx context.Context) *Invoice {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Invoice ID in the query.
// Returns a *NotSingularError when more than one Invoice ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InvoiceQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{invoice.Label}
	default:
		err = &NotSingularError{invoice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InvoiceQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Invoices.
func (iq *InvoiceQuery) All(ctx context.Context) ([]*Invoice, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Invoice, *InvoiceQuery]()
	return withInterceptors[[]*Invoice](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InvoiceQuery) AllX(ctx context.Context) []*Invoice {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Invoice IDs.
func (iq *InvoiceQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryIDs)
	if err = iq.Select(invoice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InvoiceQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InvoiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InvoiceQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InvoiceQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InvoiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InvoiceQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InvoiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InvoiceQuery) Clone() *InvoiceQuery {
	if iq == nil {
		return nil
	}
	return &InvoiceQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]invoice.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Invoice{}, iq.predicates...),
		withStatus: iq.withStatus.Clone(),
		withDebts:  iq.withDebts.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithStatus tells the query-builder to eager-load the nodes that are connected to
// the "status" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithStatus(opts ...func(*PaymentStatusQuery)) *InvoiceQuery {
	query := (&PaymentStatusClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withStatus = query
	return iq
}

// WithDebts tells the query-builder to eager-load the nodes that are connected to
// the "debts" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InvoiceQuery) WithDebts(opts ...func(*DebtQuery)) *InvoiceQuery {
	query := (&DebtClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withDebts = query
	return iq
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
//	client.Invoice.Query().
//		GroupBy(invoice.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InvoiceQuery) GroupBy(field string, fields ...string) *InvoiceGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InvoiceGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = invoice.Label
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
//	client.Invoice.Query().
//		Select(invoice.FieldCreatedAt).
//		Scan(ctx, &v)
func (iq *InvoiceQuery) Select(fields ...string) *InvoiceSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InvoiceSelect{InvoiceQuery: iq}
	sbuild.label = invoice.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InvoiceSelect configured with the given aggregations.
func (iq *InvoiceQuery) Aggregate(fns ...AggregateFunc) *InvoiceSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InvoiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !invoice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InvoiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Invoice, error) {
	var (
		nodes       = []*Invoice{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withStatus != nil,
			iq.withDebts != nil,
		}
	)
	if iq.withStatus != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Invoice).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Invoice{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withStatus; query != nil {
		if err := iq.loadStatus(ctx, query, nodes, nil,
			func(n *Invoice, e *PaymentStatus) { n.Edges.Status = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withDebts; query != nil {
		if err := iq.loadDebts(ctx, query, nodes,
			func(n *Invoice) { n.Edges.Debts = []*Debt{} },
			func(n *Invoice, e *Debt) { n.Edges.Debts = append(n.Edges.Debts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InvoiceQuery) loadStatus(ctx context.Context, query *PaymentStatusQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *PaymentStatus)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Invoice)
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
func (iq *InvoiceQuery) loadDebts(ctx context.Context, query *DebtQuery, nodes []*Invoice, init func(*Invoice), assign func(*Invoice, *Debt)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Invoice)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Debt(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(invoice.DebtsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.invoice_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "invoice_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "invoice_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *InvoiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InvoiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(invoice.Table, invoice.Columns, sqlgraph.NewFieldSpec(invoice.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invoice.FieldID)
		for i := range fields {
			if fields[i] != invoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InvoiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(invoice.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = invoice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InvoiceGroupBy is the group-by builder for Invoice entities.
type InvoiceGroupBy struct {
	selector
	build *InvoiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InvoiceGroupBy) Aggregate(fns ...AggregateFunc) *InvoiceGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InvoiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, ent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceQuery, *InvoiceGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InvoiceGroupBy) sqlScan(ctx context.Context, root *InvoiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InvoiceSelect is the builder for selecting fields of Invoice entities.
type InvoiceSelect struct {
	*InvoiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InvoiceSelect) Aggregate(fns ...AggregateFunc) *InvoiceSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InvoiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, ent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InvoiceQuery, *InvoiceSelect](ctx, is.InvoiceQuery, is, is.inters, v)
}

func (is *InvoiceSelect) sqlScan(ctx context.Context, root *InvoiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
