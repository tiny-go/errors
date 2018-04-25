package errors

import (
	"errors"
	"net/http"
)

var _ Error = BadRequest{}

// BadRequest is a 400 HTTP error.
type BadRequest struct {
	error
}

// Code returns HTTP staus code.
func (e BadRequest) Code() int {
	return http.StatusBadRequest
}

// NewBadRequest is a constructor func for 400 HTTP error.
func NewBadRequest(cause string) BadRequest {
	return BadRequest{errors.New(cause)}
}
