package errors

import (
	"fmt"
	"net/http"
)

var _ Error = Forbidden("")

// Forbidden is a 403 HTTP error.
type Forbidden string

func (e Forbidden) Error() string { return string(e) }

// Code returns HTTP status code.
func (e Forbidden) Code() int { return http.StatusForbidden }

// NewForbidden is a constructor func for 403 HTTP error.
func NewForbidden(cause error) Forbidden { return Forbidden(cause.Error()) }

// Forbiddenf returns formatted Forbidden error.
func Forbiddenf(format string, args ...interface{}) Forbidden {
	return Forbidden(fmt.Sprintf(format, args...))
}
