package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")

	// defer use database connection
	db, _ := sql.Open("driverName", "connection string")
	defer db.Close()

	rows, _ := db.Query("some sql query here")
	defer rows.Close()
}
