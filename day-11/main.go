package main

import (
	"log"
	"microservice/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting server")

	server := app.NewServer()
	err := server.Run()
	if err != nil {
		log.Printf(err.Error())
	}
}
