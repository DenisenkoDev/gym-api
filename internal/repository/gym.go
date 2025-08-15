package repository

import (
	"context"
	"fmt"
	"gym-api/db"
	"gym-api/ent"
	"gym-api/ent/gym"
)

// получение зала и всех связанных сущностей
func GetGymWithRelatedEntities(ctx context.Context, gymID int) (*ent.Gym, error) {

	gym, err := db.Client.Gym.Query().
		Where(gym.ID(gymID)).
		WithVisitors(func(uq *ent.UserQuery) {
			uq.WithAbonements(func(aq *ent.AbonementQuery) {
				aq.WithPayments()
			})
		}).
		WithManagers(func(uq *ent.UserQuery) {

		}).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query gym with relations: %v", err)
	}

	return gym, nil
}
