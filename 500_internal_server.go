package errors

import "net/http"

var _ Error = InternalServer{}

// InternalServer is a 500 HTTP error.
type InternalServer struct{ error }

// Code returns HTTP status code.
func (e InternalServer) Code() int { return http.StatusInternalServerError }

// NewInternalServer is a constructor func for 500 HTTP error.
func NewInternalServer(cause error) InternalServer { return InternalServer{cause} }
