package errors

import (
	"fmt"
	"net/http"
)

var _ Error = NotImplemented("")

// NotImplemented is a 501 HTTP error.
type NotImplemented string

func (e NotImplemented) Error() string { return string(e) }

// Code returns HTTP status code.
func (e NotImplemented) Code() int { return http.StatusNotImplemented }

// NewNotImplemented is a constructor func for 501 HTTP error.
func NewNotImplemented(cause error) NotImplemented { return NotImplemented(cause.Error()) }

// NotImplementedf returns formatted NotImplemented error.
func NotImplementedf(format string, args ...interface{}) NotImplemented {
	return NotImplemented(fmt.Sprintf(format, args...))
}
