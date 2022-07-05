package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewMethodNotAllowed(t *testing.T) {
	t.Run("NewMethodNotAllowed", func(t *testing.T) {
		err := errors.NewMethodNotAllowed(errors.New("MethodNotAllowed"))

		if err.Code() != http.StatusMethodNotAllowed {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "MethodNotAllowed" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestMethodNotAllowedf(t *testing.T) {
	t.Run("MethodNotAllowedf", func(t *testing.T) {
		err := errors.MethodNotAllowedf("MethodNotAllowed: %d", http.StatusMethodNotAllowed)

		if err.Code() != http.StatusMethodNotAllowed {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "MethodNotAllowed: 405" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
