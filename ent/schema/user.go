package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`json:"id"`),
		field.String("name").MaxLen(25).NotEmpty().StructTag(`json:"name"`),
		field.String("email").NotEmpty().Unique().Match(regexp.MustCompile("^(?P<name>[a-zA-Z0-9.!#$%&'*+/=?^_ \x60{|}~-]+)@(?P<domain>[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*)$")).StructTag(`json:"email"`),
		field.String("password").NotEmpty().Sensitive(),
		field.Enum("type").
			Values("ADMIN", "USER").Default("USER").StructTag(`json:"user_type"`),
		field.Bool("blocked").Default(false).StructTag(`json:"blocked"`),
		field.Int("total_orders").Default(0).StructTag(`json:"total_orders"`),
		field.Time("created_at").
			Default(time.Now).StructTag(`json:"created_at"`),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).StructTag(`json:"updated_at"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("orders", Order.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("categories", Category.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("sub_categories", SubCategory.Type),
	}
}
