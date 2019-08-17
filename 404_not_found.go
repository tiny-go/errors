package errors

import (
	"fmt"
	"net/http"
)

var _ Error = NotFound("")

// NotFound is a 404 HTTP error.
type NotFound string

func (e NotFound) Error() string { return string(e) }

// Code returns HTTP status code.
func (e NotFound) Code() int { return http.StatusNotFound }

// NewNotFound is a constructor func for 404 HTTP error.
func NewNotFound(cause error) NotFound { return NotFound(cause.Error()) }

// NotFoundf returns formatted NotFound error.
func NotFoundf(format string, args ...interface{}) NotFound {
	return NotFound(fmt.Sprintf(format, args...))
}
