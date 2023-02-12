// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nanaminakano/snippeter/ent/predicate"
	"github.com/nanaminakano/snippeter/ent/snippet"
)

// SnippetQuery is the builder for querying Snippet entities.
type SnippetQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Snippet
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SnippetQuery builder.
func (sq *SnippetQuery) Where(ps ...predicate.Snippet) *SnippetQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SnippetQuery) Limit(limit int) *SnippetQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SnippetQuery) Offset(offset int) *SnippetQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SnippetQuery) Unique(unique bool) *SnippetQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SnippetQuery) Order(o ...OrderFunc) *SnippetQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// First returns the first Snippet entity from the query.
// Returns a *NotFoundError when no Snippet was found.
func (sq *SnippetQuery) First(ctx context.Context) (*Snippet, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{snippet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SnippetQuery) FirstX(ctx context.Context) *Snippet {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Snippet ID from the query.
// Returns a *NotFoundError when no Snippet ID was found.
func (sq *SnippetQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{snippet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SnippetQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Snippet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Snippet entity is found.
// Returns a *NotFoundError when no Snippet entities are found.
func (sq *SnippetQuery) Only(ctx context.Context) (*Snippet, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{snippet.Label}
	default:
		return nil, &NotSingularError{snippet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SnippetQuery) OnlyX(ctx context.Context) *Snippet {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Snippet ID in the query.
// Returns a *NotSingularError when more than one Snippet ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SnippetQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{snippet.Label}
	default:
		err = &NotSingularError{snippet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SnippetQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Snippets.
func (sq *SnippetQuery) All(ctx context.Context) ([]*Snippet, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Snippet, *SnippetQuery]()
	return withInterceptors[[]*Snippet](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SnippetQuery) AllX(ctx context.Context) []*Snippet {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Snippet IDs.
func (sq *SnippetQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err := sq.Select(snippet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SnippetQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SnippetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SnippetQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SnippetQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SnippetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SnippetQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SnippetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SnippetQuery) Clone() *SnippetQuery {
	if sq == nil {
		return nil
	}
	return &SnippetQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Snippet{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Language string `json:"language,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Snippet.Query().
//		GroupBy(snippet.FieldLanguage).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SnippetQuery) GroupBy(field string, fields ...string) *SnippetGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SnippetGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = snippet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Language string `json:"language,omitempty"`
//	}
//
//	client.Snippet.Query().
//		Select(snippet.FieldLanguage).
//		Scan(ctx, &v)
func (sq *SnippetQuery) Select(fields ...string) *SnippetSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SnippetSelect{SnippetQuery: sq}
	sbuild.label = snippet.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SnippetSelect configured with the given aggregations.
func (sq *SnippetQuery) Aggregate(fns ...AggregateFunc) *SnippetSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SnippetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !snippet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SnippetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Snippet, error) {
	var (
		nodes = []*Snippet{}
		_spec = sq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Snippet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Snippet{config: sq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sq *SnippetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SnippetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   snippet.Table,
			Columns: snippet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: snippet.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, snippet.FieldID)
		for i := range fields {
			if fields[i] != snippet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SnippetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(snippet.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = snippet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SnippetGroupBy is the group-by builder for Snippet entities.
type SnippetGroupBy struct {
	selector
	build *SnippetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SnippetGroupBy) Aggregate(fns ...AggregateFunc) *SnippetGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SnippetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SnippetQuery, *SnippetGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SnippetGroupBy) sqlScan(ctx context.Context, root *SnippetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SnippetSelect is the builder for selecting fields of Snippet entities.
type SnippetSelect struct {
	*SnippetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SnippetSelect) Aggregate(fns ...AggregateFunc) *SnippetSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SnippetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SnippetQuery, *SnippetSelect](ctx, ss.SnippetQuery, ss, ss.inters, v)
}

func (ss *SnippetSelect) sqlScan(ctx context.Context, root *SnippetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
