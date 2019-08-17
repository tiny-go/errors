package errors

import (
	"fmt"
	"net/http"
)

var _ Error = MethodNotAllowed("")

// MethodNotAllowed is a 404 HTTP error.
type MethodNotAllowed string

func (e MethodNotAllowed) Error() string { return string(e) }

// Code returns HTTP status code.
func (e MethodNotAllowed) Code() int { return http.StatusMethodNotAllowed }

// NewMethodNotAllowed is a constructor func for 405 HTTP error.
func NewMethodNotAllowed(cause error) MethodNotAllowed { return MethodNotAllowed(cause.Error()) }

// MethodNotAllowedf returns formatted MethodNotAllowed error.
func MethodNotAllowedf(format string, args ...interface{}) MethodNotAllowed {
	return MethodNotAllowed(fmt.Sprintf(format, args...))
}
