package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Account struct {
	ID                   uint64 `gorm:"primaryKey"`
	Birthday             string `json:"birthday"`
	Email                string `json:"email"`
	FirstName            string `json:"firstName"`
	Password             string `json:"password"`
	Phone                string `json:"phone"`
	SocialSecurityNumber string `json:"socialSecurityNumber"`
	Surname              string `json:"surname"`
}

func CreateAccount(c *gin.Context) {
	var account Account
	err := c.BindJSON(&account)
	fmt.Println(err)
	//store into the database
	defer db.Close()
	db.Create(&account)
	//return data
	c.JSON(200, gin.H{
		"id":                   account.ID,
		"socialSecurityNumber": account.SocialSecurityNumber,
		"firstName":            account.FirstName,
		"surname":              account.Surname,
		"birthday":             account.Birthday,
		"phone":                account.Phone,
		"email":                account.Email,
	})
}
