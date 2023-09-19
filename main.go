package main

import (
	"hugsforthebugs/bank-backend/controller"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.InitDB()
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/account/:ID", controller.GetAccount)

	router.Run("localhost:8080")
}
