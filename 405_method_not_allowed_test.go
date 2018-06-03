package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewMethodNotAllowed(t *testing.T) {
	t.Run("Check MethodNotAllowed status code", func(t *testing.T) {
		var err error
		err = NewMethodNotAllowed(errors.New("MethodNotAllowed"))
		typed, ok := err.(MethodNotAllowed)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusMethodNotAllowed {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "MethodNotAllowed" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
