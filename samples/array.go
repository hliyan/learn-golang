package main

import (
	"fmt"
)

func main() {
	fmt.Println("arrays")
	primes := [5]int{1, 2, 3, 5}

	l := len(primes)

	for i := 0; i < l; i++ {
		fmt.Println(primes[i])
	}

	for index, value := range primes {
		fmt.Printf("%d = %d\n", index, value)
	}
}
