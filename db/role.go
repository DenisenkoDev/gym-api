package db

import (
	"context"
	"gym-api/ent"
	"gym-api/ent/managerrole"
	"gym-api/ent/userrole"
)

func CreateUserRole(client *ent.Client, ctx context.Context, gym_id, user_id int, user_role userrole.UserRole) error {

	_, err := client.UserRole.
		Create().
		SetUserRole(user_role).
		SetUserID(user_id).
		SetGymID(gym_id).
		Save(ctx)
	return err
}

func CreateManagerRole(client *ent.Client, ctx context.Context, gym_id, user_id int, manager_role managerrole.ManagerRole) error {

	_, err := client.ManagerRole.
		Create().
		SetManagerRole(manager_role).
		SetUserID(user_id).
		SetGymID(gym_id).
		Save(ctx)
	return err
}
