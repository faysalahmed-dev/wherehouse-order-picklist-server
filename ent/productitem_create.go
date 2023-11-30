// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/productitem"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/subcategory"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/google/uuid"
)

// ProductItemCreate is the builder for creating a ProductItem entity.
type ProductItemCreate struct {
	config
	mutation *ProductItemMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pic *ProductItemCreate) SetName(s string) *ProductItemCreate {
	pic.mutation.SetName(s)
	return pic
}

// SetAmount sets the "amount" field.
func (pic *ProductItemCreate) SetAmount(s string) *ProductItemCreate {
	pic.mutation.SetAmount(s)
	return pic
}

// SetUnitType sets the "unit_type" field.
func (pic *ProductItemCreate) SetUnitType(s string) *ProductItemCreate {
	pic.mutation.SetUnitType(s)
	return pic
}

// SetCreatedAt sets the "created_at" field.
func (pic *ProductItemCreate) SetCreatedAt(t time.Time) *ProductItemCreate {
	pic.mutation.SetCreatedAt(t)
	return pic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pic *ProductItemCreate) SetNillableCreatedAt(t *time.Time) *ProductItemCreate {
	if t != nil {
		pic.SetCreatedAt(*t)
	}
	return pic
}

// SetUpdatedAt sets the "updated_at" field.
func (pic *ProductItemCreate) SetUpdatedAt(t time.Time) *ProductItemCreate {
	pic.mutation.SetUpdatedAt(t)
	return pic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pic *ProductItemCreate) SetNillableUpdatedAt(t *time.Time) *ProductItemCreate {
	if t != nil {
		pic.SetUpdatedAt(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *ProductItemCreate) SetID(u uuid.UUID) *ProductItemCreate {
	pic.mutation.SetID(u)
	return pic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pic *ProductItemCreate) SetNillableID(u *uuid.UUID) *ProductItemCreate {
	if u != nil {
		pic.SetID(*u)
	}
	return pic
}

// SetSubCategoriesID sets the "sub_categories" edge to the SubCategory entity by ID.
func (pic *ProductItemCreate) SetSubCategoriesID(id uuid.UUID) *ProductItemCreate {
	pic.mutation.SetSubCategoriesID(id)
	return pic
}

// SetNillableSubCategoriesID sets the "sub_categories" edge to the SubCategory entity by ID if the given value is not nil.
func (pic *ProductItemCreate) SetNillableSubCategoriesID(id *uuid.UUID) *ProductItemCreate {
	if id != nil {
		pic = pic.SetSubCategoriesID(*id)
	}
	return pic
}

// SetSubCategories sets the "sub_categories" edge to the SubCategory entity.
func (pic *ProductItemCreate) SetSubCategories(s *SubCategory) *ProductItemCreate {
	return pic.SetSubCategoriesID(s.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pic *ProductItemCreate) SetUserID(id uuid.UUID) *ProductItemCreate {
	pic.mutation.SetUserID(id)
	return pic
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pic *ProductItemCreate) SetNillableUserID(id *uuid.UUID) *ProductItemCreate {
	if id != nil {
		pic = pic.SetUserID(*id)
	}
	return pic
}

// SetUser sets the "user" edge to the User entity.
func (pic *ProductItemCreate) SetUser(u *User) *ProductItemCreate {
	return pic.SetUserID(u.ID)
}

// Mutation returns the ProductItemMutation object of the builder.
func (pic *ProductItemCreate) Mutation() *ProductItemMutation {
	return pic.mutation
}

// Save creates the ProductItem in the database.
func (pic *ProductItemCreate) Save(ctx context.Context) (*ProductItem, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *ProductItemCreate) SaveX(ctx context.Context) *ProductItem {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *ProductItemCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *ProductItemCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *ProductItemCreate) defaults() {
	if _, ok := pic.mutation.CreatedAt(); !ok {
		v := productitem.DefaultCreatedAt()
		pic.mutation.SetCreatedAt(v)
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		v := productitem.DefaultUpdatedAt()
		pic.mutation.SetUpdatedAt(v)
	}
	if _, ok := pic.mutation.ID(); !ok {
		v := productitem.DefaultID()
		pic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *ProductItemCreate) check() error {
	if _, ok := pic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ProductItem.name"`)}
	}
	if v, ok := pic.mutation.Name(); ok {
		if err := productitem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ProductItem.name": %w`, err)}
		}
	}
	if _, ok := pic.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "ProductItem.amount"`)}
	}
	if v, ok := pic.mutation.Amount(); ok {
		if err := productitem.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "ProductItem.amount": %w`, err)}
		}
	}
	if _, ok := pic.mutation.UnitType(); !ok {
		return &ValidationError{Name: "unit_type", err: errors.New(`ent: missing required field "ProductItem.unit_type"`)}
	}
	if v, ok := pic.mutation.UnitType(); ok {
		if err := productitem.UnitTypeValidator(v); err != nil {
			return &ValidationError{Name: "unit_type", err: fmt.Errorf(`ent: validator failed for field "ProductItem.unit_type": %w`, err)}
		}
	}
	if _, ok := pic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ProductItem.created_at"`)}
	}
	if _, ok := pic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ProductItem.updated_at"`)}
	}
	return nil
}

func (pic *ProductItemCreate) sqlSave(ctx context.Context) (*ProductItem, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *ProductItemCreate) createSpec() (*ProductItem, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductItem{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(productitem.Table, sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeUUID))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pic.mutation.Name(); ok {
		_spec.SetField(productitem.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pic.mutation.Amount(); ok {
		_spec.SetField(productitem.FieldAmount, field.TypeString, value)
		_node.Amount = value
	}
	if value, ok := pic.mutation.UnitType(); ok {
		_spec.SetField(productitem.FieldUnitType, field.TypeString, value)
		_node.UnitType = value
	}
	if value, ok := pic.mutation.CreatedAt(); ok {
		_spec.SetField(productitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pic.mutation.UpdatedAt(); ok {
		_spec.SetField(productitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pic.mutation.SubCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productitem.SubCategoriesTable,
			Columns: []string{productitem.SubCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sub_category_product_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productitem.UserTable,
			Columns: []string{productitem.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_product_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductItemCreateBulk is the builder for creating many ProductItem entities in bulk.
type ProductItemCreateBulk struct {
	config
	err      error
	builders []*ProductItemCreate
}

// Save creates the ProductItem entities in the database.
func (picb *ProductItemCreateBulk) Save(ctx context.Context) ([]*ProductItem, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*ProductItem, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *ProductItemCreateBulk) SaveX(ctx context.Context) []*ProductItem {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *ProductItemCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *ProductItemCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
