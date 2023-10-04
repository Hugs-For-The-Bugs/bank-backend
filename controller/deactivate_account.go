package controller

import (
	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeactivateAccount(c *gin.Context) {
	var account model.Account
	session := sessions.Default(c)
	id := session.Get("id")
	result := util.DB.Select("balance").Find(&account).Where("id = ?", id).Scan(&account)
	if result.Error == nil && result.RowsAffected == 1 {
		if account.Balance == 0 {
			util.DB.Model(&account).Where("id = ?", id).Update("active", 0)
			session.Clear()
			err := session.Save()
			if err != nil {
				return
			}
			util.SuccessResponse(c, account)
		} else {
			util.BadRequestResponse(c, "Your balance is not 0, cannot deactivate account.")
		}
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
