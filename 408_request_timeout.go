package errors

import (
	"fmt"
	"net/http"
)

var _ Error = RequestTimeout("")

// RequestTimeout is a 404 HTTP error.
type RequestTimeout string

func (e RequestTimeout) Error() string { return string(e) }

// Code returns HTTP status code.
func (e RequestTimeout) Code() int { return http.StatusRequestTimeout }

// NewRequestTimeout is a constructor func for 408 HTTP error.
func NewRequestTimeout(cause error) RequestTimeout { return RequestTimeout(cause.Error()) }

// RequestTimeoutf returns formatted RequestTimeout error.
func RequestTimeoutf(format string, args ...interface{}) RequestTimeout {
	return RequestTimeout(fmt.Sprintf(format, args...))
}
