/*
DAY 9: Pointers (don't panic, they're easy)

$ mkdir day-09
$ cd day-09
$ vi main.go # copy the code below
$ go run main.go
*/

package main

import "fmt"

// Answer is just a struct used to demonstrate pointers
type Answer struct {
	value int
}

func main() {
	answer := Answer{value: 42}
	answerAgain := answer
	answer.value = 43

	fmt.Println("Wihout pointers, things get passed by value")

	fmt.Println(answer)      // output: {43}
	fmt.Println(answerAgain) // output: {42} (surprise, JS and Java folks!)

	fmt.Println("You want to pass by reference? Use pointers")

	pAnswer := &answer
	answer.value = 44
	fmt.Println(answer)   // output: {44}
	fmt.Println(*pAnswer) // output: {44}   (whew!)
}

/*
EXPLANATION:
- In Go, when you pass/assign a variable, you always get a COPY, never a reference
- This is true for both primitive types and object types (i.e. structs)
- Might be surprising to JS and JAVA devs, who expect objects to be passed/assigned by reference
- In Go, if you want to pass/assign a struct without creating a copy, you have to use a POINTER
- A pointer is just the memory address of another variable
- "&"" is the pointer operator, and you can read "&p" as "address of p"
- "*" is the deferencing operator, and you can read "*p" as "target of p"
*/
