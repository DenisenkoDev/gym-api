package routes

import (
	"fmt"
	"gym-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getGym — обработчик для получения информации о зале по ID
func getGym(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println("idStr: ", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		return
	}

	// Получаем контекст из запроса
	ctx := c.Request.Context()

	gym, err := repository.GetGymWithRelatedEntities(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{"gym": gym})
}
