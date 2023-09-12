package controller

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserForm struct {
	ID                   uint64 `gorm:"primary"`
	SocialSecurityNumber string `json:"socialSecurityNumber"`
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

	if userForm.SocialSecurityNumber != "admin" || userForm.Password != "admin" {
		c.JSON(200, gin.H{
			"id": userForm.ID,
		})
		return
	}
	session := sessions.Default(c)
	session.Set("user", userForm.SocialSecurityNumber)
	session.Save()
	c.JSON(200, gin.H{
		"id": userForm.ID,
	})
}
