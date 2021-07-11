package database

import (
	"fmt"

	"studydemo.com/book/go微服务/10questions/pkg/models"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("sqlite3", "questions.db")
	if err != nil {
		fmt.Println("Status: ", err)
	}
	// defer DB.Close()
	DB.Debug()
	DB.LogMode(true)
	DB.AutoMigrate(&models.User{}, &models.Questions{}, &models.Answer{}, &models.Tag{})
}
