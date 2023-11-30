package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type ProductItem struct {
	ent.Schema
}

// Fields of the User.
func (ProductItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.String("name").NotEmpty().StructTag(`json:"name"`),
		field.String("amount").NotEmpty().StructTag(`json:"amount"`),
		field.String("unit_type").NotEmpty().StructTag(`json:"unit_type"`),
		field.Time("created_at").
			Default(time.Now).StructTag(`json:"created_at"`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).StructTag(`json:"updated_at"`),
	}
}

// Edges of the User.
func (ProductItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sub_categories", SubCategory.Type).
			Ref("product_items").
			Unique(),
		edge.From("user", User.Type).
			Ref("product_items").
			Unique(),
	}
}
