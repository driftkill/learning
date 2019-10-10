package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connecting...")

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/testdb")

	if err !=nil {
		panic(err.Error())
	}

	fmt.Println("Connected")

	defer db.Close()
}