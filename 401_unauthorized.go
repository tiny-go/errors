package errors

import (
	"fmt"
	"net/http"
)

var _ Error = Unauthorized("")

// Unauthorized is a 401 HTTP error.
type Unauthorized string

func (e Unauthorized) Error() string { return string(e) }

// Code returns HTTP status code.
func (e Unauthorized) Code() int { return http.StatusUnauthorized }

// NewUnauthorized is a constructor func for 401 HTTP error.
func NewUnauthorized(cause error) Unauthorized { return Unauthorized(cause.Error()) }

// Unauthorizedf returns formatted Unauthorized error.
func Unauthorizedf(format string, args ...interface{}) Unauthorized {
	return Unauthorized(fmt.Sprintf(format, args...))
}
