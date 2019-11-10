package main

import (
	"fmt"
	"strconv"
)

func main() {
	ages := map[string]int{
		"John":  30,
		"James": 31,
	}

	for key, value := range ages {
		fmt.Println(key + " = " + strconv.Itoa(value))
	}

	ages["Bill"] = 40
	if billAge, billExists := ages["Bill"]; billExists {
		fmt.Println("Bill is in the map and he is " + strconv.Itoa(billAge) + " years old.")
	}
}
