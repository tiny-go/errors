package errors

import "net/http"

var _ Error = NotFound{}

// NotFound is a 404 HTTP error.
type NotFound struct{ error }

// Code returns HTTP staus code.
func (e NotFound) Code() int { return http.StatusNotFound }

// NewNotFound is a constructor func for 404 HTTP error.
func NewNotFound(cause error) NotFound { return NotFound{cause} }
