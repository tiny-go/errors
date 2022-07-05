package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewNotFound(t *testing.T) {
	t.Run("NewNotFound", func(t *testing.T) {
		err := errors.NewNotFound(errors.New("NotFound"))

		if err.Code() != http.StatusNotFound {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotFound" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestNotFoundf(t *testing.T) {
	t.Run("NotFoundf", func(t *testing.T) {
		err := errors.NotFoundf("NotFound: %d", http.StatusNotFound)

		if err.Code() != http.StatusNotFound {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotFound: 404" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
