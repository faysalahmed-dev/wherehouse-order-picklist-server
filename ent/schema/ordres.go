package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Order struct {
	ent.Schema
}

// Fields of the User.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.Enum("status").Values("PICKED", "UNPICKED").Default("UNPICKED").StructTag(`json:"status"`),
		field.Time("created_at").
			Default(time.Now).StructTag(`json:"created_at"`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).StructTag(`json:"updated_at"`),
	}
}

// Edges of the User.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_items", ProductItem.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
	}
}
