package db

import (
	"context"
	"gym-api/ent"
)

func CreateAbonementType(client *ent.Client, ctx context.Context, gym_id int, name string) (ab *ent.AbonementType, err error) {

	// Создаем пользователя
	ab, err = client.AbonementType.Create().
		SetName(name).
		SetGymID(gym_id).
		Save(ctx)
	return
}
