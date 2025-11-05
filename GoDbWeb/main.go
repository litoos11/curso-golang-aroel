package main

import (
	"fmt"
	"godbmysql/db"
	"godbmysql/models"
)

func main() {
	db.Connect()
	// db.Ping()
	// fmt.Println(db.ExistsTable("users"))
	// db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("anita", "example", "anita@litoos11@x.com")
	// users := models.ListUsers()
	// user := models.GetUser(2)
	fmt.Println(user)
	// user.Username = "Lili"
	// user.Save()
	// user.Delete()
	// db.TruncateTable("users")
	db.Close()
}
