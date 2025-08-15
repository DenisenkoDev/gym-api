package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.String("link_photo").Optional(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("confirmed_at").Optional(),
		field.Bool("is_confirmed").Default(false),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("abonement", Abonement.Type).Unique().Required(),
	}
}
