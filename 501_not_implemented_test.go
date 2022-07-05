package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewNotImplemented(t *testing.T) {
	t.Run("NewNotImplemented", func(t *testing.T) {
		err := errors.NewNotImplemented(errors.New("NotImplemented"))

		if err.Code() != http.StatusNotImplemented {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotImplemented" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestNotImplementedf(t *testing.T) {
	t.Run("NotImplemented", func(t *testing.T) {
		err := errors.NotImplementedf("NotImplemented: %d", http.StatusNotImplemented)

		if err.Code() != http.StatusNotImplemented {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotImplemented: 501" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
