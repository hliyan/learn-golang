/*
DAY 4: Standard lib, multiple return values, defer

$ cd learn-golang
$ mkdir day-04
$ cd day-04
$ vi stdlib.go # copy this code here
$ go run stdlib.go
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var resp, err = http.Get("http://example.com/posts/1")
	if err != nil {
		fmt.Println("Unable to get from example.com")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unable to read response from example.com")
		return
	}
	fmt.Println(string(body))
}

/*
EXPLANATION
- line 13: Imports can be grouped together
           Go has a rich standard library -- you will rarely need frameworks
           The http package: // docs: golang.org/pkg/net/http/
- line 20: Go functions can return multiple values!
           Standard go error handling practices is: `return result, err`
- line 25: The defer keyword tells go to execute the statement just before exiting the function
           Very nifty to do housekeeping/resource release type work
*/
