package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UsageMode holds the schema definition for the UsageMode entity.
type UsageMode struct {
	ent.Schema
}

func (UsageMode) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("mode").Values("test", "paid").Default("test"),
		field.Time("created_at"),
		field.Time("paid_activated_at").Optional(),
	}
}

func (UsageMode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
	}
}
