package tests

import (
	"gym-api/internal/config"
	"testing"
)

func TestMain(m *testing.M) {
	// Инициализация конфигурации перед запуском тестов
	config.LoadConfig()

	// Выполнение тестов
	m.Run()
}
