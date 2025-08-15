package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type ManagerRole struct {
	ent.Schema
}

// Fields of the User.
func (ManagerRole) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("manager_role").
			Values("director", "sales_manager", "trainer", "instructor", "administrator", "cleaner", "massage_therapist"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updated_at").Optional(),
		field.Bool("is_close").Default(false),
	}
}

// Edges of the User.
func (ManagerRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("gym", Gym.Type).Unique().Required(),
		edge.To("user", User.Type).Unique().Required(),
	}
}
