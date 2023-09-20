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
	router.POST("/api/login", controller.Login)
	r := router.Group("/api")
	{
		r.Use(middleware.LoginCheck)
		r.POST("/account", controller.CreateAccount)
		r.GET("/account", controller.GetAccount)
	}

	// Start the server
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
