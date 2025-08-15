package routes

import (
	"gym-api/internal/config"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Структура для токена
type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

// Middleware для проверки JWT
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		secretKey := config.Cfg.JWTSecret

		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует, требуется авторизация"})
			c.Abort()
			return
		}

		// Проверяем, начинается ли заголовок с "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Проверяем валидность токена
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный или истекший токен"})
			c.Abort()
			return
		}

		// Если токен валиден, сохраняем данные в контексте
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки токена"})
			c.Abort()
			return
		}

		c.Set("login", claims.Login)
		c.Next()
	}
}
