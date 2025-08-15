package config

import (
	"fmt"
	"os"
)

// Config - структура для хранения конфигурации
// Config - structure to hold configuration
type Config struct {
	Host          string
	Port          string
	User          string
	Password      string
	DBName        string
	JWTSecret     string
	ConnectServer string
}

var Cfg Config

// LoadConfig - загрузка конфигурации из переменных окружения
// LoadConfig - load configuration from environment variables
func LoadConfig() Config {
	Cfg = Config{
		Host:      getEnv("PG_HOST", "127.0.0.1"),
		Port:      getEnv("PG_PORT", "5432"),
		User:      getEnv("PG_USER", "dess"),
		Password:  getEnv("PG_PASSWORD", "2128506"),
		DBName:    getEnv("PG_GYM_NAME", "gym"),
		JWTSecret: getEnv("GYM_JWT_SEKRET_KEY", ""),
	}

	Cfg.createConnServ()

	fmt.Println(Cfg.ConnectServer)
	return Cfg
}

// getEnv - вспомогательная функция для получения переменных окружения
// с fallback значением по умолчанию
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	} else {
		return value
	}
}

// генерация строки подключения
func (a *Config) createConnServ() *Config {
	a.ConnectServer = fmt.Sprintf("host=%s port=%s user=%s dbname='%s' password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, "disable")
	return a
}
