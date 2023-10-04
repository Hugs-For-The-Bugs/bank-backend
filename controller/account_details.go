package controller

import (
	"github.com/gin-contrib/sessions"
	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	//TO-DO: Fetch the user data from the database
	var account model.Account
	session := sessions.Default(c)
	id := session.Get("id")
	result := util.DB.First(&account, "id =?", id)

	if result.Error == nil && result.RowsAffected == 1 {
		util.SuccessResponse(c, account)
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
