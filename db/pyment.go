package db

import (
	"context"
	"gym-api/ent"
)

func CreatePyment(client *ent.Client, ctx context.Context, amount float64, abonement_id int) (err error) {
	_, err = client.Payment.Create().
		SetAmount(amount).
		SetAbonementID(abonement_id).
		Save(ctx)
	return
}
