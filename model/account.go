package model

type Account struct {
	ID                   uint64 `gorm:"primaryKey"`
	Birthday             string `json:"birthday"`
	Email                string `json:"email"`
	FirstName            string `json:"firstName"`
	Password             string `json:"password"`
	Phone                string `json:"phone"`
	SocialSecurityNumber string `json:"socialSecurityNumber"`
	Surname              string `json:"surname"`
	Balance              int    `gorm:"default:0" json:"balance"`
	Active               bool   `gorm:"default:false" json:"active"`
}
