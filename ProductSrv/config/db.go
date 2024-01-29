package config

import (
	"fmt"
	"my-package/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDataBaseSqlite(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	db.AutoMigrate(&models.Product{})

	var count int
	db.Raw("SELECT COUNT(*) FROM products").Scan(&count)
	if count == 0 {
		fmt.Println("create user admin...")
		db.Create(&models.Product{
			Name:     "p1",
			Price:    100,
			Category: "c1",
		})
	}
	fmt.Println("number row of product : ", count)
	return db
}
