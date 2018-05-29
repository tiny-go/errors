package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewUnauthorized(t *testing.T) {
	t.Run("Check Unauthorized status code", func(t *testing.T) {
		var err error
		err = NewUnauthorized(errors.New("Unauthorized"))
		typed, ok := err.(Unauthorized)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusUnauthorized {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "Unauthorized" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
