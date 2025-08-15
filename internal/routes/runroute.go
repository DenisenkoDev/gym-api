package routes

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// localhost:8085/api/gym/15
// localhost:8085/api/user/13

func RunRoutes() {
	// Инициализация Гин
	r := gin.Default()

	// Включаем CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Лучше указать свой origin: http://localhost:8080 и т.д.
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Настраиваем маршруты
	r.POST("/api/register", register) // Новый маршрут для регистрации
	r.POST("/api/auth", authenticate) // Маршрут для авторизации
	// r.GET("/api/gym:id", authMiddleware(), getGym)
	// r.GET("/api/user:id", authMiddleware(), getUser)

	// test
	r.GET("/api/gym/:id", getGym)
	r.POST("/api/visitor", getVisitor)

	// Запуск сервера
	if err := r.Run(":8085"); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
