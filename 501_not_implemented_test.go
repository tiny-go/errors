package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewNotImplemented(t *testing.T) {
	t.Run("Check NotImplemented status code", func(t *testing.T) {
		var err error
		err = NewNotImplemented(errors.New("NotImplemented"))
		typed, ok := err.(NotImplemented)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusNotImplemented {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "NotImplemented" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
