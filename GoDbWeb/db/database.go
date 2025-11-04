package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/dbname
const url = "root:Nomelase@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database successfully")
	db = connection
}

func Close() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CreateTable(schema string, table string) {
	if !ExistsTable(table) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func ExistsTable(table string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	result, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func TruncateTable(table string) {
	sql := fmt.Sprintf("TRUNCATE %s", table)
	Exec(sql)
}
