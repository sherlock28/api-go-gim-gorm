package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sherlock28/api-go-gim-gorm/api/database"
	"github.com/sherlock28/api-go-gim-gorm/api/handlers"
)

func main() {
	r := gin.Default()

	DB := database.ConnectDatabase()
	h := handlers.New(DB)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"name": "api-go-gim-gorm", "version": "1.0.0"})
	})

	r.GET("/books", h.GetAllBooks)

	r.Run()
}
