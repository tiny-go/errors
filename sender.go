package errors

import "net/http"

// Logger is a simple error logger.
type Logger interface{ Println(...interface{}) }

// Sender is a function that sends an error through the provided response writer.
type Sender func(http.ResponseWriter, interface{})

// Middleware is a chainable sender wrapper func.
type Middleware func(Sender) Sender

// WithLogger wrapps provided sender with logger.
//
// Example: errors.LoggerWrapper(logger)(errors.Send)(w, err)
func WithLogger(logger Logger) Middleware {
	return func(sender Sender) Sender {
		return func(w http.ResponseWriter, err interface{}) {
			logger.Println(err)
			sender(w, err)
		}
	}
}

// ResponseLogger provides another (object oriented) approach to log the error.
//
// Example: (&errors.ResponseLogger{Logger: logger}).Log(errors.Send)(w, err)
type ResponseLogger struct{ Logger }

// Log creates Sender middleware func.
func (rl *ResponseLogger) Log(sender Sender) Sender {
	return func(w http.ResponseWriter, err interface{}) {
		rl.Println(err)
		sender(w, err)
	}
}
