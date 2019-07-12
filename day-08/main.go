/*
DAY 8: Using a third party package!

$ mkdir day-08
$ cd day-08
$ go mod init day-08
$ vi main.go # copy the code below
$ go get
$ go build
$ ./day-08
$ curl http://localhost:8080/hello/world
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + c.Param("name"),
		})
	})

	r.Run()
}

/*
EXPLANATION:
- Go's equivalent of Express.js: https://github.com/gin-gonic/gin
- Go's equivalent of npm init: go mod init <package-name>
- Go's equivalent of package.json: go.mod (created by go mod init)
- Go's equivalent of npm install: go get (it reads your import statements directly)
- Go's equivalent of package-lock.json: go.sum (created by go get)
- When you run the sample, don't forget to checkout the server console output!!
- The Gin API is pretty intuitive - check out the docs.
*/
