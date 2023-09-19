package controller

import (
	"hugsforthebugs/bank-backend/util"
	"strconv"

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

func GetAccount(c *gin.Context) {
	//TO-DO: Fetch the user data from the database
	var account Account
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
