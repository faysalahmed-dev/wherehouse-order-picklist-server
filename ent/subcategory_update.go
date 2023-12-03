// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/category"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/productitem"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/subcategory"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/google/uuid"
)

// SubCategoryUpdate is the builder for updating SubCategory entities.
type SubCategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *SubCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SubCategoryUpdate builder.
func (scu *SubCategoryUpdate) Where(ps ...predicate.SubCategory) *SubCategoryUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetName sets the "name" field.
func (scu *SubCategoryUpdate) SetName(s string) *SubCategoryUpdate {
	scu.mutation.SetName(s)
	return scu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableName(s *string) *SubCategoryUpdate {
	if s != nil {
		scu.SetName(*s)
	}
	return scu
}

// SetDescriptions sets the "descriptions" field.
func (scu *SubCategoryUpdate) SetDescriptions(s string) *SubCategoryUpdate {
	scu.mutation.SetDescriptions(s)
	return scu
}

// SetNillableDescriptions sets the "descriptions" field if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableDescriptions(s *string) *SubCategoryUpdate {
	if s != nil {
		scu.SetDescriptions(*s)
	}
	return scu
}

// SetValue sets the "value" field.
func (scu *SubCategoryUpdate) SetValue(s string) *SubCategoryUpdate {
	scu.mutation.SetValue(s)
	return scu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableValue(s *string) *SubCategoryUpdate {
	if s != nil {
		scu.SetValue(*s)
	}
	return scu
}

// SetCreatedAt sets the "created_at" field.
func (scu *SubCategoryUpdate) SetCreatedAt(t time.Time) *SubCategoryUpdate {
	scu.mutation.SetCreatedAt(t)
	return scu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableCreatedAt(t *time.Time) *SubCategoryUpdate {
	if t != nil {
		scu.SetCreatedAt(*t)
	}
	return scu
}

// SetUpdatedAt sets the "updated_at" field.
func (scu *SubCategoryUpdate) SetUpdatedAt(t time.Time) *SubCategoryUpdate {
	scu.mutation.SetUpdatedAt(t)
	return scu
}

// AddProductItemIDs adds the "product_items" edge to the ProductItem entity by IDs.
func (scu *SubCategoryUpdate) AddProductItemIDs(ids ...uuid.UUID) *SubCategoryUpdate {
	scu.mutation.AddProductItemIDs(ids...)
	return scu
}

// AddProductItems adds the "product_items" edges to the ProductItem entity.
func (scu *SubCategoryUpdate) AddProductItems(p ...*ProductItem) *SubCategoryUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scu.AddProductItemIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (scu *SubCategoryUpdate) SetCategoryID(id uuid.UUID) *SubCategoryUpdate {
	scu.mutation.SetCategoryID(id)
	return scu
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableCategoryID(id *uuid.UUID) *SubCategoryUpdate {
	if id != nil {
		scu = scu.SetCategoryID(*id)
	}
	return scu
}

// SetCategory sets the "category" edge to the Category entity.
func (scu *SubCategoryUpdate) SetCategory(c *Category) *SubCategoryUpdate {
	return scu.SetCategoryID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (scu *SubCategoryUpdate) SetUserID(id uuid.UUID) *SubCategoryUpdate {
	scu.mutation.SetUserID(id)
	return scu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableUserID(id *uuid.UUID) *SubCategoryUpdate {
	if id != nil {
		scu = scu.SetUserID(*id)
	}
	return scu
}

// SetUser sets the "user" edge to the User entity.
func (scu *SubCategoryUpdate) SetUser(u *User) *SubCategoryUpdate {
	return scu.SetUserID(u.ID)
}

// Mutation returns the SubCategoryMutation object of the builder.
func (scu *SubCategoryUpdate) Mutation() *SubCategoryMutation {
	return scu.mutation
}

// ClearProductItems clears all "product_items" edges to the ProductItem entity.
func (scu *SubCategoryUpdate) ClearProductItems() *SubCategoryUpdate {
	scu.mutation.ClearProductItems()
	return scu
}

// RemoveProductItemIDs removes the "product_items" edge to ProductItem entities by IDs.
func (scu *SubCategoryUpdate) RemoveProductItemIDs(ids ...uuid.UUID) *SubCategoryUpdate {
	scu.mutation.RemoveProductItemIDs(ids...)
	return scu
}

// RemoveProductItems removes "product_items" edges to ProductItem entities.
func (scu *SubCategoryUpdate) RemoveProductItems(p ...*ProductItem) *SubCategoryUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scu.RemoveProductItemIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (scu *SubCategoryUpdate) ClearCategory() *SubCategoryUpdate {
	scu.mutation.ClearCategory()
	return scu
}

// ClearUser clears the "user" edge to the User entity.
func (scu *SubCategoryUpdate) ClearUser() *SubCategoryUpdate {
	scu.mutation.ClearUser()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SubCategoryUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SubCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SubCategoryUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SubCategoryUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *SubCategoryUpdate) defaults() {
	if _, ok := scu.mutation.UpdatedAt(); !ok {
		v := subcategory.UpdateDefaultUpdatedAt()
		scu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *SubCategoryUpdate) check() error {
	if v, ok := scu.mutation.Name(); ok {
		if err := subcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SubCategory.name": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Descriptions(); ok {
		if err := subcategory.DescriptionsValidator(v); err != nil {
			return &ValidationError{Name: "descriptions", err: fmt.Errorf(`ent: validator failed for field "SubCategory.descriptions": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Value(); ok {
		if err := subcategory.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "SubCategory.value": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scu *SubCategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubCategoryUpdate {
	scu.modifiers = append(scu.modifiers, modifiers...)
	return scu
}

func (scu *SubCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(subcategory.Table, subcategory.Columns, sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeUUID))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.Name(); ok {
		_spec.SetField(subcategory.FieldName, field.TypeString, value)
	}
	if value, ok := scu.mutation.Descriptions(); ok {
		_spec.SetField(subcategory.FieldDescriptions, field.TypeString, value)
	}
	if value, ok := scu.mutation.Value(); ok {
		_spec.SetField(subcategory.FieldValue, field.TypeString, value)
	}
	if value, ok := scu.mutation.CreatedAt(); ok {
		_spec.SetField(subcategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := scu.mutation.UpdatedAt(); ok {
		_spec.SetField(subcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if scu.mutation.ProductItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedProductItemsIDs(); len(nodes) > 0 && !scu.mutation.ProductItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.ProductItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.UserTable,
			Columns: []string{subcategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.UserTable,
			Columns: []string{subcategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// SubCategoryUpdateOne is the builder for updating a single SubCategory entity.
type SubCategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SubCategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (scuo *SubCategoryUpdateOne) SetName(s string) *SubCategoryUpdateOne {
	scuo.mutation.SetName(s)
	return scuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableName(s *string) *SubCategoryUpdateOne {
	if s != nil {
		scuo.SetName(*s)
	}
	return scuo
}

// SetDescriptions sets the "descriptions" field.
func (scuo *SubCategoryUpdateOne) SetDescriptions(s string) *SubCategoryUpdateOne {
	scuo.mutation.SetDescriptions(s)
	return scuo
}

// SetNillableDescriptions sets the "descriptions" field if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableDescriptions(s *string) *SubCategoryUpdateOne {
	if s != nil {
		scuo.SetDescriptions(*s)
	}
	return scuo
}

// SetValue sets the "value" field.
func (scuo *SubCategoryUpdateOne) SetValue(s string) *SubCategoryUpdateOne {
	scuo.mutation.SetValue(s)
	return scuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableValue(s *string) *SubCategoryUpdateOne {
	if s != nil {
		scuo.SetValue(*s)
	}
	return scuo
}

// SetCreatedAt sets the "created_at" field.
func (scuo *SubCategoryUpdateOne) SetCreatedAt(t time.Time) *SubCategoryUpdateOne {
	scuo.mutation.SetCreatedAt(t)
	return scuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableCreatedAt(t *time.Time) *SubCategoryUpdateOne {
	if t != nil {
		scuo.SetCreatedAt(*t)
	}
	return scuo
}

// SetUpdatedAt sets the "updated_at" field.
func (scuo *SubCategoryUpdateOne) SetUpdatedAt(t time.Time) *SubCategoryUpdateOne {
	scuo.mutation.SetUpdatedAt(t)
	return scuo
}

// AddProductItemIDs adds the "product_items" edge to the ProductItem entity by IDs.
func (scuo *SubCategoryUpdateOne) AddProductItemIDs(ids ...uuid.UUID) *SubCategoryUpdateOne {
	scuo.mutation.AddProductItemIDs(ids...)
	return scuo
}

// AddProductItems adds the "product_items" edges to the ProductItem entity.
func (scuo *SubCategoryUpdateOne) AddProductItems(p ...*ProductItem) *SubCategoryUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scuo.AddProductItemIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (scuo *SubCategoryUpdateOne) SetCategoryID(id uuid.UUID) *SubCategoryUpdateOne {
	scuo.mutation.SetCategoryID(id)
	return scuo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableCategoryID(id *uuid.UUID) *SubCategoryUpdateOne {
	if id != nil {
		scuo = scuo.SetCategoryID(*id)
	}
	return scuo
}

// SetCategory sets the "category" edge to the Category entity.
func (scuo *SubCategoryUpdateOne) SetCategory(c *Category) *SubCategoryUpdateOne {
	return scuo.SetCategoryID(c.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (scuo *SubCategoryUpdateOne) SetUserID(id uuid.UUID) *SubCategoryUpdateOne {
	scuo.mutation.SetUserID(id)
	return scuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableUserID(id *uuid.UUID) *SubCategoryUpdateOne {
	if id != nil {
		scuo = scuo.SetUserID(*id)
	}
	return scuo
}

// SetUser sets the "user" edge to the User entity.
func (scuo *SubCategoryUpdateOne) SetUser(u *User) *SubCategoryUpdateOne {
	return scuo.SetUserID(u.ID)
}

// Mutation returns the SubCategoryMutation object of the builder.
func (scuo *SubCategoryUpdateOne) Mutation() *SubCategoryMutation {
	return scuo.mutation
}

// ClearProductItems clears all "product_items" edges to the ProductItem entity.
func (scuo *SubCategoryUpdateOne) ClearProductItems() *SubCategoryUpdateOne {
	scuo.mutation.ClearProductItems()
	return scuo
}

// RemoveProductItemIDs removes the "product_items" edge to ProductItem entities by IDs.
func (scuo *SubCategoryUpdateOne) RemoveProductItemIDs(ids ...uuid.UUID) *SubCategoryUpdateOne {
	scuo.mutation.RemoveProductItemIDs(ids...)
	return scuo
}

// RemoveProductItems removes "product_items" edges to ProductItem entities.
func (scuo *SubCategoryUpdateOne) RemoveProductItems(p ...*ProductItem) *SubCategoryUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scuo.RemoveProductItemIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (scuo *SubCategoryUpdateOne) ClearCategory() *SubCategoryUpdateOne {
	scuo.mutation.ClearCategory()
	return scuo
}

// ClearUser clears the "user" edge to the User entity.
func (scuo *SubCategoryUpdateOne) ClearUser() *SubCategoryUpdateOne {
	scuo.mutation.ClearUser()
	return scuo
}

// Where appends a list predicates to the SubCategoryUpdate builder.
func (scuo *SubCategoryUpdateOne) Where(ps ...predicate.SubCategory) *SubCategoryUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SubCategoryUpdateOne) Select(field string, fields ...string) *SubCategoryUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SubCategory entity.
func (scuo *SubCategoryUpdateOne) Save(ctx context.Context) (*SubCategory, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SubCategoryUpdateOne) SaveX(ctx context.Context) *SubCategory {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SubCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SubCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *SubCategoryUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdatedAt(); !ok {
		v := subcategory.UpdateDefaultUpdatedAt()
		scuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *SubCategoryUpdateOne) check() error {
	if v, ok := scuo.mutation.Name(); ok {
		if err := subcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SubCategory.name": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Descriptions(); ok {
		if err := subcategory.DescriptionsValidator(v); err != nil {
			return &ValidationError{Name: "descriptions", err: fmt.Errorf(`ent: validator failed for field "SubCategory.descriptions": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Value(); ok {
		if err := subcategory.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "SubCategory.value": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (scuo *SubCategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SubCategoryUpdateOne {
	scuo.modifiers = append(scuo.modifiers, modifiers...)
	return scuo
}

func (scuo *SubCategoryUpdateOne) sqlSave(ctx context.Context) (_node *SubCategory, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(subcategory.Table, subcategory.Columns, sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeUUID))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subcategory.FieldID)
		for _, f := range fields {
			if !subcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.Name(); ok {
		_spec.SetField(subcategory.FieldName, field.TypeString, value)
	}
	if value, ok := scuo.mutation.Descriptions(); ok {
		_spec.SetField(subcategory.FieldDescriptions, field.TypeString, value)
	}
	if value, ok := scuo.mutation.Value(); ok {
		_spec.SetField(subcategory.FieldValue, field.TypeString, value)
	}
	if value, ok := scuo.mutation.CreatedAt(); ok {
		_spec.SetField(subcategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.UpdatedAt(); ok {
		_spec.SetField(subcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if scuo.mutation.ProductItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedProductItemsIDs(); len(nodes) > 0 && !scuo.mutation.ProductItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.ProductItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subcategory.ProductItemsTable,
			Columns: []string{subcategory.ProductItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if scuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.UserTable,
			Columns: []string{subcategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.UserTable,
			Columns: []string{subcategory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(scuo.modifiers...)
	_node = &SubCategory{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
