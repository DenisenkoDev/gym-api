package db

import (
	"context"
	"gym-api/ent"
	"gym-api/internal/cryp"
)

func CreateUser(client *ent.Client, ctx context.Context, password, first_name, last_name, phone, mail, address string) (user *ent.User, err error) {

	hash, err := cryp.GenerateHashFromPassword(password)

	if err != nil {
		return
	}

	// Создаем пользователя
	user, err = client.User.
		Create().
		SetFirstName(first_name).
		SetLastName(last_name).
		SetPhone(phone).
		SetMail(mail).
		SetAddress(address).
		Save(ctx)

	if err != nil {
		return
	}

	_, err = client.Credential.Create().
		SetPasswordHash(hash).
		SetUserID(user.ID).
		Save(ctx)

	return
}

func AddFamily(client *ent.Client, ctx context.Context, user_id, famyly_id int) error {
	_, err := client.User.UpdateOneID(user_id).AddFamilyMemberIDs(famyly_id).Save(ctx)
	return err
}
