package controller

import (
	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"

	"fmt"

	"github.com/gin-gonic/gin"
)

func EditAccount(c *gin.Context) {
	var account model.Account
	err := c.BindJSON(&account)
	if err != nil {
		fmt.Println(err)
	}

	session := sessions.Default(c)
	id := session.Get("id")

	result := util.DB.Model(&model.Account{}).Where("id = ?", id).Updates(&account)
	if result.Error == nil && result.RowsAffected == 1 {
		util.SuccessResponse(c, account)
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
