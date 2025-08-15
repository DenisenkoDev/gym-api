package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("mail").Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.String("phone"),
		field.String("address").Optional(),
		field.String("description").Optional(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Пользователь как посетитель залов
		edge.From("visitor_gym", Gym.Type).
			Ref("visitors"),

		// Менеджер
		edge.From("manager_gym", Gym.Type).
			Ref("managers"),

		// personal_treiner
		edge.From("personal_treiner", Abonement.Type).Ref("coach"),

		// Владелец может владеть многими залами
		edge.To("owned_gyms", Gym.Type),

		// абонемент
		edge.To("abonements", Abonement.Type),

		// список ролей
		edge.From("user_roles", UserRole.Type).Ref("user"),

		// Роли manager
		edge.From("manager_roles", ManagerRole.Type).Ref("user"),

		// family
		edge.To("family_members", User.Type).From("family_of"),

		// credential
		edge.From("credential", Credential.Type).Ref("user").Unique(),

		// usage_mode
		edge.From("usage_mode", UsageMode.Type).Ref("user").Unique(),
	}
}
