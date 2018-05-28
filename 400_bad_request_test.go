package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewBadRequest(t *testing.T) {
	t.Run("Check Bad Request status code", func(t *testing.T) {
		var err error
		err = NewBadRequest(errors.New("BadRequest"))
		badRequest, ok := err.(BadRequest)
		if !ok {
			t.Error("error has invalid type")
		}
		if badRequest.Code() != http.StatusBadRequest {
			t.Error("error has invalid status code")
		}
		if badRequest.Error() != "BadRequest" {
			t.Error("error has invalid message")
		}
	})
}
