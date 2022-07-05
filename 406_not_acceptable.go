package errors

import (
	"fmt"
	"net/http"
)

var _ Error = NotAcceptable("")

// NotAcceptable is a 406 HTTP error.
type NotAcceptable string

func (e NotAcceptable) Error() string { return string(e) }

// Code returns HTTP status code.
func (e NotAcceptable) Code() int { return http.StatusNotAcceptable }

// NewMethodNotAllowed is a constructor func for 406 HTTP error.
func NewNotAcceptable(cause error) NotAcceptable { return NotAcceptable(cause.Error()) }

// NotAcceptablef returns formatted NotAcceptable error.
func NotAcceptablef(format string, args ...interface{}) NotAcceptable {
	return NotAcceptable(fmt.Sprintf(format, args...))
}
