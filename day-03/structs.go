/*
DAY 3: Structs and methods on structs

$ cd learn-golang
$ mkdir day-03
$ cd day-03
$ vi stucts.go # copy this code here
$ go run structs.go
*/

package main

import "fmt"

// define a struct
type rect struct {
	width, height int
	name          string
}

// define a method named perim() on rect type
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

// define another method
func (r *rect) expand(i int) {
	r.width *= i
	r.height *= i
}

func main() {
	// use {} notation to instantiate struct
	var r = rect{
		width:  20,
		height: 5,
		name:   "bar", // <- yes, that comma is not a typo
	}

	fmt.Println("Perimiter:", r.perim())
	r.expand(2) // since expand() receives a POINTER to the original rect "r", it modifies the original
	fmt.Println("Perimeter after expanding:", r.perim())
}

/*
EXPLANATION
- Go has no classes or inheritance (don't panic, it's a good thing!)
- If you want to logically group variables, you need to use structs
- You can use above notation to define methods on structs
- Note the named variable instead of the "this" keyword in such methods
- By default, Go variables are always passed by value (i.e when you pass something
	between functions, you're  passing copies)
- If you want to reference the original object, you need to pass a pointer (*)
- TRY: try removing the * from the expand method and see what happens!S
*/
