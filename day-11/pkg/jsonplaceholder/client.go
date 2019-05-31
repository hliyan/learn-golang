package jsonplaceholder

import (
	"encoding/json"
	"io/ioutil"
	"microservice/pkg/errors"
	"net/http"
)

// Todo represents a todo object from the rest api
type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Client is a client for jsonplaceholder.typicode.com REST API
type Client struct {
	isMocked bool
	baseURL  string
	client   *http.Client
}

// NewClient creates a new Client
func NewClient() *Client {
	return &Client{
		isMocked: false,
		baseURL:  "http://jsonplaceholder.typicode.com",
		client:   &http.Client{},
	}
}

// EnableMockMode enables/disables mock mode
func (c *Client) EnableMockMode(enable bool) {
	c.isMocked = enable
}

// FetchTodo fetches a Todo from /todos/{id}
func (c *Client) FetchTodo(id string) (*Todo, error) {
	if c.isMocked {
		return &Todo{ID: 1, UserID: 2, Title: "Hello world", Completed: false}, nil
	}

	resp, err := c.client.Get(c.baseURL + "/todos/" + id)
	if err != nil {
		return nil, errors.New("Could not reach server")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading server response")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("?").Code(400)
	}

	todo := &Todo{}
	err = json.Unmarshal(body, todo)
	if err != nil {
		return nil, errors.New("Server did not return valid Todo data")
	}

	return todo, nil
}
