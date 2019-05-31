package errors

// import "microservice/pkg/errors"
// USAGE: e := errors.New("Cannot reticulate splines").Code(123)

import "fmt"

// BasicError represents a basic error with an error code and a message
type BasicError struct {
	code    int
	message string
}

// Error is implemented by BasicError so that it can be used like the built in error type
func (e *BasicError) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.message)
}

// New creates a new BasicError
func New(message string) *BasicError {
	return &BasicError{code: 0, message: message}
}

// Code sets the code, which will otherwise be 0
func (e *BasicError) Code(code int) *BasicError {
	e.code = code
	return e
}
