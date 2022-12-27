package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sherlock28/api-go-gim-gorm/api/controllers"
	"github.com/sherlock28/api-go-gim-gorm/api/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)

	r.Run()
}
