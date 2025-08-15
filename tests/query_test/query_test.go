package querytest

import (
	"context"
	"gym-api/ent"
	"gym-api/internal/config"
	"gym-api/internal/repository"
	"testing"
)

// go test -v ./tests/query_test

func TestData(t *testing.T) {

	// load config
	config.LoadConfig()
	// Создаем тестовый клиент
	client, err := ent.Open("postgres", config.Cfg.ConnectServer)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer client.Close()

	// id user
	id := 3

	us, err := repository.GetUserGymAndRoleFromID(context.Background(), client, id)

	repository.PrintUserInfo(us, err)

	// return

	viz, err := repository.GetVizitorFromID(context.Background(), client, id)

	repository.PrintVisitorInfo(viz, err)

	viz, err = repository.GetOwnerFromID(context.Background(), client, id)

	repository.PrintOwnerInfo(viz, err)

	viz, err = repository.GetManagerFromID(context.Background(), client, id)

	repository.PrintManagerInfo(viz, err)

}
