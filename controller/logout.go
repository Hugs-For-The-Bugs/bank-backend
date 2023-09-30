package controller

import (
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		util.ServerErrorResponse(c, err.Error())
		return
	}

	util.SuccessResponse(c, nil)
	util.SuccessResponse(c, "Logout successfully")
}
