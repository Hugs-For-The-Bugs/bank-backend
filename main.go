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

	// Use session middleware
	router.Use(sessions.Sessions("SESSIONID", cookie.NewStore([]byte("secret"))))

	router.POST("/api/login", controller.Login)

	r := router.Group("/api")
	{
		r.Use(middleware.AuthCheck)

		//r.GET("/account/:id", controller.GetAccount)

	}

	// Start the server
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func someOtherHandler(c *gin.Context) {
	// Your logic here
	c.JSON(200, gin.H{
		"message": "Handled by someOtherHandler",
	})
}
