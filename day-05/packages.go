/*
DAY 5: Packages

$ cd learn-golang
$ mkdir day-05; cd day-05
$ vi packages.go <- COPY packages.go here
$ mkdir util; cd util
$ vi fact.go <- COPY util/fact.go here
$ cd ..
$ go run packages.go
*/

package main

import "./util"
import "fmt"

func main() {
	fact := util.Fact(5)
	fmt.Println(fact)
}

/*
EXPLANATION
- Think of a package as a collection of functions
- A package named "util" must be in a directory named "util"
- Package names are lowercase
- Go has no public/private keyword, instead:
- Identifiers that start with an upper case letter get exported (public)!
- E.g. PublicFunction, privateFunction
*/
