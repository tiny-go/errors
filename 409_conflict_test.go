package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewConflict(t *testing.T) {
	t.Run("Check Conflict status code", func(t *testing.T) {
		var err error
		err = NewConflict(errors.New("Conflict"))
		typed, ok := err.(Conflict)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusConflict {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "Conflict" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
