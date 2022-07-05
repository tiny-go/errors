package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewUnauthorized(t *testing.T) {
	t.Run("NewUnauthorized", func(t *testing.T) {
		err := errors.NewUnauthorized(errors.New("Unauthorized"))

		if err.Code() != http.StatusUnauthorized {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Unauthorized" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestUnauthorizedf(t *testing.T) {
	t.Run("Unauthorizedf", func(t *testing.T) {
		err := errors.Unauthorizedf("Unauthorized: %d", http.StatusUnauthorized)

		if err.Code() != http.StatusUnauthorized {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Unauthorized: 401" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
