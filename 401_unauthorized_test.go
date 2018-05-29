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
		unauthorized, ok := err.(Unauthorized)
		if !ok {
			t.Fatalf("error has invalid type: %T", err)
		}
		if unauthorized.Code() != http.StatusUnauthorized {
			t.Errorf("error has invalid status code: %d", unauthorized.Code())
		}
		if unauthorized.Error() != "Unauthorized" {
			t.Errorf("error has invalid message: %q", unauthorized.Error())
		}
	})
}
