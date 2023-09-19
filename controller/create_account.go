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
	// store into the database
	result := util.DB.Create(&account)
	// return data
	if result.Error == nil && result.RowsAffected == 1 {
		util.SuccessResponse(c, account)
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
