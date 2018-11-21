package errors

import "net/http"

var _ Error = Conflict{}

// Conflict is a 409 HTTP error.
type Conflict struct{ error }

// Code returns HTTP status code.
func (e Conflict) Code() int { return http.StatusConflict }

// NewConflict is a constructor func for 409 HTTP error.
func NewConflict(cause error) Conflict { return Conflict{cause} }
