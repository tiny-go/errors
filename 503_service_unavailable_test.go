package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewServiceUnavailable(t *testing.T) {
	t.Run("NewServiceUnavailable", func(t *testing.T) {
		err := errors.NewServiceUnavailable(errors.New("ServiceUnavailable"))

		if err.Code() != http.StatusServiceUnavailable {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "ServiceUnavailable" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestServiceUnavailablef(t *testing.T) {
	t.Run("ServiceUnavailable", func(t *testing.T) {
		err := errors.ServiceUnavailablef("ServiceUnavailable: %d", http.StatusServiceUnavailable)

		if err.Code() != http.StatusServiceUnavailable {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "ServiceUnavailable: 503" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
