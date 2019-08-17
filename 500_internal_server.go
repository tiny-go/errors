package errors

import (
	"fmt"
	"net/http"
)

var _ Error = InternalServer("")

// InternalServer is a 500 HTTP error.
type InternalServer string

func (e InternalServer) Error() string { return string(e) }

// Code returns HTTP status code.
func (e InternalServer) Code() int { return http.StatusInternalServerError }

// NewInternalServer is a constructor func for 500 HTTP error.
func NewInternalServer(cause error) InternalServer { return InternalServer(cause.Error()) }

// InternalServerf returns formatted InternalServer error.
func InternalServerf(format string, args ...interface{}) InternalServer {
	return InternalServer(fmt.Sprintf(format, args...))
}
