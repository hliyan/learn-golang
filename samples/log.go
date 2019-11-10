package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error opening file: " + err.Error())
		os.Exit(0)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("Test")
}
