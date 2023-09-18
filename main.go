package main

import (
	"hugsforthebugs/bank-backend/controller"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-gonic/gin"
)

func main() {
	util.InitDB()
	router := gin.Default()

	router.POST("/account", controller.CreateAccount)

	router.Run("localhost:8080")
}
