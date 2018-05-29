package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewForbidden(t *testing.T) {
	t.Run("Check Forbidden status code", func(t *testing.T) {
		var err error
		err = NewForbidden(errors.New("Forbidden"))
		typed, ok := err.(Forbidden)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusForbidden {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "Forbidden" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
