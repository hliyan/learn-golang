package app

import (
	"microservice/pkg/jsonplaceholder"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server wraps the gin server to create a proxy server for jsonplaceholder.typicode.com/todos API
// using the jsonplaceholder client package in this project
type Server struct {
	router *gin.Engine
	client *jsonplaceholder.Client
}

// NewServer creates a new Server instance
func NewServer() *Server {
	client := jsonplaceholder.NewClient()
	client.EnableMockMode(false)

	s := &Server{}
	r := gin.Default()
	s.router = r
	s.client = client

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "0.1",
		})
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		todo, err := client.FetchTodo(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, todo)
	})

	return s
}

// Run creates the REST API
func (s *Server) Run() error {
	return s.router.Run()
}
