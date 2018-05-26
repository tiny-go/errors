package errors

import "net/http"

var _ Error = Forbidden{}

// Forbidden is a 403 HTTP error.
type Forbidden struct{ error }

// Code returns HTTP staus code.
func (e Forbidden) Code() int { return http.StatusForbidden }

// NewForbidden is a constructor func for 403 HTTP error.
func NewForbidden(cause error) Forbidden { return Forbidden{cause} }
