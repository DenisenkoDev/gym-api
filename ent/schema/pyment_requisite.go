package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PaymentRequisite schema
type PaymentRequisite struct {
	ent.Schema
}

func (PaymentRequisite) Fields() []ent.Field {
	return []ent.Field{
		field.String("bank_kard_1").Optional(),
		field.String("bank_kard_2").Optional(),
		field.String("bank_kard_3").Optional(),
		field.String("bank_kard_4").Optional(),
		field.String("bank_kard_5").Optional(),
		field.String("name_bank_kard_1").Optional(),
		field.String("name_bank_kard_2").Optional(),
		field.String("name_bank_kard_3").Optional(),
		field.String("name_bank_kard_4").Optional(),
		field.String("name_bank_kard_5").Optional(),
		field.String("name_bank").Optional(),
		field.String("iban").Optional(),            // UA IBAN
		field.String("edrpou").Optional(),          // ЕДРПОУ / ИНН
		field.String("receiver_name").Optional(),   // Название получателя
		field.String("payment_purpose").Optional(), // Назначение платежа
		field.String("mfo").Optional(),             // MFO код
		field.String("account_number").Optional(),  // старые счета
	}
}

func (PaymentRequisite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("gym", Gym.Type).Unique().Required(),
	}
}
