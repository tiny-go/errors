package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewForbidden(t *testing.T) {
	t.Run("NewForbidden", func(t *testing.T) {
		err := errors.NewForbidden(errors.New("Forbidden"))

		if err.Code() != http.StatusForbidden {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Forbidden" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestForbiddenf(t *testing.T) {
	t.Run("Forbiddenf", func(t *testing.T) {
		err := errors.Forbiddenf("Forbidden: %d", http.StatusForbidden)

		if err.Code() != http.StatusForbidden {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Forbidden: 403" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
