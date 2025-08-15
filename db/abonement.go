package db

import (
	"context"
	"fmt"
	"gym-api/ent"
	"time"
)

func CreateAbonement(
	client *ent.Client, ctx context.Context,
	user_id, type_id, gym_id int, prise float64, duration_months int,
	coach_id *int,
) (ab *ent.Abonement, err error) {

	abonementType, err := client.AbonementType.Get(ctx, type_id)
	if err != nil {
		return nil, fmt.Errorf("failed to find abonement type: %w", err)
	}

	builder := client.Abonement.Create().
		SetName(abonementType.Name).
		SetPrice(prise).
		SetUserID(user_id).
		SetTypeID(type_id).
		SetGymID(gym_id).
		SetDurationMonths(duration_months).
		SetExpirationDate(time.Now().AddDate(0, duration_months, 0))

	if coach_id != nil {
		builder = builder.SetCoachID(*coach_id)
	}

	ab, err = builder.Save(ctx)
	return
}
