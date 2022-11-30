package errors

import (
	"fmt"
	"net/http"
)

var _ Error = ServiceUnavailable("")

// ServiceUnavailable is a 503 HTTP error.
type ServiceUnavailable string

func (e ServiceUnavailable) Error() string { return string(e) }

// Code returns HTTP status code.
func (e ServiceUnavailable) Code() int { return http.StatusServiceUnavailable }

// NewServiceUnavailable is a constructor func for 503 HTTP error.
func NewServiceUnavailable(cause error) ServiceUnavailable { return ServiceUnavailable(cause.Error()) }

// ServiceUnavailablef returns formatted ServiceUnavailable error.
func ServiceUnavailablef(format string, args ...interface{}) ServiceUnavailable {
	return ServiceUnavailable(fmt.Sprintf(format, args...))
}
