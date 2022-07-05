package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewConflict(t *testing.T) {
	t.Run("NewConflict", func(t *testing.T) {
		err := errors.NewConflict(errors.New("Conflict"))

		if err.Code() != http.StatusConflict {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Conflict" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestConflictf(t *testing.T) {
	t.Run("Conflictf", func(t *testing.T) {
		err := errors.Conflictf("Conflict: %d", http.StatusConflict)

		if err.Code() != http.StatusConflict {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "Conflict: 409" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
