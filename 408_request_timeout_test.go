package errors_test

import (
	"net/http"
	"testing"

	"github.com/tiny-go/errors"
)

func TestNewRequestTimeout(t *testing.T) {
	t.Run("NewRequestTimeout", func(t *testing.T) {
		err := errors.NewRequestTimeout(errors.New("RequestTimeout"))

		if err.Code() != http.StatusRequestTimeout {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "RequestTimeout" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}

func TestRequestTimeoutf(t *testing.T) {
	t.Run("RequestTimeoutf", func(t *testing.T) {
		err := errors.RequestTimeoutf("RequestTimeout: %d", http.StatusRequestTimeout)

		if err.Code() != http.StatusRequestTimeout {
			t.Errorf("error has invalid status code: %d", err.Code())
		}

		if err.Error() != "RequestTimeout: 408" {
			t.Errorf("error has invalid message: %q", err.Error())
		}
	})
}
