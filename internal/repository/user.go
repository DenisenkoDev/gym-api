package repository

import (
	"context"
	"fmt"
	"gym-api/db"
	"gym-api/ent"
	"gym-api/ent/abonement"
	"gym-api/ent/managerrole"
	"gym-api/ent/user"
	"gym-api/ent/userrole"
)

// получение хеш пользователя по логину
func GetLoginAndPasswordHash(ctx context.Context, login string) (string, string, bool) {
	u, err := db.Client.User.
		Query().
		Where(user.MailEQ(login)).
		WithCredential().
		Select(user.FieldMail).
		Only(ctx)

	if err != nil {
		return "", "", false
	}

	ph := u.Edges.Credential.PasswordHash

	return u.Mail, ph, true
}

func GetUserGymAndRoleFromID(ctx context.Context, client *ent.Client, id int) (us *ent.User, err error) {

	us, err = client.User.Query().
		Where(user.ID(id)).
		WithOwnedGyms(func(mg *ent.GymQuery) {
			mg.WithUserRoles(func(urq *ent.UserRoleQuery) {
				urq.Where(userrole.HasUserWith(user.IDEQ(id)))
			})
		}).
		WithManagerGym(func(mg *ent.GymQuery) {
			mg.WithUserRoles(func(urq *ent.UserRoleQuery) {
				urq.Where(userrole.HasUserWith(user.IDEQ(id)))
			})
		}).
		WithVisitorGym(func(mg *ent.GymQuery) {
			mg.WithUserRoles(func(urq *ent.UserRoleQuery) {
				urq.Where(userrole.HasUserWith(user.IDEQ(id)))
			})
		}).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load UserGymAndRole with gyms: %w", err)
	}

	if us == nil || len(us.Edges.OwnedGyms) == 0 {
		return nil, fmt.Errorf("user is not a UserGymAndRole or has no gyms")
	}

	return
}

func GetOwnerFromID(ctx context.Context, client *ent.Client, id int) (owner *ent.User, err error) {

	owner, err = client.User.Query().
		Where(user.ID(id)).
		WithOwnedGyms(func(gq *ent.GymQuery) {
			gq.WithVisitors(func(vg *ent.UserQuery) {
				vg.WithUserRoles().
					WithAbonements(func(ag *ent.AbonementQuery) {
						ag.WithPayments()
						ag.WithType()
					})
			})

			gq.WithManagers(func(mg *ent.UserQuery) {
				mg.WithUserRoles()
			})

		}).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load owner with gyms: %w", err)
	}

	if owner == nil || len(owner.Edges.OwnedGyms) == 0 {
		return nil, fmt.Errorf("user is not a owner or has no gyms")
	}

	return
}

func GetVizitorFromID(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {

	visitor, err := client.User.Query().
		Where(user.ID(id)).
		WithVisitorGym(func(gq *ent.GymQuery) {
			gq.WithUserRoles(func(urq *ent.UserRoleQuery) {
				urq.Where(userrole.HasUserWith(user.IDEQ(id)))
			})
			gq.WithManagerRoles(func(urq *ent.ManagerRoleQuery) {
				urq.Where(managerrole.HasUserWith(user.IDEQ(id)))
			})
			gq.WithAbonements(func(ag *ent.AbonementQuery) {
				ag.Where(
					abonement.HasUserWith(user.IDEQ(id)),
				)
				ag.WithPayments()
				ag.WithCoach()
			})
		}).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load visitor: %w", err)
	}

	if visitor == nil || len(visitor.Edges.VisitorGym) == 0 {
		return nil, fmt.Errorf("user is not a visitor or has no gyms")
	}

	return visitor, nil
}

func GetManagerFromID(ctx context.Context, client *ent.Client, id int) (manager *ent.User, err error) {

	manager, err = client.User.Query().
		Where(user.ID(id)).
		WithManagerGym(func(gq *ent.GymQuery) {
			gq.WithVisitors(func(vg *ent.UserQuery) {
				vg.WithUserRoles().
					WithAbonements(func(ag *ent.AbonementQuery) {
						ag.WithPayments()
						ag.WithType()
					})
			})

			gq.WithManagers(func(mg *ent.UserQuery) {
				mg.WithUserRoles()
			})
		}).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to load owner with gyms: %w", err)
	}

	if manager == nil || len(manager.Edges.ManagerGym) == 0 {
		return nil, fmt.Errorf("user is not a manager or has no gyms")
	}

	return
}
