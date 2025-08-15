package data_test

import (
	"context"
	"fmt"
	"gym-api/db"
	"gym-api/ent"
	"gym-api/ent/managerrole"
	"gym-api/ent/userrole"
	"gym-api/internal/config"
	"testing"

	_ "github.com/lib/pq"
)

// go test -v ./tests/data_test

func TestData(t *testing.T) {
	// load config
	config.LoadConfig()
	// Создаем тестовый клиент
	client, err := ent.Open("postgres", config.Cfg.ConnectServer)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Создаем пользователя owner
	tanchik, err := db.CreateUser(client, ctx, "Aa111111", "Сергей", "Танчик", "+380997776644", "tanchik@gmail.com", "Сумы, Gorky Park")
	if err != nil {
		fmt.Println("tanchik", err)
		return
	}

	olia, err := db.CreateUser(client, ctx, "Aa111111", "Оля", "Танчик", "+380997776633", "tanchik111@gmail.com", "Сумы, Gorky Park")
	if err != nil {
		fmt.Println("tanchik", err)
		return
	}

	dess, err := db.CreateUser(client, ctx, "Aa111111", "Александр", "Денисенко", "+380660021121", "20dess20@gmail.com", "Сумы, Лушпы 5/5 кв 66")
	if err != nil {
		fmt.Println("dess", err)
		return
	}

	vikki, err := db.CreateUser(client, ctx, "Aa111111", "Виктория", "Денисенко", "+380995553040", "20vikki20@gmail.com", "Сумы, Лушпы 5/5 кв 44")
	if err != nil {
		fmt.Println(6, err)
		return
	}

	// добавляем вику в семью
	err = db.AddFamily(client, ctx, dess.ID, vikki.ID)
	if err != nil {
		fmt.Println("set famyly", err)
		return
	}

	// Создаем новый спортзал
	gym, err := db.CreateGym(client, ctx, tanchik.ID, "Танчик жим", "+380997776644", "20dess20@gmail.com", "www.gym.com", "456 Elm St")

	if err != nil {
		fmt.Println(2, err)
		return
	}

	gym1, err := db.CreateGym(client, ctx, dess.ID, "Dess жим", "+380997772244", "hgyturiruu@gmail.com", "www.gym.com", "456 Elm St")

	if err != nil {
		fmt.Println(2, err)
		return
	}

	// привязываем visitor к залу
	err = db.AddVisitor(client, ctx, gym.ID, dess.ID)
	if err != nil {
		fmt.Println(4, err)
		return
	}

	err = db.AddVisitor(client, ctx, gym.ID, vikki.ID)
	if err != nil {
		fmt.Println(4, err)
		return
	}

	// привязываем manager к залу
	err = db.AddManager(client, ctx, gym.ID, vikki.ID)
	if err != nil {
		fmt.Println(41, err)
		return
	}

	err = db.AddManager(client, ctx, gym.ID, olia.ID)
	if err != nil {
		fmt.Println(42, err)
		return
	}

	err = db.AddManager(client, ctx, gym.ID, tanchik.ID)
	if err != nil {
		fmt.Println(42, err)
		return
	}

	// создаем роль пользователей
	err = db.CreateUserRole(client, ctx, gym1.ID, dess.ID, userrole.UserRoleOwner)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, dess.ID, userrole.UserRoleVisitor)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, tanchik.ID, userrole.UserRoleOwner)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, tanchik.ID, userrole.UserRoleManager)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, vikki.ID, userrole.UserRoleManager)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, vikki.ID, userrole.UserRoleVisitor)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	err = db.CreateUserRole(client, ctx, gym.ID, olia.ID, userrole.UserRoleManager)
	if err != nil {
		fmt.Println(5, err)
		return
	}

	// создаем роль для director
	err = db.CreateManagerRole(client, ctx, gym.ID, vikki.ID, managerrole.ManagerRoleDirector)
	if err != nil {
		fmt.Println(51, err)
		return
	}

	// создаем роль для trainer
	err = db.CreateManagerRole(client, ctx, gym.ID, olia.ID, managerrole.ManagerRoleTrainer)
	if err != nil {
		fmt.Println(51, err)
		return
	}

	err = db.CreateManagerRole(client, ctx, gym.ID, tanchik.ID, managerrole.ManagerRoleTrainer)
	if err != nil {
		fmt.Println(51, err)
		return
	}

	err = db.CreateManagerRole(client, ctx, gym.ID, tanchik.ID, managerrole.ManagerRoleTrainer)
	if err != nil {
		fmt.Println(51, err)
		return
	}

	// create abonement type
	abType, err := db.CreateAbonementType(client, ctx, gym.ID, "Тренажерный зал")
	if err != nil {
		fmt.Println(71, err)
		return
	}

	abType1, err := db.CreateAbonementType(client, ctx, gym.ID, "Массаж")
	if err != nil {
		fmt.Println(72, err)
		return
	}

	// create aabonement
	abnm, err := db.CreateAbonement(client, ctx, dess.ID, abType.ID, gym.ID, 1000, 1, &olia.ID)
	if err != nil {
		fmt.Println(8, err)
		return
	}

	// create aabonement1
	abnm1, err := db.CreateAbonement(client, ctx, dess.ID, abType1.ID, gym.ID, 2000, 1, nil)
	if err != nil {
		fmt.Println(9, err)
		return
	}

	// create pyment
	err = db.CreatePyment(client, ctx, 2000.00, abnm.ID)
	if err != nil {
		fmt.Println(10, err)
		return
	}

	err = db.CreatePyment(client, ctx, 500.00, abnm1.ID)

	if err != nil {
		fmt.Println(11, err)
		return
	}

	fmt.Println("test create data completed")
}
