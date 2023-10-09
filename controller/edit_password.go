package controller

import (
	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"

	"fmt"

	"github.com/gin-gonic/gin"
)

func EditPassword(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	//fmt.Print(id)      //1
	//fmt.Print(account) //{0  john@email.com John1  1234567890  Doe1 0 false}

	//check if active
	var account model.Account
	util.DB.First(&account).Where("id = ?", id)
	if !account.Active {
		util.BadRequestResponse(c, "Account Inactive!")
		return
	}

	//account active
	err := c.BindJSON(&account)
	if err != nil {
		fmt.Println(err)
	}

	//Hash password
	account.Password = util.HashPassword(account.Password)

	//only update the fields (password)
	result := util.DB.Model(&model.Account{}).Where("id = ?", id).Updates(model.Account{
		Password: account.Password,
	})

	if result.Error == nil {
		if result.RowsAffected == 1 {
			util.SuccessResponse(c, account)
		} else if result.RowsAffected == 0 {
			util.BadRequestResponse(c, "Password not changed!")
		}
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
