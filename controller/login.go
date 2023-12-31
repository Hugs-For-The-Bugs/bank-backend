package controller

import (
	"fmt"
	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	SocialSecurityNumber string `json:"socialSecurityNumber"`
	Password             string `json:"password"`
	Active               bool   `json:"active"`
}

type responseData struct {
	ID uint64 `json:"id"`
}

// Login login and store session
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	var account model.Account
	err := c.BindJSON(&loginRequest)
	if err != nil {
		fmt.Println(err)
	}

	result := util.DB.Where(map[string]interface{}{
		"social_security_number": loginRequest.SocialSecurityNumber,
	}).Find(&account)

	if result.RowsAffected == 0 {
		fmt.Println(result)
		util.BadRequestResponse(c, "Account not found")
	} else if result.Error != nil {
		util.ServerErrorResponse(c, result.Error.Error())
	} else if !util.CheckPasswordHash(loginRequest.Password, account.Password) {
		util.BadRequestResponse(c, "Password incorrect")
	} else if !account.Active {
		util.BadRequestResponse(c, "Your account is not activated")
	} else {
		session := sessions.Default(c)
		session.Set("id", account.ID)
		err := session.Save()
		if err != nil {
			return
		}
		responseData := responseData{
			ID: account.ID,
		}
		util.SuccessResponse(c, responseData)
	}

}
