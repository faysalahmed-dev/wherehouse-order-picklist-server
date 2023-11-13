// Code generated by ent, DO NOT EDIT.

package subcategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldName, v))
}

// Descriptions applies equality check predicate on the "descriptions" field. It's identical to DescriptionsEQ.
func Descriptions(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldDescriptions, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldValue, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContainsFold(FieldName, v))
}

// DescriptionsEQ applies the EQ predicate on the "descriptions" field.
func DescriptionsEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldDescriptions, v))
}

// DescriptionsNEQ applies the NEQ predicate on the "descriptions" field.
func DescriptionsNEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldDescriptions, v))
}

// DescriptionsIn applies the In predicate on the "descriptions" field.
func DescriptionsIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldDescriptions, vs...))
}

// DescriptionsNotIn applies the NotIn predicate on the "descriptions" field.
func DescriptionsNotIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldDescriptions, vs...))
}

// DescriptionsGT applies the GT predicate on the "descriptions" field.
func DescriptionsGT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldDescriptions, v))
}

// DescriptionsGTE applies the GTE predicate on the "descriptions" field.
func DescriptionsGTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldDescriptions, v))
}

// DescriptionsLT applies the LT predicate on the "descriptions" field.
func DescriptionsLT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldDescriptions, v))
}

// DescriptionsLTE applies the LTE predicate on the "descriptions" field.
func DescriptionsLTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldDescriptions, v))
}

// DescriptionsContains applies the Contains predicate on the "descriptions" field.
func DescriptionsContains(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContains(FieldDescriptions, v))
}

// DescriptionsHasPrefix applies the HasPrefix predicate on the "descriptions" field.
func DescriptionsHasPrefix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasPrefix(FieldDescriptions, v))
}

// DescriptionsHasSuffix applies the HasSuffix predicate on the "descriptions" field.
func DescriptionsHasSuffix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasSuffix(FieldDescriptions, v))
}

// DescriptionsEqualFold applies the EqualFold predicate on the "descriptions" field.
func DescriptionsEqualFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEqualFold(FieldDescriptions, v))
}

// DescriptionsContainsFold applies the ContainsFold predicate on the "descriptions" field.
func DescriptionsContainsFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContainsFold(FieldDescriptions, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldContainsFold(FieldValue, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.SubCategory {
	return predicate.SubCategory(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasOrders applies the HasEdge predicate on the "orders" edge.
func HasOrders() predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrdersWith applies the HasEdge predicate on the "orders" edge with a given conditions (other predicates).
func HasOrdersWith(preds ...predicate.Order) predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := newOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.SubCategory {
	return predicate.SubCategory(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SubCategory) predicate.SubCategory {
	return predicate.SubCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SubCategory) predicate.SubCategory {
	return predicate.SubCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SubCategory) predicate.SubCategory {
	return predicate.SubCategory(sql.NotPredicates(p))
}
