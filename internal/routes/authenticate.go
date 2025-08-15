package routes

import (
	"context"
	"fmt"
	"gym-api/db"
	"gym-api/internal/config"
	"gym-api/internal/cryp"
	"gym-api/internal/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Обработчик авторизации
func authenticate(c *gin.Context) {

	ctx := context.Background()
	secretKey := config.Cfg.JWTSecret

	var creds struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Валидация JSON из тела запроса
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Получаем хеш из БД
	_, hash, ok := repository.GetLoginAndPasswordHash(ctx, creds.Login)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	// Сравниваем хеши паролей
	if cryp.ComparePasswordWithHash(creds.Password, hash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}

	// Генерируем новый JWT токен
	expirationTime := time.Now().Add(24 * 360 * time.Hour)
	claims := &Claims{
		Login: creds.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	// Отправляем токен клиенту
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// Временная структура для десериализации JSON из запроса
type RegisterInput struct {
	Login     string `json:"login" binding:"required"`
	Password  string `json:"password" binding:"required"` // Чистый пароль из запроса
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Mail      string `json:"mail" binding:"required,email"`
}

// Функция регистрации нового пользователя
func register(c *gin.Context) {

	fmt.Println("run register")

	var input RegisterInput

	// Валидация входных данных (уже встроена в binding)
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	ctx := context.Background()
	fmt.Println(input)

	// Возвращаем успешный ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
	})

	return

	_, err := db.CreateUser(db.Client, ctx, input.Password, input.FirstName, input.LastName, "", input.Login, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания пользователя"})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"mail":    input.Mail,
	})
}

// Структура для запроса
type BodyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

func GetBodyRequest(c *gin.Context) (r BodyRequest, err error) {
	// Пробуем распарсить JSON
	if err = c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON: " + err.Error()})
		return
	}
	return
}
