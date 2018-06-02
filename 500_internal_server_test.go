package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestNewInternalServer(t *testing.T) {
	t.Run("Check InternalServer status code", func(t *testing.T) {
		var err error
		err = NewInternalServer(errors.New("InternalServer"))
		typed, ok := err.(InternalServer)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if typed.Code() != http.StatusInternalServerError {
			t.Errorf("error has invalid status code: %d", typed.Code())
		}
		if typed.Error() != "InternalServer" {
			t.Errorf("error has invalid message: %q", typed.Error())
		}
	})
}
