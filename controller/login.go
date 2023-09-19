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
}

// Login login and store session
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	var account model.Account
	err := c.BindJSON(&loginRequest)
	if err != nil {
		fmt.Println(err)
	}

	session := sessions.Default(c)
	sessionPassword := session.Get("password")
	if sessionPassword != nil && sessionPassword == loginRequest.Password {

		result := util.DB.Where(map[string]interface{}{
			"social_security_number": loginRequest.SocialSecurityNumber,
		}).Find(&account)

		if result.RowsAffected == 0 {
			c.JSON(400, gin.H{})
		} else if result.Error != nil {
			c.JSON(500, gin.H{})
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"id":   account.ID,
			})
		}
	} else {
		result := util.DB.Where(map[string]interface{}{
			"social_security_number": loginRequest.SocialSecurityNumber,
		}).Find(&account)

		if result.RowsAffected == 0 {
			fmt.Println(result)
			c.JSON(400, gin.H{
				"msg": "user not found",
			})
		} else if result.Error != nil {
			c.JSON(500, gin.H{
				"msg": result.Error.Error(),
			})
		} else if account.Password != loginRequest.Password {
			c.JSON(400, gin.H{
				"msg": "password error",
			})
		} else {
			session.Set("password", account.Password)
			session.Save()
			c.JSON(200, gin.H{
				"id": account.ID,
			})
		}
	}

}
