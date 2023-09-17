package main

import (
	"hugsforthebugs/bank-backend/controller"
	"hugsforthebugs/bank-backend/middleware"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	util.InitDB()
	router := gin.Default()
	router.Use(sessions.Sessions("SESSIONID", cookie.NewStore([]byte("secret"))))
	//login
	router.POST("/login", controller.Login)
	r := router.Group("/api")
	{
		r.Use(middleware.AuthCheck)
	}

	// test api
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run("localhost:8080")
}
