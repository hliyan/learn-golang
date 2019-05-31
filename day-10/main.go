package main

import (
	"log"
	"day-10/pkg/jpclient"
)

// a simple proxy server for: 
// http://jsonplaceholder.typicode.com/todos/{id}
func main() {
	client := jpclient.NewClient()
	client.EnableMockMode(true)

	todo, err := client.FetchTodo("1")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(todo)
}

/*
EXPLANATION:
- line 5:  
If your go.mod file defines the current package as "day-10",
and your package is inside a directory called pkg, the package
can be imported as "day-10/pkg/<packagename>".

- line 10: 
The Go way to construct structs is using a package level 
constructor function, usually named packagename.NewObjectName().
DON'T use structs to group functions like you do in Java!
Use packages to group functions.
Use structs only to logically group variables.
*/