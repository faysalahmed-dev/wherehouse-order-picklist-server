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
type Category struct {
	ent.Schema
}

// Fields of the User.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.String("name").NotEmpty().StructTag(`json:"name"`),
		field.String("value").NotEmpty().StructTag(`json:"value",omitempty`),
		field.Time("created_at").
			Default(time.Now).StructTag(`json:"created_at",omitempty`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).StructTag(`json:"updated_at",omitempty`),
	}
}

// Edges of the User.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sub_categories", SubCategory.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("user", User.Type).
			Ref("categories").
			Unique(),
	}
}
