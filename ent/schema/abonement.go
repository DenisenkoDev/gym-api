package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Abonement holds the schema definition for the Abonement entity.
type Abonement struct {
	ent.Schema
}

// Fields of the Abonement.
func (Abonement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Float("price").Positive(),
		field.Int("duration_months").Positive().Default(1),
		field.Time("expiration_date").Default(func() time.Time {
			return time.Now().AddDate(0, 1, 0)
		}), // срок истечения
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updated_at").Optional(),
		field.Time("paid_until").Optional(),
		field.Bool("is_active").Default(true),
		field.Bool("is_paid").Default(false),
	}
}

// Edges of the Abonement.
func (Abonement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payments", Payment.Type).Ref("abonement"),
		edge.From("user", User.Type).Ref("abonements").Unique().Required(),
		edge.To("coach", User.Type).Unique(),
		edge.To("gym", Gym.Type).Unique().Required(),
		edge.From("type", AbonementType.Type).Ref("abonement").Unique().Required(),
	}
}
