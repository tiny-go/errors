package errors

import (
	"fmt"
	"net/http"
)

var _ Error = BadRequest("")

// BadRequest is a 400 HTTP error.
type BadRequest string

func (e BadRequest) Error() string { return string(e) }

// Code returns HTTP status code.
func (e BadRequest) Code() int { return http.StatusBadRequest }

// NewBadRequest is a constructor func for 400 HTTP error.
func NewBadRequest(cause error) BadRequest { return BadRequest(cause.Error()) }

// BadRequestf returns formatted BadRequest error.
func BadRequestf(format string, args ...interface{}) BadRequest {
	return BadRequest(fmt.Sprintf(format, args...))
}
