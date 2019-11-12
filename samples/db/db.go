package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select id, name, age from test")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", id, name, age)
	}
}
