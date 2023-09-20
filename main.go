package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"hugsforthebugs/bank-backend/controller"
	"hugsforthebugs/bank-backend/middleware"
	"hugsforthebugs/bank-backend/util"
)

func main() {
	util.InitDB()           // Initialize the database
	router := gin.Default() // Create a new Gin router

	// Use your custom Middleware
	router.Use(middleware.Middleware())

	// Use session middleware
	router.Use(sessions.Sessions("SESSIONID", cookie.NewStore([]byte("secret"))))

	// Special route with additional middleware and a handler
	router.GET("/middleware/", middleware.Middleware(), someOtherHandler)

	// Define your other routes and their handlers
	// Login route
	router.POST("/api/login", controller.Login)

	// Test API route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start the server
	router.Run("localhost:8080")
}

func someOtherHandler(c *gin.Context) {
	// Your logic here
	c.JSON(200, gin.H{
		"message": "Handled by someOtherHandler",
	})
}
