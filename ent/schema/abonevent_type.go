package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AbonementType struct {
	ent.Schema
}

func (AbonementType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

func (AbonementType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("gym", Gym.Type).Unique().Required(),
		edge.To("abonement", Abonement.Type),
	}
}
