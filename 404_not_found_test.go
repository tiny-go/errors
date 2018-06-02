package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewNotFound(t *testing.T) {
	t.Run("Check NotFound status code", func(t *testing.T) {
		var err error
		err = NewNotFound(errors.New("NotFound"))
		typed, ok := err.(NotFound)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusNotFound {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "NotFound" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
