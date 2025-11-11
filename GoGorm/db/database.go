package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:Nomelase@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error connecting to database:", err)
		panic(err)
	} else {
		fmt.Println("Connected to the database successfully")
		return db
	}
}()
