/*
DAY 2: Variables and functions

$ cd learn-golang
$ mkdir day-02
$ cd day-02
$ vi variables-functions.go # copy this code here
$ go run variables-functions.go
*/

package main

import "fmt"

func main() {
	var foo = 3
	bar := "Hello"
	greet(bar, foo)
}

func greet(greeting string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(greeting)
	}
}

/*
EXPLANATION
- Go is statically typed, but can infer type using the first assigned value (e.g. foo)
- The x := y is a shorthand for var x = y. it declares the variable and assigns value
- Notice that there are no ";" at the end of lines. they're added by the compiler
- Go function arguments declarations start with name and end with type!
- string and int are built in types
- if and for statements are similar to C, Java and JavaScript, but the "()" are usually unnecessary
*/
