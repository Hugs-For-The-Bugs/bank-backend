package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"
)

func CreateAccount(c *gin.Context) {
	var account model.Account
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
