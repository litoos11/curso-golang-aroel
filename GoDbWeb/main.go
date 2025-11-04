package main

import (
	"fmt"
	"godbmysql/db"
)

func main() {
	db.Connect()
	// db.Ping()
	fmt.Println(db.ExistsTable("users"))
	// db.CreateTable(models.UserSchema, "users")
	db.TruncateTable("users")
	db.Close()
}
