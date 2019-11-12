package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Testing DB inserts...")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("insert into test(name, age) values (?, ?)")
	if err != nil {
		panic(err)
	}

	stmt.Exec("paul", 90)
	if err != nil {
		panic(err)
	}
}
