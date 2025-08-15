package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gym-api/ent" // Импортируем сгенерированный клиент
	"gym-api/internal/config"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

var Client *ent.Client

// NewEntClient создает подключение к БД и возвращает клиент Ent
func InitEntClient() {
	// Подключаемся через драйвер Ent
	driver, err := sql.Open("postgres", config.Cfg.ConnectServer)
	if err != nil {
		fmt.Println("ошибка подключения к БД: %w", err)
		os.Exit(1)
	}

	// Пинг с замером времени
	start := time.Now()
	if err := driver.DB().Ping(); err != nil {
		fmt.Println("ошибка пинга к БД: %w", err)
		os.Exit(2)
	}
	log.Printf("Пинг к БД выполнен успешно за %v\n", time.Since(start))

	// Возвращаем клиента Ent
	Client = ent.NewClient(ent.Driver(driver))

	// Автоматически создаём или обновляем схему
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Printf("Ошибка при обновлении схемы: %v", err)
	}

	log.Println("Схема успешно обновлена")
}

func CloseEntClient() {
	Client.Close()
}
