package errors

import (
	"errors"
	"net/http"
)

var _ Error = InternalServer{}

// InternalServer is a 500 HTTP error.
type InternalServer struct {
	error
}

// Code returns HTTP staus code.
func (e InternalServer) Code() int {
	return http.StatusInternalServerError
}

// NewInternalServer is a constructor func for 500 HTTP error.
func NewInternalServer(cause string) InternalServer {
	return InternalServer{errors.New(cause)}
}
