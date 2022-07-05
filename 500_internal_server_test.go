package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewInternalServer(t *testing.T) {
	t.Run("NewInternalServer", func(t *testing.T) {
		err := errors.NewInternalServer(errors.New("InternalServer"))

		if err.Code() != http.StatusInternalServerError {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "InternalServer" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestInternalServerf(t *testing.T) {
	t.Run("InternalServerf", func(t *testing.T) {
		err := errors.InternalServerf("InternalServer: %d", http.StatusInternalServerError)

		if err.Code() != http.StatusInternalServerError {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "InternalServer: 500" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
