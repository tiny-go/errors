package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewBadRequest(t *testing.T) {
	t.Run("NewBadRequest", func(t *testing.T) {
		err := errors.NewBadRequest(errors.New("BadRequest"))

		if err.Code() != http.StatusBadRequest {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "BadRequest" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestBadRequestf(t *testing.T) {
	t.Run("BadRequestf", func(t *testing.T) {
		err := errors.BadRequestf("BadRequest: %d", http.StatusBadRequest)

		if err.Code() != http.StatusBadRequest {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "BadRequest: 400" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
