package errors

import (
	"fmt"
	"net/http"
)

var _ Error = Conflict("")

// Conflict is a 409 HTTP error.
type Conflict string

func (e Conflict) Error() string { return string(e) }

// Code returns HTTP status code.
func (e Conflict) Code() int { return http.StatusConflict }

// NewConflict is a constructor func for 409 HTTP error.
func NewConflict(cause error) Conflict { return Conflict(cause.Error()) }

// Conflictf returns formatted Conflict error.
func Conflictf(format string, args ...interface{}) Conflict {
	return Conflict(fmt.Sprintf(format, args...))
}
