package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	log.Println("User created!")
}
