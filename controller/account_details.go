package controller

import (
	"strconv"

	"hugsforthebugs/bank-backend/model"
	"hugsforthebugs/bank-backend/util"

	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	//TO-DO: Fetch the user data from the database
	var account model.Account
	id, _ := strconv.Atoi(c.Param("ID"))
	result := util.DB.First(&account, "id =?", id)
	if result.Error == nil && result.RowsAffected == 1 {
		c.JSON(200, gin.H{
			"id":                   account.ID,
			"socialSecurityNumber": account.SocialSecurityNumber,
			"firstName":            account.FirstName,
			"surname":              account.Surname,
			"birthday":             account.Birthday,
			"phone":                account.Phone,
			"email":                account.Email,
			"balance":              account.Balance,
		})
	} else {
		c.JSON(400, gin.H{})
	}
}
