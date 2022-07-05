package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewNotAcceptable(t *testing.T) {
	t.Run("NewNotAcceptable", func(t *testing.T) {
		err := errors.NewNotAcceptable(errors.New("NotAcceptable"))

		if err.Code() != http.StatusNotAcceptable {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotAcceptable" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestNotAcceptablef(t *testing.T) {
	t.Run("NotAcceptablef", func(t *testing.T) {
		err := errors.NotAcceptablef("NotAcceptable: %d", http.StatusNotAcceptable)

		if err.Code() != http.StatusNotAcceptable {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "NotAcceptable: 406" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
