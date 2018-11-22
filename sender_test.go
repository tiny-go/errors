package errors

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"testing"
)

var (
	_ Sender     = Send
	_ Middleware = WithLogger(nil)
	_ Middleware = (*ResponseLogger)(nil).Log
)

func TestWithLogger(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	logger := log.New(buff, "", 0)

	t.Run("test if the response is logged by logger wrapper", func(t *testing.T) {
		WithLogger(logger)(func(http.ResponseWriter, interface{}) {})(nil, errors.New("failed"))
		if buff.String() != "failed\n" {
			t.Errorf("log output was expected to contain error message %q but instead of %q", "failed\n", buff.String())
		}
	})
}

func TestResponseLogger(t *testing.T) {
	buff := bytes.NewBuffer([]byte{})
	logger := log.New(buff, "", 0)

	t.Run("test if the response is logged by ResponseLogger", func(t *testing.T) {
		(&ResponseLogger{logger}).Log(func(http.ResponseWriter, interface{}) {})(nil, errors.New("failed"))
		if buff.String() != "failed\n" {
			t.Errorf("log output was expected to contain error message %q but instead of %q", "failed\n", buff.String())
		}
	})
}
