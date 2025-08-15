package db

import (
	"context"
	"fmt"
	"gym-api/ent"
	"gym-api/ent/user"
	"time"
)

func CreateGym(client *ent.Client, ctx context.Context, id_owner int, name, phone, mail, site, address string) (gym *ent.Gym, err error) {
	// Fetch user with UsageMode and existing gyms
	user, err := client.User.
		Query().
		Where(user.IDEQ(id_owner)).
		WithUsageMode().
		WithOwnedGyms().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	usage := user.Edges.UsageMode
	isFirstGym := len(user.Edges.OwnedGyms) == 0

	if usage == nil {
		if isFirstGym {
			// Create test usage mode
			_, err = client.UsageMode.
				Create().
				SetMode("test").
				SetUserID(user.ID).
				SetCreatedAt(time.Now()).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to create usage mode: %w", err)
			}
		} else {
			return nil, fmt.Errorf("usage mode not found and this is not the first gym")
		}
	} else {
		if usage.Mode == "test" && !isFirstGym {
			return nil, fmt.Errorf("test mode allows only one gym to be created")
		}
	}

	gym, err = client.Gym.
		Create().
		SetName(name).
		SetPhone(phone).
		SetMail(mail).
		SetWebSite(site).
		SetAddress(address).
		SetOwnerID(id_owner).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create gym: %w", err)
	}

	return gym, nil
}

func AddVisitor(client *ent.Client, ctx context.Context, gym_id, visitor_id int) error {
	_, err := client.Gym.UpdateOneID(gym_id).AddVisitorIDs(visitor_id).Save(ctx)
	return err
}

func AddManager(client *ent.Client, ctx context.Context, gym_id, manager_id int) error {
	_, err := client.Gym.UpdateOneID(gym_id).AddManagerIDs(manager_id).Save(ctx)
	return err
}
