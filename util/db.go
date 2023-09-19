package util

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Hostname string
	Port     string
	Username string
	Password string
	DBName   string
}

var DB *gorm.DB

func InitDB() {
	db := DBConfig{
		Hostname: "localhost",
		Port:     "3306",
		Username: "root",
		Password: "Aa@123456789",
		DBName:   "bank",
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Hostname,
		db.Port,
		db.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
