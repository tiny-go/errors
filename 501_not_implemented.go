package errors

import "net/http"

var _ Error = NotImplemented{}

// NotImplemented is a 501 HTTP error.
type NotImplemented struct{ error }

// Code returns HTTP status code.
func (e NotImplemented) Code() int { return http.StatusNotImplemented }

// NewNotImplemented is a constructor func for 501 HTTP error.
func NewNotImplemented(cause error) NotImplemented { return NotImplemented{cause} }
