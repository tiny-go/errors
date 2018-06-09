package errors

import "net/http"

var _ Error = MethodNotAllowed{}

// MethodNotAllowed is a 404 HTTP error.
type MethodNotAllowed struct{ error }

// Code returns HTTP status code.
func (e MethodNotAllowed) Code() int { return http.StatusMethodNotAllowed }

// NewMethodNotAllowed is a constructor func for 405 HTTP error.
func NewMethodNotAllowed(cause error) MethodNotAllowed { return MethodNotAllowed{cause} }
