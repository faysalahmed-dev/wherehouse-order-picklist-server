package schema

import (
	"time"

	"entgo.io/ent"
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
		field.String("name").NotEmpty().StructTag(`json:"name"`),
		field.String("amount").NotEmpty().StructTag(`json:"amount"`),
		field.String("unit_type").NotEmpty().StructTag(`json:"unit_type"`),
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
		edge.From("sub_categories", SubCategory.Type).
			Ref("orders").
			Unique(),
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
	}
}
