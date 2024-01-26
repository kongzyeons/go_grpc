package config

import (
	"fmt"
	"my-package/models"
	"my-package/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDataBaseSqlite(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.User{})

	var count int
	db.Raw("SELECT COUNT(*) FROM users").Scan(&count)
	if count == 0 {
		fmt.Println("create user admin...")
		db.Create(&models.User{
			Username: "admin",
			Password: utils.HashPassword("admin"),
		})
	}
	fmt.Println("number row of users : ", count)
	return db
}
