package errors

import "net/http"

var _ Error = Unauthorized{}

// Unauthorized is a 401 HTTP error.
type Unauthorized struct{ error }

// Code returns HTTP status code.
func (e Unauthorized) Code() int { return http.StatusUnauthorized }

// NewUnauthorized is a constructor func for 401 HTTP error.
func NewUnauthorized(cause error) Unauthorized { return Unauthorized{cause} }
