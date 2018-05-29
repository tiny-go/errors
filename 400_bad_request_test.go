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
		typed, ok := err.(BadRequest)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusBadRequest {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "BadRequest" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
