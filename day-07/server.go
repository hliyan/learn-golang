/*
DAY 7: Building

$ cd learn-golang
$ cp -r day-06 day-07 # we're building the code from the last day
$ cd day-07
$ go build # this creates an executable named day-07
$ ./day-07 # run the server
$ curl http://localhost:8080/hello?who=world # on separate terminal

It's that simple! For more options:

$ go help build
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
