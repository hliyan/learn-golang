package main

import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root:hnl54321@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("connected to db")

	var name = ""
	var owner = ""

	rows, err := db.Query("select name, owner from pet")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name, &owner)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, owner)
	}
	
	defer db.Close()
}