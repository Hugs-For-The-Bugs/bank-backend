package controller

import (
	"fmt"

	"hugsforthebugs/bank-backend/util"

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
	Balance              *int   `gorm:"default:0"`
}

func CreateAccount(c *gin.Context) {
	var account Account
	err := c.BindJSON(&account)
	fmt.Println(err)
	//store into the database
	//util.DB.AutoMigrate(&Account{})
	result := util.DB.Create(&account)
	//return data
	if result.Error == nil && result.RowsAffected == 1 {
		c.JSON(200, gin.H{
			"id":                   account.ID,
			"socialSecurityNumber": account.SocialSecurityNumber,
			"firstName":            account.FirstName,
			"surname":              account.Surname,
			"birthday":             account.Birthday,
			"phone":                account.Phone,
			"email":                account.Email,
		})
	} else {
		c.JSON(400, gin.H{})
	}
}
