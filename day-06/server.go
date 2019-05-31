/*
DAY 6: The day we write a simple REST API!

$ cd learn-golang
$ mkdir day-06; cd day-06
$ vi server.go <- COPY code here
$ go run server.go
$ curl http://localhost:8080/hello?who=world # on separate terminal
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		type Greeting struct {
			Text 		string `json:"text"`
			Error		string `json:"error,omitempty"`
		}

		who := r.URL.Query().Get("who")
		res := Greeting{}
		if who == "" {
			res.Error = "Greet who?"
		} else {
			res.Text = fmt.Sprintf("Hello %s", who)
		}

		jsonRes, _ := json.Marshal(res)
		io.WriteString(w, string(jsonRes))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

/*
EXPLANATION
- line 22:
  Tell the http module to use a function to handle a route
- line 23:
	Since Go is statically typed, we need structs to represent JSON
- line 24:
	Text attribute needs to start with an upper case letter, otherwise it would be private
	So we need a separate annotation to tell go to "marshal" (encode) it as "text" in JSON
- The rest you can learn from package docs
*/