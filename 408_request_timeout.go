package errors

import "net/http"

var _ Error = RequestTimeout{}

// RequestTimeout is a 404 HTTP error.
type RequestTimeout struct{ error }

// Code returns HTTP status code.
func (e RequestTimeout) Code() int { return http.StatusRequestTimeout }

// NewRequestTimeout is a constructor func for 408 HTTP error.
func NewRequestTimeout(cause error) RequestTimeout { return RequestTimeout{cause} }
