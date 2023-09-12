package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// test api
	router.GET("/t", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run("localhost:8080")
}
