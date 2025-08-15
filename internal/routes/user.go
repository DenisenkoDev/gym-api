package routes

import (
	"encoding/json"
	"fmt"
	"gym-api/db"
	"gym-api/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getGym — обработчик для получения информации о зале по ID
func getVisitor(c *gin.Context) {

	fmt.Println("RUN API VISITOR")

	body, err := GetBodyRequest(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(1, err)
		return
	}

	fmt.Println(body)

	id, err := strconv.Atoi(body.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid gym ID"})
		fmt.Println(2, err)
		return
	}

	// Получаем контекст из запроса
	ctx := c.Request.Context()

	vis, err := repository.GetVizitorFromID(ctx, db.Client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(3, err)
		return
	}

	r, _ := json.Marshal(vis)

	fmt.Println(string(r))

	c.JSON(http.StatusOK, gin.H{"visitor": vis})
}
