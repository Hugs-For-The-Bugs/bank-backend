package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"
)

func CreateAccount(c *gin.Context) {
	var account model.Account
	err := c.BindJSON(&account)
	fmt.Println(err)
	account.Password = util.HashPassword(account.Password)
	result := util.DB.Create(&account)
	fmt.Println(result)
	if result.RowsAffected == 1 {
		util.SuccessResponse(c, account)
	} else {
		util.BadRequestResponse(c, "Account already exists")
	}
}
