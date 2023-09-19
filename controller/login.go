package controller

import (
	"fmt"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	SocialSecurityNumber string `json:"socialSecurityNumber"`
	Password             string `json:"password"`
}

type Account struct {
	ID                   uint `gorm:"primarykey"`
	SocialSecurityNumber string
	Password             string
}

// Login login and store session
func Login(c *gin.Context) {
	var loginRequest LoginRequest
	var user Account
	err := c.BindJSON(&loginRequest)
	if err != nil {
		fmt.Println(err)
	}

	session := sessions.Default(c)
	sessionPassword := session.Get("password")
	if sessionPassword != nil && sessionPassword == loginRequest.Password {

		result := util.DB.Where(map[string]interface{}{
			"social_security_number": loginRequest.SocialSecurityNumber,
		}).Find(&user)
		if result.RowsAffected == 0 {
			fmt.Println(result)
			return
		} else if result.Error != nil {
			fmt.Println(result.Error.Error())
			return
		} else {
			c.JSON(200, gin.H{
				"code": 0,
				"id":   user.ID,
			})
			return
		}

	}

	result := util.DB.Where(map[string]interface{}{
		"social_security_number": loginRequest.SocialSecurityNumber,
	}).Find(&user)

	if result.RowsAffected == 0 {
		fmt.Println(result)
		c.JSON(400, gin.H{
			"msg": "user not found",
		})
	} else if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": result.Error.Error(),
		})
	} else if user.Password != loginRequest.Password {
		c.JSON(400, gin.H{
			"msg": "password error",
		})
	} else {
		session.Set("password", user.Password)
		session.Save()
		c.JSON(200, gin.H{
			"id": user.ID,
		})
	}
}
