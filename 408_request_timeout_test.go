package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewRequestTimeout(t *testing.T) {
	t.Run("Check RequestTimeout status code", func(t *testing.T) {
		var err error
		err = NewRequestTimeout(errors.New("RequestTimeout"))
		typed, ok := err.(RequestTimeout)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusRequestTimeout {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "RequestTimeout" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
