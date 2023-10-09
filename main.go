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
	router.POST("/api/account", controller.CreateAccount)
	r := router.Group("/api")
	{
		r.Use(middleware.LoginCheck)
		r.GET("/logout", controller.Logout)
		r.GET("/account", controller.GetAccount)
		r.PUT("/account", controller.EditAccount)
		r.DELETE("/account", controller.DeactivateAccount)
		r.POST("/transaction", controller.CreateTransaction)
		r.GET("/transaction", controller.GetTransactions)
		r.PUT("/password", controller.EditPassword)
	}

	// Start the server
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
