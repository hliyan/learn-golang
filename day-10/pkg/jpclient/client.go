package jpclient

// USAGE:
// 	import "microservice/pkg/jpclient"
// 	client := jsonplaceholder.NewClient()
// 	todo, err := client.FetchTodo("1")
// 	if err != nil {
// 		log.Println(err)
// 	} else {
//		log.Println(todo)
//  }

import (
	"errors"
)

// Todo represents a todo object from the rest api
type Todo struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Client is a client for http://jsonplaceholder.typicode.com/todos API
type Client struct {
	isMocked bool   // if true, returns mock results for test purposes
	baseURL  string // currently hard coded
}

// NewClient creates a new Client
func NewClient() *Client {
	return &Client{isMocked: false, baseURL: "http://jsonplaceholder.typicode.com"}
}

// EnableMockMode enables/disables mock mode
func (c *Client) EnableMockMode(enable bool) {
	c.isMocked = enable
}

// FetchTodo fetches a Todo from /todos/{id}
func (c *Client) FetchTodo(id string) (*Todo, error) {
	if c.isMocked {
		return &Todo{ID: "1", UserID: "john", Title: "Hello world", Completed: false}, nil
	}
	// TODO: implement
	return nil, errors.New("Not implemented")
}

/*
EXPLANATIONS:
- Everything above is just putting together what
	you have already learned.
- But do notice how simple the Go way of writing
  doc comments are: // <Identifier> is <purpose>
*/
