package controller

import (
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
}

func GetAccount(c *gin.Context) {
	//TO-DO: Fetch the user data from the database
	//result :=FetchUserData()
	id, _ := strconv.Atoi(c.Param("ID"))
	c.JSON(200, gin.H{
		"id":                   id,
		"socialSecurityNumber": "20000101-0000",
		"firstName":            "Lambo",
		"surname":              "Zhuang",
		"birthday":             "20000101",
		"phone":                "0123456789",
		"email":                "example@gmail.com",
	})
}
