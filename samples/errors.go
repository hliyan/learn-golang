package main

import "fmt"

type myError struct {
	Code    int
	Message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func test() (int, error) {
	return 0, &myError{Code: 1, Message: "Not implemented"}
}

func main() {
	if _, err := test(); err != nil {
		fmt.Println(err)
	}
}
