// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/DeluxeOwl/kala-go/ent/migrate"

	"github.com/DeluxeOwl/kala-go/ent/permission"
	"github.com/DeluxeOwl/kala-go/ent/relation"
	"github.com/DeluxeOwl/kala-go/ent/subject"
	"github.com/DeluxeOwl/kala-go/ent/tuple"
	"github.com/DeluxeOwl/kala-go/ent/typeconfig"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Relation is the client for interacting with the Relation builders.
	Relation *RelationClient
	// Subject is the client for interacting with the Subject builders.
	Subject *SubjectClient
	// Tuple is the client for interacting with the Tuple builders.
	Tuple *TupleClient
	// TypeConfig is the client for interacting with the TypeConfig builders.
	TypeConfig *TypeConfigClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Permission = NewPermissionClient(c.config)
	c.Relation = NewRelationClient(c.config)
	c.Subject = NewSubjectClient(c.config)
	c.Tuple = NewTupleClient(c.config)
	c.TypeConfig = NewTypeConfigClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Permission: NewPermissionClient(cfg),
		Relation:   NewRelationClient(cfg),
		Subject:    NewSubjectClient(cfg),
		Tuple:      NewTupleClient(cfg),
		TypeConfig: NewTypeConfigClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Permission: NewPermissionClient(cfg),
		Relation:   NewRelationClient(cfg),
		Subject:    NewSubjectClient(cfg),
		Tuple:      NewTupleClient(cfg),
		TypeConfig: NewTypeConfigClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Permission.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Permission.Use(hooks...)
	c.Relation.Use(hooks...)
	c.Subject.Use(hooks...)
	c.Tuple.Use(hooks...)
	c.TypeConfig.Use(hooks...)
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Create returns a create builder for Permission.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id int) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PermissionClient) DeleteOneID(id int) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{
		config: c.config,
	}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id int) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id int) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTypeconfig queries the typeconfig edge of a Permission.
func (c *PermissionClient) QueryTypeconfig(pe *Permission) *TypeConfigQuery {
	query := &TypeConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(typeconfig.Table, typeconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, permission.TypeconfigTable, permission.TypeconfigColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelations queries the relations edge of a Permission.
func (c *PermissionClient) QueryRelations(pe *Permission) *RelationQuery {
	query := &RelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, permission.RelationsTable, permission.RelationsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// RelationClient is a client for the Relation schema.
type RelationClient struct {
	config
}

// NewRelationClient returns a client for the Relation from the given config.
func NewRelationClient(c config) *RelationClient {
	return &RelationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `relation.Hooks(f(g(h())))`.
func (c *RelationClient) Use(hooks ...Hook) {
	c.hooks.Relation = append(c.hooks.Relation, hooks...)
}

// Create returns a create builder for Relation.
func (c *RelationClient) Create() *RelationCreate {
	mutation := newRelationMutation(c.config, OpCreate)
	return &RelationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Relation entities.
func (c *RelationClient) CreateBulk(builders ...*RelationCreate) *RelationCreateBulk {
	return &RelationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Relation.
func (c *RelationClient) Update() *RelationUpdate {
	mutation := newRelationMutation(c.config, OpUpdate)
	return &RelationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RelationClient) UpdateOne(r *Relation) *RelationUpdateOne {
	mutation := newRelationMutation(c.config, OpUpdateOne, withRelation(r))
	return &RelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RelationClient) UpdateOneID(id int) *RelationUpdateOne {
	mutation := newRelationMutation(c.config, OpUpdateOne, withRelationID(id))
	return &RelationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Relation.
func (c *RelationClient) Delete() *RelationDelete {
	mutation := newRelationMutation(c.config, OpDelete)
	return &RelationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RelationClient) DeleteOne(r *Relation) *RelationDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RelationClient) DeleteOneID(id int) *RelationDeleteOne {
	builder := c.Delete().Where(relation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RelationDeleteOne{builder}
}

// Query returns a query builder for Relation.
func (c *RelationClient) Query() *RelationQuery {
	return &RelationQuery{
		config: c.config,
	}
}

// Get returns a Relation entity by its id.
func (c *RelationClient) Get(ctx context.Context, id int) (*Relation, error) {
	return c.Query().Where(relation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RelationClient) GetX(ctx context.Context, id int) *Relation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTypeconfig queries the typeconfig edge of a Relation.
func (c *RelationClient) QueryTypeconfig(r *Relation) *TypeConfigQuery {
	query := &TypeConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(typeconfig.Table, typeconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, relation.TypeconfigTable, relation.TypeconfigColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelTypeconfigs queries the rel_typeconfigs edge of a Relation.
func (c *RelationClient) QueryRelTypeconfigs(r *Relation) *TypeConfigQuery {
	query := &TypeConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(typeconfig.Table, typeconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, relation.RelTypeconfigsTable, relation.RelTypeconfigsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a Relation.
func (c *RelationClient) QueryPermissions(r *Relation) *PermissionQuery {
	query := &PermissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, relation.PermissionsTable, relation.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTuples queries the tuples edge of a Relation.
func (c *RelationClient) QueryTuples(r *Relation) *TupleQuery {
	query := &TupleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relation.Table, relation.FieldID, id),
			sqlgraph.To(tuple.Table, tuple.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, relation.TuplesTable, relation.TuplesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RelationClient) Hooks() []Hook {
	return c.hooks.Relation
}

// SubjectClient is a client for the Subject schema.
type SubjectClient struct {
	config
}

// NewSubjectClient returns a client for the Subject from the given config.
func NewSubjectClient(c config) *SubjectClient {
	return &SubjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subject.Hooks(f(g(h())))`.
func (c *SubjectClient) Use(hooks ...Hook) {
	c.hooks.Subject = append(c.hooks.Subject, hooks...)
}

// Create returns a create builder for Subject.
func (c *SubjectClient) Create() *SubjectCreate {
	mutation := newSubjectMutation(c.config, OpCreate)
	return &SubjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Subject entities.
func (c *SubjectClient) CreateBulk(builders ...*SubjectCreate) *SubjectCreateBulk {
	return &SubjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Subject.
func (c *SubjectClient) Update() *SubjectUpdate {
	mutation := newSubjectMutation(c.config, OpUpdate)
	return &SubjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubjectClient) UpdateOne(s *Subject) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubject(s))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubjectClient) UpdateOneID(id int) *SubjectUpdateOne {
	mutation := newSubjectMutation(c.config, OpUpdateOne, withSubjectID(id))
	return &SubjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Subject.
func (c *SubjectClient) Delete() *SubjectDelete {
	mutation := newSubjectMutation(c.config, OpDelete)
	return &SubjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubjectClient) DeleteOne(s *Subject) *SubjectDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubjectClient) DeleteOneID(id int) *SubjectDeleteOne {
	builder := c.Delete().Where(subject.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubjectDeleteOne{builder}
}

// Query returns a query builder for Subject.
func (c *SubjectClient) Query() *SubjectQuery {
	return &SubjectQuery{
		config: c.config,
	}
}

// Get returns a Subject entity by its id.
func (c *SubjectClient) Get(ctx context.Context, id int) (*Subject, error) {
	return c.Query().Where(subject.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubjectClient) GetX(ctx context.Context, id int) *Subject {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryType queries the type edge of a Subject.
func (c *SubjectClient) QueryType(s *Subject) *TypeConfigQuery {
	query := &TypeConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subject.Table, subject.FieldID, id),
			sqlgraph.To(typeconfig.Table, typeconfig.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subject.TypeTable, subject.TypeColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubjectClient) Hooks() []Hook {
	return c.hooks.Subject
}

// TupleClient is a client for the Tuple schema.
type TupleClient struct {
	config
}

// NewTupleClient returns a client for the Tuple from the given config.
func NewTupleClient(c config) *TupleClient {
	return &TupleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tuple.Hooks(f(g(h())))`.
func (c *TupleClient) Use(hooks ...Hook) {
	c.hooks.Tuple = append(c.hooks.Tuple, hooks...)
}

// Create returns a create builder for Tuple.
func (c *TupleClient) Create() *TupleCreate {
	mutation := newTupleMutation(c.config, OpCreate)
	return &TupleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tuple entities.
func (c *TupleClient) CreateBulk(builders ...*TupleCreate) *TupleCreateBulk {
	return &TupleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tuple.
func (c *TupleClient) Update() *TupleUpdate {
	mutation := newTupleMutation(c.config, OpUpdate)
	return &TupleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TupleClient) UpdateOne(t *Tuple) *TupleUpdateOne {
	mutation := newTupleMutation(c.config, OpUpdateOne, withTuple(t))
	return &TupleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TupleClient) UpdateOneID(id int) *TupleUpdateOne {
	mutation := newTupleMutation(c.config, OpUpdateOne, withTupleID(id))
	return &TupleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tuple.
func (c *TupleClient) Delete() *TupleDelete {
	mutation := newTupleMutation(c.config, OpDelete)
	return &TupleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TupleClient) DeleteOne(t *Tuple) *TupleDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TupleClient) DeleteOneID(id int) *TupleDeleteOne {
	builder := c.Delete().Where(tuple.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TupleDeleteOne{builder}
}

// Query returns a query builder for Tuple.
func (c *TupleClient) Query() *TupleQuery {
	return &TupleQuery{
		config: c.config,
	}
}

// Get returns a Tuple entity by its id.
func (c *TupleClient) Get(ctx context.Context, id int) (*Tuple, error) {
	return c.Query().Where(tuple.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TupleClient) GetX(ctx context.Context, id int) *Tuple {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubject queries the subject edge of a Tuple.
func (c *TupleClient) QuerySubject(t *Tuple) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tuple.Table, tuple.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tuple.SubjectTable, tuple.SubjectColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelation queries the relation edge of a Tuple.
func (c *TupleClient) QueryRelation(t *Tuple) *RelationQuery {
	query := &RelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tuple.Table, tuple.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tuple.RelationTable, tuple.RelationColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryResource queries the resource edge of a Tuple.
func (c *TupleClient) QueryResource(t *Tuple) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tuple.Table, tuple.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tuple.ResourceTable, tuple.ResourceColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TupleClient) Hooks() []Hook {
	return c.hooks.Tuple
}

// TypeConfigClient is a client for the TypeConfig schema.
type TypeConfigClient struct {
	config
}

// NewTypeConfigClient returns a client for the TypeConfig from the given config.
func NewTypeConfigClient(c config) *TypeConfigClient {
	return &TypeConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `typeconfig.Hooks(f(g(h())))`.
func (c *TypeConfigClient) Use(hooks ...Hook) {
	c.hooks.TypeConfig = append(c.hooks.TypeConfig, hooks...)
}

// Create returns a create builder for TypeConfig.
func (c *TypeConfigClient) Create() *TypeConfigCreate {
	mutation := newTypeConfigMutation(c.config, OpCreate)
	return &TypeConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TypeConfig entities.
func (c *TypeConfigClient) CreateBulk(builders ...*TypeConfigCreate) *TypeConfigCreateBulk {
	return &TypeConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TypeConfig.
func (c *TypeConfigClient) Update() *TypeConfigUpdate {
	mutation := newTypeConfigMutation(c.config, OpUpdate)
	return &TypeConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TypeConfigClient) UpdateOne(tc *TypeConfig) *TypeConfigUpdateOne {
	mutation := newTypeConfigMutation(c.config, OpUpdateOne, withTypeConfig(tc))
	return &TypeConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TypeConfigClient) UpdateOneID(id int) *TypeConfigUpdateOne {
	mutation := newTypeConfigMutation(c.config, OpUpdateOne, withTypeConfigID(id))
	return &TypeConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TypeConfig.
func (c *TypeConfigClient) Delete() *TypeConfigDelete {
	mutation := newTypeConfigMutation(c.config, OpDelete)
	return &TypeConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TypeConfigClient) DeleteOne(tc *TypeConfig) *TypeConfigDeleteOne {
	return c.DeleteOneID(tc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TypeConfigClient) DeleteOneID(id int) *TypeConfigDeleteOne {
	builder := c.Delete().Where(typeconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TypeConfigDeleteOne{builder}
}

// Query returns a query builder for TypeConfig.
func (c *TypeConfigClient) Query() *TypeConfigQuery {
	return &TypeConfigQuery{
		config: c.config,
	}
}

// Get returns a TypeConfig entity by its id.
func (c *TypeConfigClient) Get(ctx context.Context, id int) (*TypeConfig, error) {
	return c.Query().Where(typeconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TypeConfigClient) GetX(ctx context.Context, id int) *TypeConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRelations queries the relations edge of a TypeConfig.
func (c *TypeConfigClient) QueryRelations(tc *TypeConfig) *RelationQuery {
	query := &RelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(typeconfig.Table, typeconfig.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, typeconfig.RelationsTable, typeconfig.RelationsColumn),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPermissions queries the permissions edge of a TypeConfig.
func (c *TypeConfigClient) QueryPermissions(tc *TypeConfig) *PermissionQuery {
	query := &PermissionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(typeconfig.Table, typeconfig.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, typeconfig.PermissionsTable, typeconfig.PermissionsColumn),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubjects queries the subjects edge of a TypeConfig.
func (c *TypeConfigClient) QuerySubjects(tc *TypeConfig) *SubjectQuery {
	query := &SubjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(typeconfig.Table, typeconfig.FieldID, id),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, typeconfig.SubjectsTable, typeconfig.SubjectsColumn),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRelTypeconfigs queries the rel_typeconfigs edge of a TypeConfig.
func (c *TypeConfigClient) QueryRelTypeconfigs(tc *TypeConfig) *RelationQuery {
	query := &RelationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(typeconfig.Table, typeconfig.FieldID, id),
			sqlgraph.To(relation.Table, relation.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, typeconfig.RelTypeconfigsTable, typeconfig.RelTypeconfigsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(tc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TypeConfigClient) Hooks() []Hook {
	return c.hooks.TypeConfig
}
