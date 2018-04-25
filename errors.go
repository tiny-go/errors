package errors

import (
	"fmt"
	"net/http"
)

// Error interface describes errors that contain (HTTP or any other) status code.
type Error interface {
	error
	Code() int
}

// StatusError is a basic Error interface implementation.
type StatusError struct {
	error
	code int
}

// NewStatusError is a constructor func for StatusError.
func NewStatusError(code int, cause error) StatusError {
	return StatusError{error: cause, code: code}
}

// Code returns HTTP status code (to satisfy Error interface).
func (e StatusError) Code() int {
	return e.code
}

// Send provided argument as an HTTP error to the client (retrieving status code
// if available). If status code cannot be retrieved it sends error with status 500
// by default . If value does not implement any of supported error interfaces
// (error/Error) - it tries to convert the value to a string.
func Send(w http.ResponseWriter, err interface{}) {
	switch e := err.(type) {
	case nil:
		// this is not an error (ignore)
	case Error:
		// retrieve status code and error message
		http.Error(w, e.Error(), e.Code())
	case error:
		// all standard errors without codes will be sent as Internal Server Error
		http.Error(w, e.Error(), http.StatusInternalServerError)
	default:
		// everything else
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
}
