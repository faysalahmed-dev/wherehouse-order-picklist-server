// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/category"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/order"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/productitem"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/subcategory"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// ProductItem is the client for interacting with the ProductItem builders.
	ProductItem *ProductItemClient
	// SubCategory is the client for interacting with the SubCategory builders.
	SubCategory *SubCategoryClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Order = NewOrderClient(c.config)
	c.ProductItem = NewProductItemClient(c.config)
	c.SubCategory = NewSubCategoryClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Category:    NewCategoryClient(cfg),
		Order:       NewOrderClient(cfg),
		ProductItem: NewProductItemClient(cfg),
		SubCategory: NewSubCategoryClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:         ctx,
		config:      cfg,
		Category:    NewCategoryClient(cfg),
		Order:       NewOrderClient(cfg),
		ProductItem: NewProductItemClient(cfg),
		SubCategory: NewSubCategoryClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
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
	c.Category.Use(hooks...)
	c.Order.Use(hooks...)
	c.ProductItem.Use(hooks...)
	c.SubCategory.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Category.Intercept(interceptors...)
	c.Order.Intercept(interceptors...)
	c.ProductItem.Intercept(interceptors...)
	c.SubCategory.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CategoryMutation:
		return c.Category.mutate(ctx, m)
	case *OrderMutation:
		return c.Order.mutate(ctx, m)
	case *ProductItemMutation:
		return c.ProductItem.mutate(ctx, m)
	case *SubCategoryMutation:
		return c.SubCategory.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `category.Intercept(f(g(h())))`.
func (c *CategoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.Category = append(c.inters.Category, interceptors...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CategoryClient) MapCreateBulk(slice any, setFunc func(*CategoryCreate, int)) *CategoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CategoryCreateBulk{err: fmt.Errorf("calling to CategoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CategoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id uuid.UUID) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id uuid.UUID) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCategory},
		inters: c.Interceptors(),
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id uuid.UUID) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id uuid.UUID) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubCategories queries the sub_categories edge of a Category.
func (c *CategoryClient) QuerySubCategories(ca *Category) *SubCategoryQuery {
	query := (&SubCategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(subcategory.Table, subcategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.SubCategoriesTable, category.SubCategoriesColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Category.
func (c *CategoryClient) QueryUser(ca *Category) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, category.UserTable, category.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// Interceptors returns the client interceptors.
func (c *CategoryClient) Interceptors() []Interceptor {
	return c.inters.Category
}

func (c *CategoryClient) mutate(ctx context.Context, m *CategoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Category mutation op: %q", m.Op())
	}
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `order.Intercept(f(g(h())))`.
func (c *OrderClient) Intercept(interceptors ...Interceptor) {
	c.inters.Order = append(c.inters.Order, interceptors...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderClient) MapCreateBulk(slice any, setFunc func(*OrderCreate, int)) *OrderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderCreateBulk{err: fmt.Errorf("calling to OrderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id uuid.UUID) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id uuid.UUID) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrder},
		inters: c.Interceptors(),
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id uuid.UUID) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id uuid.UUID) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProduct queries the product edge of a Order.
func (c *OrderClient) QueryProduct(o *Order) *ProductItemQuery {
	query := (&ProductItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(productitem.Table, productitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, order.ProductTable, order.ProductColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a Order.
func (c *OrderClient) QueryUser(o *Order) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, order.UserTable, order.UserColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// Interceptors returns the client interceptors.
func (c *OrderClient) Interceptors() []Interceptor {
	return c.inters.Order
}

func (c *OrderClient) mutate(ctx context.Context, m *OrderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Order mutation op: %q", m.Op())
	}
}

// ProductItemClient is a client for the ProductItem schema.
type ProductItemClient struct {
	config
}

// NewProductItemClient returns a client for the ProductItem from the given config.
func NewProductItemClient(c config) *ProductItemClient {
	return &ProductItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `productitem.Hooks(f(g(h())))`.
func (c *ProductItemClient) Use(hooks ...Hook) {
	c.hooks.ProductItem = append(c.hooks.ProductItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `productitem.Intercept(f(g(h())))`.
func (c *ProductItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.ProductItem = append(c.inters.ProductItem, interceptors...)
}

// Create returns a builder for creating a ProductItem entity.
func (c *ProductItemClient) Create() *ProductItemCreate {
	mutation := newProductItemMutation(c.config, OpCreate)
	return &ProductItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ProductItem entities.
func (c *ProductItemClient) CreateBulk(builders ...*ProductItemCreate) *ProductItemCreateBulk {
	return &ProductItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ProductItemClient) MapCreateBulk(slice any, setFunc func(*ProductItemCreate, int)) *ProductItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ProductItemCreateBulk{err: fmt.Errorf("calling to ProductItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ProductItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ProductItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ProductItem.
func (c *ProductItemClient) Update() *ProductItemUpdate {
	mutation := newProductItemMutation(c.config, OpUpdate)
	return &ProductItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductItemClient) UpdateOne(pi *ProductItem) *ProductItemUpdateOne {
	mutation := newProductItemMutation(c.config, OpUpdateOne, withProductItem(pi))
	return &ProductItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductItemClient) UpdateOneID(id uuid.UUID) *ProductItemUpdateOne {
	mutation := newProductItemMutation(c.config, OpUpdateOne, withProductItemID(id))
	return &ProductItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ProductItem.
func (c *ProductItemClient) Delete() *ProductItemDelete {
	mutation := newProductItemMutation(c.config, OpDelete)
	return &ProductItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductItemClient) DeleteOne(pi *ProductItem) *ProductItemDeleteOne {
	return c.DeleteOneID(pi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProductItemClient) DeleteOneID(id uuid.UUID) *ProductItemDeleteOne {
	builder := c.Delete().Where(productitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductItemDeleteOne{builder}
}

// Query returns a query builder for ProductItem.
func (c *ProductItemClient) Query() *ProductItemQuery {
	return &ProductItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProductItem},
		inters: c.Interceptors(),
	}
}

// Get returns a ProductItem entity by its id.
func (c *ProductItemClient) Get(ctx context.Context, id uuid.UUID) (*ProductItem, error) {
	return c.Query().Where(productitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductItemClient) GetX(ctx context.Context, id uuid.UUID) *ProductItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubCategories queries the sub_categories edge of a ProductItem.
func (c *ProductItemClient) QuerySubCategories(pi *ProductItem) *SubCategoryQuery {
	query := (&SubCategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, id),
			sqlgraph.To(subcategory.Table, subcategory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productitem.SubCategoriesTable, productitem.SubCategoriesColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOrder queries the order edge of a ProductItem.
func (c *ProductItemClient) QueryOrder(pi *ProductItem) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, productitem.OrderTable, productitem.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a ProductItem.
func (c *ProductItemClient) QueryUser(pi *ProductItem) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productitem.UserTable, productitem.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductItemClient) Hooks() []Hook {
	return c.hooks.ProductItem
}

// Interceptors returns the client interceptors.
func (c *ProductItemClient) Interceptors() []Interceptor {
	return c.inters.ProductItem
}

func (c *ProductItemClient) mutate(ctx context.Context, m *ProductItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProductItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProductItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProductItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProductItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ProductItem mutation op: %q", m.Op())
	}
}

// SubCategoryClient is a client for the SubCategory schema.
type SubCategoryClient struct {
	config
}

// NewSubCategoryClient returns a client for the SubCategory from the given config.
func NewSubCategoryClient(c config) *SubCategoryClient {
	return &SubCategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `subcategory.Hooks(f(g(h())))`.
func (c *SubCategoryClient) Use(hooks ...Hook) {
	c.hooks.SubCategory = append(c.hooks.SubCategory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `subcategory.Intercept(f(g(h())))`.
func (c *SubCategoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.SubCategory = append(c.inters.SubCategory, interceptors...)
}

// Create returns a builder for creating a SubCategory entity.
func (c *SubCategoryClient) Create() *SubCategoryCreate {
	mutation := newSubCategoryMutation(c.config, OpCreate)
	return &SubCategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SubCategory entities.
func (c *SubCategoryClient) CreateBulk(builders ...*SubCategoryCreate) *SubCategoryCreateBulk {
	return &SubCategoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SubCategoryClient) MapCreateBulk(slice any, setFunc func(*SubCategoryCreate, int)) *SubCategoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SubCategoryCreateBulk{err: fmt.Errorf("calling to SubCategoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SubCategoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SubCategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SubCategory.
func (c *SubCategoryClient) Update() *SubCategoryUpdate {
	mutation := newSubCategoryMutation(c.config, OpUpdate)
	return &SubCategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubCategoryClient) UpdateOne(sc *SubCategory) *SubCategoryUpdateOne {
	mutation := newSubCategoryMutation(c.config, OpUpdateOne, withSubCategory(sc))
	return &SubCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubCategoryClient) UpdateOneID(id uuid.UUID) *SubCategoryUpdateOne {
	mutation := newSubCategoryMutation(c.config, OpUpdateOne, withSubCategoryID(id))
	return &SubCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SubCategory.
func (c *SubCategoryClient) Delete() *SubCategoryDelete {
	mutation := newSubCategoryMutation(c.config, OpDelete)
	return &SubCategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubCategoryClient) DeleteOne(sc *SubCategory) *SubCategoryDeleteOne {
	return c.DeleteOneID(sc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubCategoryClient) DeleteOneID(id uuid.UUID) *SubCategoryDeleteOne {
	builder := c.Delete().Where(subcategory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubCategoryDeleteOne{builder}
}

// Query returns a query builder for SubCategory.
func (c *SubCategoryClient) Query() *SubCategoryQuery {
	return &SubCategoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubCategory},
		inters: c.Interceptors(),
	}
}

// Get returns a SubCategory entity by its id.
func (c *SubCategoryClient) Get(ctx context.Context, id uuid.UUID) (*SubCategory, error) {
	return c.Query().Where(subcategory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubCategoryClient) GetX(ctx context.Context, id uuid.UUID) *SubCategory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProductItems queries the product_items edge of a SubCategory.
func (c *SubCategoryClient) QueryProductItems(sc *SubCategory) *ProductItemQuery {
	query := (&ProductItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, id),
			sqlgraph.To(productitem.Table, productitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, subcategory.ProductItemsTable, subcategory.ProductItemsColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCategory queries the category edge of a SubCategory.
func (c *SubCategoryClient) QueryCategory(sc *SubCategory) *CategoryQuery {
	query := (&CategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subcategory.CategoryTable, subcategory.CategoryColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUser queries the user edge of a SubCategory.
func (c *SubCategoryClient) QueryUser(sc *SubCategory) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subcategory.UserTable, subcategory.UserColumn),
		)
		fromV = sqlgraph.Neighbors(sc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubCategoryClient) Hooks() []Hook {
	return c.hooks.SubCategory
}

// Interceptors returns the client interceptors.
func (c *SubCategoryClient) Interceptors() []Interceptor {
	return c.inters.SubCategory
}

func (c *SubCategoryClient) mutate(ctx context.Context, m *SubCategoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubCategoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubCategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubCategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubCategoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SubCategory mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrders queries the orders edge of a User.
func (c *UserClient) QueryOrders(u *User) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.OrdersTable, user.OrdersColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProductItems queries the product_items edge of a User.
func (c *UserClient) QueryProductItems(u *User) *ProductItemQuery {
	query := (&ProductItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(productitem.Table, productitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ProductItemsTable, user.ProductItemsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCategories queries the categories edge of a User.
func (c *UserClient) QueryCategories(u *User) *CategoryQuery {
	query := (&CategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CategoriesTable, user.CategoriesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySubCategories queries the sub_categories edge of a User.
func (c *UserClient) QuerySubCategories(u *User) *SubCategoryQuery {
	query := (&SubCategoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(subcategory.Table, subcategory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SubCategoriesTable, user.SubCategoriesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Category, Order, ProductItem, SubCategory, User []ent.Hook
	}
	inters struct {
		Category, Order, ProductItem, SubCategory, User []ent.Interceptor
	}
)
