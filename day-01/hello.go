/*
	Day 1: Hello World

  https://golang.org/dl/ : download and install golang for your OS

	$ cd your/src/directory
	$ mkdir learning-golang
	$ cd learning-golang
	$ vi hello.go            # copy and paste this code
	$ go run hello.go        # run directly (doesn't create the compiled exec)
*/

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}

/*
NOTES:

- Packages, import and main are kind of like in Java. We'll revisit later
- "fmt" is a standard library package
- main() should always be in a package named "main"

*/
