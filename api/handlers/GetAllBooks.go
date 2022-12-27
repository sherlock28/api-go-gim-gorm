package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sherlock28/api-go-gim-gorm/api/models"
)

func (h handler) GetAllBooks(c *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
