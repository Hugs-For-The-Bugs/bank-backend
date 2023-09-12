package controller

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserForm struct {
	socialSecurityNumber string `json:"socialSecurityNumber"`
	Password             string `json:"password"`
}

// Login login and store session
func Login(c *gin.Context) {
	var userForm UserForm
	err := c.BindJSON(&userForm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userForm)

	if userForm.socialSecurityNumber != "admin" || userForm.Password != "admin" {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "Incorrect socialSecurityNumber or password",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user", userForm.socialSecurityNumber)
	session.Save()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "login successfully",
	})
}
