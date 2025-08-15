package main

import (
	"gym-api/db"
	"gym-api/internal/config"
	"gym-api/internal/routes"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"github.com/hedwigz/entviz"
	_ "github.com/lib/pq"
)

// main - application entry point
func main() {

	// инициируем конфиг
	config.LoadConfig()

	// Создаём драйвер Ent
	db.InitEntClient()
	defer db.CloseEntClient()

	generate_entviz := true

	if generate_entviz {
		err := entc.Generate("../../ent/schema", &gen.Config{}, entc.Extensions(entviz.Extension{}))
		if err != nil {
			log.Fatalf("running ent codegen: %v", err)
		}
	}

	// os.Exit(1)

	routes.RunRoutes()

}
