package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Gym holds the schema definition for the Gym entity.
type Gym struct {
	ent.Schema
}

// Fields of the Gym.
func (Gym) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("currency_id").Default(3),
		field.String("phone").Unique(),
		field.String("mail").Unique(),
		field.String("address").Optional(),
		field.String("web_site").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Gym.
func (Gym) Edges() []ent.Edge {
	return []ent.Edge{
		// Посетители зала
		edge.To("visitors", User.Type),

		// Менеджеры
		edge.To("managers", User.Type),

		// AbonementType
		edge.From("abonement_type", AbonementType.Type).Ref("gym"),

		// Владелец зала
		edge.From("owner", User.Type).Ref("owned_gyms").Unique().Required(),

		// abonements
		edge.From("abonements", Abonement.Type).Ref("gym"),

		// Роли пользователей
		edge.From("user_roles", UserRole.Type).Ref("gym"),

		// Роли manager
		edge.From("manager_roles", ManagerRole.Type).Ref("gym"),

		// payment_requisite
		edge.From("payment_requisite", PaymentRequisite.Type).Ref("gym").Unique(),
	}
}
